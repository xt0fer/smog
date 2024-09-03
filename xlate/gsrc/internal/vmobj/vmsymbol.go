package vmobj

type VMSymbol struct {
	VMObject
	value string
}

func NewVMSymbol(value string) *VMSymbol {
	return &VMSymbol{NewObjInternal(0, nil), value}
}

func (s *VMSymbol) ToString() string {
	return s.value
}

// Nuber of Signature Arguments is the number of arguments in the signature
// of the method. For example, the signature of the method 'ifTrue:' is
// 'ifTrue:', so the number of signature arguments is 1.
func (s *VMSymbol) NumberOfSignatureArguments() int {
	return s.Nfields
}

func (s *VMSymbol) IsBinarySignature() bool {
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

func (s *VMSymbol) SetNumberOfSignatureArguments(Nfields int) {
	s.Nfields = Nfields
}

func (s *VMSymbol) SetValue(value string) {
	s.value = value
}

func (s *VMSymbol) Value() string {
	return s.value
}
