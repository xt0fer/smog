// Generated from /Volumes/Terabyte/kristofer/LocalProjects/smog/OldCode/SomSom/bin/specification/SOM.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SOMParser}.
 */
public interface SOMListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SOMParser#classdef}.
	 * @param ctx the parse tree
	 */
	void enterClassdef(SOMParser.ClassdefContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#classdef}.
	 * @param ctx the parse tree
	 */
	void exitClassdef(SOMParser.ClassdefContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#superclass}.
	 * @param ctx the parse tree
	 */
	void enterSuperclass(SOMParser.SuperclassContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#superclass}.
	 * @param ctx the parse tree
	 */
	void exitSuperclass(SOMParser.SuperclassContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#instanceFields}.
	 * @param ctx the parse tree
	 */
	void enterInstanceFields(SOMParser.InstanceFieldsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#instanceFields}.
	 * @param ctx the parse tree
	 */
	void exitInstanceFields(SOMParser.InstanceFieldsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#classFields}.
	 * @param ctx the parse tree
	 */
	void enterClassFields(SOMParser.ClassFieldsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#classFields}.
	 * @param ctx the parse tree
	 */
	void exitClassFields(SOMParser.ClassFieldsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#method}.
	 * @param ctx the parse tree
	 */
	void enterMethod(SOMParser.MethodContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#method}.
	 * @param ctx the parse tree
	 */
	void exitMethod(SOMParser.MethodContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#pattern}.
	 * @param ctx the parse tree
	 */
	void enterPattern(SOMParser.PatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#pattern}.
	 * @param ctx the parse tree
	 */
	void exitPattern(SOMParser.PatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#unaryPattern}.
	 * @param ctx the parse tree
	 */
	void enterUnaryPattern(SOMParser.UnaryPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#unaryPattern}.
	 * @param ctx the parse tree
	 */
	void exitUnaryPattern(SOMParser.UnaryPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#binaryPattern}.
	 * @param ctx the parse tree
	 */
	void enterBinaryPattern(SOMParser.BinaryPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#binaryPattern}.
	 * @param ctx the parse tree
	 */
	void exitBinaryPattern(SOMParser.BinaryPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#keywordPattern}.
	 * @param ctx the parse tree
	 */
	void enterKeywordPattern(SOMParser.KeywordPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#keywordPattern}.
	 * @param ctx the parse tree
	 */
	void exitKeywordPattern(SOMParser.KeywordPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#methodBlock}.
	 * @param ctx the parse tree
	 */
	void enterMethodBlock(SOMParser.MethodBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#methodBlock}.
	 * @param ctx the parse tree
	 */
	void exitMethodBlock(SOMParser.MethodBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#unarySelector}.
	 * @param ctx the parse tree
	 */
	void enterUnarySelector(SOMParser.UnarySelectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#unarySelector}.
	 * @param ctx the parse tree
	 */
	void exitUnarySelector(SOMParser.UnarySelectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#binarySelector}.
	 * @param ctx the parse tree
	 */
	void enterBinarySelector(SOMParser.BinarySelectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#binarySelector}.
	 * @param ctx the parse tree
	 */
	void exitBinarySelector(SOMParser.BinarySelectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(SOMParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(SOMParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#keyword}.
	 * @param ctx the parse tree
	 */
	void enterKeyword(SOMParser.KeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#keyword}.
	 * @param ctx the parse tree
	 */
	void exitKeyword(SOMParser.KeywordContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#argument}.
	 * @param ctx the parse tree
	 */
	void enterArgument(SOMParser.ArgumentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#argument}.
	 * @param ctx the parse tree
	 */
	void exitArgument(SOMParser.ArgumentContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#blockContents}.
	 * @param ctx the parse tree
	 */
	void enterBlockContents(SOMParser.BlockContentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#blockContents}.
	 * @param ctx the parse tree
	 */
	void exitBlockContents(SOMParser.BlockContentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#localDefs}.
	 * @param ctx the parse tree
	 */
	void enterLocalDefs(SOMParser.LocalDefsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#localDefs}.
	 * @param ctx the parse tree
	 */
	void exitLocalDefs(SOMParser.LocalDefsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#blockBody}.
	 * @param ctx the parse tree
	 */
	void enterBlockBody(SOMParser.BlockBodyContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#blockBody}.
	 * @param ctx the parse tree
	 */
	void exitBlockBody(SOMParser.BlockBodyContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#result}.
	 * @param ctx the parse tree
	 */
	void enterResult(SOMParser.ResultContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#result}.
	 * @param ctx the parse tree
	 */
	void exitResult(SOMParser.ResultContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(SOMParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(SOMParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#assignation}.
	 * @param ctx the parse tree
	 */
	void enterAssignation(SOMParser.AssignationContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#assignation}.
	 * @param ctx the parse tree
	 */
	void exitAssignation(SOMParser.AssignationContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#assignments}.
	 * @param ctx the parse tree
	 */
	void enterAssignments(SOMParser.AssignmentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#assignments}.
	 * @param ctx the parse tree
	 */
	void exitAssignments(SOMParser.AssignmentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(SOMParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(SOMParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#evaluation}.
	 * @param ctx the parse tree
	 */
	void enterEvaluation(SOMParser.EvaluationContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#evaluation}.
	 * @param ctx the parse tree
	 */
	void exitEvaluation(SOMParser.EvaluationContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#primary}.
	 * @param ctx the parse tree
	 */
	void enterPrimary(SOMParser.PrimaryContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#primary}.
	 * @param ctx the parse tree
	 */
	void exitPrimary(SOMParser.PrimaryContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#variable}.
	 * @param ctx the parse tree
	 */
	void enterVariable(SOMParser.VariableContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#variable}.
	 * @param ctx the parse tree
	 */
	void exitVariable(SOMParser.VariableContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#messages}.
	 * @param ctx the parse tree
	 */
	void enterMessages(SOMParser.MessagesContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#messages}.
	 * @param ctx the parse tree
	 */
	void exitMessages(SOMParser.MessagesContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#unaryMessage}.
	 * @param ctx the parse tree
	 */
	void enterUnaryMessage(SOMParser.UnaryMessageContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#unaryMessage}.
	 * @param ctx the parse tree
	 */
	void exitUnaryMessage(SOMParser.UnaryMessageContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#binaryMessage}.
	 * @param ctx the parse tree
	 */
	void enterBinaryMessage(SOMParser.BinaryMessageContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#binaryMessage}.
	 * @param ctx the parse tree
	 */
	void exitBinaryMessage(SOMParser.BinaryMessageContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#binaryOperand}.
	 * @param ctx the parse tree
	 */
	void enterBinaryOperand(SOMParser.BinaryOperandContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#binaryOperand}.
	 * @param ctx the parse tree
	 */
	void exitBinaryOperand(SOMParser.BinaryOperandContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#keywordMessage}.
	 * @param ctx the parse tree
	 */
	void enterKeywordMessage(SOMParser.KeywordMessageContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#keywordMessage}.
	 * @param ctx the parse tree
	 */
	void exitKeywordMessage(SOMParser.KeywordMessageContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#formula}.
	 * @param ctx the parse tree
	 */
	void enterFormula(SOMParser.FormulaContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#formula}.
	 * @param ctx the parse tree
	 */
	void exitFormula(SOMParser.FormulaContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#nestedTerm}.
	 * @param ctx the parse tree
	 */
	void enterNestedTerm(SOMParser.NestedTermContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#nestedTerm}.
	 * @param ctx the parse tree
	 */
	void exitNestedTerm(SOMParser.NestedTermContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(SOMParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(SOMParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalArray}.
	 * @param ctx the parse tree
	 */
	void enterLiteralArray(SOMParser.LiteralArrayContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalArray}.
	 * @param ctx the parse tree
	 */
	void exitLiteralArray(SOMParser.LiteralArrayContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalNumber}.
	 * @param ctx the parse tree
	 */
	void enterLiteralNumber(SOMParser.LiteralNumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalNumber}.
	 * @param ctx the parse tree
	 */
	void exitLiteralNumber(SOMParser.LiteralNumberContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalDecimal}.
	 * @param ctx the parse tree
	 */
	void enterLiteralDecimal(SOMParser.LiteralDecimalContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalDecimal}.
	 * @param ctx the parse tree
	 */
	void exitLiteralDecimal(SOMParser.LiteralDecimalContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#negativeDecimal}.
	 * @param ctx the parse tree
	 */
	void enterNegativeDecimal(SOMParser.NegativeDecimalContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#negativeDecimal}.
	 * @param ctx the parse tree
	 */
	void exitNegativeDecimal(SOMParser.NegativeDecimalContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalInteger}.
	 * @param ctx the parse tree
	 */
	void enterLiteralInteger(SOMParser.LiteralIntegerContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalInteger}.
	 * @param ctx the parse tree
	 */
	void exitLiteralInteger(SOMParser.LiteralIntegerContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalDouble}.
	 * @param ctx the parse tree
	 */
	void enterLiteralDouble(SOMParser.LiteralDoubleContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalDouble}.
	 * @param ctx the parse tree
	 */
	void exitLiteralDouble(SOMParser.LiteralDoubleContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalSymbol}.
	 * @param ctx the parse tree
	 */
	void enterLiteralSymbol(SOMParser.LiteralSymbolContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalSymbol}.
	 * @param ctx the parse tree
	 */
	void exitLiteralSymbol(SOMParser.LiteralSymbolContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#literalString}.
	 * @param ctx the parse tree
	 */
	void enterLiteralString(SOMParser.LiteralStringContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#literalString}.
	 * @param ctx the parse tree
	 */
	void exitLiteralString(SOMParser.LiteralStringContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#selector}.
	 * @param ctx the parse tree
	 */
	void enterSelector(SOMParser.SelectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#selector}.
	 * @param ctx the parse tree
	 */
	void exitSelector(SOMParser.SelectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#keywordSelector}.
	 * @param ctx the parse tree
	 */
	void enterKeywordSelector(SOMParser.KeywordSelectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#keywordSelector}.
	 * @param ctx the parse tree
	 */
	void exitKeywordSelector(SOMParser.KeywordSelectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#string}.
	 * @param ctx the parse tree
	 */
	void enterString(SOMParser.StringContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#string}.
	 * @param ctx the parse tree
	 */
	void exitString(SOMParser.StringContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#nestedBlock}.
	 * @param ctx the parse tree
	 */
	void enterNestedBlock(SOMParser.NestedBlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#nestedBlock}.
	 * @param ctx the parse tree
	 */
	void exitNestedBlock(SOMParser.NestedBlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#blockPattern}.
	 * @param ctx the parse tree
	 */
	void enterBlockPattern(SOMParser.BlockPatternContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#blockPattern}.
	 * @param ctx the parse tree
	 */
	void exitBlockPattern(SOMParser.BlockPatternContext ctx);
	/**
	 * Enter a parse tree produced by {@link SOMParser#blockArguments}.
	 * @param ctx the parse tree
	 */
	void enterBlockArguments(SOMParser.BlockArgumentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SOMParser#blockArguments}.
	 * @param ctx the parse tree
	 */
	void exitBlockArguments(SOMParser.BlockArgumentsContext ctx);
}