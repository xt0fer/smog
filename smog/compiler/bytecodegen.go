package compiler

import (
	"github.com/xt0fer/smog"
)

type BytecodeGenerator struct {
}

func (bc *BytecodeGenerator) EmitPOP(mgenc *MethodGenerator) {
	mgenc.AddBytecode(smog.POP)
}

func (bc *BytecodeGenerator) EmitPUSHARGUMENT(mgenc *MethodGenerator, index byte) {
	mgenc.AddBytecode(smog.PUSHARGUMENT)
	mgenc.AddBytecode(index)
}

func (bc *BytecodeGenerator) EmitRETURNLOCAL(mgenc *MethodGenerator) {
	mgenc.AddBytecode(smog.RETURNLOCAL)
}

func (bc *BytecodeGenerator) EmitRETURNNONLOCAL(mgenc *MethodGenerator) {
	mgenc.AddBytecode(smog.RETURNNONLOCAL)
}

func (bc *BytecodeGenerator) EmitDUP(mgenc *MethodGenerator) {
	mgenc.AddBytecode(smog.DUP)
}	

func (bc *BytecodeGenerator) EmitPUSHBLOCK(mgenc *MethodGenerator, blockMethod *smog.Method) {
	mgenc.AddBytecode(smog.PUSHBLOCK)
	mgenc.AddBytecode(mgenc.findLiteralIndex(blockMethod))
}

func (bc *BytecodeGenerator) EmitPUSHLOCAL(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(smog.PUSHLOCAL)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPUSHFIELD(mgenc *MethodGenerator, fieldName string) {
	mgenc.AddBytecode(smog.PUSHFIELD)
	mgenc.AddBytecode(mgenc.findLiteralIndex(fieldName))
}

func (bc *BytecodeGenerator) EmitPUSHGLOBAL(mgenc *MethodGenerator, global string) {
	mgenc.AddBytecode(smog.PUSHGLOBAL)
	mgenc.AddBytecode(mgenc.findLiteralIndex(global))
}

func (bc *BytecodeGenerator) EmitPOPARGUMENT(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(smog.POPARGUMENT)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPOPLOCAL(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(smog.POPLOCAL)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPOPFIELD(mgenc *MethodGenerator, fieldName string) {
	mgenc.AddBytecode(smog.POPFIELD)
	mgenc.AddBytecode(mgenc.findLiteralIndex(fieldName))
}

func (bc *BytecodeGenerator) EmitSUPERSEND(mgenc *MethodGenerator, msg string) {
	mgenc.AddBytecode(smog.SUPERSEND)
	mgenc.AddBytecode(mgenc.findLiteralIndex(msg))
}

func (bc *BytecodeGenerator) EmitSEND(mgenc *MethodGenerator, msg string) {
	mgenc.AddBytecode(smog.SEND)
	mgenc.AddBytecode(mgenc.findLiteralIndex(msg))
}

func (bc *BytecodeGenerator) EmitPUSHCONSTANT(mgenc *MethodGenerator, lit byte) {
	mgenc.AddBytecode(smog.PUSHCONSTANT)
	mgenc.AddBytecode(mgenc.findLiteralIndex(lit))
}
