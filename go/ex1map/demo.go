// Package ex1map works with data stored in a map
//
// For simplicity, these examples do not handle list values.
package ex1map

import (
	"fmt"
)

func GetValue(m map[string]any, path []string) any {
	if len(path) == 0 {
		return nil
	}

	for k, v := range m {
		mapVal, isMap := v.(map[string]any)

		// work on current field?
		switch {
		case k != path[0]:
			continue
		case len(path) == 1:
			return v
		case !isMap:
			return nil
		}

		// go deeper
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

		// work on current field?
		switch {
		case k != path[0]:
			continue
		case len(path) == 1:
			m[k] = val
			return
		case !isMap:
			return
		}

		// go deeper
		SetValue(mapVal, path[1:], val)
	}
}

func DepthFirstTraversal(m map[string]any) []any {
	var out []any

	for _, v := range m {
		mapVal, isMap := v.(map[string]any)

		if !isMap {
			out = append(out, v)
			continue
		}

		// go deeper
		out = append(out, DepthFirstTraversal(mapVal)...)
	}
	return out
}

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

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", DepthFirstTraversal(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(apple, []string{"skin", "blemishes"}, 4)
	fmt.Println("after update\t:", apple)
}