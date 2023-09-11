package smog

import (
	"bufio"
	"fmt"
	"unicode"
)

type BytecodeGenerator struct {
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

// private List<som.vmobjects.Object> instanceFields = new ArrayList<som.vmobjects.Object>();
// private List<som.vmobjects.Invokable> instanceMethods = new ArrayList<som.vmobjects.Invokable>();
// private List<som.vmobjects.Object> classFields = new ArrayList<som.vmobjects.Object>();
// private List<som.vmobjects.Invokable> classMethods = new ArrayList<som.vmobjects.Invokable>();
type ClassGenerator struct {
	name            *Symbol
	superName       *Symbol
	classSide       bool
	instanceFields  []*Object
	instanceMethods []Invokable
	classFields     []*Object
	classMethods    []Invokable
}

func NewClassGenerator() *ClassGenerator {
	cg := &ClassGenerator{}
	cg.instanceFields = make([]*Object, 0)
	cg.instanceMethods = make([]Invokable, 0)
	cg.classFields = make([]*Object, 0)
	cg.classMethods = make([]Invokable, 0)
	return cg
}

//	public void setName(Symbol name) {
//		this.name = name;
//	}
func (cg *ClassGenerator) SetName(name *Symbol) {
	cg.name = name
}

//	public void setSuperName(Symbol superName) {
//		this.superName = superName;
//	}
func (cg *ClassGenerator) SetSuperName(superName *Symbol) {
	cg.superName = superName
}

//	public void addInstanceMethod(som.vmobjects.Invokable meth) {
//		instanceMethods.add(meth);
//	}
func (cg *ClassGenerator) AddInstanceMethod(meth Invokable) {
	cg.instanceMethods = append(cg.instanceMethods, meth)
}

//	public void setClassSide(boolean b) {
//		classSide = b;
//	}
func (cg *ClassGenerator) SetClassSide(b bool) {
	cg.classSide = b
}

//	public void addClassMethod(som.vmobjects.Invokable meth) {
//		classMethods.add(meth);
//	}
func (cg *ClassGenerator) AddClassMethod(meth Invokable) {
	cg.classMethods = append(cg.classMethods, meth)
}

//	public void addInstanceField(Symbol field) {
//		instanceFields.add(field);
//	}
func (cg *ClassGenerator) AddInstanceField(field *Object) {
	cg.instanceFields = append(cg.instanceFields, field)
}

//	public void addClassField(Symbol field) {
//		classFields.add(field);
//	}
func (cg *ClassGenerator) AddClassField(field *Object) {
	cg.classFields = append(cg.classFields, field)
}

//	public boolean findField(String field) {
//		return (isClassSide() ? classFields : instanceFields).indexOf(Universe.symbolFor(field)) != -1;
//	}
func (cg *ClassGenerator) FindField(field string) bool {
	var fields []*Object
	if cg.classSide {
		fields = cg.classFields
	} else {
		fields = cg.instanceFields
	}
	for _, f := range fields {
		if f.GetSOMClass().Name.Name == field {
			return true
		}
	}
	return false
}

//	public boolean isClassSide() {
//		return classSide;
//	}
func (cg *ClassGenerator) IsClassSide() bool {
	return cg.classSide
}

//	public som.vmobjects.Class assemble() {
//		// build class class name
//		String ccname = name.getString() + " class";
//		// Load the super class
//		som.vmobjects.Class superClass = Universe.loadClass(superName);
//		// Allocate the class of the resulting class
//		som.vmobjects.Class resultClass = Universe.newClass(Universe.metaclassClass);
//		// Initialize the class of the resulting class
//		resultClass.setInstanceFields(Universe.newArray(classFields));
//		resultClass.setInstanceInvokables(Universe.newArray(classMethods));
//		resultClass.setName(Universe.symbolFor(ccname));
//		som.vmobjects.Class superMClass = superClass.getSOMClass();
//		resultClass.setSuperClass(superMClass);
//		// Allocate the resulting class
//		som.vmobjects.Class result = Universe.newClass(resultClass);
//		// Initialize the resulting class
//		result.setInstanceFields(Universe.newArray(instanceFields));
//		result.setInstanceInvokables(Universe.newArray(instanceMethods));
//		result.setName(name);
//		result.setSuperClass(superClass);
//		return result;
//	}
func (cg *ClassGenerator) Assemble() *Class {
	// build class class name
	ccname := cg.name.String() + " class"
	u := GetUniverse()
	// Load the super class
	superClass := u.LoadClass(cg.superName)
	// Allocate the class of the resulting class
	resultClass := u.NewClass(u.metaclassClass)
	// Initialize the class of the resulting class
	resultClass.SetInstanceFields(u.NewArray(cg.classFields))
	resultClass.SetInstanceInvokables(u.NewArray(cg.classMethods))
	resultClass.SetName(u.SymbolFor(ccname))
	superMClass := superClass.GetSOMClass()
	resultClass.SetSuperClass(superMClass)
	// Allocate the resulting class
	result := u.NewClass(resultClass)
	// Initialize the resulting class
	result.SetInstanceFields(u.NewArray(instanceFields))
	result.SetInstanceInvokables(u.NewArray(instanceMethods))
	result.SetName(cg.name)
	result.SetSuperClass(superClass)
	return result

}

// public void assembleSystemClass(som.vmobjects.Class systemClass) {
// 	systemClass.setInstanceInvokables(Universe.newArray(instanceMethods));
// 	systemClass.setInstanceFields(Universe.newArray(instanceFields));
// 	// class-bound == class-instance-bound
// 	som.vmobjects.Class superMClass = systemClass.getSOMClass();
// 	superMClass.setInstanceInvokables(Universe.newArray(classMethods));
// 	superMClass.setInstanceFields(Universe.newArray(classFields));
// }

// the NEWARRAY issue, need s special NewArray in Universe to handle these initializations
func (cg *ClassGenerator) AssembleSystemClass(systemClass *Class) {
	u := GetUniverse()
	systemClass.SetInstanceInvokables(u.NewArray(cg.instanceMethods))
	systemClass.SetInstanceFields(u.NewArray(cg.instanceFields))
	// class-bound == class-instance-bound
	superMClass := systemClass.GetSOMClass()
	superMClass.SetInstanceInvokables(u.NewArray(cg.classMethods))
	superMClass.SetInstanceFields(u.NewArray(cg.classFields))
}

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
		fmt.Printf("%s>>%s = ", cl.GetName().String(), inv.GetSignature().String())
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

// private static final String SEPARATOR = "----";
// private static final String PRIMITIVE = "primitive";
const SEPARATOR = "----"
const PRIMITIVE = "primitive"

type Lexer struct {
	lineNumber int
	infile     *bufio.Reader
	sym        Token
	symc       rune
	text       string
	peekDone   bool
	nextSym    Token
	nextSymc   rune
	nextText   string
	buf        string
	bufp       int
}

//     private int lineNumber;
//     private BufferedReader infile;
//     private Token sym;
//     private char symc;
//     private StringBuffer text;
//     private boolean peekDone;
//     private Token nextSym;
//     private char nextSymc;
//     private StringBuffer nextText;
//     private String buf;
//     private int bufp;

//	protected Lexer(Reader reader) {
//	    infile = new BufferedReader(reader);
//	    peekDone = false;
//	    buf = "";
//	    text = new StringBuffer();
//	    bufp = 0;
//	    lineNumber = 0;
//	}
func NewLexer(reader io.Reader) *Lexer {
	lexer := &Lexer{}
	lexer.infile = bufio.NewReader(reader)
	lexer.peekDone = false
	lexer.buf = ""
	lexer.text = ""
	lexer.bufp = 0
	lexer.lineNumber = 0
	return lexer
}

func (l *Lexer) GetSym() Token {
	if l.peekDone {
		l.peekDone = false
		l.sym = l.nextSym
		l.symc = l.nextSymc
		l.text = l.nextText
		return l.sym
	}
	//         do {
	//             if(!hasMoreInput()) {
	//                 sym = Token.NONE;
	//                 symc = '\0';
	//                 text = new StringBuffer(symc);
	//                 return sym;
	//             }
	//             skipWhiteSpace();
	//             skipComment();
	//         } while(endOfBuffer() || Character.isWhitespace(currentChar()) || currentChar() == '"');
	for l.hasMoreInput() {
		if !l.hasMoreInput() {
			l.sym = NONE
			l.symc = 0
			l.text = string(l.symc)
			return l.sym
		}
		l.skipWhiteSpace()
		l.skipComment()
		if l.endOfBuffer() || l.currentChar() == '"' {
			continue
		}
		//break
	}

	switch l.currentChar() {
	case '\'':
		l.sym = STString
		l.symc = 0
		l.text = ""
		//             do {
		//                 text.append(bufchar(++bufp));
		//             } while(currentChar() != '\'');
		for l.currentChar() != '\'' {
			l.text += string(l.buf[l.bufp])
			l.bufp++
		}
		l.text += string(l.bufchar(l.bufp))
		l.bufp++
	case '[':
		l.match(NewBlock)
	case ']':
		l.match(EndBlock)
	case '(':
		l.match(NewTerm)
	case ')':
		l.match(EndTerm)
	case '#':
		l.match(Pound)
	case '^':
		l.match(Exit)
	case '.':
		l.match(Period)
	case '-':
		if l.startsWith(SEPARATOR) {
			l.text = ""
			for l.currentChar() == '-' {
				l.text += string(l.bufchar(l.bufp))
				l.bufp++
			}
			l.sym = Separator
		} else {
			l.bufp++
			l.sym = Minus
			l.symc = '-'
			l.text = "-"
		}
	default:
		break
	}

	if l.isOperator(l.currentChar()) {
		if l.isOperator(l.bufchar(l.bufp + 1)) {
			l.sym = OperatorSequence
			l.symc = 0
			l.text = ""
			for l.isOperator(l.currentChar()) {
				l.text += string(l.bufchar(l.bufp))
				l.bufp++
			}
		} else {
			switch l.currentChar() {
			case '~':
				l.match(Not)
			case '&':
				l.match(And)
			case '|':
				l.match(Or)
			case '*':
				l.match(Star)
			case '/':
				l.match(Div)
			case '\\':
				l.match(Mod)
			case '+':
				l.match(Plus)
			case '=':
				l.match(Equal)
			case '>':
				l.match(More)
			case '<':
				l.match(Less)
			case ',':
				l.match(Comma)
			case '@':
				l.match(At)
			case '%':
				l.match(Per)
			default:
				break
			}
		}
	}

	if l.startsWith(PRIMITIVE) {
		l.bufp += len(PRIMITIVE)
		l.sym = Primitive
		l.symc = 0
		l.text = PRIMITIVE
	}

	if unicode.IsLetter(l.currentChar()) {
		l.symc = 0
		l.text = ""
		for unicode.IsLetter(l.currentChar()) ||
			unicode.IsDigit(l.currentChar()) || l.currentChar() == '_' {
			l.text += string(l.bufchar(l.bufp))
			l.bufp++
		}
		l.sym = Identifier
		if l.bufchar(l.bufp) == ':' {
			l.sym = Keyword
			l.bufp++
			l.text += ":"
			if unicode.IsLetter(l.currentChar()) {
				l.sym = KeywordSequence
				for unicode.IsLetter(l.currentChar()) || l.currentChar() == ':' {
					l.text += string(l.bufchar(l.bufp))
					l.bufp++
				}
			}
		}
	}
	if unicode.IsDigit(l.currentChar()) {
		l.sym = Integer
		l.symc = 0
		l.text = ""
		for unicode.IsDigit(l.currentChar()) {
			l.text += string(l.bufchar(l.bufp))
			l.bufp++
		}
	} else {
		l.sym = NONE
		l.symc = l.currentChar()
		l.text = string(l.symc)
	}

	return l.sym
}

//     protected Token peek() {
//         Token saveSym = sym;
//         char saveSymc = symc;
//         StringBuffer saveText = new StringBuffer(text);
//         if(peekDone)
//             throw new IllegalStateException("SOM lexer: cannot peek twice!");
//         getSym();
//         nextSym = sym;
//         nextSymc = symc;
//         nextText = new StringBuffer(text);
//         sym = saveSym;
//         symc = saveSymc;
//         text = saveText;
//         peekDone = true;
//         return nextSym;
//     }

func (l *Lexer) peek() Token {
	saveSym := l.sym
	saveSymc := l.symc
	saveText := ""
	if l.peekDone {
		panic("SOM lexer: cannot peek twice!")
	}
	l.GetSym()
	l.nextSym = l.sym
	l.nextSymc = l.symc
	l.nextText = l.text
	l.sym = saveSym
	l.symc = saveSymc
	l.text = saveText
	l.peekDone = true
	return l.nextSym
}

//	protected String getText() {
//	    return text.toString();
//	}
func (l *Lexer) getText() string {
	return l.text
}

//	protected String getNextText() {
//	    return nextText.toString();
//	}
func (l *Lexer) getNextText() string {
	return l.nextText
}

//	protected String getRawBuffer() {
//	    return buf;
//	}
func (l *Lexer) getRawBuffer() string {
	return l.buf
}

//	protected int getCurrentLineNumber() {
//	    return lineNumber;
//	}
func (l *Lexer) getCurrentLineNumber() int {
	return l.lineNumber
}

//	private int fillBuffer() {
//	    try {
//	        if(!infile.ready())
//	            return -1;
//	        buf = infile.readLine();
//	        if(buf == null)
//	            return -1;
//	        ++lineNumber;
//	        bufp = 0;
//	        return buf.length();
//	    } catch(IOException ioe) {
//	        throw new IllegalStateException("Error reading from input: " + ioe.toString());
//	    }
//	}
func (l *Lexer) fillBuffer() int {
	buf, err := l.infile.ReadString('\n')
	if err != nil {
		return -1
	}
	l.lineNumber++
	l.bufp = 0
	return len(buf)
}

//	private boolean hasMoreInput() {
//	    while(endOfBuffer())
//	        if(fillBuffer() == -1)
//	            return false;
//	    return true;
//	}
func (l *Lexer) hasMoreInput() bool {
	for l.endOfBuffer() {
		if l.fillBuffer() == -1 {
			return false
		}
	}
	return true
}

//	private void skipWhiteSpace() {
//	    while(Character.isWhitespace(currentChar())) {
//	        bufp++;
//	        while(endOfBuffer())
//	            if(fillBuffer() == -1)
//	                return;
//	    }
//	}
func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.currentChar()) {
		l.bufp++
		for l.endOfBuffer() {
			if l.fillBuffer() == -1 {
				return
			}
		}
	}
}

