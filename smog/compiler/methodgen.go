package compiler

import (
	"github.com/xt0fer/smog"
)

// private ClassGenerationContext holderGenc;
// private MethodGenerationContext outerGenc;
// private boolean blockMethod;
// private som.vmobjects.Symbol signature;
// private List<String> arguments = new ArrayList<String>();
// private boolean primitive;
// private List<String> locals = new ArrayList<String>();
// private List<som.vmobjects.Object> literals = new ArrayList<som.vmobjects.Object>();
// private boolean finished;
// private Vector<Byte> bytecode = new Vector<Byte>();
type MethodGenerator struct {
	HolderGenc  *ClassGenerator
	OuterGenc   *MethodGenerator
	BlockMethod bool
	Signature   *smog.Symbol
	Arguments   []string
	Primitive   bool
	Locals      []string
	Literals    []*smog.Object
	Finished    bool
	Bytecode    []byte
}

//	public void setHolder(ClassGenerationContext cgenc) {
//		holderGenc = cgenc;
//	}
func (mg *MethodGenerator) SetHolder(cgenc *ClassGenerator) {
	mg.HolderGenc = cgenc
}

//	public void addArgument(String arg) {
//		arguments.add(arg);
//	}
func (mg *MethodGenerator) AddArgument(arg string) {
	mg.Arguments = append(mg.Arguments, arg)
}

//	public boolean isPrimitive() {
//		return primitive;
//	}
func (mg *MethodGenerator) IsPrimitive() bool {
	return mg.Primitive
}

//	public Invokable assemblePrimitive() {
//		return Primitive.getEmptyPrimitive(signature.getString());
//	}
func (mg *MethodGenerator) AssemblePrimitive() *smog.Invokable {
	return smog.NewEmptyPrimitive(mg.Signature.GetString())
}

//	public Method assemble() {
//		// create a method instance with the given number of bytecodes and literals
//		int numLiterals = literals.size();
//		Method meth = Universe.newMethod(signature, bytecode.size(), numLiterals);
//		// populate the fields that are immediately available
//		int numLocals = locals.size();
//		meth.setNumberOfLocals(numLocals);
//		meth.setMaximumNumberOfStackElements(computeStackDepth());
//		// copy literals into the method
//		int i = 0;
//		for(som.vmobjects.Object l : literals)
//			meth.setIndexableField(i++, l);
//		// copy bytecodes into method
//		i = 0;
//		for(byte bc : bytecode)
//			meth.setBytecode(i++, bc);
//		// return the method - the holder field is to be set later on!
//		return meth;
//	}
func (mg *MethodGenerator) Assemble() *smog.Method {
	// create a method instance with the given number of bytecodes and literals
	numLiterals := len(mg.Literals)
	meth := smog.NewMethod(mg.signature, len(mg.Bytecode), numLiterals)
	// populate the fields that are immediately available
	numLocals := len(mg.Locals)
	meth.SetNumberOfLocals(numLocals)
	meth.SetMaximumNumberOfStackElements(mg.computeStackDepth())
	// copy literals into the method
	i := 0
	for _, l := range mg.Literals {
		meth.SetIndexableField(i, l)
		i++
	}
	// copy bytecodes into method
	i = 0
	for _, bc := range mg.Bytecode {
		meth.SetBytecode(i, bc)
		i++
	}
	// return the method - the holder field is to be set later on!
	return meth
}

// private int computeStackDepth() {
// 	int depth = 0;
// 	int maxDepth = 0;
// 	int i = 0;

// 	while(i < bytecode.size()) {
// 		switch(bytecode.elementAt(i)) {
// 			case halt             :          i++;    break;
// 			case dup              : depth++; i++;    break;
// 			case push_local       :
// 			case push_argument    : depth++; i += 3; break;
// 			case push_field       :
// 			case push_block       :
// 			case push_constant    :
// 			case push_global      : depth++; i += 2; break;
// 			case pop              : depth--; i++;    break;
// 			case pop_local        :
// 			case pop_argument     : depth--; i += 3; break;
// 			case pop_field        : depth--; i += 2; break;
// 			case send             :
// 			case super_send       : {
// 				// these are special: they need to look at the number of
// 				// arguments (extractable from the signature)
// 				som.vmobjects.Symbol sig = (som.vmobjects.Symbol)literals.get(bytecode.elementAt(i+1));

