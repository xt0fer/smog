package vm

type Array struct {
	IndexableFields []*Object
}

// func NewArray() *Array {
// 	return &Array{}
// }

func NewArray(n int) *Array {
	a := &Array{}
	a.IndexableFields = make([]*Object, n)
	return a
}

// frame is a subtype of array

func (a *Array) getIndexableField(index int) *Object {
	// Get the indexable field with the given index
	return a.IndexableFields[index]
}

func (a *Array) setIndexableField(index int, value *Object) {
	// Set the indexable field with the given index to the given value
	a.IndexableFields[index] = value
}

func (a *Array) getNumberOfIndexableFields() int {
	// Get the number of indexable fields in this array
	return len(a.IndexableFields)
}

func (a *Array) setNumberOfIndexableFields(value int) {
	// Allocate a new array of indexable fields
	a.IndexableFields = make([]*Object, value) //new Object[value];

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
