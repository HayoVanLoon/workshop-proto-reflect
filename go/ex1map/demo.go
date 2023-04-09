// Package ex1map works with data stored in a map
//
// For simplicity, these examples do not handle list values.
package ex1map

import (
	"fmt"
)

func Create() map[string]any {
	apple := make(map[string]any)

	brand := "granny-smith"
	apple["brand"] = brand
	age := 42
	apple["age"] = age

	skin := make(map[string]any)
	skin["colour"] = "green"
	skin["blemishes"] = 3
	apple["skin"] = skin

	return apple
}

func GetValue(m map[string]any, path []string) any {
	if len(path) == 0 {
		return nil
	}

	// follow the path
	v, ok := m[path[0]]
	if !ok {
		return nil
	}

	// are we there yet?
	if len(path) == 1 {
		return v
	}

	// try to go deeper
	mapVal, isMap := v.(map[string]any)
	if !isMap {
		return nil
	}
	return GetValue(mapVal, path[1:])
}

func SetValue(m map[string]any, path []string, val any) {
	if len(path) == 0 {
		return
	}

	// follow the path
	v, ok := m[path[0]]

	// are we there yet?
	if len(path) == 1 {
		m[path[0]] = val
		return
	}

	// try to go deeper
	var mapVal map[string]any
	if ok {
		var isMap bool
		if mapVal, isMap = v.(map[string]any); !isMap {
			return
		}
	} else {
		mapVal = make(map[string]any)
	}
	m[path[0]] = mapVal
	SetValue(mapVal, path[1:], val)
}

func Traverse(m map[string]any) []any {
	var out []any

	for _, v := range m {
		// get field info
		mapVal, isMap := v.(map[string]any)

		if !isMap {
			// operate on simple value
			out = append(out, v)
		} else {
			// operate on nested object
			out = append(out, Traverse(mapVal)...)
		}
	}
	return out
}

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", Traverse(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(apple, []string{"skin", "blemishes"}, 4)
	fmt.Println("after update\t:", apple)
}
