package smog

import (
	"fmt"
	"io"
	"strconv"
)

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
		singleOpSyms:        []Token{tNot, tAnd, tOr, tStar, tDiv, tMod, tPlus, tEqual, tMore, tLess, tComma, tAt, tPer, NONE},
		binaryOpSyms:        []Token{tOr, tComma, tMinus, tEqual, tNot, tAnd, tOr, tStar, tDiv, tMod, tPlus, tEqual, tMore, tLess, tComma, tAt, tPer, NONE},
		keywordSelectorSyms: []Token{tKeyword, tKeywordSequence},
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

func (p *Parser) classdef(cgenc *ClassGenerator) {
	cgenc.SetName(GetUniverse().SymbolFor(p.text))
	p.expect(tIdentifier)
	p.expect(tEqual)
	if p.sym == tIdentifier {
		cgenc.SetSuperName(GetUniverse().SymbolFor(p.text))
		p.accept(tIdentifier)
	} else {
		cgenc.SetSuperName(GetUniverse().SymbolFor("Object"))
	}
	p.expect(tNewTerm)
	cgenc.instanceFields()
	for p.sym == tIdentifier || p.sym == tKeyword || p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
		mgenc := NewMethodGenerator()
		mgenc.SetHolder(cgenc)
		mgenc.AddArgument("self")

		mgenc.Method()

		if mgenc.IsPrimitive() {
			cgenc.AddInstanceMethod(mgenc.AssemblePrimitive())
		} else {
			cgenc.AddInstanceMethod(mgenc.Assemble())
		}
		if p.accept(tSeparator) {
			cgenc.SetClassSide(true)
			cgenc.classFields()
			for p.sym == tIdentifier || p.sym == tKeyword || p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
				mgenc := NewMethodGenerator()
				mgenc.SetHolder(cgenc)
				mgenc.AddArgument("self")
				mgenc.Method()
				if mgenc.IsPrimitive() {
					cgenc.AddClassMethod(mgenc.AssemblePrimitive())
				} else {
					cgenc.AddClassMethod(mgenc.Assemble())
				}
			}
		}
		p.expect(tEndTerm)
	}
}

//	private boolean symIn(List<Symbol> ss) {
//		return ss.contains(sym);
//	}
func (p *Parser) symIn(ss []Token) bool {
	for _, s := range ss {
		if s == p.sym {
			return true
		}
	}
	return false
}

//	private boolean accept(Symbol s) {
//		if(sym == s) {
//			GETSYM();
//			return true;
//		}
//		return false;
//	}
func (p *Parser) accept(s Token) bool {
	if p.sym == s {
		p.lexer.GetSym()
		return true
	}
	return false
}

//	private boolean acceptOneOf(List<Symbol> ss) {
//		if(symIn(ss)) {
//			GETSYM();
//			return true;
//		}
//		return false;
//	}
func (p *Parser) acceptOneOf(ss []Token) bool {
	if p.symIn(ss) {
		p.GetSym()
		return true
	}
	return false
}

//	private boolean expect(Symbol s) {
//		if(accept(s))
//			return true;
//		StringBuffer err = new StringBuffer("Error: unexpected symbol in line " + lexer.getCurrentLineNumber() +
//			". Expected " + s.toString() + ", but found " + sym.toString());
//		if(printableSymbol())
//			err.append(" (" + text + ")");
//		err.append(": " + lexer.getRawBuffer());
//		throw new IllegalStateException(err.toString());
//	}
func (p *Parser) expect(s Token) {
	if p.accept(s) {
		return
	}
	err := fmt.Sprintf("Error: unexpected symbol in line %d. Expected %s, but found %s", p.lexer.getCurrentLineNumber(), s, p.sym)
	if p.printableSymbol() {
		err += fmt.Sprintf(" (%s)", p.text)
	}
	err += fmt.Sprintf(": %s", p.lexer.getRawBuffer())
	panic(err)
}

//	private boolean expectOneOf(List<Symbol> ss) {
//		if(acceptOneOf(ss))
//			return true;
//		StringBuffer err = new StringBuffer("Error: unexpected symbol in line " + lexer.getCurrentLineNumber() +
//			". Expected one of ");
//		for(Symbol s : ss)
//			err.append(s.toString() + ", ");
//		err.append("but found " + sym.toString());
//		if(printableSymbol())
//			err.append(" (" + text + ")");
//		err.append(": " + lexer.getRawBuffer());
//		throw new IllegalStateException(err.toString());
//	}
func (p *Parser) expectOneOf(ss []Token) {
	if p.acceptOneOf(ss) {
		return
	}
	err := fmt.Sprintf("Error: unexpected symbol in line %d. Expected one of ", p.lexer.getCurrentLineNumber())
	for _, s := range ss {
		err += fmt.Sprintf("%s, ", s)
	}
	err += fmt.Sprintf("but found %s", p.sym)
	if p.printableSymbol() {
		err += fmt.Sprintf(" (%s)", p.text)
	}
	err += fmt.Sprintf(": %s", p.lexer.getRawBuffer())
	panic(err)
}

