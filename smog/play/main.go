package main

func main() {

	t1 := NewVMInteger(4)
	t2 := NewVMInteger(6)
	op := NewMessage("+", t2)

	result := t1.Send(op)
	println(result.Value())
}

type VMObject struct {
	Clazz  *VMObject
	Kind   string
	Fields []*VMObject // local vars (any object) index of field is same as index of Class.InstanceFields
	N      int32
}

func NewVMObject(cls string) *VMObject {
	no := &VMObject{Fields: make([]*VMObject, 0)}
	no.Kind = cls
	return no
}

func (o *VMObject) ClassOf() string {
	// TODO
	return o.Kind
}

func (o *VMObject) Send(m *Message) *VMObject {
	// TODO
	return o
}

func (o *VMObject) Value() string {
	return o.N
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
	m.Args = args
	return m
}