//	private void skipComment() {
//	    if(currentChar() == '"') {
//	        do {
//	            bufp++;
//	            while(endOfBuffer())
//	                if(fillBuffer() == -1)
//	                    return;
//	        } while(currentChar() != '"');
//	        bufp++;
//	    }
//	}
func (l *Lexer) skipComment() {
	if l.currentChar() == '"' {
		for l.currentChar() != '"' {
			l.bufp++
			for l.endOfBuffer() {
				if l.fillBuffer() == -1 {
					return
				}
			}
		}
		l.bufp++
	}
}

//	private char currentChar() {
//	    return bufchar(bufp);
//	}
func (l *Lexer) currentChar() rune {
	return l.bufchar(l.bufp)
}

//	private boolean endOfBuffer() {
//	    return bufp >= buf.length();
//	}
func (l *Lexer) endOfBuffer() bool {
	return l.bufp >= len(l.buf)
}

//	private boolean isOperator(char c) {
//	    return c == '~' || c == '&' || c == '|' || c == '*' || c == '/' ||
//	        c == '\\' || c == '+' || c == '=' || c == '>' || c == '<' ||
//	        c == ',' || c == '@' || c == '%';
//	}
func (l *Lexer) isOperator(c rune) bool {
	return c == '~' || c == '&' || c == '|' || c == '*' || c == '/' ||
		c == '\\' || c == '+' || c == '=' || c == '>' || c == '<' ||
		c == ',' || c == '@' || c == '%'
}

