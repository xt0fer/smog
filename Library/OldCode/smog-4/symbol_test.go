package smog

import (
	"testing"
)

func TestNewSymbol(t *testing.T) {
	s := NewSymbol("foo")

	if s.Name != "foo" {
		t.Errorf("TestNewSymbol Fail")
	}
}
func TestSymbolGet(t *testing.T) {
	s := NewSymbol("foo")

	if s.Name != "foo" {
		t.Errorf("TestNewSymbol Fail")
	}
}
func TestSymbolSet(t *testing.T) {
	s := NewSymbol("foo")

	if s.Name != "foo" {
		t.Errorf("TestNewSymbol Fail")
	}
}