//	private void instanceFields(ClassGenerationContext cgenc) {
//		if(accept(Or)) {
//			while(sym == Identifier) {
//				String var = variable();
//				cgenc.addInstanceField(Universe.symbolFor(var));
//			}
//			expect(Or);
//		}
//	}
func (p *Parser) instanceFields(cgenc *ClassGenerator) {
	if p.accept(tOr) {
		for p.sym == tIdentifier {
			variable := p.variable()
			cgenc.AddInstanceField(GetUniverse().SymbolFor(variable))
		}
		p.expect(tOr)
	}
}

//	private void classFields(ClassGenerationContext cgenc) {
//		if(accept(Or)) {
//			while(sym == Identifier) {
//				String var = variable();
//				cgenc.addClassField(Universe.symbolFor(var));
//			}
//			expect(Or);
//		}
//	}
func (p *Parser) classFields(cgenc *ClassGenerator) {
	if p.accept(tOr) {
		for p.sym == tIdentifier {
			variable := p.variable()
			cgenc.AddClassField(GetUniverse().SymbolFor(variable))
		}
		p.expect(tOr)
	}
}

//	private void method(MethodGenerationContext mgenc) {
//		pattern(mgenc);
//		expect(Equal);
//		if(sym == Primitive) {
//			mgenc.setPrimitive(true);
//			primitiveBlock();
//		} else
//			methodBlock(mgenc);
//	}
func (p *Parser) method(mgenc *MethodGenerator) {
	p.pattern(mgenc)
	p.expect(tEqual)
	if p.sym == tPrimitive {
		mgenc.SetPrimitive(true)
		p.primitiveBlock()
	} else {
		p.methodBlock(mgenc)
	}
}

//	private void primitiveBlock() {
//		expect(Primitive);
//	}
func (p *Parser) primitiveBlock() {
	p.expect(tPrimitive)
}

//	private void pattern(MethodGenerationContext mgenc) {
//		switch(sym) {
//		case Identifier:
//			unaryPattern(mgenc);
//			break;
//		case Keyword:
//			keywordPattern(mgenc);
//			break;
//		default:
//			binaryPattern(mgenc);
//			break;
//		}
//	}
func (p *Parser) pattern(mgenc *MethodGenerator) {
	switch p.sym {
	case tIdentifier:
		p.unaryPattern(mgenc)
	case tKeyword:
		p.keywordPattern(mgenc)
	default:
		p.binaryPattern(mgenc)
	}
}

//	private void unaryPattern(MethodGenerationContext mgenc) {
//		mgenc.setSignature(unarySelector());
//	}
func (p *Parser) unaryPattern(mgenc *MethodGenerator) {
	mgenc.SetSignature(p.unarySelector())
}

//	private void binaryPattern(MethodGenerationContext mgenc) {
//		mgenc.setSignature(binarySelector());
//		mgenc.addArgumentIfAbsent(argument());
//	}
func (p *Parser) binaryPattern(mgenc *MethodGenerator) {
	mgenc.SetSignature(p.binarySelector())
	mgenc.AddArgumentIfAbsent(p.argument())
}

// private void keywordPattern(MethodGenerationContext mgenc) {
// 	StringBuffer kw = new StringBuffer();
// 	do {
// 		kw.append(keyword());
// 		mgenc.addArgumentIfAbsent(argument());
// 	} while(sym == Keyword);

//		mgenc.setSignature(Universe.symbolFor(kw.toString()));
//	}
func (p *Parser) keywordPattern(mgenc *MethodGenerator) {
	kw := ""
	for p.sym == tKeyword {
		kw += p.keyword()
		mgenc.AddArgumentIfAbsent(p.argument())
	}

	mgenc.SetSignature(GetUniverse().SymbolFor(kw))
}

