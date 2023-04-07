package withreflect

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"strings"
)

type Apple struct {
	Brand string
	Age   int32
	Skin  AppleSkin
}

type AppleSkin struct {
	Colour    string
	Blemishes int32
}

func GetValue(m any, path []string) any {
	if len(path) == 0 {
		return nil
	}

	mr := reflect.TypeOf(m)
	fs := reflect.VisibleFields(mr)

	for i := 0; i < len(fs); i += 1 {
		v := reflect.ValueOf(m).FieldByIndex([]int{i})

		switch {
		case strings.ToLower(fs[i].Name) != path[0]:
			continue
		case len(path) == 1:
			return v
		case v.Type().Kind() != reflect.Struct:
			return nil
		}

		vm := v.Interface()
		return GetValue(vm, path[1:])
	}
	return protoreflect.Value{}
}

func SetValue(m any, path []string, val reflect.Value) {
	if len(path) == 0 {
		return
	}

	t := reflect.TypeOf(m).Elem()
	mp := reflect.ValueOf(m).Elem()

	for i := 0; i < t.NumField(); i += 1 {
		fd := t.Field(i)
		v := mp.Field(i)

		switch {
		case strings.ToLower(fd.Name) != path[0]:
			continue
		case len(path) == 1:
			v.Set(val)
			return
		case v.Type().Kind() != reflect.Struct:
			return
		}

		vm := v.Addr().Interface()
		SetValue(vm, path[1:], val)
	}
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
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(&apple, []string{"skin", "blemishes"}, reflect.ValueOf(int32(4)))
	fmt.Println("after update\t:", apple)
}
