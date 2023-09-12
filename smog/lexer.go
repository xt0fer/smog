package smog

import (
	"bufio"
	"io"
	"unicode"
)

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
				l.match(tNot)
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
