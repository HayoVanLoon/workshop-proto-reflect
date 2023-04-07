package withproto

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb "workshop/lib/workshop/v1"
)

func GetValue(m proto.Message, path []string) protoreflect.Value {
	if len(path) == 0 {
		return protoreflect.Value{}
	}

	mp := m.ProtoReflect()
	d := mp.Descriptor()

	for i := 0; i < d.Fields().Len(); i += 1 {
		fd := d.Fields().Get(i)
		v := mp.Get(fd)

		switch {
		case string(fd.Name()) != path[0]:
			continue
		case len(path) == 1:
			return v
		case fd.Kind() != protoreflect.MessageKind:
			return protoreflect.Value{}
		}

		vm := v.Message().Interface()
		return GetValue(vm, path[1:])
	}
	return protoreflect.Value{}
}

func SetValue(m proto.Message, path []string, val protoreflect.Value) {
	if len(path) == 0 {
		return
	}

	mp := m.ProtoReflect()
	d := mp.Descriptor()
	for i := 0; i < d.Fields().Len(); i += 1 {
		fd := d.Fields().Get(i)
		v := mp.Get(fd)

		switch {
		case string(fd.Name()) != path[0]:
			continue
		case len(path) == 1:
			mp.Set(fd, val)
			return
		case fd.Kind() != protoreflect.MessageKind:
			return
		}

		vm := v.Message().Interface()
		SetValue(vm, path[1:], val)
	}
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
		if fd.Kind() == protoreflect.MessageKind && fd.Cardinality() != protoreflect.Repeated {
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