//	private void methodBlock(MethodGenerationContext mgenc) {
//		expect(NewTerm);
//		blockContents(mgenc);
//		// if no return has been generated so far, we can be sure there was no .
//		// terminating the last expression, so the last expression's value must be
//		// popped off the stack and a ^self be generated
//		if(!mgenc.isFinished()) {
//			bcGen.emitPOP(mgenc);
//			bcGen.emitPUSHARGUMENT(mgenc, (byte) 0, (byte) 0);
//			bcGen.emitRETURNLOCAL(mgenc);
//			mgenc.setFinished();
//		}
func (p *Parser) methodBlock(mgenc *MethodGenerator) {
	p.expect(tNewTerm)
	p.blockContents(mgenc)
	// if no return has been generated so far, we can be sure there was no .
	// terminating the last expression, so the last expression's value must be
	// popped off the stack and a ^self be generated
	if !mgenc.IsFinished() {
		p.bcGen.emitPOP(mgenc)
		p.bcGen.emitPUSHARGUMENT(mgenc, 0, 0)
		p.bcGen.emitRETURNLOCAL(mgenc)
		mgenc.SetFinished()
	}

	expect(EndTerm)
}

//	private som.vmobjects.Symbol unarySelector() {
//		return Universe.symbolFor(identifier());
//	}
func (p *Parser) unarySelector() *Symbol {
	return GetUniverse().SymbolFor(p.identifier())
}

//	private som.vmobjects.Symbol binarySelector() {
//		String s = new String(text);
//		if(accept(Or))
//			;
//		else if(accept(Comma))
//			;
//		else if(accept(Minus))
//			;
//		else if(accept(Equal))
//			;
//		else if(acceptOneOf(singleOpSyms))
//			;
//		else if(accept(OperatorSequence))
//			;
//		else
//			expect(NONE);
//		return Universe.symbolFor(s);
//	}
func (p *Parser) binarySelector() *Symbol {
	s := p.text
	if p.accept(tOr) {
	} else if p.accept(tComma) {
	} else if p.accept(tMinus) {
	} else if p.accept(tEqual) {
	} else if p.acceptOneOf(p.singleOpSyms) {
	} else if p.accept(tOperatorSequence) {
	} else {
		p.expect(NONE)
	}
	return GetUniverse().SymbolFor(s)
}

//	private String identifier() {
//		String s = new String(text);
//		if(accept(Primitive))
//			; // text is set
//		else
//			expect(Identifier);
//		return s;
//	}
func (p *Parser) identifier() string {
	s := p.text
	if p.accept(tPrimitive) {
	}
	p.expect(tIdentifier)
	return s
}

//	private String keyword() {
//		String s = new String(text);
//		expect(Keyword);
//		return s;
//	}
func (p *Parser) keyword() string {
	s := p.text
	p.expect(tKeyword)
	return s
}

//	private String argument() {
//		return variable();
//	}
func (p *Parser) argument() string {
	return p.variable()
}

//	private void blockContents(MethodGenerationContext mgenc) {
//		if(accept(Or)) {
//			locals(mgenc);
//			expect(Or);
//		}
//		blockBody(mgenc, false);
//	}
func (p *Parser) blockContents(mgenc *MethodGenerator) {
	if p.accept(tOr) {
		p.locals(mgenc)
		p.expect(tOr)
	}
	p.blockBody(mgenc, false)
}

//	private void locals(MethodGenerationContext mgenc) {
//		while(sym == Identifier)
//			mgenc.addLocalIfAbsent(variable());
//	}
func (p *Parser) locals(mgenc *MethodGenerator) {
	for p.sym == tIdentifier {
		mgenc.AddLocalIfAbsent(p.variable())
	}
}

//	private void blockBody(MethodGenerationContext mgenc, boolean seenPeriod) {
//		if(accept(Exit))
//			result(mgenc);
//		else if(sym == EndBlock) {
//			if(seenPeriod) {
//				// a POP has been generated which must be elided (blocks always
//				// return the value of the last expression, regardless of whether it
//				// was terminated with a . or not)
//				mgenc.removeLastBytecode();
//			}
//			bcGen.emitRETURNLOCAL(mgenc);
//			mgenc.setFinished();
//		} else if(sym == EndTerm) {
//			// it does not matter whether a period has been seen, as the end of the
//			// method has been found (EndTerm) - so it is safe to emit a "return
//			// self"
//			bcGen.emitPUSHARGUMENT(mgenc, (byte) 0, (byte) 0);
//			bcGen.emitRETURNLOCAL(mgenc);
//			mgenc.setFinished();
//		} else {
//			expression(mgenc);
//			if(accept(Period)) {
//				bcGen.emitPOP(mgenc);
//				blockBody(mgenc, true);
//			}
//		}
//	}
func (p *Parser) blockBody(mgenc *MethodGenerator, seenPeriod bool) {
	if p.accept(tExit) {
		p.result(mgenc)
	} else if p.sym == tEndBlock {
		if seenPeriod {
			// a POP has been generated which must be elided (blocks always
			// return the value of the last expression, regardless of whether it
			// was terminated with a . or not)
			mgenc.RemoveLastBytecode()
		}
		p.bcGen.emitRETURNLOCAL(mgenc)
		mgenc.SetFinished()
	} else if p.sym == tEndTerm {
		// it does not matter whether a period has been seen, as the end of the
		// method has been found (EndTerm) - so it is safe to emit a "return
		// self"
		p.bcGen.emitPUSHARGUMENT(mgenc, 0, 0)
		p.bcGen.emitRETURNLOCAL(mgenc)
		mgenc.SetFinished()
	} else {
		p.expression(mgenc)
		if p.accept(tPeriod) {
			p.bcGen.emitPOP(mgenc)
			p.blockBody(mgenc, true)
		}
	}
}

