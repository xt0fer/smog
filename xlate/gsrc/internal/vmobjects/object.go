package vmobjects

// Object is base for the vmobjects
type Object interface {
	ToString() string
	LookupField(name string) Object
	SetField(name string, obj Object)
	RemoveField(name string)
	NumberOfFields() int
	Fields() map[string]Object
	SetFields(fields map[string]Object)

}

type ObjectClass struct {
	Fields map[string]Object
}

func NewObjectClass() *ObjectClass {
	return &ObjectClass{make(map[string]Object)}
}

func (oc *ObjectClass) LookupField(name string) Object {
	return oc.Fields[name]
}

func (oc *ObjectClass) SetField(name string, obj Object) {
	oc.Fields[name] = obj
}

func (oc *ObjectClass) RemoveField(name string) {
	delete(oc.Fields, name)
}

func (oc *ObjectClass) NumberOfFields() int {
	return len(oc.Fields)
}

func (oc *ObjectClass) Fields() map[string]Object {
	return oc.Fields
}

func (oc *ObjectClass) SetFields(fields map[string]Object) {
	oc.Fields = fields
}

func (oc *ObjectClass) ToString() string {
	return "ObjectClass"
}


public Object()
{
  // Set the number of fields to the default value
  setNumberOfFields(getDefaultNumberOfFields());
}

public Object(int numberOfFields)
{
  // Set the number of fields to the given value
  setNumberOfFields(numberOfFields);
}

public Class getSOMClass()
{
  // Get the class of this object by reading the field with class index
  return (Class) getField(classIndex);
}

public void setClass(Class value)
{
  // Set the class of this object by writing to the field with class index
  setField(classIndex, value);
}

public Symbol getFieldName(int index)
{
  // Get the name of the field with the given index
  return getSOMClass().getInstanceFieldName(index);
}

public int getFieldIndex(Symbol name)
{
  // Get the index for the field with the given name
  return getSOMClass().lookupFieldIndex(name);
}

public int getNumberOfFields()
{
  // Get the number of fields in this object
  return fields.length;
}

public void setNumberOfFields(int value)
{
  // Allocate a new array of fields
  fields = new Object[value];
  
  // Clear each and every field by putting nil into them
  for (int i = 0; i < getNumberOfFields(); i++) {
	setField(i, Universe.nilObject);
  }
}

public int getDefaultNumberOfFields()
{
  // Return the default number of fields in an object
  return numberOfObjectFields;
}

public void send(java.lang.String selectorString, Object[] arguments) 
{
  // Turn the selector string into a selector
  Symbol selector = Universe.symbolFor(selectorString);

  // Push the receiver onto the stack
  Interpreter.getFrame().push(this);

  // Push the arguments onto the stack
  for(Object arg : arguments)
	  Interpreter.getFrame().push(arg);
  
  // Lookup the invokable 
  Invokable invokable = getSOMClass().lookupInvokable(selector);

  // Invoke the invokable
  invokable.invoke(Interpreter.getFrame());
}

public Object getField(int index)
{
  // Get the field with the given index
  return fields[index];
}

public void setField(int index, Object value)
{
  // Set the field with the given index to the given value
  fields[index] = value;
}

public static void _assert(boolean value)
{
  // Delegate to universal assertion routine
  Universe._assert(value);
}

