package vmobj

// Object is base for the vmobjects
type VMObjectI interface {
	ToString() string
	GetField(i int) VMObjectI
	SetField(i int, obj VMObjectI)
	NumberOfFields() int
	GetClass() *VMObject
}

type VMObject struct {
	Clazz   *VMClass
	Fields  []VMObjectI
	Nfields int
}

func NewObject(nfields int, clazz *VMClass) *VMObject {
	return &VMObject{Clazz: clazz,
		Fields:  make([]VMObjectI, nfields),
		Nfields: nfields}
}

func NewObjInternal(nfields int, clazz *VMClass) VMObject {
	return VMObject{Clazz: clazz,
		Fields:  make([]VMObjectI, nfields),
		Nfields: nfields}
}

func (oc *VMObject) SetClazz(c *VMClass) {
	oc.Clazz = c
}

func (oc *VMObject) GetClazz() *VMClass {
	return oc.Clazz
}
func (oc *VMObject) GetClazzName() string {
	return oc.Clazz.ToString()
}
func (oc *VMObject) NumberOfFields() int {
	return len(oc.Fields)
}

func (oc *VMObject) ToString() string {
	return "VMObject"
}

func (oc *VMObject) GetField(i int) VMObjectI {
	return oc.Fields[i]
}

func (oc *VMObject) SetField(i int, obj VMObjectI) {
	oc.Fields[i] = obj
}

func (oc *VMObject) GetClass() *VMClass {
	return oc.Clazz
}

// public Object()
// {
//   // Set the number of fields to the default value
//   setNumberOfFields(getDefaultNumberOfFields());
// }

// public Object(int numberOfFields)
// {
//   // Set the number of fields to the given value
//   setNumberOfFields(numberOfFields);
// }

// public Class getSOMClass()
// {
//   // Get the class of this object by reading the field with class index
//   return (Class) getField(classIndex);
// }

// public void setClass(Class value)
// {
//   // Set the class of this object by writing to the field with class index
//   setField(classIndex, value);
// }

// public Symbol getFieldName(int index)
// {
//   // Get the name of the field with the given index
//   return getSOMClass().getInstanceFieldName(index);
// }

// public int getFieldIndex(Symbol name)
// {
//   // Get the index for the field with the given name
//   return getSOMClass().lookupFieldIndex(name);
// }

// public int getNumberOfFields()
// {
//   // Get the number of fields in this object
//   return fields.length;
// }

// public void setNumberOfFields(int value)
// {
//   // Allocate a new array of fields
//   fields = new Object[value];

//   // Clear each and every field by putting nil into them
//   for (int i = 0; i < getNumberOfFields(); i++) {
// 	setField(i, Universe.nilObject);
//   }
// }

// public int getDefaultNumberOfFields()
// {
//   // Return the default number of fields in an object
//   return numberOfObjectFields;
// }

// public void send(java.lang.String selectorString, Object[] arguments)
// {
//   // Turn the selector string into a selector
//   Symbol selector = Universe.symbolFor(selectorString);

//   // Push the receiver onto the stack
//   Interpreter.getFrame().push(this);

//   // Push the arguments onto the stack
//   for(Object arg : arguments)
// 	  Interpreter.getFrame().push(arg);

//   // Lookup the invokable
//   Invokable invokable = getSOMClass().lookupInvokable(selector);

//   // Invoke the invokable
//   invokable.invoke(Interpreter.getFrame());
// }

// public Object getField(int index)
// {
//   // Get the field with the given index
//   return fields[index];
// }

// public void setField(int index, Object value)
// {
//   // Set the field with the given index to the given value
//   fields[index] = value;
// }

// public static void _assert(boolean value)
// {
//   // Delegate to universal assertion routine
//   Universe._assert(value);
// }
