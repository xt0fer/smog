package smog

type BytecodeGenerator struct {
}

func NewBytecodeGenerator() *BytecodeGenerator {
	return &BytecodeGenerator{}
}

func (bc *BytecodeGenerator) EmitPOP(mgenc *MethodGenerator) {
	mgenc.AddBytecode(POP)
}

func (bc *BytecodeGenerator) EmitPUSHARGUMENT(mgenc *MethodGenerator, index byte) {
	mgenc.AddBytecode(PUSHARGUMENT)
	mgenc.AddBytecode(index)
}

func (bc *BytecodeGenerator) EmitRETURNLOCAL(mgenc *MethodGenerator) {
	mgenc.AddBytecode(RETURNLOCAL)
}

func (bc *BytecodeGenerator) EmitRETURNNONLOCAL(mgenc *MethodGenerator) {
	mgenc.AddBytecode(RETURNNONLOCAL)
}

func (bc *BytecodeGenerator) EmitDUP(mgenc *MethodGenerator) {
	mgenc.AddBytecode(DUP)
}

func (bc *BytecodeGenerator) EmitPUSHBLOCK(mgenc *MethodGenerator, blockMethod *Method) {
	mgenc.AddBytecode(PUSHBLOCK)
	mgenc.AddBytecode(mgenc.findLiteralIndex(blockMethod))
}

func (bc *BytecodeGenerator) EmitPUSHLOCAL(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(PUSHLOCAL)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPUSHFIELD(mgenc *MethodGenerator, fieldName string) {
	mgenc.AddBytecode(PUSHFIELD)
	mgenc.AddBytecode(mgenc.findLiteralIndex(fieldName))
}

func (bc *BytecodeGenerator) EmitPUSHGLOBAL(mgenc *MethodGenerator, global string) {
	mgenc.AddBytecode(PUSHGLOBAL)
	mgenc.AddBytecode(mgenc.findLiteralIndex(global))
}

func (bc *BytecodeGenerator) EmitPOPARGUMENT(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(POPARGUMENT)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPOPLOCAL(mgenc *MethodGenerator, idx byte, ctx byte) {
	mgenc.AddBytecode(POPLOCAL)
	mgenc.AddBytecode(idx)
	mgenc.AddBytecode(ctx)
}

func (bc *BytecodeGenerator) EmitPOPFIELD(mgenc *MethodGenerator, fieldName string) {
	mgenc.AddBytecode(POPFIELD)
	mgenc.AddBytecode(mgenc.findLiteralIndex(fieldName))
}

func (bc *BytecodeGenerator) EmitSUPERSEND(mgenc *MethodGenerator, msg string) {
	mgenc.AddBytecode(SUPERSEND)
	mgenc.AddBytecode(mgenc.findLiteralIndex(msg))
}

func (bc *BytecodeGenerator) EmitSEND(mgenc *MethodGenerator, msg string) {
	mgenc.AddBytecode(SEND)
	mgenc.AddBytecode(mgenc.findLiteralIndex(msg))
}

func (bc *BytecodeGenerator) EmitPUSHCONSTANT(mgenc *MethodGenerator, lit byte) {
	mgenc.AddBytecode(PUSHCONSTANT)
	mgenc.AddBytecode(mgenc.findLiteralIndex(lit))
}
