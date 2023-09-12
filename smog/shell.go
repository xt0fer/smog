package smog

type Shell struct{}

func (s *Shell) Start()                       {}
func (s *Shell) SetBootstrapMethod(m *Method) {}
func (s *Shell) GetBootstrapMethod() *Method  { return nil }