// 				depth -= sig.getNumberOfSignatureArguments();

// 				depth++; // return value
// 				i += 2;
// 				break;
// 			}
// 			case return_local     :
// 			case return_non_local :          i++;    break;
// 			default               :
// 				throw new IllegalStateException("Illegal bytecode " + bytecode.elementAt(i));
// 		}

// 		if(depth > maxDepth)
// 			maxDepth = depth;
// 	}

//		return maxDepth;
//	}
func (mg *MethodGenerator) computeStackDepth() int {
	depth := 0
	maxDepth := 0
	i := 0

	for i < len(mg.Bytecode) {
		switch mg.Bytecode[i] {
		case smog.Halt:
			i++
		case smog.Dup:
			depth++
			i++
		case smog.PushLocal, smog.PushArgument:
			depth++
			i += 3
		case smog.PushField, smog.PushBlock, smog.PushConstant, smog.PushGlobal:
			depth++
			i += 2
		case smog.Pop:
			depth--
			i++
		case smog.PopLocal, smog.PopArgument:
			depth--
			i += 3
		case smog.PopField:
			depth--
			i += 2
		case smog.Send, smog.SuperSend:
			// these are special: they need to look at the number of
			// arguments (extractable from the signature)
			sig := mg.Literals[mg.Bytecode[i+1]].(*smog.Symbol)

			depth -= sig.GetNumberOfSignatureArguments()

			depth++ // return value
			i += 2
		case smog.ReturnLocal, smog.ReturnNonLocal:
			i++
		default:
			panic("Illegal bytecode")
		}

		if depth > maxDepth {
			maxDepth = depth
		}
	}

	return maxDepth
}

//	public void setPrimitive(boolean prim) {
//		primitive = prim;
//	}
func (mg *MethodGenerator) SetPrimitive(prim bool) {
	mg.Primitive = prim
}

//	public void setSignature(Symbol sig) {
//		signature = sig;
//	}
func (mg *MethodGenerator) SetSignature(sig *smog.Symbol) {
	mg.Signature = sig
}

//	public boolean addArgumentIfAbsent(String arg) {
//		if (locals.indexOf(arg) != -1)
//			return false;
//		arguments.add(arg);
//		return true;
//	}
func (mg *MethodGenerator) AddArgumentIfAbsent(arg string) bool {
	if mg.Locals.indexOf(arg) != -1 {
		return false
	}
	mg.Arguments = append(mg.Arguments, arg)
	return true
}

//	public boolean isFinished() {
//		return finished;
//	}
func (mg *MethodGenerator) IsFinished() bool {
	return mg.Finished
}

//	public void setFinished(boolean finished) {
//		this.finished = finished;
//	}
func (mg *MethodGenerator) SetFinished(finished bool) {
	mg.Finished = finished
}

//	public boolean addLocalIfAbsent(String local) {
//		if(locals.indexOf(local) != -1)
//			return false;
//		locals.add(local);
//		return true;
//	}
func (mg *MethodGenerator) AddLocalIfAbsent(local string) bool {
	if mg.Locals.indexOf(local) != -1 {
		return false
	}
	mg.Locals = append(mg.Locals, local)
	return true
}

//	public void addLocal(String local) {
//		locals.add(local);
//	}
func (mg *MethodGenerator) AddLocal(local string) {
	mg.Locals = append(mg.Locals, local)
}

//	public void removeLastBytecode() {
//		bytecode.removeElementAt(bytecode.size()-1);
//	}
func (mg *MethodGenerator) RemoveLastBytecode() {
	mg.Bytecode = mg.Bytecode[:len(mg.Bytecode)-1]
}

