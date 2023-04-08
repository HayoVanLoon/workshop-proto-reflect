package ex1map_test

import (
	"testing"
	"workshop/ex1map"

	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	type args struct {
		m    map[string]any
		path []string
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"happy 1st level",
			args{ex1map.Create(), []string{"brand"}},
			"granny-smith",
		},
		{
			"happy 2nd level",
			args{ex1map.Create(), []string{"skin", "blemishes"}},
			3,
		},
		{
			"not found",
			args{ex1map.Create(), []string{"skin", "punctures"}},
			nil,
		},
		{
			"empty",
			args{ex1map.Create(), nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex1map.GetValue(tt.args.m, tt.args.path)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestSetValue(t *testing.T) {
	type args struct {
		m     map[string]any
		path  []string
		value any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			"happy 1st level",
			args{ex1map.Create(), []string{"brand"}, "elstar"},
			func() map[string]any {
				a := ex1map.Create()
				a["brand"] = "elstar"
				return a
			}(),
		},
		{
			"happy 2nd level",
			args{ex1map.Create(), []string{"skin", "blemishes"}, 4},
			func() map[string]any {
				a := ex1map.Create()
				a["skin"].(map[string]any)["blemishes"] = 4
				return a
			}(),
		},
		{
			"add",
			args{ex1map.Create(), []string{"skin", "punctures"}, 5},
			func() map[string]any {
				a := ex1map.Create()
				a["skin"].(map[string]any)["punctures"] = 5
				return a
			}(),
		},
		{
			"empty",
			args{ex1map.Create(), nil, 1},
			ex1map.Create(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.args.m
			ex1map.SetValue(actual, tt.args.path, tt.args.value)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestDepthFirstTraversal(t *testing.T) {
	type args struct {
		m map[string]any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			"empty",
			args{nil},
			[]any(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex1map.Traverse(tt.args.m)
			require.Equal(t, tt.want, actual)
		})
	}
}
