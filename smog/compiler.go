package smog

import "fmt"

type Compiler struct {
	Parser *Parser
}

func NewCompiler() *Compiler {
	c := &Compiler{}
	c.Parser = NewParser()
	return c
}

func (c *Compiler) CompileClassInFile(path string, file string, systemClass *Class) *Class {
	return c.Parser.CompileClass(path, file, systemClass)
}

func (c *Compiler) CompileClassInString(stmt string, systemClass *Class) *Class {
	return c.Parser.CompileClassString(stmt, systemClass)
}

// public static som.vmobjects.Class compileClass(String stmt, som.vmobjects.Class systemClass) {
// 	return new SourcecodeCompiler().compileClassString(stmt, systemClass);
// }

//	private som.vmobjects.Class compile(String path, String file, som.vmobjects.Class systemClass) throws IOException {
//		som.vmobjects.Class result = systemClass;
//		String fname = path + Universe.fileSeparator + file + ".som";
//		parser = new Parser(new FileReader(fname));
//		result = compile(systemClass);
//		som.vmobjects.Symbol cname = result.getName();
//		String cnameC = cname.getString();
//		if(file != cnameC)
//			throw new IllegalStateException("File name " + file + " does not match class name " + cnameC);
//		return result;
//	}
func (c *Compiler) Compile(path string, file string, systemClass *Class) *Class {
	result := systemClass
	fname := path + "/" + file + ".som"
	c.Parser = NewParser(io.FileReader(fname))
	result = c.Parser.Compile(systemClass)
	cname := result.GetName()
	cnameC := cname.GetString()
	if file != cnameC {
		panic("File name " + file + " does not match class name " + cnameC)
	}
	return result
}

//	private som.vmobjects.Class compileClassString(String stream, som.vmobjects.Class systemClass) {
//		parser = new Parser(new StringReader(stream));
//		som.vmobjects.Class result = compile(systemClass);
//		return result;
//	}
func (c *Compiler) CompileClassString(stream string, systemClass *Class) *Class {
	c.Parser = NewParser(io.StringReader(stream))
	result := c.Parser.Compile(systemClass)
	return result
}

//	private som.vmobjects.Class compile(som.vmobjects.Class systemClass) {
//		ClassGenerationContext cgc = new ClassGenerationContext();
//		som.vmobjects.Class result = systemClass;
//		parser.classdef(cgc);
//		if(systemClass == null)
//			result = cgc.assemble();
//		else
//			cgc.assembleSystemClass(result);
//		return result;
//	}
func (c *Compiler) Compile(systemClass *Class) *Class {
	cgc := NewClassGenerator()
	result := systemClass
	c.Parser.Classdef(cgc)
	if systemClass == nil {
		result = cgc.Assemble()
	} else {
		cgc.AssembleSystemClass(result)
	}
	return result
}

func Dump(cl *Class) {
	for i := 0; i < cl.GetNumberOfInstanceInvokables(); i++ {
		inv := cl.GetInstanceInvokable(i)
		// output header and skip if the Invokable is a Primitive
		fmt.Printf("%s>>%s = ", cl.GetName().ToString(), inv.GetSignature().ToString())
		if inv.IsPrimitive() {
			fmt.Println("<primitive>")
			continue
		}
		// output actual method
		dumpMethod(inv.(*Method), "\t")
	}
}

func dumpMethod(m *Method, indent string) {
	fmt.Println("(")
	// output stack information
	fmt.Printf("%s<%d locals, %d stack, %d bc_count>\n", indent, m.GetNumberOfLocals(), m.GetMaximumNumberOfStackElements(), m.GetNumberOfBytecodes())
	// output bytecodes
	for b := 0; b < m.GetNumberOfBytecodes(); b += GetBytecodeLength(m.GetBytecode(b)) {
		fmt.Print(indent)
		// bytecode index
		if b < 10 {
			fmt.Print(" ")
		}
		if b < 100 {
			fmt.Print(" ")
		}
		fmt.Printf(" %d:", b)
		// mnemonic
		bytecode := m.GetBytecode(b)
		fmt.Printf("%s  ", BytecodeNames[bytecode])
		// parameters (if any)
		if GetBytecodeLength(bytecode) == 1 {
			fmt.Println()
			continue
		}
		switch bytecode {
		case PUSHLOCAL:
			fmt.Printf("local: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case PUSHARGUMENT:
			fmt.Printf("argument: %d, context %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case PUSHFIELD:
			fmt.Printf("(index: %d) field: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*Symbol).ToString())
		case PUSHBLOCK:
			fmt.Printf("block: (index: %d) ", m.GetBytecode(b+1))
			dumpMethod(m.GetConstant(b).(*Method), indent+"\t")
		case PUSHCONSTANT:
			constant := m.GetConstant(b)
			fmt.Printf("(index: %d) value: (%s) %s\n", m.GetBytecode(b+1), constant.GetSOMClass().GetName().ToString(), constant.ToString())
		case PUSHGLOBAL:
			fmt.Printf("(index: %d) value: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*Symbol).ToString())
		case POPLOCAL:
			fmt.Printf("local: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case POPARGUMENT:
			fmt.Printf("argument: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case POPFIELD:
			fmt.Printf("(index: %d) field: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*Symbol).ToString())
		case SEND:
			fmt.Printf("(index: %d) signature: %s\n", m.GetBytecode(b+1), m.GetConstant(b).(*Symbol).ToString())
		case SUPERSEND:
			fmt.Printf("(index: %d) signature: %s\n", m.GetBytecode(b+1), m.GetConstant(b).(*Symbol).ToString())
		default:
			fmt.Println("<incorrect bytecode>")
		}
	}
	fmt.Println(indent + ")")
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
	tInteger
	tNot
	tAnd
	tOr
	tStar
	tDiv
	tMod
	tPlus
	tMinus
	tEqual
	tMore
	tLess
	tComma
	tAt
	tPer
	tNewBlock
	tEndBlock
	tColon
	tPeriod
	tExit
	tAssign
	tNewTerm
	tEndTerm
	tPound
	tPrimitive
	tSeparator
	tSTString
	tIdentifier
	tKeyword
	tKeywordSequence
	tOperatorSequence
)
