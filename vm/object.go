package vm

import "github.com/xt0fer/smog/interpreter"

type Object struct {
	ClassIndex           int
	SOMClass             *Class
	NumberOfObjectFields int
	Fields               []*Object
}

func NewObject(name string, args int) *Object {
	ns := &Object{}
	if args == -1 {
		ns.NumberOfObjectFields = ns.getDefaultNumberOfFields()
	} else {
		ns.NumberOfObjectFields = args
	}
	ns.Fields = make([]*Object, ns.NumberOfObjectFields)
	return ns
}

func (o *Object) getSOMClass() *Class {
	return o.SOMClass
}

func (o *Object) setClass(c *Class) {
	o.SOMClass = c
}

func (o *Object) getFieldName(index int) *Symbol {
	// Get the name of the field with the given index
	return o.SOMClass.getInstanceFieldName(index)
}

func (o *Object) getFieldIndex(name Symbol) {
	// Get the index for the field with the given name
	return o.getSOMClass().lookupFieldIndex(name)
}

func (o *Object) getNumberOfFields() int {
	// Get the number of fields in this object
	return len(o.Fields)
}

func (o *Object) setNumberOfFields(value int) {
	// Allocate a new array of fields
	o.Fields = make([]*Object, value)

	// Clear each and every field by putting nil into them
	for i := 0; i < o.getNumberOfFields(); i++ {
		o.setField(i, GetUniverse().NilObject)
	}
}

func (o *Object) getDefaultNumberOfFields() int {
	// Return the default number of fields in an object
	return o.NumberOfObjectFields
}

func (o *Object) send(interpreter interpreter.Interpreter, selectorStr string, arguments []*Object) {
	// Turn the selector string into a selector
	selector := GetUniverse().symbolFor(selectorStr)

	// Push the receiver onto the stack
	interpreter.getFrame().push(o)

	// Push the arguments onto the stack
	for arg := range arguments {
		interpreter.getFrame().push(arg)
	}
	// Lookup the invokable
	invokable := o.getSOMClass().lookupInvokable(selector)

	// Invoke the invokable
	invokable.invoke(interpreter.getFrame())
}

func (o *Object) getField(index int) *Object {
	// Get the field with the given index
	return o.Fields[index]
}

func (o *Object) setField(index int, value *Object) {
	// Set the field with the given index to the given value
	o.Fields[index] = value
}

// func (o *Object) _assert(boolean value) {
// 	GetUniverse()._assert(value)
// }
