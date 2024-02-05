package vm

import (
	"testing"
)

func TestNewSymbol(t *testing.T) {
	s := NewSymbol("foo", 0)

	if s.Name != "foo" || s.Nargs != 0 {
		t.Errorf("TestNewSymbol Fail")
	}
}
func TestSymbolGet(t *testing.T) {
	s := NewSymbol("foo", 0)

	if s.Name != "foo" || s.Nargs != 0 {
		t.Errorf("TestNewSymbol Fail")
	}
}
func TestSymbolSet(t *testing.T) {
	s := NewSymbol("foo", 0)

	if s.Name != "foo" || s.Nargs != 0 {
		t.Errorf("TestNewSymbol Fail")
	}
}
