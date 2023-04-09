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

	var v tree.Value
	for _, kv := range m.Children() {
		if kv.Key == path[0] {
			v = kv.Value
			break
		}
	}

	// are we there yet?
	if len(path) == 1 {
		return v
	}

	// try to go deeper
	if v.Tree() == nil {
		return nil
	}
	return GetValue(v.Tree(), path[1:])
}

func SetValue(m tree.Tree, path []string, val tree.Value) {
	if len(path) == 0 {
		return
	}

	var v tree.Value
	for _, kv := range m.Children() {
		if kv.Key == path[0] {
			v = kv.Value
			break
		}
	}

	// are we there yet?
	if len(path) == 1 {
		// check type
		if v != nil && v.Type() != val.Type() {
			return
		}
		m.Set(path[0], val)
		return
	}

	// try to go deeper
	if v == nil {
		v = tree.ValueOfMessage(tree.NewTree())
		m.Set(path[0], v)
	} else if v.Type() != tree.ValueTypeTree {
		return
	}
	SetValue(v.Tree(), path[1:], val)
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
