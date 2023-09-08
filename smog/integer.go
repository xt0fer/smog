package smog

import "fmt"

type Integer struct {
	embeddedInt int
}

func NewInteger(i int) *Integer {
	return &Integer{embeddedInt: i}
}

func (i *Integer) getEmbeddedInteger() int {
	return i.embeddedInt
}

func (i *Integer) setEmbeddedInteger(value int) {
	i.embeddedInt = value
}

func (s *Integer) Integer() string {
	return fmt.Sprintf("%v", s.embeddedInt)
}