//	private void match(Token s) {
//	    sym = s;
//	    symc = currentChar();
//	    text = new StringBuffer("" + symc);
//	    bufp++;
//	}
func (l *Lexer) match(s Token) {
	l.sym = s
	l.symc = l.currentChar()
	l.text = string(l.symc)
	l.bufp++
}

//	private char bufchar(int p) {
//	    return p >= buf.length() ? '\0' : buf.charAt(p);
//	}
func (l *Lexer) bufchar(p int) rune {
	if p >= len(l.buf) {
		return 0
	}
	return rune(l.buf[p])
}

// function startsWith looks in buf at bufp and returns true if string equals SEPARATOR
// l.buf.startsWith(SEPARATOR, l.bufp)
func (l *Lexer) startsWith(s string) bool {
	return l.buf[l.bufp:len(s)] == s
}

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
	Signature   *Symbol
	Arguments   []string
	Primitive   bool
	Locals      []string
	Literals    []*Object
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
func (mg *MethodGenerator) AssemblePrimitive() *Invokable {
	return NewEmptyPrimitive(mg.Signature.GetString())
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
func (mg *MethodGenerator) Assemble() *Method {
	// create a method instance with the given number of bytecodes and literals
	numLiterals := len(mg.Literals)
	meth := NewMethod(mg.signature, len(mg.Bytecode), numLiterals)
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
		case Halt:
			i++
		case Dup:
			depth++
			i++
		case PushLocal, PushArgument:
			depth++
			i += 3
		case PushField, PushBlock, PushConstant, PushGlobal:
			depth++
			i += 2
		case Pop:
			depth--
			i++
		case PopLocal, PopArgument:
			depth--
			i += 3
		case PopField:
			depth--
			i += 2
		case Send, SuperSend:
			// these are special: they need to look at the number of
			// arguments (extractable from the signature)
			sig := mg.Literals[mg.Bytecode[i+1]].(*Symbol)

			depth -= sig.GetNumberOfSignatureArguments()

			depth++ // return value
			i += 2
		case ReturnLocal, ReturnNonLocal:
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
func (mg *MethodGenerator) SetSignature(sig *Symbol) {
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
func (mg *MethodGenerator) AddLiteralIfAbsent(lit *Object) bool {
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
func (mg *MethodGenerator) AddLiteral(lit *Object) {
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
func (mg *MethodGenerator) FindLiteralIndex(lit *Object) byte {
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
func (mg *MethodGenerator) GetSignature() *Symbol {
	return mg.Signature
}

// private Lexer lexer;
// private Symbol sym;
// private String text;
// private Symbol nextSym;
// private BytecodeGenerator bcGen;
type Parser struct {
	lexer               *Lexer
	sym                 Token
	text                string
	nextSym             Token
	bcGen               *BytecodeGenerator
	singleOpSyms        []Token
	binaryOpSyms        []Token
	keywordSelectorSyms []Token
}

// private static final List<Symbol> singleOpSyms = new ArrayList<Symbol>();

// private static final List<Symbol> binaryOpSyms = new ArrayList<Symbol>();
// private static final List<Symbol> keywordSelectorSyms = new ArrayList<Symbol>();

//	static {
//		for(Symbol s : new Symbol[] { Not, And, Or, Star, Div, Mod, Plus, Equal, More, Less, Comma, At, Per, NONE })
//			singleOpSyms.add(s);
//		for(Symbol s : new Symbol[] {
//			Or, Comma, Minus, Equal, Not, And, Or, Star, Div, Mod, Plus, Equal, More, Less, Comma, At, Per, NONE
//		})
//			binaryOpSyms.add(s);
//		for(Symbol s : new Symbol[] { Keyword, KeywordSequence })
//			keywordSelectorSyms.add(s);
//	}
func NewParser(reader io.Reader) *Parser {
	np := &Parser{
		lexer:               NewLexer(reader),
		sym:                 NONE,
		text:                "",
		nextSym:             NONE,
		bcGen:               NewBytecodeGenerator(),
		singleOpSyms:        []Token{Not, And, Or, Star, Div, Mod, Plus, Equal, More, Less, Comma, At, Per, NONE},
		binaryOpSyms:        []Token{Or, Comma, Minus, Equal, Not, And, Or, Star, Div, Mod, Plus, Equal, More, Less, Comma, At, Per, NONE},
		keywordSelectorSyms: []Token{Keyword, KeywordSequence},
	}
	return np
}

//	public Parser(Reader reader) {
//		sym = NONE;
//		lexer = new Lexer(reader);
//		bcGen = new BytecodeGenerator();
//		nextSym = NONE;
//		GETSYM();
//	}
func (p *Parser) Parse() {
	p.classdef()
}

// public void classdef(ClassGenerationContext cgenc) {
// 	cgenc.setName(Universe.symbolFor(text));
// 	expect(Identifier);
// 	expect(Equal);

// 	if(sym == Identifier) {
// 		cgenc.setSuperName(Universe.symbolFor(text));
// 		accept(Identifier);
// 	} else
// 		cgenc.setSuperName(Universe.symbolFor("Object"));

// 	expect(NewTerm);
// 	instanceFields(cgenc);
// 	while(sym == Identifier || sym == Keyword || sym == OperatorSequence || symIn(binaryOpSyms)) {
// 		MethodGenerationContext mgenc = new MethodGenerationContext();
// 		mgenc.setHolder(cgenc);
// 		mgenc.addArgument("self");

// 		method(mgenc);

// 		if(mgenc.isPrimitive())
// 			cgenc.addInstanceMethod(mgenc.assemblePrimitive());
// 		else
// 			cgenc.addInstanceMethod(mgenc.assemble());
// 	}

// 	if(accept(Separator)) {
// 		cgenc.setClassSide(true);
// 		classFields(cgenc);
// 		while(sym == Identifier || sym == Keyword || sym == OperatorSequence || symIn(binaryOpSyms)) {
// 			MethodGenerationContext mgenc = new MethodGenerationContext();
// 			mgenc.setHolder(cgenc);
// 			mgenc.addArgument("self");

// 			method(mgenc);

// 			if(mgenc.isPrimitive())
// 				cgenc.addClassMethod(mgenc.assemblePrimitive());
// 			else
// 				cgenc.addClassMethod(mgenc.assemble());
// 		}
// 	}
// 	expect(EndTerm);
// }

// private boolean symIn(List<Symbol> ss) {
// 	return ss.contains(sym);
// }

// private boolean accept(Symbol s) {
// 	if(sym == s) {
// 		GETSYM();
// 		return true;
// 	}
// 	return false;
// }

// private boolean acceptOneOf(List<Symbol> ss) {
// 	if(symIn(ss)) {
// 		GETSYM();
// 		return true;
// 	}
// 	return false;
// }

// private boolean expect(Symbol s) {
// 	if(accept(s))
// 		return true;
// 	StringBuffer err = new StringBuffer("Error: unexpected symbol in line " + lexer.getCurrentLineNumber() +
// 		". Expected " + s.toString() + ", but found " + sym.toString());
// 	if(printableSymbol())
// 		err.append(" (" + text + ")");
// 	err.append(": " + lexer.getRawBuffer());
// 	throw new IllegalStateException(err.toString());
// }

// private boolean expectOneOf(List<Symbol> ss) {
// 	if(acceptOneOf(ss))
// 		return true;
// 	StringBuffer err = new StringBuffer("Error: unexpected symbol in line " + lexer.getCurrentLineNumber() +
// 		". Expected one of ");
// 	for(Symbol s : ss)
// 		err.append(s.toString() + ", ");
// 	err.append("but found " + sym.toString());
// 	if(printableSymbol())
// 		err.append(" (" + text + ")");
// 	err.append(": " + lexer.getRawBuffer());
// 	throw new IllegalStateException(err.toString());
// }

// private void instanceFields(ClassGenerationContext cgenc) {
// 	if(accept(Or)) {
// 		while(sym == Identifier) {
// 			String var = variable();
// 			cgenc.addInstanceField(Universe.symbolFor(var));
// 		}
// 		expect(Or);
// 	}
// }

// private void classFields(ClassGenerationContext cgenc) {
// 	if(accept(Or)) {
// 		while(sym == Identifier) {
// 			String var = variable();
// 			cgenc.addClassField(Universe.symbolFor(var));
// 		}
// 		expect(Or);
// 	}
// }

// private void method(MethodGenerationContext mgenc) {
// 	pattern(mgenc);
// 	expect(Equal);
// 	if(sym == Primitive) {
// 		mgenc.setPrimitive(true);
// 		primitiveBlock();
// 	} else
// 		methodBlock(mgenc);
// }

// private void primitiveBlock() {
// 	expect(Primitive);
// }

// private void pattern(MethodGenerationContext mgenc) {
// 	switch(sym) {
// 	case Identifier:
// 		unaryPattern(mgenc);
// 		break;
// 	case Keyword:
// 		keywordPattern(mgenc);
// 		break;
// 	default:
// 		binaryPattern(mgenc);
// 		break;
// 	}
// }

// private void unaryPattern(MethodGenerationContext mgenc) {
// 	mgenc.setSignature(unarySelector());
// }

// private void binaryPattern(MethodGenerationContext mgenc) {
// 	mgenc.setSignature(binarySelector());
// 	mgenc.addArgumentIfAbsent(argument());
// }

// private void keywordPattern(MethodGenerationContext mgenc) {
// 	StringBuffer kw = new StringBuffer();
// 	do {
// 		kw.append(keyword());
// 		mgenc.addArgumentIfAbsent(argument());
// 	} while(sym == Keyword);

// 	mgenc.setSignature(Universe.symbolFor(kw.toString()));
// }

// private void methodBlock(MethodGenerationContext mgenc) {
// 	expect(NewTerm);
// 	blockContents(mgenc);
// 	// if no return has been generated so far, we can be sure there was no .
// 	// terminating the last expression, so the last expression's value must be
// 	// popped off the stack and a ^self be generated
// 	if(!mgenc.isFinished()) {
// 		bcGen.emitPOP(mgenc);
// 		bcGen.emitPUSHARGUMENT(mgenc, (byte) 0, (byte) 0);
// 		bcGen.emitRETURNLOCAL(mgenc);
// 		mgenc.setFinished();
// 	}

// 	expect(EndTerm);
// }

// private som.vmobjects.Symbol unarySelector() {
// 	return Universe.symbolFor(identifier());
// }

// private som.vmobjects.Symbol binarySelector() {
// 	String s = new String(text);

// 	if(accept(Or))
// 		;
// 	else if(accept(Comma))
// 		;
// 	else if(accept(Minus))
// 		;
// 	else if(accept(Equal))
// 		;
// 	else if(acceptOneOf(singleOpSyms))
// 		;
// 	else if(accept(OperatorSequence))
// 		;
// 	else
// 		expect(NONE);

// 	return Universe.symbolFor(s);
// }

// private String identifier() {
// 	String s = new String(text);
// 	if(accept(Primitive))
// 		; // text is set
// 	else
// 		expect(Identifier);

// 	return s;
// }

// private String keyword() {
// 	String s = new String(text);
// 	expect(Keyword);

// 	return s;
// }

// private String argument() {
// 	return variable();
// }

// private void blockContents(MethodGenerationContext mgenc) {
// 	if(accept(Or)) {
// 		locals(mgenc);
// 		expect(Or);
// 	}
// 	blockBody(mgenc, false);
// }

// private void locals(MethodGenerationContext mgenc) {
// 	while(sym == Identifier)
// 		mgenc.addLocalIfAbsent(variable());
// }

// private void blockBody(MethodGenerationContext mgenc, boolean seenPeriod) {
// 	if(accept(Exit))
// 		result(mgenc);
// 	else if(sym == EndBlock) {
// 		if(seenPeriod) {
// 			// a POP has been generated which must be elided (blocks always
// 			// return the value of the last expression, regardless of whether it
// 			// was terminated with a . or not)
// 			mgenc.removeLastBytecode();
// 		}
// 		bcGen.emitRETURNLOCAL(mgenc);
// 		mgenc.setFinished();
// 	} else if(sym == EndTerm) {
// 		// it does not matter whether a period has been seen, as the end of the
// 		// method has been found (EndTerm) - so it is safe to emit a "return
// 		// self"
// 		bcGen.emitPUSHARGUMENT(mgenc, (byte) 0, (byte) 0);
// 		bcGen.emitRETURNLOCAL(mgenc);
// 		mgenc.setFinished();
// 	} else {
// 		expression(mgenc);
// 		if(accept(Period)) {
// 			bcGen.emitPOP(mgenc);
// 			blockBody(mgenc, true);
// 		}
// 	}
// }

// private void result(MethodGenerationContext mgenc) {
// 	expression(mgenc);

// 	if(mgenc.isBlockMethod())
// 		bcGen.emitRETURNNONLOCAL(mgenc);
// 	else
// 		bcGen.emitRETURNLOCAL(mgenc);

// 	mgenc.setFinished(true);
// 	accept(Period);
// }

// private void expression(MethodGenerationContext mgenc) {
// 	PEEK();
// 	if(nextSym == Assign)
// 		assignation(mgenc);
// 	else
// 		evaluation(mgenc);
// }

// private void assignation(MethodGenerationContext mgenc) {
// 	List<String> l = new ArrayList<String>();

// 	assignments(mgenc, l);
// 	evaluation(mgenc);

// 	for(int i = 1; i <= l.size(); i++)
// 		bcGen.emitDUP(mgenc);
// 	for(String s : l)
// 		genPopVariable(mgenc, s);
// }

// private void assignments(MethodGenerationContext mgenc, List<String> l) {
// 	if(sym == Identifier) {
// 		l.add(assignment(mgenc));
// 		PEEK();
// 		if(nextSym == Assign)
// 			assignments(mgenc, l);
// 	}
// }

// private String assignment(MethodGenerationContext mgenc) {
// 	String v = variable();
// 	som.vmobjects.Symbol var = Universe.symbolFor(v);
// 	mgenc.addLiteralIfAbsent(var);

// 	expect(Assign);

// 	return v;
// }

// private void evaluation(MethodGenerationContext mgenc) {
// 	// single: superSend
// 	Single<Boolean> si = new Single<Boolean>(false);

// 	primary(mgenc, si);
// 	if(sym == Identifier || sym == Keyword || sym == OperatorSequence || symIn(binaryOpSyms)) {
// 		messages(mgenc, si);
// 	}
// }

// private void primary(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	superSend.set(false);
// 	switch(sym) {
// 		case Identifier: {
// 			String v = variable();
// 			if(v.equals("super")) {
// 				superSend.set(true);
// 				// sends to super push self as the receiver
// 				v = "self";
// 			}

// 			genPushVariable(mgenc, v);
// 			break;
// 		}
// 		case NewTerm:
// 			nestedTerm(mgenc);
// 			break;
// 		case NewBlock: {
// 			MethodGenerationContext bgenc = new MethodGenerationContext();
// 			bgenc.setIsBlockMethod(true);
// 			bgenc.setHolder(mgenc.getHolder());
// 			bgenc.setOuter(mgenc);

// 			nestedBlock(bgenc);

// 			som.vmobjects.Method blockMethod = bgenc.assemble();
// 			mgenc.addLiteral(blockMethod);
// 			bcGen.emitPUSHBLOCK(mgenc, blockMethod);
// 			break;
// 		}
// 		default:
// 			literal(mgenc);
// 			break;
// 	}
// }

// private String variable() {
// 	return identifier();
// }

// private void messages(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	if(sym == Identifier) {
// 		do {
// 			// only the first message in a sequence can be a super send
// 			unaryMessage(mgenc, superSend);
// 			superSend.set(false);
// 		} while(sym == Identifier);

// 		while(sym == OperatorSequence || symIn(binaryOpSyms)) {
// 			binaryMessage(mgenc, new Single<Boolean>(false));
// 		}

// 		if(sym == Keyword) {
// 			keywordMessage(mgenc, new Single<Boolean>(false));
// 		}
// 	} else if(sym == OperatorSequence || symIn(binaryOpSyms)) {
// 		do {
// 			// only the first message in a sequence can be a super send
// 			binaryMessage(mgenc, superSend);
// 			superSend.set(false);
// 		} while(sym == OperatorSequence || symIn(binaryOpSyms));

// 		if(sym == Keyword) {
// 			keywordMessage(mgenc, new Single<Boolean>(false));
// 		}
// 	} else
// 		keywordMessage(mgenc, superSend);
// }

// private void unaryMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	som.vmobjects.Symbol msg = unarySelector();
// 	mgenc.addLiteralIfAbsent(msg);

// 	if(superSend.get())
// 		bcGen.emitSUPERSEND(mgenc, msg);
// 	else
// 		bcGen.emitSEND(mgenc, msg);
// }

// private void binaryMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	som.vmobjects.Symbol msg = binarySelector();
// 	mgenc.addLiteralIfAbsent(msg);

// 	binaryOperand(mgenc, new Single<Boolean>(false));

// 	if(superSend.get())
// 		bcGen.emitSUPERSEND(mgenc, msg);
// 	else
// 		bcGen.emitSEND(mgenc, msg);
// }

// private void binaryOperand(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	primary(mgenc, superSend);

// 	while(sym == Identifier)
// 		unaryMessage(mgenc, superSend);
// }

// private void keywordMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
// 	StringBuffer kw = new StringBuffer();
// 	do {
// 		kw.append(keyword());
// 		formula(mgenc);
// 	} while(sym == Keyword);

// 	som.vmobjects.Symbol msg = Universe.symbolFor(kw.toString());

// 	mgenc.addLiteralIfAbsent(msg);

// 	if(superSend.get())
// 		bcGen.emitSUPERSEND(mgenc, msg);
// 	else
// 		bcGen.emitSEND(mgenc, msg);
// }

// private void formula(MethodGenerationContext mgenc) {
// 	Single<Boolean> superSend = new Single<Boolean>(false);
// 	binaryOperand(mgenc, superSend);

// 	// only the first message in a sequence can be a super send
// 	if(sym == OperatorSequence || symIn(binaryOpSyms))
// 		binaryMessage(mgenc, superSend);
// 	while(sym == OperatorSequence || symIn(binaryOpSyms))
// 		binaryMessage(mgenc, new Single<Boolean>(false));
// }

// private void nestedTerm(MethodGenerationContext mgenc) {
// 	expect(NewTerm);
// 	expression(mgenc);
// 	expect(EndTerm);
// }

// private void literal(MethodGenerationContext mgenc) {
// 	switch(sym) {
// 	case Pound:
// 		literalSymbol(mgenc);
// 		break;
// 	case STString:
// 		literalString(mgenc);
// 		break;
// 	default:
// 		literalNumber(mgenc);
// 		break;
// 	}
// }

// private void literalNumber(MethodGenerationContext mgenc) {
// 	int val;
// 	if(sym == Minus)
// 		val = negativeDecimal();
// 	else
// 		val = literalDecimal();

// 	som.vmobjects.Integer lit = Universe.newInteger(val);
// 	mgenc.addLiteralIfAbsent(lit);
// 	bcGen.emitPUSHCONSTANT(mgenc, lit);
// }

// private int literalDecimal() {
// 	return literalInteger();
// }

// private int negativeDecimal() {
// 	expect(Minus);
// 	return -literalInteger();
// }

// private int literalInteger() {
// 	int i = java.lang.Integer.parseInt(text);
// 	expect(Integer);
// 	return i;
// }

// private void literalSymbol(MethodGenerationContext mgenc) {
// 	som.vmobjects.Symbol symb;
// 	expect(Pound);
// 	if(sym == STString) {
// 		String s = string();
// 		symb = Universe.symbolFor(s);
// 	} else
// 		symb = selector();
// 	mgenc.addLiteralIfAbsent(symb);
// 	bcGen.emitPUSHCONSTANT(mgenc, symb);
// }

// private void literalString(MethodGenerationContext mgenc) {
// 	String s = string();

// 	som.vmobjects.String str = Universe.newString(s);
// 	mgenc.addLiteralIfAbsent(str);

// 	bcGen.emitPUSHCONSTANT(mgenc, str);
// }

// private som.vmobjects.Symbol selector() {
// 	if(sym == OperatorSequence || symIn(singleOpSyms))
// 		return binarySelector();
// 	else if(sym == Keyword || sym == KeywordSequence)
// 		return keywordSelector();
// 	else
// 		return unarySelector();
// }

// private som.vmobjects.Symbol keywordSelector() {
// 	String s = new String(text);
// 	expectOneOf(keywordSelectorSyms);
// 	som.vmobjects.Symbol symb = Universe.symbolFor(s);
// 	return symb;
// }

// private String string() {
// 	String s = new String(text);
// 	expect(STString);
// 	return s;
// }

// private void nestedBlock(MethodGenerationContext mgenc) {
// 	mgenc.addArgumentIfAbsent("$block self");

// 	expect(NewBlock);
// 	if(sym == Colon)
// 		blockPattern(mgenc);

// 	// generate Block signature
// 	String blockSig = "$block method";
// 	int argSize = mgenc.getNumberOfArguments();
// 	for(int i = 1; i < argSize; i++)
// 		blockSig += ":";

// 	mgenc.setSignature(Universe.symbolFor(blockSig));

// 	blockContents(mgenc);

// 	// if no return has been generated, we can be sure that the last expression
// 	// in the block was not terminated by ., and can generate a return
// 	if(!mgenc.isFinished()) {
// 		bcGen.emitRETURNLOCAL(mgenc);
// 		mgenc.setFinished(true);
// 	}

// 	expect(EndBlock);
// }

// private void blockPattern(MethodGenerationContext mgenc) {
// 	blockArguments(mgenc);
// 	expect(Or);
// }

// private void blockArguments(MethodGenerationContext mgenc) {
// 	do {
// 		expect(Colon);
// 		mgenc.addArgumentIfAbsent(argument());
// 	} while(sym == Colon);
// }

// private void genPushVariable(MethodGenerationContext mgenc, String var) {
// 	// The purpose of this function is to find out whether the variable to be
// 	// pushed on the stack is a local variable, argument, or object field. This
// 	// is done by examining all available lexical contexts, starting with the
// 	// innermost (i.e., the one represented by mgenc).

// 	// triplet: index, context, isArgument
// 	Triplet<Byte,Byte,Boolean> tri = new Triplet<Byte,Byte,Boolean>((byte) 0, (byte) 0, false);

// 	if(mgenc.findVar(var, tri)) {
// 		if(tri.getZ())
// 			bcGen.emitPUSHARGUMENT(mgenc, tri.getX(), tri.getY());
// 		else
// 			bcGen.emitPUSHLOCAL(mgenc, tri.getX(), tri.getY());
// 	} else if(mgenc.findField(var)) {
// 		som.vmobjects.Symbol fieldName = Universe.symbolFor(var);
// 		mgenc.addLiteralIfAbsent(fieldName);
// 		bcGen.emitPUSHFIELD(mgenc, fieldName);
// 	} else {
// 		som.vmobjects.Symbol global = Universe.symbolFor(var);
// 		mgenc.addLiteralIfAbsent(global);
// 		bcGen.emitPUSHGLOBAL(mgenc, global);
// 	}
// }

// private void genPopVariable(MethodGenerationContext mgenc, String var) {
// 	// The purpose of this function is to find out whether the variable to be
// 	// popped off the stack is a local variable, argument, or object field. This
// 	// is done by examining all available lexical contexts, starting with the
// 	// innermost (i.e., the one represented by mgenc).

// 	// triplet: index, context, isArgument
// 	Triplet<Byte,Byte,Boolean> tri = new Triplet<Byte,Byte,Boolean>((byte) 0, (byte) 0, false);

// 	if(mgenc.findVar(var, tri)) {
// 		if(tri.getZ())
// 			bcGen.emitPOPARGUMENT(mgenc, tri.getX(), tri.getY());
// 		else
// 			bcGen.emitPOPLOCAL(mgenc, tri.getX(), tri.getY());
// 	} else
// 		bcGen.emitPOPFIELD(mgenc, Universe.symbolFor(var));
// }

// private void GETSYM() {
// 	sym = lexer.getSym();
// 	text = lexer.getText();
// }

// private void PEEK() {
// 	nextSym = lexer.peek();
// }

// private boolean printableSymbol() {
// 	return sym == Integer || sym.compareTo(STString) >= 0;
// }

// }

type Shell struct{}

func (s *Shell) Start()                       {}
func (s *Shell) SetBootstrapMethod(m *Method) {}
func (s *Shell) GetBootstrapMethod() *Method  { return nil }


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
