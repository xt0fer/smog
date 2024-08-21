package vm

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
	InvokablesTable         map[*Symbol]Invokable
	SuperClassIndex         int
	NameIndex               int
	InstanceFieldsIndex     int
	InstanceInvokablesIndex int
	NumberOfClassFields     int
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
	c.InvokablesTable = make(map[*Symbol]Invokable)
	return c
}

func (c *Class) getSuperClass() interface{} {
	// Get the super class by reading the field with super class index
	return c.getField(c.SuperClassIndex)
}

func (c *Class) setSuperClass(value *Class) {
	// Set the super class by writing to the field with super class index
	c.setField(c.SuperClassIndex, value)
}

func (c *Class) hasSuperClass() bool {
	// Check whether or not this class has a super class
	return (c.getField(c.SuperClassIndex) != GetUniverse().NilObject)
}

func (c *Class) getName() interface{} {
	// Get the name of this class by reading the field with name index
	return c.getField(c.NameIndex)
}

func (c *Class) setName(value interface{}) {
	// Set the name of this class by writing to the field with name index
	c.setField(c.NameIndex, value)
}

func (c *Class) getInstanceFields() interface{} {
	// Get the instance fields by reading the field with the instance fields index
	return c.getField(c.InstanceFieldsIndex)
}

func (c *Class) setInstanceFields(value interface{}) {
	// Set the instance fields by writing to the field with the instance fields index
	c.setField(c.InstanceFieldsIndex, value)
}

func (c *Class) getInstanceInvokables() interface{} {
	// Get the instance invokables by reading the field with the instance invokables index
	return c.getField(c.InstanceInvokablesIndex)
}

func (c *Class) setInstanceInvokables(value interface{}) {
	// Set the instance invokables by writing to the field with the instance invokables index
	c.setField(c.InstanceInvokablesIndex, value)

	// Make sure this class is the holder of all invokables in the array
	for i := 0; i < c.getNumberOfInstanceInvokables(); i++ {
		c.getInstanceInvokable(i).setHolder(c)
	}
}

func (c *Class) getNumberOfInstanceInvokables() int {
	// Return the number of instance invokables in this class
	return len(c.InvokablesTable)
}

func (c *Class) getInstanceInvokable(index int) Invokable {
	// Get the instance invokable with the given index
	return c.getInstanceInvokables().getIndexableField(index)
}

func (c *Class) setInstanceInvokable(index int, value Invokable) {
	// Set this class as the holder of the given invokable
	value.SetHolder(c)

	// Set the instance method with the given index to the given value
	c.getInstanceInvokables().setIndexableField(index, value)
}

func (c *Class) getDefaultNumberOfFields() int {
	// Return the default number of fields in a class
	return c.NumberOfClassFields
}

func (c *Class) lookupInvokable(signature *Symbol) Invokable {

	// Lookup invokable and return if found
	invokable := c.InvokablesTable[signature]
	if invokable != nil {
		return invokable
	}

	// Lookup invokable with given signature in array of instance invokables
	for i := 0; i < c.getNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		invokable = c.getInstanceInvokable(i)

		// Return the invokable if the signature matches
		if invokable.GetSignature() == signature {
			c.InvokablesTable[signature] = invokable
			return invokable
		}
	}

	// Traverse the super class chain by calling lookup on the super class
	if c.hasSuperClass() {
		superc, _ := c.getSuperClass().(Class)
		invokable = superc.lookupInvokable(signature)
		if invokable != nil {
			c.InvokablesTable[signature] = invokable
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
	for i := 0; i < c.getNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		invokable := c.getInstanceInvokable(i)

		// Replace the invokable with the given one if the signature matches
		if invokable.getSignature() == value.getSignature() {
			c.setInstanceInvokable(i, value)
			return false
		}
	}

	// Append the given method to the array of instance methods
	c.setInstanceInvokables(getInstanceInvokables().copyAndExtendWith(value))
	return true
}

func (c *Class) addInstancePrimitive(value *Primitive) {
	if c.addInstanceInvokable(value) {
		fmt.println("Warning: Primitive " + value.getSignature().ToString())
		fmt.println(" is not in class definition for class " + c.getName().ToString())
	}
}

func (c *Class) getInstanceFieldName(index int) *Symbol {
	// Get the name of the instance field with the given index
	if index >= c.getNumberOfSuperInstanceFields() {
		// Adjust the index to account for fields defined in the super class
		index -= c.getNumberOfSuperInstanceFields()

		// Return the symbol representing the instance fields name
		return c.getInstanceFields().getIndexableField(index)
	} else {
		// Ask the super class to return the name of the instance field
		return c.getSuperClass().getInstanceFieldName(index)
	}
}

func (c *Class) getNumberOfInstanceFields() int {
	// Get the total number of instance fields in this class
	return c.getInstanceFields().getNumberOfIndexableFields() +
		c.getNumberOfSuperInstanceFields()
}

func (c *Class) getNumberOfSuperInstanceFields() int {
	// Get the total number of instance fields defined in super classes
	if c.hasSuperClass() {
		return c.getSuperClass().getNumberOfInstanceFields()
	} else {
		return 0
	}
}

func (c *Class) setInstanceFields(fields []string) {
	// Allocate an array of the right size
	instanceFields := GetUniverse().newArray(len(c.Fields))

	// Iterate through all the given fields
	for i := 0; i < len(c.Fields); i++ {
		// Insert the symbol corresponding to the given field string in the array
		instanceFields.setIndexableField(i, Universe.symbolFor(fields[i]))
	}

	// Set the instance fields of this class to the new array
	setInstanceFields(instanceFields)
}

func (c *Class) hasPrimitives() bool {
	// Lookup invokable with given signature in array of instance invokables
	for i := 0; i < c.getNumberOfInstanceInvokables(); i++ {
		// Get the next invokable in the instance invokable array
		if c.getInstanceInvokable(i).isPrimitive() {
			return true
		}
	}
	return false
}

func (c *Class) loadPrimitives() {
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

func (c *Class) replaceBytecodes() {
	fmt.Println("Class replaceBytecodes, off by one bug!")
	cnt := c.getNumberOfInstanceInvokables()
	for index := 0; index < cnt; index++ { // no pre-increment in Go
		inv := c.getInstanceInvokable(index)
		if !inv.isPrimitive() {
			met := inv
			met.replaceBytecodes()
		}
	}
}
