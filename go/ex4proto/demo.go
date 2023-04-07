// Package ex4proto works with data stored in a Protobuf message.
//
// For simplicity, these examples do not handle list or map values.
package ex4proto

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb "workshop/lib/workshop/v1"
)

// GetValue expects path in snake_case.
func GetValue(m proto.Message, path []string) protoreflect.Value {
	if len(path) == 0 {
		return protoreflect.Value{}
	}

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	// get the field
	fd := d.Fields().ByName(protoreflect.Name(path[0]))
	if fd == nil {
		return protoreflect.Value{}
	}
	v := mp.Get(fd)
	if len(path) == 1 {
		return v
	}

	// try to go deeper
	if fd.Kind() != protoreflect.MessageKind {
		return protoreflect.Value{}
	}
	vm := v.Message().Interface()
	return GetValue(vm, path[1:])
}

// SetValue expects path in snake_case.
func SetValue(m proto.Message, path []string, val protoreflect.Value) {
	if len(path) == 0 {
		return
	}

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	// get the field
	fd := d.Fields().ByName(protoreflect.Name(path[0]))
	if fd == nil || string(fd.Name()) != path[0] {
		return
	}
	if len(path) == 1 {
		mp.Set(fd, val)
		return
	}

	// try to go deeper
	if fd.Kind() != protoreflect.MessageKind {
		return
	}
	v := mp.Get(fd)
	if !v.IsValid() {
		return
	}
	vm := v.Message().Interface()
	SetValue(vm, path[1:], val)
}

func DepthFirstTraversal(m proto.Message) []protoreflect.Value {
	var out []protoreflect.Value

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	for i := 0; i < d.Fields().Len(); i += 1 {
		fd := d.Fields().Get(i)
		v := mp.Get(fd)

		if fd.Kind() != protoreflect.MessageKind {
			out = append(out, v)
			continue
		}

		// go deeper
		vm := v.Message().Interface()
		out = append(out, DepthFirstTraversal(vm)...)
	}
	return out
}

func Create() *pb.Apple {
	apple := &pb.Apple{}

	brand := "granny-smith"
	apple.Brand = brand
	age := 42
	apple.Age = int32(age)

	skin := &pb.Apple_Skin{}
	skin.Colour = "green"
	skin.Blemishes = 3
	apple.Skin = skin

	return apple
}

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", DepthFirstTraversal(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(apple, []string{"skin", "blemishes"}, protoreflect.ValueOfInt32(4))
	fmt.Println("after update\t:", apple)
}

func Apply(m proto.Message) {
	mp := m.ProtoReflect()
	mp.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if Hide(fd) {
			mp.Clear(fd)
		}
		if fd.Kind() == protoreflect.MessageKind {
			// go deeper
			Apply(v.Message().Interface())
		}
		return true
	})
}

func Hide(fd protoreflect.FieldDescriptor) bool {
	opt := fd.Options()
	if opt == nil {
		return false
	}
	x := proto.GetExtension(opt, pb.E_MyAnnotation)
	if x == (*pb.MyAnnotation)(nil) {
		return false
	}
	a, ok := x.(*pb.MyAnnotation)
	if !ok {
		return false
	}
	return a.Hide
}
