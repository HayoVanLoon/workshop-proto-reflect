package ex3reflect_test

import (
	"reflect"
	"testing"
	"workshop/ex3reflect"

	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	type args struct {
		m    ex3reflect.Apple
		path []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"happy 1st level",
			args{ex3reflect.Create(), []string{"Brand"}},
			"granny-smith",
		},
		{
			"happy 2nd level",
			args{ex3reflect.Create(), []string{"Skin", "Blemishes"}},
			int32(3),
		},
		{
			"not found",
			args{ex3reflect.Create(), []string{"Skin", "Punctures"}},
			nil,
		},
		{
			"empty",
			args{ex3reflect.Create(), nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex3reflect.GetValue(tt.args.m, tt.args.path)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestSetValue(t *testing.T) {
	type args struct {
		m     ex3reflect.Apple
		path  []string
		value reflect.Value
	}
	tests := []struct {
		name string
		args args
		want ex3reflect.Apple
	}{
		{
			"happy 1st level",
			args{ex3reflect.Create(), []string{"Brand"}, reflect.ValueOf("elstar")},
			ex3reflect.Apple{
				Brand: "elstar",
				Age:   42,
				Skin:  ex3reflect.AppleSkin{Colour: "green", Blemishes: 3},
			},
		},
		{
			"happy 2nd level",
			args{ex3reflect.Create(), []string{"Skin", "Blemishes"}, reflect.ValueOf(int32(4))},
			ex3reflect.Apple{
				Brand: "granny-smith",
				Age:   42,
				Skin:  ex3reflect.AppleSkin{Colour: "green", Blemishes: int32(4)},
			},
		},
		{
			"! add",
			args{ex3reflect.Create(), []string{"Skin", "Punctures"}, reflect.ValueOf(5)},
			ex3reflect.Apple{
				Brand: "granny-smith",
				Age:   42,
				Skin:  ex3reflect.AppleSkin{Colour: "green", Blemishes: 3},
			},
		},
		{
			"empty",
			args{ex3reflect.Apple{}, []string{"Skin", "Colour"}, reflect.ValueOf("green")},
			ex3reflect.Apple{Skin: ex3reflect.AppleSkin{Colour: "green"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.args.m
			ex3reflect.SetValue(&actual, tt.args.path, tt.args.value)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestDepthFirstTraversal(t *testing.T) {
	type args struct {
		m ex3reflect.Apple
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			"happy",
			args{ex3reflect.Create()},
			[]any{"granny-smith", int32(42), "green", int32(3)},
		},
		{
			"empty",
			args{ex3reflect.Apple{}},
			[]any{"", int32(0), "", int32(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex3reflect.Traverse(tt.args.m)
			require.Equal(t, tt.want, actual)
		})
	}
}
