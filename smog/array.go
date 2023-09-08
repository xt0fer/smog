package smog

type IsArray interface {
	getIndexableField(index int) interface{}
	setIndexableField(index int, value interface{})
	getNumberOfIndexableFields() int
	setNumberOfIndexableFields(value int)
	copyAndExtendWith(value IsObject) IsArray
	copyIndexableFieldsTo(destination IsArray)
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
	a.Object.ObjectInit("", 2)
	a.ArrayInit(n)
	return a
}

func (a *Array) ArrayInit(n int) {
	a.IndexableFields = make([]interface{}, n)
}

// frame is a subtype of array

func (a *Array) GetIndexableField(index int) interface{} {
	// Get the indexable field with the given index
	return a.IndexableFields[index]
}

func (a *Array) SetIndexableField(index int, value interface{}) {
	// Set the indexable field with the given index to the given value
	a.IndexableFields[index] = &value
}

func (a *Array) GetNumberOfIndexableFields() int {
	// Get the number of indexable fields in this array
	return len(a.IndexableFields)
}

func (a *Array) SetNumberOfIndexableFields(value int) {
	// Allocate a new array of indexable fields
	a.IndexableFields = make([]interface{}, value) //new Object[value];

	// Clear each and every field by putting nil into them
	for i := 0; i < a.GetNumberOfIndexableFields(); i++ {
		a.SetIndexableField(i, GetUniverse().NilObject)
	}
}

func (a *Array) copyAndExtendWith(value interface{}) *Array {
	// Allocate a new array which has one indexable field more than this array
	result := GetUniverse().NewArray(a.GetNumberOfIndexableFields() + 1)

	// Copy the indexable fields from this array to the new array
	a.copyIndexableFieldsTo(result)

	// Insert the given object as the last indexable field in the new array
	result.SetIndexableField(a.GetNumberOfIndexableFields(), value)

	// Return the new array
	return result
}

func (a *Array) copyIndexableFieldsTo(destination *Array) {
	// Copy all indexable fields from this array to the destination array
	for i := 0; i < a.GetNumberOfIndexableFields(); i++ {
		destination.SetIndexableField(i, a.GetIndexableField(i))
	}
}