//	private void result(MethodGenerationContext mgenc) {
//		expression(mgenc);
//		if(mgenc.isBlockMethod())
//			bcGen.emitRETURNNONLOCAL(mgenc);
//		else
//			bcGen.emitRETURNLOCAL(mgenc);
//		mgenc.setFinished(true);
//		accept(Period);
//	}
func (p *Parser) result(mgenc *MethodGenerator) {
	p.expression(mgenc)
	if mgenc.IsBlockMethod() {
		p.bcGen.emitRETURNNONLOCAL(mgenc)
	}
	p.bcGen.emitRETURNLOCAL(mgenc)
	mgenc.SetFinished(true)
	p.accept(tPeriod)
}

//	private void expression(MethodGenerationContext mgenc) {
//		PEEK();
//		if(nextSym == Assign)
//			assignation(mgenc);
//		else
//			evaluation(mgenc);
//	}
func (p *Parser) expression(mgenc *MethodGenerator) {
	p.PEEK()
	if p.nextSym == tAssign {
		p.assignation(mgenc)
	} else {
		p.evaluation(mgenc)
	}
}

//	private void assignation(MethodGenerationContext mgenc) {
//		List<String> l = new ArrayList<String>();
//		assignments(mgenc, l);
//		evaluation(mgenc);
//		for(int i = 1; i <= l.size(); i++)
//			bcGen.emitDUP(mgenc);
//		for(String s : l)
//			genPopVariable(mgenc, s);
//	}
func (p *Parser) assignation(mgenc *MethodGenerator) {
	l := []string{}
	p.assignments(mgenc, l)
	p.evaluation(mgenc)
	for i := 1; i <= len(l); i++ {
		p.bcGen.emitDUP(mgenc)
	}
	for _, s := range l {
		p.genPopVariable(mgenc, s)
	}
}

//	private void assignments(MethodGenerationContext mgenc, List<String> l) {
//		if(sym == Identifier) {
//			l.add(assignment(mgenc));
//			PEEK();
//			if(nextSym == Assign)
//				assignments(mgenc, l);
//		}
//	}
func (p *Parser) assignments(mgenc *MethodGenerator, l []string) {
	if p.sym == tIdentifier {
		l = append(l, p.assignment(mgenc))
		p.PEEK()
		if p.nextSym == tAssign {
			p.assignments(mgenc, l)
		}
	}
}

//	private String assignment(MethodGenerationContext mgenc) {
//		String v = variable();
//		som.vmobjects.Symbol var = GetUniverse().SymbolFor(v);
//		mgenc.addLiteralIfAbsent(var);
//		expect(Assign);
//		return v;
//	}
func (p *Parser) assignment(mgenc *MethodGenerator) string {
	v := p.variable()
	variable := GetUniverse().SymbolFor(v)
	mgenc.AddLiteralIfAbsent(variable)
	p.expect(tAssign)
	return v
}

//	private void evaluation(MethodGenerationContext mgenc) {
//		// single: superSend
//		Single<Boolean> si = new Single<Boolean>(false);
//		primary(mgenc, si);
//		if(sym == Identifier || sym == Keyword || sym == OperatorSequence || symIn(binaryOpSyms)) {
//			messages(mgenc, si);
//		}
//	}
func (p *Parser) evaluation(mgenc *MethodGenerator) {
	// single: superSend
	si := false
	p.primary(mgenc, &si)
	if p.sym == tIdentifier || p.sym == tKeyword || p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
		p.messages(mgenc, &si)
	}
}

