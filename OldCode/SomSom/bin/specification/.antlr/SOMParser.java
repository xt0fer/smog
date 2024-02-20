// Generated from /Volumes/Terabyte/kristofer/LocalProjects/smog/OldCode/SomSom/bin/specification/SOM.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class SOMParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Comment=1, Whitespace=2, Primitive=3, Identifier=4, Equal=5, Separator=6, 
		NewTerm=7, EndTerm=8, Or=9, Comma=10, Minus=11, Not=12, And=13, Star=14, 
		Div=15, Mod=16, Plus=17, More=18, Less=19, At=20, Per=21, OperatorSequence=22, 
		Colon=23, NewBlock=24, EndBlock=25, Pound=26, Exit=27, Period=28, Assign=29, 
		Integer=30, Double=31, Keyword=32, KeywordSequence=33, STString=34;
	public static final int
		RULE_classdef = 0, RULE_superclass = 1, RULE_instanceFields = 2, RULE_classFields = 3, 
		RULE_method = 4, RULE_pattern = 5, RULE_unaryPattern = 6, RULE_binaryPattern = 7, 
		RULE_keywordPattern = 8, RULE_methodBlock = 9, RULE_unarySelector = 10, 
		RULE_binarySelector = 11, RULE_identifier = 12, RULE_keyword = 13, RULE_argument = 14, 
		RULE_blockContents = 15, RULE_localDefs = 16, RULE_blockBody = 17, RULE_result = 18, 
		RULE_expression = 19, RULE_assignation = 20, RULE_assignments = 21, RULE_assignment = 22, 
		RULE_evaluation = 23, RULE_primary = 24, RULE_variable = 25, RULE_messages = 26, 
		RULE_unaryMessage = 27, RULE_binaryMessage = 28, RULE_binaryOperand = 29, 
		RULE_keywordMessage = 30, RULE_formula = 31, RULE_nestedTerm = 32, RULE_literal = 33, 
		RULE_literalArray = 34, RULE_literalNumber = 35, RULE_literalDecimal = 36, 
		RULE_negativeDecimal = 37, RULE_literalInteger = 38, RULE_literalDouble = 39, 
		RULE_literalSymbol = 40, RULE_literalString = 41, RULE_selector = 42, 
		RULE_keywordSelector = 43, RULE_string = 44, RULE_nestedBlock = 45, RULE_blockPattern = 46, 
		RULE_blockArguments = 47;
	private static String[] makeRuleNames() {
		return new String[] {
			"classdef", "superclass", "instanceFields", "classFields", "method", 
			"pattern", "unaryPattern", "binaryPattern", "keywordPattern", "methodBlock", 
			"unarySelector", "binarySelector", "identifier", "keyword", "argument", 
			"blockContents", "localDefs", "blockBody", "result", "expression", "assignation", 
			"assignments", "assignment", "evaluation", "primary", "variable", "messages", 
			"unaryMessage", "binaryMessage", "binaryOperand", "keywordMessage", "formula", 
			"nestedTerm", "literal", "literalArray", "literalNumber", "literalDecimal", 
			"negativeDecimal", "literalInteger", "literalDouble", "literalSymbol", 
			"literalString", "selector", "keywordSelector", "string", "nestedBlock", 
			"blockPattern", "blockArguments"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, "'primitive'", null, "'='", null, "'('", "')'", "'|'", 
			"','", "'-'", "'~'", "'&'", "'*'", "'/'", "'\\'", "'+'", "'>'", "'<'", 
			"'@'", "'%'", null, "':'", "'['", "']'", "'#'", "'^'", "'.'", "':='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Comment", "Whitespace", "Primitive", "Identifier", "Equal", "Separator", 
			"NewTerm", "EndTerm", "Or", "Comma", "Minus", "Not", "And", "Star", "Div", 
			"Mod", "Plus", "More", "Less", "At", "Per", "OperatorSequence", "Colon", 
			"NewBlock", "EndBlock", "Pound", "Exit", "Period", "Assign", "Integer", 
			"Double", "Keyword", "KeywordSequence", "STString"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "SOM.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SOMParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ClassdefContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(SOMParser.Identifier, 0); }
		public TerminalNode Equal() { return getToken(SOMParser.Equal, 0); }
		public SuperclassContext superclass() {
			return getRuleContext(SuperclassContext.class,0);
		}
		public InstanceFieldsContext instanceFields() {
			return getRuleContext(InstanceFieldsContext.class,0);
		}
		public TerminalNode EndTerm() { return getToken(SOMParser.EndTerm, 0); }
		public List<MethodContext> method() {
			return getRuleContexts(MethodContext.class);
		}
		public MethodContext method(int i) {
			return getRuleContext(MethodContext.class,i);
		}
		public TerminalNode Separator() { return getToken(SOMParser.Separator, 0); }
		public ClassFieldsContext classFields() {
			return getRuleContext(ClassFieldsContext.class,0);
		}
		public ClassdefContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classdef; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterClassdef(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitClassdef(this);
		}
	}

	public final ClassdefContext classdef() throws RecognitionException {
		ClassdefContext _localctx = new ClassdefContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_classdef);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(Identifier);
			setState(97);
			match(Equal);
			setState(98);
			superclass();
			setState(99);
			instanceFields();
			setState(103);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4303355448L) != 0)) {
				{
				{
				setState(100);
				method();
				}
				}
				setState(105);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(114);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Separator) {
				{
				setState(106);
				match(Separator);
				setState(107);
				classFields();
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4303355448L) != 0)) {
					{
					{
					setState(108);
					method();
					}
					}
					setState(113);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(116);
			match(EndTerm);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SuperclassContext extends ParserRuleContext {
		public TerminalNode NewTerm() { return getToken(SOMParser.NewTerm, 0); }
		public TerminalNode Identifier() { return getToken(SOMParser.Identifier, 0); }
		public SuperclassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_superclass; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterSuperclass(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitSuperclass(this);
		}
	}

	public final SuperclassContext superclass() throws RecognitionException {
		SuperclassContext _localctx = new SuperclassContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_superclass);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(118);
				match(Identifier);
				}
			}

			setState(121);
			match(NewTerm);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InstanceFieldsContext extends ParserRuleContext {
		public List<TerminalNode> Or() { return getTokens(SOMParser.Or); }
		public TerminalNode Or(int i) {
			return getToken(SOMParser.Or, i);
		}
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public InstanceFieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instanceFields; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterInstanceFields(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitInstanceFields(this);
		}
	}

	public final InstanceFieldsContext instanceFields() throws RecognitionException {
		InstanceFieldsContext _localctx = new InstanceFieldsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instanceFields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				setState(123);
				match(Or);
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Primitive || _la==Identifier) {
					{
					{
					setState(124);
					variable();
					}
					}
					setState(129);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(130);
				match(Or);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ClassFieldsContext extends ParserRuleContext {
		public List<TerminalNode> Or() { return getTokens(SOMParser.Or); }
		public TerminalNode Or(int i) {
			return getToken(SOMParser.Or, i);
		}
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public ClassFieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classFields; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterClassFields(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitClassFields(this);
		}
	}

	public final ClassFieldsContext classFields() throws RecognitionException {
		ClassFieldsContext _localctx = new ClassFieldsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_classFields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(133);
				match(Or);
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Primitive || _la==Identifier) {
					{
					{
					setState(134);
					variable();
					}
					}
					setState(139);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(140);
				match(Or);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodContext extends ParserRuleContext {
		public PatternContext pattern() {
			return getRuleContext(PatternContext.class,0);
		}
		public TerminalNode Equal() { return getToken(SOMParser.Equal, 0); }
		public TerminalNode Primitive() { return getToken(SOMParser.Primitive, 0); }
		public MethodBlockContext methodBlock() {
			return getRuleContext(MethodBlockContext.class,0);
		}
		public MethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterMethod(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitMethod(this);
		}
	}

	public final MethodContext method() throws RecognitionException {
		MethodContext _localctx = new MethodContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_method);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			pattern();
			setState(144);
			match(Equal);
			setState(147);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Primitive:
				{
				setState(145);
				match(Primitive);
				}
				break;
			case NewTerm:
				{
				setState(146);
				methodBlock();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PatternContext extends ParserRuleContext {
		public UnaryPatternContext unaryPattern() {
			return getRuleContext(UnaryPatternContext.class,0);
		}
		public KeywordPatternContext keywordPattern() {
			return getRuleContext(KeywordPatternContext.class,0);
		}
		public BinaryPatternContext binaryPattern() {
			return getRuleContext(BinaryPatternContext.class,0);
		}
		public PatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitPattern(this);
		}
	}

	public final PatternContext pattern() throws RecognitionException {
		PatternContext _localctx = new PatternContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_pattern);
		try {
			setState(152);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Primitive:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(149);
				unaryPattern();
				}
				break;
			case Keyword:
				enterOuterAlt(_localctx, 2);
				{
				setState(150);
				keywordPattern();
				}
				break;
			case Equal:
			case Or:
			case Comma:
			case Minus:
			case Not:
			case And:
			case Star:
			case Div:
			case Mod:
			case Plus:
			case More:
			case Less:
			case At:
			case Per:
			case OperatorSequence:
				enterOuterAlt(_localctx, 3);
				{
				setState(151);
				binaryPattern();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnaryPatternContext extends ParserRuleContext {
		public UnarySelectorContext unarySelector() {
			return getRuleContext(UnarySelectorContext.class,0);
		}
		public UnaryPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterUnaryPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitUnaryPattern(this);
		}
	}

	public final UnaryPatternContext unaryPattern() throws RecognitionException {
		UnaryPatternContext _localctx = new UnaryPatternContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_unaryPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			unarySelector();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryPatternContext extends ParserRuleContext {
		public BinarySelectorContext binarySelector() {
			return getRuleContext(BinarySelectorContext.class,0);
		}
		public ArgumentContext argument() {
			return getRuleContext(ArgumentContext.class,0);
		}
		public BinaryPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBinaryPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBinaryPattern(this);
		}
	}

	public final BinaryPatternContext binaryPattern() throws RecognitionException {
		BinaryPatternContext _localctx = new BinaryPatternContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_binaryPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			binarySelector();
			setState(157);
			argument();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordPatternContext extends ParserRuleContext {
		public List<KeywordContext> keyword() {
			return getRuleContexts(KeywordContext.class);
		}
		public KeywordContext keyword(int i) {
			return getRuleContext(KeywordContext.class,i);
		}
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public KeywordPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keywordPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterKeywordPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitKeywordPattern(this);
		}
	}

	public final KeywordPatternContext keywordPattern() throws RecognitionException {
		KeywordPatternContext _localctx = new KeywordPatternContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_keywordPattern);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(159);
				keyword();
				setState(160);
				argument();
				}
				}
				setState(164); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Keyword );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MethodBlockContext extends ParserRuleContext {
		public TerminalNode NewTerm() { return getToken(SOMParser.NewTerm, 0); }
		public TerminalNode EndTerm() { return getToken(SOMParser.EndTerm, 0); }
		public BlockContentsContext blockContents() {
			return getRuleContext(BlockContentsContext.class,0);
		}
		public MethodBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterMethodBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitMethodBlock(this);
		}
	}

	public final MethodBlockContext methodBlock() throws RecognitionException {
		MethodBlockContext _localctx = new MethodBlockContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_methodBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(NewTerm);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 20619201176L) != 0)) {
				{
				setState(167);
				blockContents();
				}
			}

			setState(170);
			match(EndTerm);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnarySelectorContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public UnarySelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unarySelector; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterUnarySelector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitUnarySelector(this);
		}
	}

	public final UnarySelectorContext unarySelector() throws RecognitionException {
		UnarySelectorContext _localctx = new UnarySelectorContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_unarySelector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BinarySelectorContext extends ParserRuleContext {
		public TerminalNode Or() { return getToken(SOMParser.Or, 0); }
		public TerminalNode Comma() { return getToken(SOMParser.Comma, 0); }
		public TerminalNode Minus() { return getToken(SOMParser.Minus, 0); }
		public TerminalNode Equal() { return getToken(SOMParser.Equal, 0); }
		public TerminalNode Not() { return getToken(SOMParser.Not, 0); }
		public TerminalNode And() { return getToken(SOMParser.And, 0); }
		public TerminalNode Star() { return getToken(SOMParser.Star, 0); }
		public TerminalNode Div() { return getToken(SOMParser.Div, 0); }
		public TerminalNode Mod() { return getToken(SOMParser.Mod, 0); }
		public TerminalNode Plus() { return getToken(SOMParser.Plus, 0); }
		public TerminalNode More() { return getToken(SOMParser.More, 0); }
		public TerminalNode Less() { return getToken(SOMParser.Less, 0); }
		public TerminalNode At() { return getToken(SOMParser.At, 0); }
		public TerminalNode Per() { return getToken(SOMParser.Per, 0); }
		public TerminalNode OperatorSequence() { return getToken(SOMParser.OperatorSequence, 0); }
		public BinarySelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binarySelector; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBinarySelector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBinarySelector(this);
		}
	}

	public final BinarySelectorContext binarySelector() throws RecognitionException {
		BinarySelectorContext _localctx = new BinarySelectorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_binarySelector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8388128L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ParserRuleContext {
		public TerminalNode Primitive() { return getToken(SOMParser.Primitive, 0); }
		public TerminalNode Identifier() { return getToken(SOMParser.Identifier, 0); }
		public IdentifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterIdentifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitIdentifier(this);
		}
	}

	public final IdentifierContext identifier() throws RecognitionException {
		IdentifierContext _localctx = new IdentifierContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_identifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(176);
			_la = _input.LA(1);
			if ( !(_la==Primitive || _la==Identifier) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordContext extends ParserRuleContext {
		public TerminalNode Keyword() { return getToken(SOMParser.Keyword, 0); }
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitKeyword(this);
		}
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_keyword);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(Keyword);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public ArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterArgument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitArgument(this);
		}
	}

	public final ArgumentContext argument() throws RecognitionException {
		ArgumentContext _localctx = new ArgumentContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_argument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(180);
			variable();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContentsContext extends ParserRuleContext {
		public BlockBodyContext blockBody() {
			return getRuleContext(BlockBodyContext.class,0);
		}
		public List<TerminalNode> Or() { return getTokens(SOMParser.Or); }
		public TerminalNode Or(int i) {
			return getToken(SOMParser.Or, i);
		}
		public LocalDefsContext localDefs() {
			return getRuleContext(LocalDefsContext.class,0);
		}
		public BlockContentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockContents; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBlockContents(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBlockContents(this);
		}
	}

	public final BlockContentsContext blockContents() throws RecognitionException {
		BlockContentsContext _localctx = new BlockContentsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_blockContents);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(186);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Or) {
				{
				setState(182);
				match(Or);
				setState(183);
				localDefs();
				setState(184);
				match(Or);
				}
			}

			setState(188);
			blockBody();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LocalDefsContext extends ParserRuleContext {
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public LocalDefsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_localDefs; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLocalDefs(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLocalDefs(this);
		}
	}

	public final LocalDefsContext localDefs() throws RecognitionException {
		LocalDefsContext _localctx = new LocalDefsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_localDefs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Primitive || _la==Identifier) {
				{
				{
				setState(190);
				variable();
				}
				}
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockBodyContext extends ParserRuleContext {
		public TerminalNode Exit() { return getToken(SOMParser.Exit, 0); }
		public ResultContext result() {
			return getRuleContext(ResultContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Period() { return getToken(SOMParser.Period, 0); }
		public BlockBodyContext blockBody() {
			return getRuleContext(BlockBodyContext.class,0);
		}
		public BlockBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBlockBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBlockBody(this);
		}
	}

	public final BlockBodyContext blockBody() throws RecognitionException {
		BlockBodyContext _localctx = new BlockBodyContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_blockBody);
		int _la;
		try {
			setState(205);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Exit:
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				match(Exit);
				setState(197);
				result();
				}
				break;
			case Primitive:
			case Identifier:
			case NewTerm:
			case Minus:
			case NewBlock:
			case Pound:
			case Integer:
			case Double:
			case STString:
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				expression();
				setState(203);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Period) {
					{
					setState(199);
					match(Period);
					setState(201);
					_errHandler.sync(this);
					_la = _input.LA(1);
					if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 20619200664L) != 0)) {
						{
						setState(200);
						blockBody();
						}
					}

					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ResultContext extends ParserRuleContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Period() { return getToken(SOMParser.Period, 0); }
		public ResultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_result; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterResult(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitResult(this);
		}
	}

	public final ResultContext result() throws RecognitionException {
		ResultContext _localctx = new ResultContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_result);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			expression();
			setState(209);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Period) {
				{
				setState(208);
				match(Period);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public AssignationContext assignation() {
			return getRuleContext(AssignationContext.class,0);
		}
		public EvaluationContext evaluation() {
			return getRuleContext(EvaluationContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitExpression(this);
		}
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expression);
		try {
			setState(213);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(211);
				assignation();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				evaluation();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignationContext extends ParserRuleContext {
		public AssignmentsContext assignments() {
			return getRuleContext(AssignmentsContext.class,0);
		}
		public EvaluationContext evaluation() {
			return getRuleContext(EvaluationContext.class,0);
		}
		public AssignationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignation; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterAssignation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitAssignation(this);
		}
	}

	public final AssignationContext assignation() throws RecognitionException {
		AssignationContext _localctx = new AssignationContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assignation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(215);
			assignments();
			setState(216);
			evaluation();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentsContext extends ParserRuleContext {
		public List<AssignmentContext> assignment() {
			return getRuleContexts(AssignmentContext.class);
		}
		public AssignmentContext assignment(int i) {
			return getRuleContext(AssignmentContext.class,i);
		}
		public AssignmentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignments; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterAssignments(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitAssignments(this);
		}
	}

	public final AssignmentsContext assignments() throws RecognitionException {
		AssignmentsContext _localctx = new AssignmentsContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_assignments);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(219); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(218);
					assignment();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(221); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,19,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AssignmentContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode Assign() { return getToken(SOMParser.Assign, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitAssignment(this);
		}
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			variable();
			setState(224);
			match(Assign);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EvaluationContext extends ParserRuleContext {
		public PrimaryContext primary() {
			return getRuleContext(PrimaryContext.class,0);
		}
		public MessagesContext messages() {
			return getRuleContext(MessagesContext.class,0);
		}
		public EvaluationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_evaluation; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterEvaluation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitEvaluation(this);
		}
	}

	public final EvaluationContext evaluation() throws RecognitionException {
		EvaluationContext _localctx = new EvaluationContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_evaluation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226);
			primary();
			setState(228);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4303355448L) != 0)) {
				{
				setState(227);
				messages();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public NestedTermContext nestedTerm() {
			return getRuleContext(NestedTermContext.class,0);
		}
		public NestedBlockContext nestedBlock() {
			return getRuleContext(NestedBlockContext.class,0);
		}
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public PrimaryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primary; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterPrimary(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitPrimary(this);
		}
	}

	public final PrimaryContext primary() throws RecognitionException {
		PrimaryContext _localctx = new PrimaryContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_primary);
		try {
			setState(234);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Primitive:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(230);
				variable();
				}
				break;
			case NewTerm:
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
				nestedTerm();
				}
				break;
			case NewBlock:
				enterOuterAlt(_localctx, 3);
				{
				setState(232);
				nestedBlock();
				}
				break;
			case Minus:
			case Pound:
			case Integer:
			case Double:
			case STString:
				enterOuterAlt(_localctx, 4);
				{
				setState(233);
				literal();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VariableContext extends ParserRuleContext {
		public IdentifierContext identifier() {
			return getRuleContext(IdentifierContext.class,0);
		}
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterVariable(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitVariable(this);
		}
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_variable);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			identifier();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MessagesContext extends ParserRuleContext {
		public List<UnaryMessageContext> unaryMessage() {
			return getRuleContexts(UnaryMessageContext.class);
		}
		public UnaryMessageContext unaryMessage(int i) {
			return getRuleContext(UnaryMessageContext.class,i);
		}
		public List<BinaryMessageContext> binaryMessage() {
			return getRuleContexts(BinaryMessageContext.class);
		}
		public BinaryMessageContext binaryMessage(int i) {
			return getRuleContext(BinaryMessageContext.class,i);
		}
		public KeywordMessageContext keywordMessage() {
			return getRuleContext(KeywordMessageContext.class,0);
		}
		public MessagesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_messages; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterMessages(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitMessages(this);
		}
	}

	public final MessagesContext messages() throws RecognitionException {
		MessagesContext _localctx = new MessagesContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_messages);
		int _la;
		try {
			setState(261);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Primitive:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(239); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(238);
					unaryMessage();
					}
					}
					setState(241); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==Primitive || _la==Identifier );
				setState(246);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8388128L) != 0)) {
					{
					{
					setState(243);
					binaryMessage();
					}
					}
					setState(248);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(250);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Keyword) {
					{
					setState(249);
					keywordMessage();
					}
				}

				}
				break;
			case Equal:
			case Or:
			case Comma:
			case Minus:
			case Not:
			case And:
			case Star:
			case Div:
			case Mod:
			case Plus:
			case More:
			case Less:
			case At:
			case Per:
			case OperatorSequence:
				enterOuterAlt(_localctx, 2);
				{
				setState(253); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(252);
					binaryMessage();
					}
					}
					setState(255); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 8388128L) != 0) );
				setState(258);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Keyword) {
					{
					setState(257);
					keywordMessage();
					}
				}

				}
				break;
			case Keyword:
				enterOuterAlt(_localctx, 3);
				{
				setState(260);
				keywordMessage();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnaryMessageContext extends ParserRuleContext {
		public UnarySelectorContext unarySelector() {
			return getRuleContext(UnarySelectorContext.class,0);
		}
		public UnaryMessageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryMessage; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterUnaryMessage(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitUnaryMessage(this);
		}
	}

	public final UnaryMessageContext unaryMessage() throws RecognitionException {
		UnaryMessageContext _localctx = new UnaryMessageContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_unaryMessage);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			unarySelector();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryMessageContext extends ParserRuleContext {
		public BinarySelectorContext binarySelector() {
			return getRuleContext(BinarySelectorContext.class,0);
		}
		public BinaryOperandContext binaryOperand() {
			return getRuleContext(BinaryOperandContext.class,0);
		}
		public BinaryMessageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryMessage; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBinaryMessage(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBinaryMessage(this);
		}
	}

	public final BinaryMessageContext binaryMessage() throws RecognitionException {
		BinaryMessageContext _localctx = new BinaryMessageContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_binaryMessage);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			binarySelector();
			setState(266);
			binaryOperand();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BinaryOperandContext extends ParserRuleContext {
		public PrimaryContext primary() {
			return getRuleContext(PrimaryContext.class,0);
		}
		public List<UnaryMessageContext> unaryMessage() {
			return getRuleContexts(UnaryMessageContext.class);
		}
		public UnaryMessageContext unaryMessage(int i) {
			return getRuleContext(UnaryMessageContext.class,i);
		}
		public BinaryOperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_binaryOperand; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBinaryOperand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBinaryOperand(this);
		}
	}

	public final BinaryOperandContext binaryOperand() throws RecognitionException {
		BinaryOperandContext _localctx = new BinaryOperandContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_binaryOperand);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			primary();
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Primitive || _la==Identifier) {
				{
				{
				setState(269);
				unaryMessage();
				}
				}
				setState(274);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordMessageContext extends ParserRuleContext {
		public List<KeywordContext> keyword() {
			return getRuleContexts(KeywordContext.class);
		}
		public KeywordContext keyword(int i) {
			return getRuleContext(KeywordContext.class,i);
		}
		public List<FormulaContext> formula() {
			return getRuleContexts(FormulaContext.class);
		}
		public FormulaContext formula(int i) {
			return getRuleContext(FormulaContext.class,i);
		}
		public KeywordMessageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keywordMessage; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterKeywordMessage(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitKeywordMessage(this);
		}
	}

	public final KeywordMessageContext keywordMessage() throws RecognitionException {
		KeywordMessageContext _localctx = new KeywordMessageContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_keywordMessage);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(278); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(275);
				keyword();
				setState(276);
				formula();
				}
				}
				setState(280); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Keyword );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FormulaContext extends ParserRuleContext {
		public BinaryOperandContext binaryOperand() {
			return getRuleContext(BinaryOperandContext.class,0);
		}
		public List<BinaryMessageContext> binaryMessage() {
			return getRuleContexts(BinaryMessageContext.class);
		}
		public BinaryMessageContext binaryMessage(int i) {
			return getRuleContext(BinaryMessageContext.class,i);
		}
		public FormulaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formula; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterFormula(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitFormula(this);
		}
	}

	public final FormulaContext formula() throws RecognitionException {
		FormulaContext _localctx = new FormulaContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_formula);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			binaryOperand();
			setState(286);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8388128L) != 0)) {
				{
				{
				setState(283);
				binaryMessage();
				}
				}
				setState(288);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NestedTermContext extends ParserRuleContext {
		public TerminalNode NewTerm() { return getToken(SOMParser.NewTerm, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode EndTerm() { return getToken(SOMParser.EndTerm, 0); }
		public NestedTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestedTerm; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterNestedTerm(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitNestedTerm(this);
		}
	}

	public final NestedTermContext nestedTerm() throws RecognitionException {
		NestedTermContext _localctx = new NestedTermContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_nestedTerm);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			match(NewTerm);
			setState(290);
			expression();
			setState(291);
			match(EndTerm);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralArrayContext literalArray() {
			return getRuleContext(LiteralArrayContext.class,0);
		}
		public LiteralSymbolContext literalSymbol() {
			return getRuleContext(LiteralSymbolContext.class,0);
		}
		public LiteralStringContext literalString() {
			return getRuleContext(LiteralStringContext.class,0);
		}
		public LiteralNumberContext literalNumber() {
			return getRuleContext(LiteralNumberContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteral(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_literal);
		try {
			setState(297);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(293);
				literalArray();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(294);
				literalSymbol();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(295);
				literalString();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(296);
				literalNumber();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralArrayContext extends ParserRuleContext {
		public TerminalNode Pound() { return getToken(SOMParser.Pound, 0); }
		public TerminalNode NewTerm() { return getToken(SOMParser.NewTerm, 0); }
		public TerminalNode EndTerm() { return getToken(SOMParser.EndTerm, 0); }
		public List<LiteralContext> literal() {
			return getRuleContexts(LiteralContext.class);
		}
		public LiteralContext literal(int i) {
			return getRuleContext(LiteralContext.class,i);
		}
		public LiteralArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalArray; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralArray(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralArray(this);
		}
	}

	public final LiteralArrayContext literalArray() throws RecognitionException {
		LiteralArrayContext _localctx = new LiteralArrayContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_literalArray);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			match(Pound);
			setState(300);
			match(NewTerm);
			setState(304);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 20468205568L) != 0)) {
				{
				{
				setState(301);
				literal();
				}
				}
				setState(306);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(307);
			match(EndTerm);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralNumberContext extends ParserRuleContext {
		public NegativeDecimalContext negativeDecimal() {
			return getRuleContext(NegativeDecimalContext.class,0);
		}
		public LiteralDecimalContext literalDecimal() {
			return getRuleContext(LiteralDecimalContext.class,0);
		}
		public LiteralNumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalNumber; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralNumber(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralNumber(this);
		}
	}

	public final LiteralNumberContext literalNumber() throws RecognitionException {
		LiteralNumberContext _localctx = new LiteralNumberContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_literalNumber);
		try {
			setState(311);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Minus:
				enterOuterAlt(_localctx, 1);
				{
				setState(309);
				negativeDecimal();
				}
				break;
			case Integer:
			case Double:
				enterOuterAlt(_localctx, 2);
				{
				setState(310);
				literalDecimal();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralDecimalContext extends ParserRuleContext {
		public LiteralIntegerContext literalInteger() {
			return getRuleContext(LiteralIntegerContext.class,0);
		}
		public LiteralDoubleContext literalDouble() {
			return getRuleContext(LiteralDoubleContext.class,0);
		}
		public LiteralDecimalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalDecimal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralDecimal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralDecimal(this);
		}
	}

	public final LiteralDecimalContext literalDecimal() throws RecognitionException {
		LiteralDecimalContext _localctx = new LiteralDecimalContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_literalDecimal);
		try {
			setState(315);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Integer:
				enterOuterAlt(_localctx, 1);
				{
				setState(313);
				literalInteger();
				}
				break;
			case Double:
				enterOuterAlt(_localctx, 2);
				{
				setState(314);
				literalDouble();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NegativeDecimalContext extends ParserRuleContext {
		public TerminalNode Minus() { return getToken(SOMParser.Minus, 0); }
		public LiteralDecimalContext literalDecimal() {
			return getRuleContext(LiteralDecimalContext.class,0);
		}
		public NegativeDecimalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_negativeDecimal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterNegativeDecimal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitNegativeDecimal(this);
		}
	}

	public final NegativeDecimalContext negativeDecimal() throws RecognitionException {
		NegativeDecimalContext _localctx = new NegativeDecimalContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_negativeDecimal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
			match(Minus);
			setState(318);
			literalDecimal();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralIntegerContext extends ParserRuleContext {
		public TerminalNode Integer() { return getToken(SOMParser.Integer, 0); }
		public LiteralIntegerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalInteger; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralInteger(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralInteger(this);
		}
	}

	public final LiteralIntegerContext literalInteger() throws RecognitionException {
		LiteralIntegerContext _localctx = new LiteralIntegerContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_literalInteger);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(320);
			match(Integer);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralDoubleContext extends ParserRuleContext {
		public TerminalNode Double() { return getToken(SOMParser.Double, 0); }
		public LiteralDoubleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalDouble; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralDouble(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralDouble(this);
		}
	}

	public final LiteralDoubleContext literalDouble() throws RecognitionException {
		LiteralDoubleContext _localctx = new LiteralDoubleContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_literalDouble);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322);
			match(Double);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralSymbolContext extends ParserRuleContext {
		public TerminalNode Pound() { return getToken(SOMParser.Pound, 0); }
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public SelectorContext selector() {
			return getRuleContext(SelectorContext.class,0);
		}
		public LiteralSymbolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalSymbol; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralSymbol(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralSymbol(this);
		}
	}

	public final LiteralSymbolContext literalSymbol() throws RecognitionException {
		LiteralSymbolContext _localctx = new LiteralSymbolContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_literalSymbol);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			match(Pound);
			setState(327);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STString:
				{
				setState(325);
				string();
				}
				break;
			case Primitive:
			case Identifier:
			case Equal:
			case Or:
			case Comma:
			case Minus:
			case Not:
			case And:
			case Star:
			case Div:
			case Mod:
			case Plus:
			case More:
			case Less:
			case At:
			case Per:
			case OperatorSequence:
			case Keyword:
			case KeywordSequence:
				{
				setState(326);
				selector();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralStringContext extends ParserRuleContext {
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public LiteralStringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalString; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterLiteralString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitLiteralString(this);
		}
	}

	public final LiteralStringContext literalString() throws RecognitionException {
		LiteralStringContext _localctx = new LiteralStringContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_literalString);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			string();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SelectorContext extends ParserRuleContext {
		public BinarySelectorContext binarySelector() {
			return getRuleContext(BinarySelectorContext.class,0);
		}
		public KeywordSelectorContext keywordSelector() {
			return getRuleContext(KeywordSelectorContext.class,0);
		}
		public UnarySelectorContext unarySelector() {
			return getRuleContext(UnarySelectorContext.class,0);
		}
		public SelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selector; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterSelector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitSelector(this);
		}
	}

	public final SelectorContext selector() throws RecognitionException {
		SelectorContext _localctx = new SelectorContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_selector);
		try {
			setState(334);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Equal:
			case Or:
			case Comma:
			case Minus:
			case Not:
			case And:
			case Star:
			case Div:
			case Mod:
			case Plus:
			case More:
			case Less:
			case At:
			case Per:
			case OperatorSequence:
				enterOuterAlt(_localctx, 1);
				{
				setState(331);
				binarySelector();
				}
				break;
			case Keyword:
			case KeywordSequence:
				enterOuterAlt(_localctx, 2);
				{
				setState(332);
				keywordSelector();
				}
				break;
			case Primitive:
			case Identifier:
				enterOuterAlt(_localctx, 3);
				{
				setState(333);
				unarySelector();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class KeywordSelectorContext extends ParserRuleContext {
		public TerminalNode Keyword() { return getToken(SOMParser.Keyword, 0); }
		public TerminalNode KeywordSequence() { return getToken(SOMParser.KeywordSequence, 0); }
		public KeywordSelectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keywordSelector; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterKeywordSelector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitKeywordSelector(this);
		}
	}

	public final KeywordSelectorContext keywordSelector() throws RecognitionException {
		KeywordSelectorContext _localctx = new KeywordSelectorContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_keywordSelector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			_la = _input.LA(1);
			if ( !(_la==Keyword || _la==KeywordSequence) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ParserRuleContext {
		public TerminalNode STString() { return getToken(SOMParser.STString, 0); }
		public StringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterString(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitString(this);
		}
	}

	public final StringContext string() throws RecognitionException {
		StringContext _localctx = new StringContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_string);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(338);
			match(STString);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NestedBlockContext extends ParserRuleContext {
		public TerminalNode NewBlock() { return getToken(SOMParser.NewBlock, 0); }
		public TerminalNode EndBlock() { return getToken(SOMParser.EndBlock, 0); }
		public BlockPatternContext blockPattern() {
			return getRuleContext(BlockPatternContext.class,0);
		}
		public BlockContentsContext blockContents() {
			return getRuleContext(BlockContentsContext.class,0);
		}
		public NestedBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nestedBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterNestedBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitNestedBlock(this);
		}
	}

	public final NestedBlockContext nestedBlock() throws RecognitionException {
		NestedBlockContext _localctx = new NestedBlockContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_nestedBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			match(NewBlock);
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(341);
				blockPattern();
				}
			}

			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 20619201176L) != 0)) {
				{
				setState(344);
				blockContents();
				}
			}

			setState(347);
			match(EndBlock);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockPatternContext extends ParserRuleContext {
		public BlockArgumentsContext blockArguments() {
			return getRuleContext(BlockArgumentsContext.class,0);
		}
		public TerminalNode Or() { return getToken(SOMParser.Or, 0); }
		public BlockPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBlockPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBlockPattern(this);
		}
	}

	public final BlockPatternContext blockPattern() throws RecognitionException {
		BlockPatternContext _localctx = new BlockPatternContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_blockPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(349);
			blockArguments();
			setState(350);
			match(Or);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockArgumentsContext extends ParserRuleContext {
		public List<TerminalNode> Colon() { return getTokens(SOMParser.Colon); }
		public TerminalNode Colon(int i) {
			return getToken(SOMParser.Colon, i);
		}
		public List<ArgumentContext> argument() {
			return getRuleContexts(ArgumentContext.class);
		}
		public ArgumentContext argument(int i) {
			return getRuleContext(ArgumentContext.class,i);
		}
		public BlockArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_blockArguments; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).enterBlockArguments(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof SOMListener ) ((SOMListener)listener).exitBlockArguments(this);
		}
	}

	public final BlockArgumentsContext blockArguments() throws RecognitionException {
		BlockArgumentsContext _localctx = new BlockArgumentsContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_blockArguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(354); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(352);
				match(Colon);
				setState(353);
				argument();
				}
				}
				setState(356); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Colon );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001\"\u0167\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0005\u0000f\b\u0000\n\u0000\f\u0000i\t\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0005\u0000n\b\u0000\n\u0000\f\u0000"+
		"q\t\u0000\u0003\u0000s\b\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0003"+
		"\u0001x\b\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0005"+
		"\u0002~\b\u0002\n\u0002\f\u0002\u0081\t\u0002\u0001\u0002\u0003\u0002"+
		"\u0084\b\u0002\u0001\u0003\u0001\u0003\u0005\u0003\u0088\b\u0003\n\u0003"+
		"\f\u0003\u008b\t\u0003\u0001\u0003\u0003\u0003\u008e\b\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u0094\b\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0003\u0005\u0099\b\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0004\b\u00a3"+
		"\b\b\u000b\b\f\b\u00a4\u0001\t\u0001\t\u0003\t\u00a9\b\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0003\u000f\u00bb\b\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0005\u0010"+
		"\u00c0\b\u0010\n\u0010\f\u0010\u00c3\t\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00ca\b\u0011\u0003\u0011\u00cc"+
		"\b\u0011\u0003\u0011\u00ce\b\u0011\u0001\u0012\u0001\u0012\u0003\u0012"+
		"\u00d2\b\u0012\u0001\u0013\u0001\u0013\u0003\u0013\u00d6\b\u0013\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0004\u0015\u00dc\b\u0015\u000b"+
		"\u0015\f\u0015\u00dd\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001"+
		"\u0017\u0003\u0017\u00e5\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0003\u0018\u00eb\b\u0018\u0001\u0019\u0001\u0019\u0001\u001a\u0004"+
		"\u001a\u00f0\b\u001a\u000b\u001a\f\u001a\u00f1\u0001\u001a\u0005\u001a"+
		"\u00f5\b\u001a\n\u001a\f\u001a\u00f8\t\u001a\u0001\u001a\u0003\u001a\u00fb"+
		"\b\u001a\u0001\u001a\u0004\u001a\u00fe\b\u001a\u000b\u001a\f\u001a\u00ff"+
		"\u0001\u001a\u0003\u001a\u0103\b\u001a\u0001\u001a\u0003\u001a\u0106\b"+
		"\u001a\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001d\u0001\u001d\u0005\u001d\u010f\b\u001d\n\u001d\f\u001d\u0112\t\u001d"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0004\u001e\u0117\b\u001e\u000b\u001e"+
		"\f\u001e\u0118\u0001\u001f\u0001\u001f\u0005\u001f\u011d\b\u001f\n\u001f"+
		"\f\u001f\u0120\t\u001f\u0001 \u0001 \u0001 \u0001 \u0001!\u0001!\u0001"+
		"!\u0001!\u0003!\u012a\b!\u0001\"\u0001\"\u0001\"\u0005\"\u012f\b\"\n\""+
		"\f\"\u0132\t\"\u0001\"\u0001\"\u0001#\u0001#\u0003#\u0138\b#\u0001$\u0001"+
		"$\u0003$\u013c\b$\u0001%\u0001%\u0001%\u0001&\u0001&\u0001\'\u0001\'\u0001"+
		"(\u0001(\u0001(\u0003(\u0148\b(\u0001)\u0001)\u0001*\u0001*\u0001*\u0003"+
		"*\u014f\b*\u0001+\u0001+\u0001,\u0001,\u0001-\u0001-\u0003-\u0157\b-\u0001"+
		"-\u0003-\u015a\b-\u0001-\u0001-\u0001.\u0001.\u0001.\u0001/\u0001/\u0004"+
		"/\u0163\b/\u000b/\f/\u0164\u0001/\u0000\u00000\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@BDFHJLNPRTVXZ\\^\u0000\u0003\u0002\u0000\u0005\u0005\t\u0016"+
		"\u0001\u0000\u0003\u0004\u0001\u0000 !\u0165\u0000`\u0001\u0000\u0000"+
		"\u0000\u0002w\u0001\u0000\u0000\u0000\u0004\u0083\u0001\u0000\u0000\u0000"+
		"\u0006\u008d\u0001\u0000\u0000\u0000\b\u008f\u0001\u0000\u0000\u0000\n"+
		"\u0098\u0001\u0000\u0000\u0000\f\u009a\u0001\u0000\u0000\u0000\u000e\u009c"+
		"\u0001\u0000\u0000\u0000\u0010\u00a2\u0001\u0000\u0000\u0000\u0012\u00a6"+
		"\u0001\u0000\u0000\u0000\u0014\u00ac\u0001\u0000\u0000\u0000\u0016\u00ae"+
		"\u0001\u0000\u0000\u0000\u0018\u00b0\u0001\u0000\u0000\u0000\u001a\u00b2"+
		"\u0001\u0000\u0000\u0000\u001c\u00b4\u0001\u0000\u0000\u0000\u001e\u00ba"+
		"\u0001\u0000\u0000\u0000 \u00c1\u0001\u0000\u0000\u0000\"\u00cd\u0001"+
		"\u0000\u0000\u0000$\u00cf\u0001\u0000\u0000\u0000&\u00d5\u0001\u0000\u0000"+
		"\u0000(\u00d7\u0001\u0000\u0000\u0000*\u00db\u0001\u0000\u0000\u0000,"+
		"\u00df\u0001\u0000\u0000\u0000.\u00e2\u0001\u0000\u0000\u00000\u00ea\u0001"+
		"\u0000\u0000\u00002\u00ec\u0001\u0000\u0000\u00004\u0105\u0001\u0000\u0000"+
		"\u00006\u0107\u0001\u0000\u0000\u00008\u0109\u0001\u0000\u0000\u0000:"+
		"\u010c\u0001\u0000\u0000\u0000<\u0116\u0001\u0000\u0000\u0000>\u011a\u0001"+
		"\u0000\u0000\u0000@\u0121\u0001\u0000\u0000\u0000B\u0129\u0001\u0000\u0000"+
		"\u0000D\u012b\u0001\u0000\u0000\u0000F\u0137\u0001\u0000\u0000\u0000H"+
		"\u013b\u0001\u0000\u0000\u0000J\u013d\u0001\u0000\u0000\u0000L\u0140\u0001"+
		"\u0000\u0000\u0000N\u0142\u0001\u0000\u0000\u0000P\u0144\u0001\u0000\u0000"+
		"\u0000R\u0149\u0001\u0000\u0000\u0000T\u014e\u0001\u0000\u0000\u0000V"+
		"\u0150\u0001\u0000\u0000\u0000X\u0152\u0001\u0000\u0000\u0000Z\u0154\u0001"+
		"\u0000\u0000\u0000\\\u015d\u0001\u0000\u0000\u0000^\u0162\u0001\u0000"+
		"\u0000\u0000`a\u0005\u0004\u0000\u0000ab\u0005\u0005\u0000\u0000bc\u0003"+
		"\u0002\u0001\u0000cg\u0003\u0004\u0002\u0000df\u0003\b\u0004\u0000ed\u0001"+
		"\u0000\u0000\u0000fi\u0001\u0000\u0000\u0000ge\u0001\u0000\u0000\u0000"+
		"gh\u0001\u0000\u0000\u0000hr\u0001\u0000\u0000\u0000ig\u0001\u0000\u0000"+
		"\u0000jk\u0005\u0006\u0000\u0000ko\u0003\u0006\u0003\u0000ln\u0003\b\u0004"+
		"\u0000ml\u0001\u0000\u0000\u0000nq\u0001\u0000\u0000\u0000om\u0001\u0000"+
		"\u0000\u0000op\u0001\u0000\u0000\u0000ps\u0001\u0000\u0000\u0000qo\u0001"+
		"\u0000\u0000\u0000rj\u0001\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000"+
		"st\u0001\u0000\u0000\u0000tu\u0005\b\u0000\u0000u\u0001\u0001\u0000\u0000"+
		"\u0000vx\u0005\u0004\u0000\u0000wv\u0001\u0000\u0000\u0000wx\u0001\u0000"+
		"\u0000\u0000xy\u0001\u0000\u0000\u0000yz\u0005\u0007\u0000\u0000z\u0003"+
		"\u0001\u0000\u0000\u0000{\u007f\u0005\t\u0000\u0000|~\u00032\u0019\u0000"+
		"}|\u0001\u0000\u0000\u0000~\u0081\u0001\u0000\u0000\u0000\u007f}\u0001"+
		"\u0000\u0000\u0000\u007f\u0080\u0001\u0000\u0000\u0000\u0080\u0082\u0001"+
		"\u0000\u0000\u0000\u0081\u007f\u0001\u0000\u0000\u0000\u0082\u0084\u0005"+
		"\t\u0000\u0000\u0083{\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000"+
		"\u0000\u0000\u0084\u0005\u0001\u0000\u0000\u0000\u0085\u0089\u0005\t\u0000"+
		"\u0000\u0086\u0088\u00032\u0019\u0000\u0087\u0086\u0001\u0000\u0000\u0000"+
		"\u0088\u008b\u0001\u0000\u0000\u0000\u0089\u0087\u0001\u0000\u0000\u0000"+
		"\u0089\u008a\u0001\u0000\u0000\u0000\u008a\u008c\u0001\u0000\u0000\u0000"+
		"\u008b\u0089\u0001\u0000\u0000\u0000\u008c\u008e\u0005\t\u0000\u0000\u008d"+
		"\u0085\u0001\u0000\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000\u008e"+
		"\u0007\u0001\u0000\u0000\u0000\u008f\u0090\u0003\n\u0005\u0000\u0090\u0093"+
		"\u0005\u0005\u0000\u0000\u0091\u0094\u0005\u0003\u0000\u0000\u0092\u0094"+
		"\u0003\u0012\t\u0000\u0093\u0091\u0001\u0000\u0000\u0000\u0093\u0092\u0001"+
		"\u0000\u0000\u0000\u0094\t\u0001\u0000\u0000\u0000\u0095\u0099\u0003\f"+
		"\u0006\u0000\u0096\u0099\u0003\u0010\b\u0000\u0097\u0099\u0003\u000e\u0007"+
		"\u0000\u0098\u0095\u0001\u0000\u0000\u0000\u0098\u0096\u0001\u0000\u0000"+
		"\u0000\u0098\u0097\u0001\u0000\u0000\u0000\u0099\u000b\u0001\u0000\u0000"+
		"\u0000\u009a\u009b\u0003\u0014\n\u0000\u009b\r\u0001\u0000\u0000\u0000"+
		"\u009c\u009d\u0003\u0016\u000b\u0000\u009d\u009e\u0003\u001c\u000e\u0000"+
		"\u009e\u000f\u0001\u0000\u0000\u0000\u009f\u00a0\u0003\u001a\r\u0000\u00a0"+
		"\u00a1\u0003\u001c\u000e\u0000\u00a1\u00a3\u0001\u0000\u0000\u0000\u00a2"+
		"\u009f\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001\u0000\u0000\u0000\u00a4"+
		"\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000\u0000\u0000\u00a5"+
		"\u0011\u0001\u0000\u0000\u0000\u00a6\u00a8\u0005\u0007\u0000\u0000\u00a7"+
		"\u00a9\u0003\u001e\u000f\u0000\u00a8\u00a7\u0001\u0000\u0000\u0000\u00a8"+
		"\u00a9\u0001\u0000\u0000\u0000\u00a9\u00aa\u0001\u0000\u0000\u0000\u00aa"+
		"\u00ab\u0005\b\u0000\u0000\u00ab\u0013\u0001\u0000\u0000\u0000\u00ac\u00ad"+
		"\u0003\u0018\f\u0000\u00ad\u0015\u0001\u0000\u0000\u0000\u00ae\u00af\u0007"+
		"\u0000\u0000\u0000\u00af\u0017\u0001\u0000\u0000\u0000\u00b0\u00b1\u0007"+
		"\u0001\u0000\u0000\u00b1\u0019\u0001\u0000\u0000\u0000\u00b2\u00b3\u0005"+
		" \u0000\u0000\u00b3\u001b\u0001\u0000\u0000\u0000\u00b4\u00b5\u00032\u0019"+
		"\u0000\u00b5\u001d\u0001\u0000\u0000\u0000\u00b6\u00b7\u0005\t\u0000\u0000"+
		"\u00b7\u00b8\u0003 \u0010\u0000\u00b8\u00b9\u0005\t\u0000\u0000\u00b9"+
		"\u00bb\u0001\u0000\u0000\u0000\u00ba\u00b6\u0001\u0000\u0000\u0000\u00ba"+
		"\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc"+
		"\u00bd\u0003\"\u0011\u0000\u00bd\u001f\u0001\u0000\u0000\u0000\u00be\u00c0"+
		"\u00032\u0019\u0000\u00bf\u00be\u0001\u0000\u0000\u0000\u00c0\u00c3\u0001"+
		"\u0000\u0000\u0000\u00c1\u00bf\u0001\u0000\u0000\u0000\u00c1\u00c2\u0001"+
		"\u0000\u0000\u0000\u00c2!\u0001\u0000\u0000\u0000\u00c3\u00c1\u0001\u0000"+
		"\u0000\u0000\u00c4\u00c5\u0005\u001b\u0000\u0000\u00c5\u00ce\u0003$\u0012"+
		"\u0000\u00c6\u00cb\u0003&\u0013\u0000\u00c7\u00c9\u0005\u001c\u0000\u0000"+
		"\u00c8\u00ca\u0003\"\u0011\u0000\u00c9\u00c8\u0001\u0000\u0000\u0000\u00c9"+
		"\u00ca\u0001\u0000\u0000\u0000\u00ca\u00cc\u0001\u0000\u0000\u0000\u00cb"+
		"\u00c7\u0001\u0000\u0000\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc"+
		"\u00ce\u0001\u0000\u0000\u0000\u00cd\u00c4\u0001\u0000\u0000\u0000\u00cd"+
		"\u00c6\u0001\u0000\u0000\u0000\u00ce#\u0001\u0000\u0000\u0000\u00cf\u00d1"+
		"\u0003&\u0013\u0000\u00d0\u00d2\u0005\u001c\u0000\u0000\u00d1\u00d0\u0001"+
		"\u0000\u0000\u0000\u00d1\u00d2\u0001\u0000\u0000\u0000\u00d2%\u0001\u0000"+
		"\u0000\u0000\u00d3\u00d6\u0003(\u0014\u0000\u00d4\u00d6\u0003.\u0017\u0000"+
		"\u00d5\u00d3\u0001\u0000\u0000\u0000\u00d5\u00d4\u0001\u0000\u0000\u0000"+
		"\u00d6\'\u0001\u0000\u0000\u0000\u00d7\u00d8\u0003*\u0015\u0000\u00d8"+
		"\u00d9\u0003.\u0017\u0000\u00d9)\u0001\u0000\u0000\u0000\u00da\u00dc\u0003"+
		",\u0016\u0000\u00db\u00da\u0001\u0000\u0000\u0000\u00dc\u00dd\u0001\u0000"+
		"\u0000\u0000\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000"+
		"\u0000\u0000\u00de+\u0001\u0000\u0000\u0000\u00df\u00e0\u00032\u0019\u0000"+
		"\u00e0\u00e1\u0005\u001d\u0000\u0000\u00e1-\u0001\u0000\u0000\u0000\u00e2"+
		"\u00e4\u00030\u0018\u0000\u00e3\u00e5\u00034\u001a\u0000\u00e4\u00e3\u0001"+
		"\u0000\u0000\u0000\u00e4\u00e5\u0001\u0000\u0000\u0000\u00e5/\u0001\u0000"+
		"\u0000\u0000\u00e6\u00eb\u00032\u0019\u0000\u00e7\u00eb\u0003@ \u0000"+
		"\u00e8\u00eb\u0003Z-\u0000\u00e9\u00eb\u0003B!\u0000\u00ea\u00e6\u0001"+
		"\u0000\u0000\u0000\u00ea\u00e7\u0001\u0000\u0000\u0000\u00ea\u00e8\u0001"+
		"\u0000\u0000\u0000\u00ea\u00e9\u0001\u0000\u0000\u0000\u00eb1\u0001\u0000"+
		"\u0000\u0000\u00ec\u00ed\u0003\u0018\f\u0000\u00ed3\u0001\u0000\u0000"+
		"\u0000\u00ee\u00f0\u00036\u001b\u0000\u00ef\u00ee\u0001\u0000\u0000\u0000"+
		"\u00f0\u00f1\u0001\u0000\u0000\u0000\u00f1\u00ef\u0001\u0000\u0000\u0000"+
		"\u00f1\u00f2\u0001\u0000\u0000\u0000\u00f2\u00f6\u0001\u0000\u0000\u0000"+
		"\u00f3\u00f5\u00038\u001c\u0000\u00f4\u00f3\u0001\u0000\u0000\u0000\u00f5"+
		"\u00f8\u0001\u0000\u0000\u0000\u00f6\u00f4\u0001\u0000\u0000\u0000\u00f6"+
		"\u00f7\u0001\u0000\u0000\u0000\u00f7\u00fa\u0001\u0000\u0000\u0000\u00f8"+
		"\u00f6\u0001\u0000\u0000\u0000\u00f9\u00fb\u0003<\u001e\u0000\u00fa\u00f9"+
		"\u0001\u0000\u0000\u0000\u00fa\u00fb\u0001\u0000\u0000\u0000\u00fb\u0106"+
		"\u0001\u0000\u0000\u0000\u00fc\u00fe\u00038\u001c\u0000\u00fd\u00fc\u0001"+
		"\u0000\u0000\u0000\u00fe\u00ff\u0001\u0000\u0000\u0000\u00ff\u00fd\u0001"+
		"\u0000\u0000\u0000\u00ff\u0100\u0001\u0000\u0000\u0000\u0100\u0102\u0001"+
		"\u0000\u0000\u0000\u0101\u0103\u0003<\u001e\u0000\u0102\u0101\u0001\u0000"+
		"\u0000\u0000\u0102\u0103\u0001\u0000\u0000\u0000\u0103\u0106\u0001\u0000"+
		"\u0000\u0000\u0104\u0106\u0003<\u001e\u0000\u0105\u00ef\u0001\u0000\u0000"+
		"\u0000\u0105\u00fd\u0001\u0000\u0000\u0000\u0105\u0104\u0001\u0000\u0000"+
		"\u0000\u01065\u0001\u0000\u0000\u0000\u0107\u0108\u0003\u0014\n\u0000"+
		"\u01087\u0001\u0000\u0000\u0000\u0109\u010a\u0003\u0016\u000b\u0000\u010a"+
		"\u010b\u0003:\u001d\u0000\u010b9\u0001\u0000\u0000\u0000\u010c\u0110\u0003"+
		"0\u0018\u0000\u010d\u010f\u00036\u001b\u0000\u010e\u010d\u0001\u0000\u0000"+
		"\u0000\u010f\u0112\u0001\u0000\u0000\u0000\u0110\u010e\u0001\u0000\u0000"+
		"\u0000\u0110\u0111\u0001\u0000\u0000\u0000\u0111;\u0001\u0000\u0000\u0000"+
		"\u0112\u0110\u0001\u0000\u0000\u0000\u0113\u0114\u0003\u001a\r\u0000\u0114"+
		"\u0115\u0003>\u001f\u0000\u0115\u0117\u0001\u0000\u0000\u0000\u0116\u0113"+
		"\u0001\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u0118\u0116"+
		"\u0001\u0000\u0000\u0000\u0118\u0119\u0001\u0000\u0000\u0000\u0119=\u0001"+
		"\u0000\u0000\u0000\u011a\u011e\u0003:\u001d\u0000\u011b\u011d\u00038\u001c"+
		"\u0000\u011c\u011b\u0001\u0000\u0000\u0000\u011d\u0120\u0001\u0000\u0000"+
		"\u0000\u011e\u011c\u0001\u0000\u0000\u0000\u011e\u011f\u0001\u0000\u0000"+
		"\u0000\u011f?\u0001\u0000\u0000\u0000\u0120\u011e\u0001\u0000\u0000\u0000"+
		"\u0121\u0122\u0005\u0007\u0000\u0000\u0122\u0123\u0003&\u0013\u0000\u0123"+
		"\u0124\u0005\b\u0000\u0000\u0124A\u0001\u0000\u0000\u0000\u0125\u012a"+
		"\u0003D\"\u0000\u0126\u012a\u0003P(\u0000\u0127\u012a\u0003R)\u0000\u0128"+
		"\u012a\u0003F#\u0000\u0129\u0125\u0001\u0000\u0000\u0000\u0129\u0126\u0001"+
		"\u0000\u0000\u0000\u0129\u0127\u0001\u0000\u0000\u0000\u0129\u0128\u0001"+
		"\u0000\u0000\u0000\u012aC\u0001\u0000\u0000\u0000\u012b\u012c\u0005\u001a"+
		"\u0000\u0000\u012c\u0130\u0005\u0007\u0000\u0000\u012d\u012f\u0003B!\u0000"+
		"\u012e\u012d\u0001\u0000\u0000\u0000\u012f\u0132\u0001\u0000\u0000\u0000"+
		"\u0130\u012e\u0001\u0000\u0000\u0000\u0130\u0131\u0001\u0000\u0000\u0000"+
		"\u0131\u0133\u0001\u0000\u0000\u0000\u0132\u0130\u0001\u0000\u0000\u0000"+
		"\u0133\u0134\u0005\b\u0000\u0000\u0134E\u0001\u0000\u0000\u0000\u0135"+
		"\u0138\u0003J%\u0000\u0136\u0138\u0003H$\u0000\u0137\u0135\u0001\u0000"+
		"\u0000\u0000\u0137\u0136\u0001\u0000\u0000\u0000\u0138G\u0001\u0000\u0000"+
		"\u0000\u0139\u013c\u0003L&\u0000\u013a\u013c\u0003N\'\u0000\u013b\u0139"+
		"\u0001\u0000\u0000\u0000\u013b\u013a\u0001\u0000\u0000\u0000\u013cI\u0001"+
		"\u0000\u0000\u0000\u013d\u013e\u0005\u000b\u0000\u0000\u013e\u013f\u0003"+
		"H$\u0000\u013fK\u0001\u0000\u0000\u0000\u0140\u0141\u0005\u001e\u0000"+
		"\u0000\u0141M\u0001\u0000\u0000\u0000\u0142\u0143\u0005\u001f\u0000\u0000"+
		"\u0143O\u0001\u0000\u0000\u0000\u0144\u0147\u0005\u001a\u0000\u0000\u0145"+
		"\u0148\u0003X,\u0000\u0146\u0148\u0003T*\u0000\u0147\u0145\u0001\u0000"+
		"\u0000\u0000\u0147\u0146\u0001\u0000\u0000\u0000\u0148Q\u0001\u0000\u0000"+
		"\u0000\u0149\u014a\u0003X,\u0000\u014aS\u0001\u0000\u0000\u0000\u014b"+
		"\u014f\u0003\u0016\u000b\u0000\u014c\u014f\u0003V+\u0000\u014d\u014f\u0003"+
		"\u0014\n\u0000\u014e\u014b\u0001\u0000\u0000\u0000\u014e\u014c\u0001\u0000"+
		"\u0000\u0000\u014e\u014d\u0001\u0000\u0000\u0000\u014fU\u0001\u0000\u0000"+
		"\u0000\u0150\u0151\u0007\u0002\u0000\u0000\u0151W\u0001\u0000\u0000\u0000"+
		"\u0152\u0153\u0005\"\u0000\u0000\u0153Y\u0001\u0000\u0000\u0000\u0154"+
		"\u0156\u0005\u0018\u0000\u0000\u0155\u0157\u0003\\.\u0000\u0156\u0155"+
		"\u0001\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157\u0159"+
		"\u0001\u0000\u0000\u0000\u0158\u015a\u0003\u001e\u000f\u0000\u0159\u0158"+
		"\u0001\u0000\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015a\u015b"+
		"\u0001\u0000\u0000\u0000\u015b\u015c\u0005\u0019\u0000\u0000\u015c[\u0001"+
		"\u0000\u0000\u0000\u015d\u015e\u0003^/\u0000\u015e\u015f\u0005\t\u0000"+
		"\u0000\u015f]\u0001\u0000\u0000\u0000\u0160\u0161\u0005\u0017\u0000\u0000"+
		"\u0161\u0163\u0003\u001c\u000e\u0000\u0162\u0160\u0001\u0000\u0000\u0000"+
		"\u0163\u0164\u0001\u0000\u0000\u0000\u0164\u0162\u0001\u0000\u0000\u0000"+
		"\u0164\u0165\u0001\u0000\u0000\u0000\u0165_\u0001\u0000\u0000\u0000(g"+
		"orw\u007f\u0083\u0089\u008d\u0093\u0098\u00a4\u00a8\u00ba\u00c1\u00c9"+
		"\u00cb\u00cd\u00d1\u00d5\u00dd\u00e4\u00ea\u00f1\u00f6\u00fa\u00ff\u0102"+
		"\u0105\u0110\u0118\u011e\u0129\u0130\u0137\u013b\u0147\u014e\u0156\u0159"+
		"\u0164";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}