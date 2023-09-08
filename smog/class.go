package smog

import "fmt"

type IsClass interface {
	getSuperClass() IsClass
	setSuperClass(value IsClass)
	hasSuperClass() bool
	getName() *Symbol
	setName(value *Symbol)
	getInstanceFields() *Array
	setInstanceFields(value *Array)
	getInstanceInvokables() *Array
	setInstanceInvokables(value *Array)
	getNumberOfInstanceInvokables() int
	getInstanceInvokable(index int) Invokable
	setInstanceInvokable(index int, value Invokable)
	getDefaultNumberOfFields() int
	lookupInvokable(signature *Symbol) Invokable
	lookupFieldIndex(fieldName *Symbol) int
	addInstanceInvokable(value Invokable) bool
	addInstancePrimitive(value *Primitive)
	getInstanceFieldName(index int) *Symbol
	getNumberOfInstanceFields() int
	getNumberOfSuperInstanceFields() int
	setInstanceFieldsFromStrings(fields []string)
	hasPrimitives() bool
	loadPrimitives()
	replaceBytecodes()
}
type Class struct {
	Object
	// Map of symbols to invokables
	SuperClass     *Class
	Name           *Symbol
	InstanceFields *Array
	Invokables     *Array
	ClassFields    *Array
}

//  field indices and number of class fields
// superClassIndex         = 1 + classIndex;
// nameIndex                = 1 + superClassIndex;
// instanceFieldsIndex     = 1 + nameIndex;
// instanceInvokablesIndex = 1 + instanceFieldsIndex;
// numberOfClassFields    = 1 + instanceInvokablesIndex;

func NewClass(numberOfFields int) *Class {
	// Initialize this class by calling the super constructor with the given value
	//super(numberOfFields);
	c := &Class{}
	c.Invokables = GetUniverse().NewArray(numberOfFields) //make(map[*Symbol]Invokable)
	return c
}

func (c *Class) GetSuperClass() *Class {
	// Get the super class by reading the field with super class index
	return c.SuperClass
}

func (c *Class) SetSuperClass(value *Class) {
	// Set the super class by writing to the field with super class index
	c.SuperClass = value
}

func (c *Class) HasSuperClass() bool {
	// Check whether or not this class has a super class
	return c != nil && c.GetSuperClass() != nil
}

func (c *Class) GetName() *Symbol {
	// Get the name of this class by reading the field with name index
	return c.Name
}

func (c *Class) SetName(value *Symbol) {
	// Set the name of this class by writing to the field with name index
	c.Name = value
}

func (c *Class) GetInstanceFields() *Array {
	// Get the instance fields by reading the field with the instance fields index
	return c.InstanceFields
}

func (c *Class) SetInstanceFields(value *Array) {
	// Set the instance fields by writing to the field with the instance fields index
	c.InstanceFields = value
}

func (c *Class) GetInstanceInvokables() *Array {
	// Get the instance invokables by reading the field with the instance invokables index
	return c.Invokables
}

func (c *Class) SetInstanceInvokables(value *Array) {
	// Set the instance invokables by writing to the field with the instance invokables index
	c.Invokables = value

	// Make sure this class is the holder of all invokables in the array
	for i := 0; i < c.GetNumberOfInstanceInvokables(); i++ {
		c.GetInstanceInvokable(i).SetHolder(c)
	}
}

func (c *Class) GetNumberOfInstanceInvokables() int {
	// Return the number of instance invokables in this class
	return c.Invokables.GetNumberOfIndexableFields()
}

func (c *Class) GetInstanceInvokable(index int) Invokable {
	// Get the instance invokable with the given index
	pi := c.Invokables.GetIndexableField(index)
	invk, ok := pi.(Invokable)
	if ok {
		return invk
	} else {
		panic("not an invokable")
	}
}

func (c *Class) setInstanceInvokable(index int, value Invokable) {
	// Set this class as the holder of the given invokable
	value.SetHolder(c)

	// Set the instance method with the given index to the given value
	c.GetInstanceInvokables().SetIndexableField(index, value)
}

func (c *Class) getDefaultNumberOfFields() int {
	// Return the default number of fields in a class
	return c.ClassFields.GetNumberOfIndexableFields()
}

