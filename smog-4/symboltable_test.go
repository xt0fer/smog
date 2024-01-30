package smog

import (
	"reflect"
	"testing"
)

func TestSymbolTable_insert0(t *testing.T) {
	st := NewSymbolTable()
	expected := NewSymbol("foo")
	st.insert(expected)
	actual := st.lookup("foo")
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("insert0() = %v, want %v", expected, actual)
	}
}
func TestSymbolTable_insert1(t *testing.T) {
	st := NewSymbolTable()
	expected := NewSymbol("foo")
	st.insert(expected)
	actual := st.lookup("f00")
	if reflect.DeepEqual(expected, actual) { // nb: we expect the DeepCopy to fail
		t.Errorf("insert1() = %v, want %v", actual, expected)
	}
}
