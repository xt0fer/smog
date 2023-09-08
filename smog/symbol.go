package smog

type Symbol struct {
	Object
	Name                       string
	NumberOfSignatureArguments int
}

func NewSymbol(name string) *Symbol {
	ns := &Symbol{}
	ns.Name = name
	ns.determineNumberOfSignatureArguments()
	return ns
}

func (s *Symbol) SetString(value string) {
	s.Name = value
	s.determineNumberOfSignatureArguments()
}

func (s *Symbol) String() string {
	return s.Name
}

func (s *Symbol) determineNumberOfSignatureArguments() {
	// Check for binary signature
	if s.isBinarySignature() {
		s.NumberOfSignatureArguments = 2
	} else {
		// Count the colons in the signature string
		numberOfColons := 0

		// Iterate through every character in the signature string
		for _, c := range s.Name {
			if c == ':' {
				numberOfColons++
			}
		}
		// The number of arguments is equal to the number of colons plus one
		s.NumberOfSignatureArguments = numberOfColons + 1
	}
}

func (s *Symbol) GetNumberOfSignatureArguments() int {
	return s.NumberOfSignatureArguments
}

func (s *Symbol) isBinarySignature() bool {
	for _, c := range s.Name {
		if c != '~' && c != '&' && c != '|' && c != '*' &&
			c != '/' && c != '@' && c != '+' && c != '-' &&
			c != '=' && c != '>' && c != '<' && c != ',' &&
			c != '%' && c != '\\' {
			return false
		}
	}
	return true
}