//	private void primary(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		superSend.set(false);
//		switch(sym) {
//			case Identifier: {
//				String v = variable();
//				if(v.equals("super")) {
//					superSend.set(true);
//					// sends to super push self as the receiver
//					v = "self";
//				}
//				genPushVariable(mgenc, v);
//				break;
//			}
//			case NewTerm:
//				nestedTerm(mgenc);
//				break;
//			case NewBlock: {
//				MethodGenerationContext bgenc = new MethodGenerationContext();
//				bgenc.setIsBlockMethod(true);
//				bgenc.setHolder(mgenc.getHolder());
//				bgenc.setOuter(mgenc);
//				nestedBlock(bgenc);
//				som.vmobjects.Method blockMethod = bgenc.assemble();
//				mgenc.addLiteral(blockMethod);
//				bcGen.emitPUSHBLOCK(mgenc, blockMethod);
//				break;
//			}
//			default:
//				literal(mgenc);
//				break;
//		}
//	}
func (p *Parser) primary(mgenc *MethodGenerator, superSend *bool) {
	*superSend = false
	switch p.sym {
	case tIdentifier:
		v := p.variable()
		if v == "super" {
			*superSend = true
			// sends to super push self as the receiver
			v = "self"
		}
		p.genPushVariable(mgenc, v)
	case tNewTerm:
		p.nestedTerm(mgenc)
	case tNewBlock:
		bgenc := NewMethodGenerator()
		bgenc.SetIsBlockMethod(true)
		bgenc.SetHolder(mgenc.GetHolder())
		bgenc.SetOuter(mgenc)
		p.nestedBlock(bgenc)
		blockMethod := bgenc.Assemble()
		mgenc.AddLiteral(blockMethod)
		p.bcGen.emitPUSHBLOCK(mgenc, blockMethod)
	default:
		p.literal(mgenc)
	}
}

//	private String variable() {
//		return identifier();
//	}
func (p *Parser) variable() string {
	return p.identifier()
}

//	private void messages(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		if(sym == Identifier) {
//			do {
//				// only the first message in a sequence can be a super send
//				unaryMessage(mgenc, superSend);
//				superSend.set(false);
//			} while(sym == Identifier);
//			while(sym == OperatorSequence || symIn(binaryOpSyms)) {
//				binaryMessage(mgenc, new Single<Boolean>(false));
//			}
//			if(sym == Keyword) {
//				keywordMessage(mgenc, new Single<Boolean>(false));
//			}
//		} else if(sym == OperatorSequence || symIn(binaryOpSyms)) {
//			do {
//				// only the first message in a sequence can be a super send
//				binaryMessage(mgenc, superSend);
//				superSend.set(false);
//			} while(sym == OperatorSequence || symIn(binaryOpSyms));
//			if(sym == Keyword) {
//				keywordMessage(mgenc, new Single<Boolean>(false));
//			}
//		} else
//			keywordMessage(mgenc, superSend);
//	}
func (p *Parser) messages(mgenc *MethodGenerator, superSend *bool) {
	if p.sym == tIdentifier {
		for p.sym == tIdentifier {
			// only the first message in a sequence can be a super send
			p.unaryMessage(mgenc, superSend)
			*superSend = false
		}
		for p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
			p.binaryMessage(mgenc, false)
		}
		if p.sym == tKeyword {
			p.keywordMessage(mgenc, false)
		}
	} else if p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
		for p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
			// only the first message in a sequence can be a super send
			p.binaryMessage(mgenc, superSend)
			*superSend = false
		}
		if p.sym == tKeyword {
			p.keywordMessage(mgenc, false)
		}
	} else {
		p.keywordMessage(mgenc, superSend)
	}
}

//	private void unaryMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		som.vmobjects.Symbol msg = unarySelector();
//		mgenc.addLiteralIfAbsent(msg);
//		if(superSend.get())
//			bcGen.emitSUPERSEND(mgenc, msg);
//		else
//			bcGen.emitSEND(mgenc, msg);
//	}
func (p *Parser) unaryMessage(mgenc *MethodGenerator, superSend *bool) {
	msg := p.unarySelector()
	mgenc.AddLiteralIfAbsent(msg)
	if *superSend {
		p.bcGen.emitSUPERSEND(mgenc, msg)
	} else {
		p.bcGen.emitSEND(mgenc, msg)
	}
}

//	private void binaryMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		som.vmobjects.Symbol msg = binarySelector();
//		mgenc.addLiteralIfAbsent(msg);
//		binaryOperand(mgenc, new Single<Boolean>(false));
//		if(superSend.get())
//			bcGen.emitSUPERSEND(mgenc, msg);
//		else
//			bcGen.emitSEND(mgenc, msg);
//	}
func (p *Parser) binaryMessage(mgenc *MethodGenerator, superSend bool) {
	msg := p.binarySelector()
	mgenc.AddLiteralIfAbsent(msg)
	p.binaryOperand(mgenc, false)
	if superSend {
		p.bcGen.emitSUPERSEND(mgenc, msg)
	} else {
		p.bcGen.emitSEND(mgenc, msg)
	}
}

//	private void binaryOperand(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		primary(mgenc, superSend);
//		while(sym == Identifier)
//			unaryMessage(mgenc, superSend);
//	}
func (p *Parser) binaryOperand(mgenc *MethodGenerator, superSend bool) {
	p.primary(mgenc, &superSend)
	for p.sym == tIdentifier {
		p.unaryMessage(mgenc, &superSend)
	}
}