func (c *Class) lookupInvokable(signature *Symbol) Invokable {
	// Lookup invokable with given signature in array of instance invokables
	for i := 0; i < c.GetNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		invokable := c.GetInstanceInvokable(i)

		// Return the invokable if the signature matches
		if invokable.GetSignature() == signature {
			return invokable
		}
	}

	// Traverse the super class chain by calling lookup on the super class
	if c.HasSuperClass() {
		superc := c.GetSuperClass()
		invokable := superc.lookupInvokable(signature)
		if invokable != nil {
			return invokable
		}
	}

	// Invokable not found
	return nil
}

func (c *Class) lookupFieldIndex(fieldName *Symbol) int {
	// Lookup field with given name in array of instance fields
	for i := c.getNumberOfInstanceFields() - 1; i >= 0; i-- {
		// Return the current index if the name matches
		if fieldName == c.getInstanceFieldName(i) {
			return i
		}
	}

	// Field not found
	return -1
}

func (c *Class) addInstanceInvokable(value Invokable) bool {
	// Add the given invokable to the array of instance invokables
	for i := 0; i < c.GetNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		invokable := c.GetInstanceInvokable(i)

		// Replace the invokable with the given one if the signature matches
		if invokable.GetSignature() == value.GetSignature() {
			c.setInstanceInvokable(i, value)
			return false
		}
	}

	// Append the given method to the array of instance methods
	c.SetInstanceInvokables(c.GetInstanceInvokables().copyAndExtendWith(value))
	return true
}

func (c *Class) addInstancePrimitive(value *Primitive) {
	if c.addInstanceInvokable(value) {
		fmt.Println("Warning: Primitive " + value.GetSignature().String())
		fmt.Println(" is not in class definition for class " + c.GetName().String())
	}
}

func (c *Class) getInstanceFieldName(index int) *Symbol {
	// Get the name of the instance field with the given index
	if index >= c.GetNumberOfSuperInstanceFields() {
		// Adjust the index to account for fields defined in the super class
		index -= c.getNumberOfSuperInstanceFields()

		// Return the symbol representing the instance fields name
		return c.GetInstanceFields().GetIndexableField(index)
	} else {
		// Ask the super class to return the name of the instance field
		return c.GetSuperClass().getInstanceFieldName(index)
	}
}

func (c *Class) GetNumberOfInstanceFields() int {
	// Get the total number of instance fields in this class
	return c.GetInstanceFields().GetNumberOfIndexableFields() +
		c.GetNumberOfSuperInstanceFields()
}

func (c *Class) GetNumberOfSuperInstanceFields() int {
	// Get the total number of instance fields defined in super classes
	if c.HasSuperClass() {
		return c.GetSuperClass().GetNumberOfInstanceFields()
	} else {
		return 0
	}
}

// Set from an array of strings
// func (c *Class) SetInstanceFields(fields []string) {
// 	// Allocate an array of the right size
// 	instanceFields := GetUniverse().NewArray(len(c.Fields))

// 	// Iterate through all the given fields
// 	for i := 0; i < len(c.Fields); i++ {
// 		// Insert the symbol corresponding to the given field string in the array
// 		instanceFields.SetIndexableField(i, GetUniverse().SymbolFor(fields[i]))
// 	}

// 	// Set the instance fields of this class to the new array
// 	c.SetInstanceFields(instanceFields)
// }

func (c *Class) HasPrimitives() bool {
	// Lookup invokable with given signature in array of instance invokables
	for i := 0; i < c.GetNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		if c.GetInstanceInvokable(i).IsPrimitive() {
			return true
		}
	}
	return false
}

func (c *Class) LoadPrimitives() {
	panic("load Primitives?")
	// Compute the class name of the Java(TM) class containing the primitives
	//  java.lang.String className = "som.primitives." + getName().getString() + "Primitives";

	//  // Try loading the primitives
	//  try {
	//    java.lang.Class<?> primitivesClass = java.lang.Class.forName(className);
	//    try {
	// 	 ((Primitives) primitivesClass.newInstance()).installPrimitivesIn(this);
	//    } catch (Exception e) {
	// 	 System.out.println("Primitives class " + className + " cannot be instantiated");
	//    }
	//  } catch (ClassNotFoundException e) {
	//    System.out.println("Primitives class " + className + " not found");
	//  }
}

func (c *Class) ReplaceBytecodes() {
	fmt.Println("Class ReplaceBytecodes, off by one bug!??")
	cnt := c.GetNumberOfInstanceInvokables()
	for index := 0; index < cnt; index++ { // no pre-increment in Go
		inv := c.GetInstanceInvokable(index)
		if !inv.IsPrimitive() {
			met := inv
			met.ReplaceBytecodes()
		}
	}
}
