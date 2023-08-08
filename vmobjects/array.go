package vmobjects

type Array struct {
	Fields []*Object

}

func NewArray() *Array {
	return &Array{}
}

// frame is a subtype of array

func (a *Array) getIndexableField(int index) *Object  {
    // Get the indexable field with the given index
    return indexableFields[index];
  }

  func (a *Array) setIndexableField(int index, Object value) {
    // Set the indexable field with the given index to the given value
    indexableFields[index] = value;
  }
  
  func (a *Array) getNumberOfIndexableFields() int {
    // Get the number of indexable fields in this array
    return indexableFields.length;
  }

  func (a *Array) setNumberOfIndexableFields(int value) {
    // Allocate a new array of indexable fields
    indexableFields = new Object[value];
    
    // Clear each and every field by putting nil into them
    for (int i := 0; i < getNumberOfIndexableFields(); i++) {
      setIndexableField(i, Universe.nilObject);
    }
  }

  func (a *Array) copyAndExtendWith(Object value) *Array {
    // Allocate a new array which has one indexable field more than this array
    Array result = Universe.newArray(getNumberOfIndexableFields() + 1);
    
    // Copy the indexable fields from this array to the new array
    copyIndexableFieldsTo(result);
    
    // Insert the given object as the last indexable field in the new array
    result.setIndexableField(getNumberOfIndexableFields(), value);
    
    // Return the new array
    return result;
  }

  func (a *Array) copyIndexableFieldsTo(Array destination) {
    // Copy all indexable fields from this array to the destination array
    for (int i := 0; i < getNumberOfIndexableFields(); i++) {
      destination.setIndexableField(i, getIndexableField(i));
    }
  }
  