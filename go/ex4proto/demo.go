// Package ex4proto works with data stored in a Protobuf message.
//
// For simplicity, these examples do not handle list or map values.
package ex4proto

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	pb "workshop/lib/workshop/v1"
)

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

// GetValue expects path in snake_case.
func GetValue(m proto.Message, path []string) any {
	if len(path) == 0 {
		return nil
	}

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	// check if path section exists; get field info
	fd := d.Fields().ByName(protoreflect.Name(path[0]))
	if fd == nil {
		return nil
	}

	// follow the path
	v := mp.Get(fd)
	if !v.IsValid() {
		return nil
	}

	// are we there yet?
	if len(path) == 1 {
		return v.Interface()
	}

	// try to go deeper
	if fd.Kind() != protoreflect.MessageKind {
		return nil
	}
	vm := v.Message().Interface()
	return GetValue(vm, path[1:])
}

// SetValue expects path in snake_case.
func SetValue(m proto.Message, path []string, value any) {
	if len(path) == 0 {
		return
	}
	val := protoreflect.ValueOf(value)

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	// check if path section exists; get field info
	fd := d.Fields().ByName(protoreflect.Name(path[0]))
	if fd == nil || string(fd.Name()) != path[0] {
		return
	}

	// are we there yet?
	if len(path) == 1 {
		mp.Set(fd, val)
		return
	}

	// try to go deeper
	if fd.Kind() != protoreflect.MessageKind {
		return
	}
	v := mp.Get(fd)
	if !v.Message().IsValid() {
		t, _ := protoregistry.GlobalTypes.FindMessageByName(fd.Message().FullName())
		v = protoreflect.ValueOfMessage(t.New())
		mp.Set(fd, v)
	}
	vm := v.Message().Interface()
	SetValue(vm, path[1:], value)
}

func Traverse(m proto.Message) []any {
	var out []any

	// "get into reflect mode"
	mp := m.ProtoReflect()
	d := mp.Descriptor()

	for i := 0; i < d.Fields().Len(); i += 1 {
		// get field info
		fd := d.Fields().Get(i)
		v := mp.Get(fd)

		if fd.Kind() != protoreflect.MessageKind {
			// operate on simple value
			out = append(out, v.Interface())
		} else {
			// go deeper
			vm := v.Message().Interface()
			out = append(out, Traverse(vm)...)
		}
	}
	return out
}

func SchemaFor(m proto.Message, maxDepth int) map[string]any {
	// "get into reflect mode"
	return schemaFor(m.ProtoReflect().Descriptor(), maxDepth)
}

func schemaFor(d protoreflect.MessageDescriptor, maxDepth int) map[string]any {
	out := make(map[string]any)

	fds := d.Fields()

	for i := 0; i < fds.Len(); i += 1 {
		// get field info
		fd := fds.Get(i)

		if fd.Kind() != protoreflect.MessageKind {
			// operate on simple field
			out[string(fd.Name())] = fd.Kind().String()
		} else {
			if maxDepth <= 0 {
				continue
			}
			// go deeper
			out[string(fd.Name())] = schemaFor(fd.Message(), maxDepth-1)
		}
	}

	return out
}

func Apply(m proto.Message) {
	// "get into reflect mode"
	mp := m.ProtoReflect()

	mp.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if Hide(fd) {
			mp.Clear(fd)
			return true
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
	return ok && a.Hide
}

func Run() {
	apple := Create()
	fmt.Println("apple\t\t:", apple)
	fmt.Println("traversal\t:", Traverse(apple))
	fmt.Println("skin.blemishes\t:", GetValue(apple, []string{"skin", "blemishes"}))
	SetValue(apple, []string{"skin", "blemishes"}, int32(4))
	fmt.Println("after update\t:", apple)
	schema := SchemaFor(apple, 10)
	fmt.Println("schema\t\t:", schema)
	Apply(apple)
	fmt.Println("after apply\t:", apple)
}
