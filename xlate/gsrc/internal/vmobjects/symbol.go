package vmobjects

type Symbol struct {
	value string
	nArgs int
}

func NewSymbol(value string, nArgs int) *Symbol {
	return &Symbol{value, nArgs}
}

func (s *Symbol) ToString() string {
	return s.value
}

func (s *Symbol) NumberOfSignatureArguments() int {
	return s.nArgs
}

func (s *Symbol) IsBinarySignature() bool {
	for _, c := range s.value {
		if c != '~' && c != '&' && c != '|' && c != '*' &&
			c != '/' && c != '@' && c != '+' && c != '-' &&
			c != '=' && c != '>' && c != '<' && c != ',' &&
			c != '%' && c != '\\' {
			return false
		}
	}
	return true
}

func (s *Symbol) SetNumberOfSignatureArguments(nArgs int) {
	s.nArgs = nArgs
}

func (s *Symbol) SetValue(value string) {
	s.value = value
}

func (s *Symbol) Value() string {
	return s.value
}