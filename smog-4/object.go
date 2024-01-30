package smog

import "log"

type IsObject interface {
	GetSOMClass() *Class
	SetClass(c *Class)
	GetFieldName(index int) *Symbol
	GetField(index int) *Object
	SetField(index int, value *Object)
	GetFieldIndex(name *Symbol)
	GetNumberOfFields() int
	SetNumberOfFields(value int)
	GetDefaultNumberOfFields() int
	Send(selectorStr string, arguments []*Object)
}

type Object struct {
	SOMClass             *Class
	NumberOfObjectFields int
	Fields               []interface{}
}

func NewObject(fields int) *Object {
	ns := &Object{}
	ns.ObjectInit(fields)
	return ns
}

func (ns *Object) ObjectInit(fields int) {
	if fields == -1 {
		ns.NumberOfObjectFields = ns.GetDefaultNumberOfFields()
	} else {
		ns.NumberOfObjectFields = fields
	}
	ns.Fields = make([]interface{}, ns.NumberOfObjectFields)
}

func (o *Object) GetSOMClass() *Class {
	return o.SOMClass
}

func (o *Object) SetClass(c *Class) {
	o.SOMClass = c
}

func (o *Object) GetFieldName(index int) *Symbol {
	// Get the name of the field with the given index
	return o.GetSOMClass().GetInstanceFieldName(index)
}

func (o *Object) GetFieldIndex(name *Symbol) int {
	// Get the index for the field with the given name
	return o.GetSOMClass().LookupFieldIndex(name)
}

func (o *Object) GetNumberOfFields() int {
	// Get the number of fields in this object
	return len(o.Fields)
}

func (o *Object) SetNumberOfFields(value int) {
	// Allocate a new array of fields
	o.Fields = make([]interface{}, value)

	// Clear each and every field by putting nil into them
	for i := 0; i < o.GetNumberOfFields(); i++ {
		o.SetField(i, GetUniverse().NilObject)
	}
}

func (o *Object) GetDefaultNumberOfFields() int {
	// Return the default number of fields in an object
	return o.NumberOfObjectFields
}

func (o *Object) Send(selectorStr string, arguments []*Object) {
	// Turn the selector string into a selector
	selector := GetUniverse().SymbolFor(selectorStr)

	if arguments == nil {
		log.Println("in Object.Send(), argument list is NIL\nProbably a bug in the interpreter.")
	}
	// Push the receiver onto the stack
	interpreter := GetInterpreter()
	interpreter.GetFrame().Push(o)

	// Push the arguments onto the stack
	for arg := range arguments {
		interpreter.GetFrame().Push(arg)
	}
	// Lookup the invokable
	invokable := o.GetSOMClass().LookupInvokable(selector)

	// Invoke the invokable
	invokable.Invoke(interpreter.GetFrame())
}

func (o *Object) GetField(index int) interface{} {
	// Get the field with the given index
	return o.Fields[index]
}

func (o *Object) SetField(index int, value interface{}) {
	// Set the field with the given index to the given value
	o.Fields[index] = value
}

// func (o *Object) _assert(value bool) {
// 	GetUniverse()._assert(value)
// }
