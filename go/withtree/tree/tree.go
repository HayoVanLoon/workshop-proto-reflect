package tree

import (
	"strconv"
	"strings"
)

type KeyValue struct {
	key   string
	value value
}

func (kv KeyValue) Key() string {
	return kv.key
}

func (kv KeyValue) Value() Value {
	return kv.value
}

func (kv KeyValue) String() string {
	return kv.key + ":" + kv.value.String()
}

type Value interface {
	Tree() Tree
	StringVal() string
	IntVal() int

	private()
}

type Tree interface {
	Children() []KeyValue

	Set(key string, value Value)

	private()
}

type message struct {
	fields []KeyValue
}

func (*message) private() {}

func (m *message) Children() []KeyValue {
	return m.fields
}

func (m *message) Set(key string, val Value) {
	v := val.(value)
	for i := 0; i < len(m.fields); i += 1 {
		if m.fields[i].key == key {
			m.fields[i].value = v
			return
		}
	}
	m.fields = append(m.fields, KeyValue{key, v})
}

func (m *message) String() string {
	var ss []string
	for _, kv := range m.fields {
		ss = append(ss, kv.String())
	}
	return "{" + strings.Join(ss, ",") + "}"
}

type value struct {
	message    *message
	string_val string
	int_val    int
}

func (value) private() {}

func (v value) Tree() Tree {
	return v.message
}

func (v value) StringVal() string {
	return v.string_val
}

func (v value) IntVal() int {
	return v.int_val
}

func (v value) String() string {
	switch {
	case v.message != nil:
		return v.message.String()
	case v.int_val != 0:
		return strconv.Itoa(v.int_val)
	case v.string_val != "":
		return v.string_val
	default:
		return "<zero>"
	}
}

func ValueOfMessage(m Tree) Value {
	return value{message: m.(*message)}
}

func ValueOfString(s string) Value {
	return value{string_val: s}
}

func ValueOfInt(i int) Value {
	return value{int_val: i}
}

func NewTree() Tree {
	return &message{}
}
