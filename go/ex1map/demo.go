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

	for k, v := range m {
		mapVal, isMap := v.(map[string]any)

		if k != path[0] {
			continue
		}
		// work on current field?
		if len(path) == 1 {
			return m[k]
		}

		// try to go deeper
		if !isMap {
			return nil
		}
		return GetValue(mapVal, path[1:])
	}
	return nil
}

func SetValue(m map[string]any, path []string, val any) {
	if len(path) == 0 {
		return
	}

	for k, v := range m {
		mapVal, isMap := v.(map[string]any)

		if k != path[0] {
			continue
		}
		// work on current field?
		if len(path) == 1 {
			m[k] = val
			return
		}

		// try to go deeper
		if !isMap {
			return
		}
		SetValue(mapVal, path[1:], val)
		return
	}
	m[path[0]] = val
}

func Traverse(m map[string]any) []any {
	var out []any

	for _, v := range m {
		mapVal, isMap := v.(map[string]any)

		if !isMap {
			out = append(out, v)
			continue
		}

		// go deeper
		out = append(out, Traverse(mapVal)...)
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
