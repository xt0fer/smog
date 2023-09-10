package compiler

import (
	"fmt"

	"github.com/xt0fer/smog"
)

type Shell struct{}

func (s *Shell) Start()                            {}
func (s *Shell) SetBootstrapMethod(m *smog.Method) {}
func (s *Shell) GetBootstrapMethod() *smog.Method  { return nil }

type SourcecodeCompiler struct{}

func SourcecodeCompileClass(cp string, className string, systemClass *smog.Class) *smog.Class {
	return nil
}

type Disassembler struct{}

func (d *Disassembler) DisassembleClass(c *smog.Class) string { return "" }
func DisassemblerDump(m *smog.Class) {
	fmt.Println("")
}

//	enum Symbol {
//	    NONE, Integer, Not, And, Or, Star, Div, Mod, Plus,
//	    Minus, Equal, More, Less, Comma, At, Per, NewBlock,
//	    EndBlock, Colon, Period, Exit, Assign, NewTerm, EndTerm, Pound,
//	    Primitive, Separator, STString, Identifier, Keyword, KeywordSequence,
//	    OperatorSequence
//	}
type Token int

const (
	NONE Token = iota
	Integer
	Not
	And
	Or
	Star
	Div
	Mod
	Plus
	Minus
	Equal
	More
	Less
	Comma
	At
	Per
	NewBlock
	EndBlock
	Colon
	Period
	Exit
	Assign
	NewTerm
	EndTerm
	Pound
	Primitive
	Separator
	STString
	Identifier
	Keyword
	KeywordSequence
	OperatorSequence
)
