package smog

import "fmt"

type Shell struct{}

func (s *Shell) Start()                       {}
func (s *Shell) SetBootstrapMethod(m *Method) {}
func (s *Shell) GetBootstrapMethod() *Method  { return nil }

type SourcecodeCompiler struct{}

func SourcecodeCompileClass(cp string, className string, systemClass *Class) *Class {
	return nil
}

type Disassembler struct{}

func (d *Disassembler) DisassembleClass(c *Class) string { return "" }
func DisassemblerDump(m *Class) {
	fmt.Println("")
}
