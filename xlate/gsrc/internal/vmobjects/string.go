package vmobjects

type String struct {
	value string
}

func NewString(value string) *String {
	return &String{value}
}

func (s *String) ToString() string {
	return s.value
}

func (s *String) Value() string {
	return s.value
}

func (s *String) SetValue(value string) {
	s.value = value
}
