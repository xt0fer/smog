package vmobjects

type Symbol struct {
	Name  string
	Nargs int
}

func NewSymbol(name string, args int) *Symbol {
	ns := &Symbol{}
	ns.Name = name
	ns.Nargs = args
	return ns
}

func (s *Symbol) SetString(value string) {
	s.Name = value
	s.determineNumberOfSignatureArguments()
}

func (s *Symbol) determineNumberOfSignatureArguments() {
	// Check for binary signature
	if s.isBinarySignature() {
		s.Nargs = 2
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
		s.Nargs = numberOfColons + 1
	}
}

func (s *Symbol) getNumberOfSignatureArguments() int {
	return s.Nargs
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
