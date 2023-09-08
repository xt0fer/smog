package smog

import "fmt"

type String struct {
	Str string
}

func NewString(s string) *String {
	return &String{Str: s}
}

func (s *String) getEmbeddedString() string {
	return s.Str
}

func (s *String) setEmbeddedString(value string) {
	s.Str = value
}

func (s *String) String() string {
	return fmt.Sprintf("%v", s.Str)
}