//	private void keywordMessage(MethodGenerationContext mgenc, Single<Boolean> superSend) {
//		StringBuffer kw = new StringBuffer();
//		do {
//			kw.append(keyword());
//			formula(mgenc);
//		} while(sym == Keyword);
//		som.vmobjects.Symbol msg = GetUniverse().SymbolFor(kw.toString());
//		mgenc.addLiteralIfAbsent(msg);
//		if(superSend.get())
//			bcGen.emitSUPERSEND(mgenc, msg);
//		else
//			bcGen.emitSEND(mgenc, msg);
//	}
func (p *Parser) keywordMessage(mgenc *MethodGenerator, superSend bool) {
	kw := ""
	for p.sym == tKeyword {
		kw += p.keyword()
		p.formula(mgenc)
	}
	msg := GetUniverse().SymbolFor(kw)
	mgenc.AddLiteralIfAbsent(msg)
	if superSend {
		p.bcGen.emitSUPERSEND(mgenc, msg)
	} else {
		p.bcGen.emitSEND(mgenc, msg)
	}
}

//	private void formula(MethodGenerationContext mgenc) {
//		Single<Boolean> superSend = new Single<Boolean>(false);
//		binaryOperand(mgenc, superSend);
//
// only the first message in a sequence can be a super send
//
//		if(sym == OperatorSequence || symIn(binaryOpSyms))
//			binaryMessage(mgenc, superSend);
//		while(sym == OperatorSequence || symIn(binaryOpSyms))
//			binaryMessage(mgenc, new Single<Boolean>(false));
//	}
func (p *Parser) formula(mgenc *MethodGenerator) {
	superSend := false
	p.binaryOperand(mgenc, superSend)
	// only the first message in a sequence can be a super send
	if p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
		p.binaryMessage(mgenc, superSend)
	}
	for p.sym == tOperatorSequence || p.symIn(p.binaryOpSyms) {
		p.binaryMessage(mgenc, false)
	}
}

//	private void nestedTerm(MethodGenerationContext mgenc) {
//		expect(NewTerm);
//		expression(mgenc);
//		expect(EndTerm);
//	}
func (p *Parser) nestedTerm(mgenc *MethodGenerator) {
	p.expect(tNewTerm)
	p.expression(mgenc)
	p.expect(tEndTerm)
}

//	private void literal(MethodGenerationContext mgenc) {
//		switch(sym) {
//		case Pound:
//			literalSymbol(mgenc);
//			break;
//		case STString:
//			literalString(mgenc);
//			break;
//		default:
//			literalNumber(mgenc);
//			break;
//		}
//	}
func (p *Parser) literal(mgenc *MethodGenerator) {
	switch p.sym {
	case tPound:
		p.literalSymbol(mgenc)
	case tSTString:
		p.literalString(mgenc)
	default:
		p.literalNumber(mgenc)
	}
}

//	private void literalNumber(MethodGenerationContext mgenc) {
//		int val;
//		if(sym == Minus)
//			val = negativeDecimal();
//		else
//			val = literalDecimal();
//		som.vmobjects.Integer lit = GetUniverse().newInteger(val);
//		mgenc.addLiteralIfAbsent(lit);
//		bcGen.emitPUSHCONSTANT(mgenc, lit);
//	}
func (p *Parser) literalNumber(mgenc *MethodGenerator) {
	var val int
	if p.sym == tMinus {
		val = p.negativeDecimal()
	} else {
		val = p.literalDecimal()
	}
	lit := GetUniverse().newInteger(val)
	mgenc.AddLiteralIfAbsent(lit)
	p.bcGen.emitPUSHCONSTANT(mgenc, lit)
}

//	private int literalDecimal() {
//		return literalInteger();
//	}
func (p *Parser) literalDecimal() int {
	return p.literalInteger()
}

//	private int negativeDecimal() {
//		expect(Minus);
//		return -literalInteger();
//	}
func (p *Parser) negativeDecimal() int {
	p.expect(tMinus)
	return -p.literalInteger()
}

//	private int literalInteger() {
//		int i = java.lang.Integer.parseInt(text);
//		expect(Integer);
//		return i;
//	}
func (p *Parser) literalInteger() int {
	i, _ := strconv.Atoi(p.text)
	p.expect(tInteger)
	return i
}

