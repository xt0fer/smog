package vmobj

type VMString struct {
	value string
}

func NewVMString(value string) *VMString {
	return &VMString{value}
}

func (s *VMString) ToString() string {
	return s.value
}

func (s *VMString) Value() string {
	return s.value
}

func (s *VMString) SetValue(value string) {
	s.value = value
}
