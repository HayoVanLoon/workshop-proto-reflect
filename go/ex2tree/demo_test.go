package ex2tree_test

import (
	"testing"
	"workshop/ex2tree"
	"workshop/ex2tree/tree"

	"github.com/stretchr/testify/require"
)

func TestGetValue(t *testing.T) {
	type args struct {
		m    tree.Tree
		path []string
	}
	tests := []struct {
		name string
		args args
		want tree.Value
	}{
		{
			"happy 1st level",
			args{ex2tree.Create(), []string{"brand"}},
			tree.ValueOfString("granny-smith"),
		},
		{
			"happy 2nd level",
			args{ex2tree.Create(), []string{"skin", "blemishes"}},
			tree.ValueOfInt(3),
		},
		{
			"not found",
			args{ex2tree.Create(), []string{"skin", "punctures"}},
			nil,
		},
		{
			"empty",
			args{ex2tree.Create(), nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex2tree.GetValue(tt.args.m, tt.args.path)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestSetValue(t *testing.T) {
	type args struct {
		m     tree.Tree
		path  []string
		value tree.Value
	}
	tests := []struct {
		name string
		args args
		want tree.Tree
	}{
		{
			"happy 1st level",
			args{ex2tree.Create(), []string{"brand"}, tree.ValueOfString("elstar")},
			func() tree.Tree {
				apple := tree.NewTree()
				apple.Set("brand", tree.ValueOfString("elstar"))
				apple.Set("age", tree.ValueOfInt(42))
				skin := tree.NewTree()
				skin.Set("colour", tree.ValueOfString("green"))
				skin.Set("blemishes", tree.ValueOfInt(3))
				apple.Set("skin", tree.ValueOfMessage(skin))
				return apple
			}(),
		},
		{
			"happy 2nd level",
			args{ex2tree.Create(), []string{"skin", "blemishes"}, tree.ValueOfInt(4)},
			func() tree.Tree {
				apple := tree.NewTree()
				apple.Set("brand", tree.ValueOfString("granny-smith"))
				apple.Set("age", tree.ValueOfInt(42))
				skin := tree.NewTree()
				skin.Set("colour", tree.ValueOfString("green"))
				skin.Set("blemishes", tree.ValueOfInt(4))
				apple.Set("skin", tree.ValueOfMessage(skin))
				return apple
			}(),
		},
		{
			"add",
			args{ex2tree.Create(), []string{"skin", "punctures"}, tree.ValueOfInt(5)},
			func() tree.Tree {
				apple := tree.NewTree()
				apple.Set("brand", tree.ValueOfString("granny-smith"))
				apple.Set("age", tree.ValueOfInt(42))
				skin := tree.NewTree()
				skin.Set("colour", tree.ValueOfString("green"))
				skin.Set("blemishes", tree.ValueOfInt(3))
				skin.Set("punctures", tree.ValueOfInt(5))
				apple.Set("skin", tree.ValueOfMessage(skin))
				return apple
			}(),
		},
		{
			"create subs",
			args{ex2tree.Create(), []string{"flavour", "sweetness"}, tree.ValueOfInt(5)},
			func() tree.Tree {
				apple := tree.NewTree()
				apple.Set("brand", tree.ValueOfString("granny-smith"))
				apple.Set("age", tree.ValueOfInt(42))
				skin := tree.NewTree()
				skin.Set("colour", tree.ValueOfString("green"))
				skin.Set("blemishes", tree.ValueOfInt(3))
				apple.Set("skin", tree.ValueOfMessage(skin))
				flavour := tree.NewTree()
				flavour.Set("sweetness", tree.ValueOfInt(5))
				apple.Set("flavour", tree.ValueOfMessage(flavour))
				return apple
			}(),
		},
		{
			"! wrong type",
			args{ex2tree.Create(), []string{"brand"}, tree.ValueOfInt(5)},
			ex2tree.Create(),
		},
		{
			"empty",
			args{ex2tree.Create(), nil, tree.ValueOfInt(1)},
			ex2tree.Create(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.args.m
			ex2tree.SetValue(actual, tt.args.path, tt.args.value)
			require.Equal(t, tt.want, actual)
		})
	}
}

func TestTraverse(t *testing.T) {
	type args struct {
		m tree.Tree
	}
	tests := []struct {
		name string
		args args
		want []tree.Value
	}{
		{
			"happy",
			args{ex2tree.Create()},
			[]tree.Value{
				tree.ValueOfString("granny-smith"),
				tree.ValueOfInt(42),
				tree.ValueOfString("green"),
				tree.ValueOfInt(3),
			},
		},
		{
			"empty",
			args{nil},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ex2tree.Traverse(tt.args.m)
			require.Equal(t, tt.want, actual)
		})
	}
}