//	private void literalSymbol(MethodGenerationContext mgenc) {
//		som.vmobjects.Symbol symb;
//		expect(Pound);
//		if(sym == STString) {
//			String s = string();
//			symb = GetUniverse().SymbolFor(s);
//		} else
//			symb = selector();
//		mgenc.addLiteralIfAbsent(symb);
//		bcGen.emitPUSHCONSTANT(mgenc, symb);
//	}
func (p *Parser) literalSymbol(mgenc *MethodGenerator) {
	var symb *Symbol
	p.expect(tPound)
	if p.sym == tSTString {
		s := p.string()
		symb = GetUniverse().SymbolFor(s)
	} else {
		symb = p.selector()
	}
	mgenc.AddLiteralIfAbsent(symb)
	p.bcGen.emitPUSHCONSTANT(mgenc, symb)
}

//	private void literalString(MethodGenerationContext mgenc) {
//		String s = string();
//		som.vmobjects.String str = GetUniverse().newString(s);
//		mgenc.addLiteralIfAbsent(str);
//		bcGen.emitPUSHCONSTANT(mgenc, str);
//	}
func (p *Parser) literalString(mgenc *MethodGenerator) {
	s := p.string()
	str := GetUniverse().newString(s)
	mgenc.AddLiteralIfAbsent(str)
	p.bcGen.emitPUSHCONSTANT(mgenc, str)
}

//	private som.vmobjects.Symbol selector() {
//		if(sym == OperatorSequence || symIn(singleOpSyms))
//			return binarySelector();
//		else if(sym == Keyword || sym == KeywordSequence)
//			return keywordSelector();
//		else
//			return unarySelector();
//	}
func (p *Parser) selector() *Symbol {
	if p.sym == tOperatorSequence || p.symIn(p.singleOpSyms) {
		return p.binarySelector()
	} else if p.sym == tKeyword || p.sym == tKeywordSequence {
		return p.keywordSelector()
	} else {
		return p.unarySelector()
	}
}

//	private som.vmobjects.Symbol keywordSelector() {
//		String s = new String(text);
//		expectOneOf(keywordSelectorSyms);
//		som.vmobjects.Symbol symb = GetUniverse().SymbolFor(s);
//		return symb;
//	}
func (p *Parser) keywordSelector() *Symbol {
	s := p.text
	p.expectOneOf(p.keywordSelectorSyms)
	symb := GetUniverse().SymbolFor(s)
	return symb
}

//	private String string() {
//		String s = new String(text);
//		expect(STString);
//		return s;
//	}
func (p *Parser) string() string {
	s := p.text
	p.expect(tSTString)
	return s
}

//	private void nestedBlock(MethodGenerationContext mgenc) {
//		mgenc.addArgumentIfAbsent("$block self");
//		expect(NewBlock);
//		if(sym == Colon)
//			blockPattern(mgenc);
//		// generate Block signature
//		String blockSig = "$block method";
//		int argSize = mgenc.getNumberOfArguments();
//		for(int i = 1; i < argSize; i++)
//			blockSig += ":";
//		mgenc.setSignature(GetUniverse().SymbolFor(blockSig));
//		blockContents(mgenc);
//		// if no return has been generated, we can be sure that the last expression
//		// in the block was not terminated by ., and can generate a return
//		if(!mgenc.isFinished()) {
//			bcGen.emitRETURNLOCAL(mgenc);
//			mgenc.setFinished(true);
//		}
//		expect(EndBlock);
//	}
func (p *Parser) nestedBlock(mgenc *MethodGenerator) {
	mgenc.AddArgumentIfAbsent("$block self")
	p.expect(tNewBlock)
	if p.sym == tColon {
		p.blockPattern(mgenc)
	}
	// generate Block signature
	blockSig := "$block method"
	argSize := mgenc.GetNumberOfArguments()
	for i := 1; i < argSize; i++ {
		blockSig += ":"
	}
	mgenc.SetSignature(GetUniverse().SymbolFor(blockSig))
	p.blockContents(mgenc)
	// if no return has been generated, we can be sure that the last expression
	// in the block was not terminated by ., and can generate a return
	if !mgenc.IsFinished() {
		p.bcGen.emitRETURNLOCAL(mgenc)
		mgenc.SetFinished(true)
	}
	p.expect(tEndBlock)
}

//	private void blockPattern(MethodGenerationContext mgenc) {
//		blockArguments(mgenc);
//		expect(Or);
//	}
func (p *Parser) blockPattern(mgenc *MethodGenerator) {
	p.blockArguments(mgenc)
	p.expect(tOr)
}

//	private void blockArguments(MethodGenerationContext mgenc) {
//		do {
//			expect(Colon);
//			mgenc.addArgumentIfAbsent(argument());
//		} while(sym == Colon);
//	}
func (p *Parser) blockArguments(mgenc *MethodGenerator) {
	for p.sym == tColon {
		p.expect(tColon)
		mgenc.AddArgumentIfAbsent(p.argument())
	}
}

