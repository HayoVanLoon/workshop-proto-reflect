package ex4proto_test

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"testing"
	"workshop/ex4proto"
	pb "workshop/lib/workshop/v1"

	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	type args struct {
		m    *pb.Apple
		path []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"happy 1st level",
			args{ex4proto.Create(), []string{"Brand"}},
			"granny-smith",
		},
		{
			"happy 2nd level",
			args{ex4proto.Create(), []string{"Skin", "Blemishes"}},
			int32(3),
		},
		{
			"not found",
			args{ex4proto.Create(), []string{"Skin", "Punctures"}},
			nil,
		},
		{
			"empty",
			args{ex4proto.Create(), nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex4proto.GetValue(tt.args.m, tt.args.path)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestSetValue(t *testing.T) {
	type args struct {
		m     *pb.Apple
		path  []string
		value protoreflect.Value
	}
	tests := []struct {
		name string
		args args
		want *pb.Apple
	}{
		{
			"happy 1st level",
			args{ex4proto.Create(), []string{"Brand"}, protoreflect.ValueOf("elstar")},
			&pb.Apple{
				Brand: "elstar",
				Age:   42,
				Skin:  &pb.Apple_Skin{Colour: "green", Blemishes: 3},
			},
		},
		{
			"happy 2nd level",
			args{ex4proto.Create(), []string{"Skin", "Blemishes"}, protoreflect.ValueOf(int32(4))},
			&pb.Apple{
				Brand: "granny-smith",
				Age:   42,
				Skin:  &pb.Apple_Skin{Colour: "green", Blemishes: int32(4)},
			},
		},
		{
			"! add",
			args{ex4proto.Create(), []string{"Skin", "Punctures"}, protoreflect.ValueOf(5)},
			&pb.Apple{
				Brand: "granny-smith",
				Age:   42,
				Skin:  &pb.Apple_Skin{Colour: "green", Blemishes: 3},
			},
		},
		{
			"empty",
			args{&pb.Apple{}, []string{"Skin", "Colour"}, protoreflect.ValueOf("green")},
			&pb.Apple{Skin: &pb.Apple_Skin{Colour: "green"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.args.m
			ex4proto.SetValue(actual, tt.args.path, tt.args.value)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestDepthFirstTraversal(t *testing.T) {
	type args struct {
		m *pb.Apple
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			"happy",
			args{ex4proto.Create()},
			[]any{"granny-smith", int32(42), "green", int32(3)},
		},
		{
			"empty",
			args{&pb.Apple{}},
			[]any{"", int32(0), "", int32(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex4proto.Traverse(tt.args.m)
			require.Equal(t, tt.want, actual)
		})
	}
}