//	public boolean isBlockMethod() {
//		return blockMethod;
//	}
func (mg *MethodGenerator) IsBlockMethod() bool {
	return mg.BlockMethod
}

//	public void setFinished() {
//		finished = true;
//	}
func (mg *MethodGenerator) Finish() {
	mg.Finished = true
}

//	public boolean addLiteralIfAbsent(som.vmobjects.Object lit) {
//		if(literals.indexOf(lit) != -1)
//			return false;
//		literals.add(lit);
//		return true;
//	}
func (mg *MethodGenerator) AddLiteralIfAbsent(lit *smog.Object) bool {
	if mg.Literals.indexOf(lit) != -1 {
		return false
	}
	mg.Literals = append(mg.Literals, lit)
	return true
}

//	public void setIsBlockMethod(boolean isBlock) {
//		blockMethod = isBlock;
//	}
func (mg *MethodGenerator) SetIsBlockMethod(isBlock bool) {
	mg.BlockMethod = isBlock
}

//	public ClassGenerationContext getHolder() {
//		return holderGenc;
//	}
func (mg *MethodGenerator) GetHolder() *ClassGenerator {
	return mg.HolderGenc
}

//	public void setOuter(MethodGenerationContext mgenc) {
//		outerGenc = mgenc;
//	}
func (mg *MethodGenerator) SetOuter(mgenc *MethodGenerator) {
	mg.OuterGenc = mgenc
}

//	public void addLiteral(som.vmobjects.Object lit) {
//		literals.add(lit);
//	}
func (mg *MethodGenerator) AddLiteral(lit *smog.Object) {
	mg.Literals = append(mg.Literals, lit)
}

//	public boolean findVar(String var, Triplet<Byte, Byte, Boolean> tri) {
//		// triplet: index, context, isArgument
//		tri.setX((byte) locals.indexOf(var));
//		if(tri.getX() == -1) {
//			tri.setX((byte) arguments.indexOf(var));
//			if(tri.getX() == -1) {
//				if(outerGenc == null)
//					return false;
//				else {
//					tri.setY((byte) (tri.getY() + 1));
//					return outerGenc.findVar(var, tri);
//				}
//			} else
//				tri.setZ(true);
//		}
//		return true;
//	}
func (mg *MethodGenerator) FindVar(v string, tri *Triplet) bool {
	// triplet: index, context, isArgument
	tri.SetX(byte(mg.Locals.indexOf(v)))
	if tri.GetX() == -1 {
		tri.SetX(byte(mg.Arguments.indexOf(v)))
		if tri.GetX() == -1 {
			if mg.OuterGenc == nil {
				return false
			} else {
				tri.SetY(byte(tri.GetY() + 1))
				return mg.OuterGenc.FindVar(v, tri)
			}
		} else {
			tri.SetZ(true)
		}
	}
	return true
}

//	public boolean findField(String field) {
//		return holderGenc.findField(field);
//	}
func (mg *MethodGenerator) FindField(field string) bool {
	return mg.HolderGenc.FindField(field)
}

//	public int getNumberOfArguments() {
//		return arguments.size();
//	}
func (mg *MethodGenerator) GetNumberOfArguments() int {
	return len(mg.Arguments)
}

//	public void addBytecode(byte code) {
//		bytecode.add(code);
//	}
func (mg *MethodGenerator) AddBytecode(code byte) {
	mg.Bytecode = append(mg.Bytecode, code)
}

//	public byte findLiteralIndex(som.vmobjects.Object lit) {
//		return (byte) literals.indexOf(lit);
//	}
func (mg *MethodGenerator) FindLiteralIndex(lit *smog.Object) byte {
	return byte(mg.Literals.indexOf(lit))
}

//	public MethodGenerationContext getOuter() {
//		return outerGenc;
//	}
func (mg *MethodGenerator) GetOuter() *MethodGenerator {
	return mg.OuterGenc
}

//	public som.vmobjects.Symbol getSignature() {
//		return signature;
//	}
func (mg *MethodGenerator) GetSignature() *smog.Symbol {
	return mg.Signature
}
