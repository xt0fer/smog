package main

import (
	"log"
	"strconv"
)

// Universe items

var (
	MethodRegistry = initClassRegistry()
	NILClass       *VMObject
	NILObject      *VMObject
	IntegerClass   *VMObject
)

func init() {
	log.Println("init()")
	NILClass = &VMObject{Clazz: nil, Kind: "NIL"}
	NILObject = &VMObject{Clazz: NILClass, Kind: "NIL"}
	//
	IntegerClass = NewClazz("Integer")
	IntegerClass.addBinaryMethod("+", (*VMObject).doIntegerAdd)
	// IntegerClass.InstanceFields = []*VMObject{NewVMObject("Integer")}
}

// Main
func main() {

	t1 := NewVMInteger(4)
	t2 := NewVMInteger(6)

	result := t1.Send(NewMessage("+", t2))
	result.Print()
}

type Sender interface {
	ClassOf() string
	Send(*Message) *VMObject
	Print()
}

type Primitive interface {
	//	ClassOf() string
	GetValue() interface{}
}

type ClassMethodRegistry struct {
	Methods map[string]func(*VMObject, ...*VMObject) *VMObject
}
type ClassRegistry struct {
	Classes map[string]ClassMethodRegistry
}

func initClassRegistry() *ClassRegistry {
	return &ClassRegistry{Classes: make(map[string]ClassMethodRegistry)}
}
func addClassMethodRegistry(reg *ClassRegistry, cls string) {
	reg.Classes[cls] = ClassMethodRegistry{Methods: make(map[string]func(*VMObject, ...*VMObject) *VMObject)}
}
func isClassMethodRegistry(reg *ClassRegistry, cls string) bool {
	_, ok := reg.Classes[cls]
	return ok
}
func (reg *ClassRegistry) getMethod(cls string, method string) func(*VMObject, ...*VMObject) *VMObject {
	return reg.Classes[cls].Methods[method]
}

type VMObject struct {
	Clazz  *VMObject
	Kind   string
	Fields []*VMObject // local vars (any object) index of field is same as index of Class.InstanceFields
	N      int32
}

func (o *VMObject) addBinaryMethod(selector string, method func(*VMObject, ...*VMObject) *VMObject) {
	if isClassMethodRegistry(MethodRegistry, o.Kind) {
		MethodRegistry.Classes[o.Kind].Methods[selector] = method
	} else {
		addClassMethodRegistry(MethodRegistry, o.Kind)
		MethodRegistry.Classes[o.Kind].Methods[selector] = method
	}
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
	// 1. lookup method in class
	// msgsig := m.Selector
	// methodToCall := o.getMethod(msgsig)
	// // 2. call method with args
	// result := methodToCall(o, m.Args...)
	// 3. return result
	return NILObject
}

func (o *VMObject) Print() {
	if o == NILObject {
		log.Println("Nil")
		return
	}
	log.Println(o.Kind, o.Fields[0])
}
func (o *VMObject) getMethod(selector string) func(*VMObject, ...*VMObject) *VMObject {
	clazz := o.Clazz
	for clazz != nil {
		clazzName := clazz.Kind
		if method := MethodRegistry.getMethod(clazzName, selector); method != nil {
			return method
		}
		clazz = clazz.Clazz
	}
	return nil // maybe send SelectorNotFound()??
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

func (o *VMInteger) GetValue() interface{} {
	return o.Value
}

func (o *VMObject) doIntegerAdd(args ...*VMObject) *VMObject {
	// 1. get value of arg
	// 2. add to self
	// 3. return new object
	if len(args) != 1 {
		return NILObject
	}
	// switch v := (Primitive ) args[0].(type) {
	// case *VMInteger:
	// 	return NewVMInteger(o.Value + v.Value)
	// 	// case float64:
	// 	// 	return strconv.FormatFloat(float64(v), 'g', -1, 32)
	// }
	return NILObject
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
