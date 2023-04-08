// Package ex2tree  works with data stored in a tree.
//
// For simplicity, these examples do not handle list or map values.
package ex2tree

import (
	"fmt"

	"workshop/ex2tree/tree"
)

func Create() tree.Tree {
	apple := tree.NewTree()

	brand := tree.ValueOfString("granny-smith")
	apple.Set("brand", brand)
	age := tree.ValueOfInt(42)
	apple.Set("age", age)

	skin := tree.NewTree()
	skin.Set("colour", tree.ValueOfString("green"))
	skin.Set("blemishes", tree.ValueOfInt(3))
	apple.Set("skin", tree.ValueOfMessage(skin))

	return apple
}

func GetValue(m tree.Tree, path []string) tree.Value {
	if len(path) == 0 {
		return nil
	}

	for _, kv := range m.Children() {
		k := kv.Key
		v := kv.Value

		// work with current field?
		switch {
		case k != path[0]:
			continue
		case len(path) == 1:
			return v
		case v.Tree() == nil:
			return nil
		}

		// go deeper
		return GetValue(v.Tree(), path[1:])
	}
	return nil
}

func SetValue(m tree.Tree, path []string, val tree.Value) {
	if len(path) == 0 {
		return
	}

	for _, kv := range m.Children() {
		k := kv.Key
		v := kv.Value

		// work on current field?
		switch {
		case k != path[0]:
			continue
		case len(path) == 1:
			m.Set(k, val)
			return
		case v.Tree() == nil:
			return
		}

		// go deeper
		SetValue(v.Tree(), path[1:], val)
	}
}

func Traverse(m tree.Tree) []tree.Value {
	if m == nil {
		return nil
	}
	var out []tree.Value

	for _, kv := range m.Children() {
		v := kv.Value

		if v.Type() != tree.ValueTypeTree {
			out = append(out, v)
			continue
		}

		// go deeper
		out = append(out, Traverse(v.Tree())...)
	}
	return out
}

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", Traverse(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(apple, []string{"skin", "blemishes"}, tree.ValueOfInt(4))
	fmt.Println("after update\t:", apple)
}
