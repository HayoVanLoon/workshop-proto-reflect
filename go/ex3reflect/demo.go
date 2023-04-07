// Package ex3reflect works with data stored in a struct.
//
// For simplicity, these examples do not handle list or map values.
package ex3reflect

import (
	"fmt"
	"reflect"
	"strconv"
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
func SetValue(m any, path []string, val reflect.Value) {
	if len(path) == 0 {
		return
	}

	// "get into reflect mode"
	t := reflect.TypeOf(m)
	mv := reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
		mv = mv.Elem()
	}

	// get the field
	if _, ok := t.FieldByName(path[0]); !ok {
		return
	}
	v := mv.FieldByName(path[0])
	if len(path) == 1 {
		v.Set(val)
		return
	}

	if v.Type().Kind() != reflect.Struct {
		return
	}
	// go deeper
	vm := v.Addr().Interface()
	SetValue(vm, path[1:], val)
}

func DepthFirstTraversal(m any) []any {
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
		out = append(out, DepthFirstTraversal(vm)...)
	}
	return out
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

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", DepthFirstTraversal(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"Skin", "Blemishes"}))
	SetValue(&apple, []string{"Skin", "Blemishes"}, reflect.ValueOf(int32(4)))
	fmt.Println("after update\t:", apple)
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