//	private void genPushVariable(MethodGenerationContext mgenc, String var) {
//		// The purpose of this function is to find out whether the variable to be
//		// pushed on the stack is a local variable, argument, or object field. This
//		// is done by examining all available lexical contexts, starting with the
//		// innermost (i.e., the one represented by mgenc).
//		// triplet: index, context, isArgument
//		Triplet<Byte,Byte,Boolean> tri = new Triplet<Byte,Byte,Boolean>((byte) 0, (byte) 0, false);
//		if(mgenc.findVar(var, tri)) {
//			if(tri.getZ())
//				bcGen.emitPUSHARGUMENT(mgenc, tri.getX(), tri.getY());
//			else
//				bcGen.emitPUSHLOCAL(mgenc, tri.getX(), tri.getY());
//		} else if(mgenc.findField(var)) {
//			som.vmobjects.Symbol fieldName = GetUniverse().SymbolFor(var);
//			mgenc.addLiteralIfAbsent(fieldName);
//			bcGen.emitPUSHFIELD(mgenc, fieldName);
//		} else {
//			som.vmobjects.Symbol global = GetUniverse().SymbolFor(var);
//			mgenc.addLiteralIfAbsent(global);
//			bcGen.emitPUSHGLOBAL(mgenc, global);
//		}
//	}
func (p *Parser) genPushVariable(mgenc *MethodGenerator, varName string) {
	// The purpose of this function is to find out whether the variable to be
	// pushed on the stack is a local variable, argument, or object field. This
	// is done by examining all available lexical contexts, starting with the
	// innermost (i.e., the one represented by mgenc).
	// triplet: index, context, isArgument
	tri := p.findVar(mgenc, varName)
	if tri != nil {
		if tri.Z {
			p.bcGen.emitPUSHARGUMENT(mgenc, tri.X, tri.Y)
		} else {
			p.bcGen.emitPUSHLOCAL(mgenc, tri.X, tri.Y)
		}
	} else if mgenc.FindField(varName) {
		fieldName := GetUniverse().SymbolFor(varName)
		mgenc.AddLiteralIfAbsent(fieldName)
		p.bcGen.emitPUSHFIELD(mgenc, fieldName)
	} else {
		global := GetUniverse().SymbolFor(varName)
		mgenc.AddLiteralIfAbsent(global)
		p.bcGen.emitPUSHGLOBAL(mgenc, global)
	}
}

//	private void genPopVariable(MethodGenerationContext mgenc, String var) {
//		// The purpose of this function is to find out whether the variable to be
//		// popped off the stack is a local variable, argument, or object field. This
//		// is done by examining all available lexical contexts, starting with the
//		// innermost (i.e., the one represented by mgenc).
//		// triplet: index, context, isArgument
//		Triplet<Byte,Byte,Boolean> tri = new Triplet<Byte,Byte,Boolean>((byte) 0, (byte) 0, false);
//		if(mgenc.findVar(var, tri)) {
//			if(tri.getZ())
//				bcGen.emitPOPARGUMENT(mgenc, tri.getX(), tri.getY());
//			else
//				bcGen.emitPOPLOCAL(mgenc, tri.getX(), tri.getY());
//		} else
//			bcGen.emitPOPFIELD(mgenc, GetUniverse().SymbolFor(var));
//	}
func (p *Parser) genPopVariable(mgenc *MethodGenerator, varName string) {
	// The purpose of this function is to find out whether the variable to be
	// popped off the stack is a local variable, argument, or object field. This
	// is done by examining all available lexical contexts, starting with the
	// innermost (i.e., the one represented by mgenc).
	// triplet: index, context, isArgument
	tri := p.findVar(mgenc, varName)
	if tri != nil {
		if tri.Z {
			p.bcGen.emitPOPARGUMENT(mgenc, tri.X, tri.Y)
		} else {
			p.bcGen.emitPOPLOCAL(mgenc, tri.X, tri.Y)
		}
	} else {
		p.bcGen.emitPOPFIELD(mgenc, GetUniverse().SymbolFor(varName))
	}
}

//	private void GETSYM() {
//		sym = lexer.getSym();
//		text = lexer.getText();
//	}
func (p *Parser) GETSYM() {
	p.sym, p.text = p.lexer.GetSym()
}

//	private void PEEK() {
//		nextSym = lexer.peek();
//	}
func (p *Parser) PEEK() {
	p.nextSym = p.lexer.Peek()
}

//	private boolean printableSymbol() {
//		return sym == Integer || sym.compareTo(STString) >= 0;
//	}
func (p *Parser) printableSymbol() bool {
	return p.sym == tInteger || p.sym >= tSTString
}

// }
