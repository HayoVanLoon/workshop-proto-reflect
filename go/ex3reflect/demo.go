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

	// get the field
	v := mv.FieldByName(path[0])
	if !v.IsValid() {
		return nil
	}
	if len(path) == 1 {
		return v.Interface()
	}

	// try to go deeper
	if v.Type().Kind() != reflect.Struct {
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

	// get the field
	fd, ok := t.FieldByName(path[0])
	if !ok {
		return
	}
	v := mv.FieldByName(path[0])
	if len(path) == 1 {
		v.Set(val)
		return
	}

	// try to go deeper
	if v.Type().Kind() != reflect.Struct {
		return
	}
	var vm any
	if v.IsValid() {
		vm = v.Addr().Interface()
	} else {
		// might work for non-initialised pointer values; untested
		vm = reflect.New(fd.Type)
	}
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
		v := mv.Field(i)

		if v.Type().Kind() != reflect.Struct {
			out = append(out, v.Interface())
			continue
		}

		// go deeper
		vm := v.Interface()
		out = append(out, Traverse(vm)...)
	}
	return out
}

func SchemaFor(m any) map[string]any {
	return schemaFor(reflect.TypeOf(m))
}

func schemaFor(t reflect.Type) map[string]any {
	out := make(map[string]any)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i += 1 {
		fd := t.Field(i)
		var v any
		if fd.Type.Kind() != reflect.Struct {
			v = fd.Type.Kind().String()
		} else {
			v = schemaFor(fd.Type)
		}
		out[strings.ToLower(fd.Name)] = v
	}
	return out
}

func Apply(m any) {
	// "get into reflect mode"
	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	for i := 0; i < t.NumField(); i += 1 {
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
	schema := SchemaFor(apple)
	fmt.Println("schema\t\t:", schema)
	Apply(&apple)
	fmt.Printf("after apply\t: %+v\n", apple)
}
