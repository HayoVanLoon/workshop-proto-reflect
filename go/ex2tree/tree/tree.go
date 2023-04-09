package tree

import (
	"strconv"
	"strings"
)

type KeyValue struct {
	Key   string
	Value Value
}

func (kv KeyValue) String() string {
	return kv.Key + ":" + kv.Value.(value).String()
}

type Value interface {
	Type() ValueType
	Tree() Tree
	StringVal() string
	IntVal() int

	private()
}

type Tree interface {
	// Children returns a list of child nodes, ordered by insertion time.
	Children() []KeyValue

	Set(key string, value Value)

	private()
}

type message struct {
	Fields []KeyValue
}

func (*message) private() {}

func (m *message) Children() []KeyValue {
	return m.Fields
}

func (m *message) Set(key string, val Value) {
	v := val.(value)
	for i := 0; i < len(m.Fields); i += 1 {
		if m.Fields[i].Key == key {
			m.Fields[i].Value = v
			return
		}
	}
	m.Fields = append(m.Fields, KeyValue{key, v})
}

func (m *message) String() string {
	var ss []string
	for _, kv := range m.Fields {
		ss = append(ss, kv.String())
	}
	return "{" + strings.Join(ss, ",") + "}"
}

type ValueType int

const (
	ValueTypeUnknown ValueType = iota
	ValueTypeTree
	ValueTypeString
	ValueTypeInt
)

type value struct {
	Typ_       ValueType
	Message    *message
	StringVal_ string
	IntVal_    int
}

func (value) private() {}

func (v value) Type() ValueType {
	return v.Typ_
}

func (v value) Tree() Tree {
	return v.Message
}

func (v value) StringVal() string {
	return v.StringVal_
}

func (v value) IntVal() int {
	return v.IntVal_
}

func (v value) String() string {
	switch v.Typ_ {
	case ValueTypeTree:
		return v.Message.String()
	case ValueTypeInt:
		return strconv.Itoa(v.IntVal_)
	case ValueTypeString:
		return v.StringVal_
	default:
		return "<nil>"
	}
}

func ValueOfMessage(m Tree) Value {
	return value{Typ_: ValueTypeTree, Message: m.(*message)}
}

func ValueOfString(s string) Value {
	return value{Typ_: ValueTypeString, StringVal_: s}
}

func ValueOfInt(i int) Value {
	return value{Typ_: ValueTypeInt, IntVal_: i}
}

func NewTree() Tree {
	return &message{}
}
