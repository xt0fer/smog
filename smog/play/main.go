package main

func main() {

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

type ClassMap struct {
	Classes map[string]*VMObject
}

type VMInteger struct {
	VMObject
	Value int32
}

func NewVMInteger(value int32) *VMInteger {
	no := &VMInteger{Value: value}
	no.Kind = "Integer"
	return no
}
