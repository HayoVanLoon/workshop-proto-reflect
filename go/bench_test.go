package main

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"reflect"
	"testing"
	"workshop/withmap"
	"workshop/withproto"
	"workshop/withreflect"
	"workshop/withtree"
	"workshop/withtree/tree"
)

func Benchmark(b *testing.B) {
	path := [][]string{
		{"age"},
		{"skin", "blemishes"},
	}

	b.Run("with map", func(b *testing.B) {
		m := withmap.Create()
		for i := 0; i < b.N; i += 1 {
			withmap.SetValue(m, path[0], i)
			withmap.SetValue(m, path[1], i)
		}
	})

	b.Run("with struct", func(b *testing.B) {
		m := withtree.Create()
		for i := 0; i < b.N; i += 1 {
			withtree.SetValue(m, path[0], tree.ValueOfInt(i))
			withtree.SetValue(m, path[1], tree.ValueOfInt(i))
		}
	})

	b.Run("with reflect", func(b *testing.B) {
		m := withreflect.Create()
		for i := 0; i < b.N; i += 1 {
			withreflect.SetValue(&m, path[0], reflect.ValueOf(int32(i)))
			withreflect.SetValue(&m, path[1], reflect.ValueOf(int32(i)))
		}
	})

	b.Run("with proto reflect", func(b *testing.B) {
		m := withproto.Create()
		for i := 0; i < b.N; i += 1 {
			withproto.SetValue(m, path[0], protoreflect.ValueOf(int32(i)))
			withproto.SetValue(m, path[1], protoreflect.ValueOf(int32(i)))
		}
	})
}
