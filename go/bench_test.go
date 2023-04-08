package main_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
	"workshop/ex1map"
	"workshop/ex2tree"
	"workshop/ex2tree/tree"
	"workshop/ex3reflect"
	"workshop/ex4proto"
)

func BenchmarkGetValue(b *testing.B) {
	b.Run("map", func(b *testing.B) {
		m := ex1map.Create()
		for i := 0; i < b.N; i += 1 {
			ex1map.GetValue(m, path[0])
			ex1map.GetValue(m, path[1])
		}
	})

	b.Run("tree", func(b *testing.B) {
		m := ex2tree.Create()
		for i := 0; i < b.N; i += 1 {
			ex2tree.GetValue(m, path[0])
			ex2tree.GetValue(m, path[1])
		}
	})

	b.Run("reflect", func(b *testing.B) {
		m := ex3reflect.Create()
		for i := 0; i < b.N; i += 1 {
			ex3reflect.GetValue(m, pathPascal[0])
			ex3reflect.GetValue(m, pathPascal[1])
		}
	})

	b.Run("reflect pointer", func(b *testing.B) {
		m := ex3reflect.Create()
		for i := 0; i < b.N; i += 1 {
			ex3reflect.GetValue(&m, pathPascal[0])
			ex3reflect.GetValue(&m, pathPascal[1])
		}
	})

	b.Run("proto", func(b *testing.B) {
		m := ex4proto.Create()
		for i := 0; i < b.N; i += 1 {
			ex4proto.GetValue(m, path[0])
			ex4proto.GetValue(m, path[1])
		}
	})
}

func BenchmarkSetValue(b *testing.B) {
	b.Run("map", func(b *testing.B) {
		m := ex1map.Create()
		for i := 0; i < b.N; i += 1 {
			ex1map.SetValue(m, path[0], i)
			ex1map.SetValue(m, path[1], i)
		}
	})

	b.Run("tree", func(b *testing.B) {
		m := ex2tree.Create()
		for i := 0; i < b.N; i += 1 {
			v := tree.ValueOfInt(i)
			ex2tree.SetValue(m, path[0], v)
			ex2tree.SetValue(m, path[1], v)
		}
	})

	b.Run("reflect", func(b *testing.B) {
		m := ex3reflect.Create()
		for i := 0; i < b.N; i += 1 {
			v := int32(i)
			ex3reflect.SetValue(&m, pathPascal[0], v)
			ex3reflect.SetValue(&m, pathPascal[1], v)
		}
	})

	b.Run("proto", func(b *testing.B) {
		m := ex4proto.Create()
		for i := 0; i < b.N; i += 1 {
			v := int32(i)
			ex4proto.SetValue(m, path[0], v)
			ex4proto.SetValue(m, path[1], v)
		}
	})
}

func BenchmarkSerialisation(b *testing.B) {
	b.Run("JSON map", func(b *testing.B) {
		m := ex1map.Create()
		for i := 0; i < b.N; i += 1 {
			_, _ = json.Marshal(m)
		}
	})

	b.Run("JSON tree", func(b *testing.B) {
		m := ex2tree.Create()
		for i := 0; i < b.N; i += 1 {
			_, _ = json.Marshal(m)
		}
	})

	b.Run("JSON struct", func(b *testing.B) {
		m := ex3reflect.Create()
		for i := 0; i < b.N; i += 1 {
			_, _ = json.Marshal(m)
		}
	})

	b.Run("Protobuf proto", func(b *testing.B) {
		m := ex4proto.Create()
		for i := 0; i < b.N; i += 1 {
			_, _ = proto.Marshal(m)
		}
	})

	bs, _ := json.Marshal(ex1map.Create())
	fmt.Printf("map: %dB\n", len(bs))
	bs, _ = json.Marshal(ex2tree.Create())
	fmt.Printf("tree: %dB\n", len(bs))
	bs, _ = json.Marshal(ex3reflect.Create())
	fmt.Printf("struct: %dB\n", len(bs))
	bs, _ = proto.Marshal(ex4proto.Create())
	fmt.Printf("proto %dB\n", len(bs))
}

func BenchmarkApply(b *testing.B) {
	b.Run("reflect", func(b *testing.B) {
		m := ex3reflect.Create()
		for i := 0; i < b.N; i += 1 {
			ex3reflect.Apply(&m)
		}
	})

	b.Run("proto", func(b *testing.B) {
		m := ex4proto.Create()
		for i := 0; i < b.N; i += 1 {
			ex4proto.Apply(m)
		}
	})
}

var path = [][]string{
	{"age"},
	{"skin", "blemishes"},
	{"skin", "smell"},
}
var pathPascal = [][]string{
	{"Age"},
	{"Skin", "Blemishes"},
	{"Skin", "Smell"},
}
