// Package ex3reflect works with data stored in a struct.
//
// For simplicity, these examples do not handle list or map values.
package ex3reflect

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Apple struct {
	Brand string
	Age   int32
	Skin  AppleSkin
}

type AppleSkin struct {
	Colour    string
	Blemishes int32 `hide:"true"`
}

func Create() Apple {
	apple := Apple{}

	brand := "granny-smith"
	apple.Brand = brand
	age := 42
	apple.Age = int32(age)

	skin := AppleSkin{}
	skin.Colour = "green"
	skin.Blemishes = 3
	apple.Skin = skin

	return apple
}

// GetValue expects path in PascalCase.
func GetValue(m any, path []string) any {
	if len(path) == 0 {
		return nil
	}

	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	// check if path section exists; get field info
	ft, ok := t.FieldByName(path[0])
	if !ok {
		return nil
	}

	// follow the path
	v := mv.FieldByName(path[0])
	if !v.IsValid() {
		return nil
	}

	// are we there yet?
	if len(path) == 1 {
		return v.Interface()
	}

	// try to go deeper
	if ft.Type.Kind() != reflect.Struct {
		return nil
	}
	vm := v.Interface()
	return GetValue(vm, path[1:])
}

// SetValue expects path in PascalCase.
func SetValue(m any, path []string, value any) {
	if len(path) == 0 {
		return
	}
	val := reflect.ValueOf(value)

	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	// check if path section exists; get field info
	ft, ok := t.FieldByName(path[0])
	if !ok {
		return
	}

	// are we there yet?
	if len(path) == 1 {
		v := mv.FieldByName(path[0])
		v.Set(val)
		return
	}

	// try to go deeper
	if ft.Type.Kind() != reflect.Struct {
		return
	}
	v := mv.FieldByName(path[0])
	if !v.IsValid() {
		// might work for non-initialised pointer values; untested
		v = reflect.New(ft.Type)
	}
	vm := v.Addr().Interface()
	SetValue(vm, path[1:], value)
}

func Traverse(m any) []any {
	var out []any

	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	for i := 0; i < t.NumField(); i += 1 {
		// get field info
		ft := t.Field(i)
		v := mv.Field(i)

		if ft.Type.Kind() != reflect.Struct {
			// operate on simple value
			out = append(out, v.Interface())
		} else {
			// go deeper
			vm := v.Interface()
			out = append(out, Traverse(vm)...)
		}
	}
	return out
}

func SchemaFor(m any, maxDepth int) map[string]any {
	// "get into reflect mode"
	return schemaFor(reflect.TypeOf(m), maxDepth)
}

func schemaFor(t reflect.Type, maxDepth int) map[string]any {
	out := make(map[string]any)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i += 1 {
		// get field info
		fd := t.Field(i)

		if fd.Type.Kind() != reflect.Struct {
			// operate on simple field
			out[strings.ToLower(fd.Name)] = fd.Type.Kind().String()
		} else {
			if maxDepth <= 0 {
				continue
			}
			// go deeper
			out[strings.ToLower(fd.Name)] = schemaFor(fd.Type, maxDepth-1)
		}
	}
	return out
}

func Apply(m any) {
	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	for i := 0; i < t.NumField(); i += 1 {
		// get field info
		fd := t.Field(i)
		v := mv.Field(i)

		if Hide(fd) {
			v.SetZero()
			continue
		}
		if fd.Type.Kind() == reflect.Struct {
			// go deeper
			vm := v.Addr().Interface()
			Apply(vm)
		}
	}
}

func Hide(fd reflect.StructField) bool {
	b, _ := strconv.ParseBool(fd.Tag.Get("hide"))
	return b
}

func Run() {
	apple := Create()
	fmt.Printf("apple\t\t: %+v\n", apple)
	fmt.Println("traversal\t:", Traverse(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"Skin", "Blemishes"}))
	SetValue(&apple, []string{"Skin", "Blemishes"}, int32(4))
	fmt.Printf("after update\t: %+v\n", apple)
	schema := SchemaFor(apple, 10)
	fmt.Println("schema\t\t:", schema)
	Apply(&apple)
	fmt.Printf("after apply\t: %+v\n", apple)
}
