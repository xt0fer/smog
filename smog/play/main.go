package main

import (
	"fmt"
	"strconv"
)

// Universe items

var NILClass *VMObject
var NILObject *VMObject

func init() {
	NILClass = &VMObject{Clazz: nil, Kind: "NIL"}
	NILObject = &VMObject{Clazz: NILClass, Kind: "NIL"}
	//
	IntegerClass := NewClazz("Integer")
	// IntegerClass.InstanceFields = []*VMObject{NewVMObject("Integer")}
}

// Main
func main() {

	t1 := NewVMInteger(4)
	t2 := NewVMInteger(6)
	op := NewMessage("+", t2)

	result := t1.Send(op)
	result.Print()
}

type Sender interface {
	ClassOf() string
	Send(*Message) *VMObject
	Print()
}

type Primitive interface {
	ClassOf() string
	Value() interface{}
}

type VMObject struct {
	Clazz  *VMObject
	Kind   string
	Fields []*VMObject // local vars (any object) index of field is same as index of Class.InstanceFields
	N      int32
}
// remember Stringer below, when pondering the dynamic type lookup and method dispatch

func NewVMObject(cls string) *VMObject {
	no := &VMObject{Fields: make([]*VMObject, 0)}
	no.Kind = cls
	return no
}

func NewClazz(cls string) *VMObject {
	no := &VMObject{Fields: make([]*VMObject, 0)}
	no.Clazz = NILObject
	no.Kind = cls
	return no
}



func (o *VMObject) ClassOf() string {
	return o.Clazz.Kind
}

func (o *VMObject) Send(m *Message) *VMObject {

	// TODO
	return o
}

func (o *VMObject) Print() {
	fmt.Println(o.Kind, o.Fields[0])
}

type ClassMap struct {
	Classes map[string]*VMObject
}

type VMInteger struct {
	VMObject
	Value int32
}

func NewVMInteger(value int32) *VMObject {
	no := &VMObject{Fields: make([]*VMObject, 0)}
	no.Fields = append(no.Fields, &VMObject{Kind: "Integer", N: value})
	no.Kind = "Integer"
	return no
}

type Message struct {
	Selector string
	Args     []*VMObject
}

func NewMessage(s string, args ...*VMObject) (m *Message) {
	m = &Message{
		Selector: s,
	}
	m.Args = append(m.Args, args...)
	return m
}

/*
* And example of using a dynamic type lookup
 */
type Stringer interface {
	String() string
}

func AsString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(float64(v), 'g', -1, 32)
	}
	return "???"
}
