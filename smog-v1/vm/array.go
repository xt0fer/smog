package vm

type IsArray interface {
getIndexableField(index int) interface{};
setIndexableField(index int, value interface{});
getNumberOfIndexableFields() int;
setNumberOfIndexableFields(value int);
copyAndExtendWith(value IsObject) IsArray;
copyIndexableFieldsTo(destination IsArray);
}

type Array struct {
	Object
	IndexableFields []interface{}
}

// func NewArray() *Array {
// 	return &Array{}
// }

func NewArray(n int) *Array {
	a := &Array{}
	a.Object.ObjectInit("",2)
	a.ArrayInit(n)
	return a
}

func (a *Array) ArrayInit(n int) {
	a.IndexableFields = make([]interface{}, n)
}

// frame is a subtype of array

func (a *Array) getIndexableField(index int) interface{} {
	// Get the indexable field with the given index
	return a.IndexableFields[index]
}

func (a *Array) setIndexableField(index int, value interface{}) {
	// Set the indexable field with the given index to the given value
	a.IndexableFields[index] = &value
}

func (a *Array) getNumberOfIndexableFields() int {
	// Get the number of indexable fields in this array
	return len(a.IndexableFields)
}

func (a *Array) setNumberOfIndexableFields(value int) {
	// Allocate a new array of indexable fields
	a.IndexableFields = make([]interface{}, value) //new Object[value];

	// Clear each and every field by putting nil into them
	for i := 0; i < a.getNumberOfIndexableFields(); i++ {
		a.setIndexableField(i, GetUniverse().NilObject)
	}
}

func (a *Array) copyAndExtendWith(value *Object) *Array {
	// Allocate a new array which has one indexable field more than this array
	result := GetUniverse().NewArray(a.getNumberOfIndexableFields() + 1)

	// Copy the indexable fields from this array to the new array
	a.copyIndexableFieldsTo(result)

	// Insert the given object as the last indexable field in the new array
	result.setIndexableField(a.getNumberOfIndexableFields(), value)

	// Return the new array
	return result
}

func (a *Array) copyIndexableFieldsTo(destination *Array) {
	// Copy all indexable fields from this array to the destination array
	for i := 0; i < a.getNumberOfIndexableFields(); i++ {
		destination.setIndexableField(i, a.getIndexableField(i))
	}
}
