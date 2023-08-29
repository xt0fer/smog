//
//  Parser.swift
//  smog
//
//  Created by Kristofer Younger on 8/28/23.
//

import Foundation

class Parser {
    
    //Parser = (
    //  | lexer sym text nextSym filename cgenc universe bcGen |
    var lexer: Lexer
    var     sym: LexerToken
    var     text: String
    var     nextSym: LexerToken
    var     filename: String
    var     cgenc: ClassGenerationContext
    var     universe: Universe
    var     bcGen: BytecodeGenerator.Type
    
    //  initializeWith: aString for: aFilename in: aUniverse = (
    //    filename := aFilename.
    //    lexer := Lexer new: aString.
    //    universe := aUniverse.
    //    cgenc := ClassGenerationContext new: universe.
    //    self takeSymbolFromLexer.
    
    //    "This is just a convenient abbreviation."
    //    bcGen := BytecodeGenerator.
    //  )
    init(_ source: String, for fileName: String, in universe: Universe) {
        self.filename = fileName
        self.lexer = Lexer(source)
        self.universe = universe
        self.cgenc = ClassGenerationContext(universe)
        self.takeSymbolFromLexer()
        self.bcGen = BytecodeGenerator.self
    }
    
    //  takeSymbolFromLexer = (
    //    sym := lexer sym.
    //    text := lexer text.
    //    nextSym := #none.
    //  )
    func takeSymbolFromLexer() {
        self.sym = lexer.sym
        self.text = lexer.text
        self.nextSym = .none
    }
    
    //  classdef = (
    //    cgenc name: (universe symbolFor: text).
    //    self expect: #identifier.
    //    self expect: #equal.
    
    //    self superclass.
    
    //    self expect: #newTerm.
    //    self classBody.
    
    //    (self accept: #separator) ifTrue: [
    //      cgenc startClassSide.
    //      self classBody ].
    
    //    self expect: #endTerm.
    //    ^ cgenc
    //  )
    func classDef() -> ClassGenerationContext {
        self.cgenc.name = universe.symbolFor(text)
        self.expect(.identifier)
        self.expect(.equal)
        
        self.superclass()
        
        self.expect(.newTerm)
        self.classBody()
        
        if self.accept(.separator) {
            cgenc.startClassSide()
            self.classBody()
        }
        self.expect(.endTerm)
        return cgenc
    }
    
    //  classBody = (
    //    self fields.
    //    [self symIsMethod] whileTrue: [
    //       | mgenc |
    //       mgenc := MethodGenerationContext new: cgenc.
    //       mgenc addArgument: 'self'.
    //       self method: mgenc.
    //       cgenc addMethod: (mgenc assemble: universe) ].
    //  )
    func classDef() {
        self.fields()
        while self.symIsMethod() {
            var mgenc = MethodGenerationContext(cgenc)
            mgenc.addArgument("self")
            self.method(mgenc)
            cgenc.addMethod(mgenc.assemble(universe))
        }
    }
    //  superclass = (
    //    | superName |
    //    sym == #identifier
    //      ifTrue: [
    //        superName := universe symbolFor: text.
    //        self accept: #identifier ]
    //      ifFalse: [
    //        superName := universe symbolFor: 'Object' ].
    
    //    cgenc superName: superName.
    
    //    superName string = 'nil' ifFalse: [
    //      self initializeFromSuperClass: superName ].
    //  )
    
    //  initializeFromSuperClass: superName = (
    //    | superClass |
    //    superClass := universe loadClass: superName.
    //    superClass == nil ifTrue: [
    //      self error: 'Was not able to load super class: ' + superName string + ' in ' + filename ].
    //    cgenc instanceFieldsOfSuper: superClass instanceFields.
    //    cgenc classFieldsOfSuper: superClass somClass instanceFields.
    //  )
    
    //  fields = (
    //    (self accept: #or) ifTrue: [
    //      [sym == #identifier] whileTrue: [
    //        | var |
    //        var := self variable.
    //        cgenc addField: (universe symbolFor: var) ].
    //      self expect: #or ]
    //  )
    
    //  method: mgenc = (
    //    self pattern: mgenc.
    //    self expect: #equal.
    
    //    sym == #primitive
    //      ifTrue: [
    //        mgenc markAsPrimitive.
    //        self primBlock ]
    //      ifFalse: [
    //        self methodBlock: mgenc ]
    //  )
    
    //  primBlock = (
    //    self expect: #primitive
    //  )
    
    //  pattern: mgenc = (
    //    sym == #identifier ifTrue: [
    //      ^ self unaryPattern: mgenc ].
    //    sym == #keyword ifTrue: [
    //      ^ self keywordPattern: mgenc ].
    //    self binaryPattern: mgenc
    //  )
    
    //  unaryPattern: mgenc = (
    //    mgenc signature: self unarySelector
    //  )
    
    //  binaryPattern: mgenc = (
    //    mgenc signature: self binarySelector.
    //    mgenc addArgumentIfAbsent: self argument
    //  )
    
    //  keywordPattern: mgenc = (
    //    | kw |
    //    kw := ''.
    
    //    [sym == #keyword] whileTrue: [
    //      kw := kw + self keyword.
    //      mgenc addArgumentIfAbsent: self argument ].
    
    //    mgenc signature: (universe symbolFor: kw)
    //  )
    
    //  methodBlock: mgenc = (
    //    self expect: #newTerm.
    
    //    self blockContents: mgenc.
    
    //    " if no return has been generated so far, we can be sure there was no . (dot)
    //      terminating the last expression, so the last expression's value must be
    //      popped off the stack and a ^self be generated "
    //    mgenc isFinished ifFalse: [
    //      bcGen emitPop: mgenc.
    //      bcGen emit: mgenc pushArgument: 1 in: 0.
    //      bcGen emitReturnLocal: mgenc.
    //      mgenc markAsFinished ].
    
    //    self expect: #endTerm.
    //  )
    
    //  blockContents: mgenc = (
    //    (self accept: #or) ifTrue: [
    //      self locals: mgenc.
    //      self expect: #or ].
    //    self blockBody: mgenc sawPeriod: false
    //  )
    
    //  locals: mgenc = (
    //    [sym == #identifier] whileTrue: [
    //      mgenc addLocalIfAbsent: self variable ]
    //  )
    
    //  blockBody: mgenc sawPeriod: seenPeriod = (
    //    (self accept: #exit) ifTrue: [
    //      ^ self result: mgenc ].
    
    //    sym == #endBlock ifTrue: [
    //      seenPeriod ifTrue: [
    //        "a POP has been generated which must be elided (blocks always
    //         return the value of the last expression, regardless of
    //         whether it was terminated with a . or not)"
    //        mgenc removeLastBytecode ].
    
    //      (mgenc isBlockMethod and: [ mgenc hasBytecodes not ]) ifTrue: [
    //        | nilSym |
    //        "if the block is empty, we need to return nil"
    //        nilSym := universe symbolFor: 'nil'.
    //        bcGen emit: mgenc pushGlobal: nilSym. ].
    
    //      bcGen emitReturnLocal: mgenc.
    //      mgenc markAsFinished.
    //      ^ self ].
    
    //    sym == #endTerm ifTrue: [
    //      "it does not matter whether a period has been seen, as the end of
    //       the method has been found (EndTerm) - so it is safe to emit a
    //       'return self'"
    //      bcGen emit: mgenc pushArgument: 1 in: 0.
    //      bcGen emitReturnLocal: mgenc.
    //      mgenc markAsFinished.
    //      ^ self ].
    
    //    self expression: mgenc.
    //    (self accept: #period) ifTrue: [
    //      bcGen emitPop: mgenc.
    //      self blockBody: mgenc sawPeriod: true ]
    //  )
    
    //  unarySelector = (
    //    ^ universe symbolFor: self identifier
    //  )
    func unarySelector() -> SSymbol {
        return universe.symbolFor(self.identifier())
    }
    
    //  binarySelector = (
    //    | s |
    //    s := text.
    
    //    (self accept: #operatorSequence) or: [
    //    (self acceptOneOf: Parser singleOpSyms) or: [
    //        self expect: #none ] ].
    
    //    ^ universe symbolFor: s
    //  )
    func binarySelector() -> SSymbol {
        let s = self.text
        if self.accept(.operatorSequence) ||
            self.acceptOneOf(Parser.singleOpSyms) ||
            self.expect(.none) { }
        return universe.symbolFor(s)
    }

    
    //  variable = (
    //    ^ self identifier
    //  )
    func variable() -> String {
        return self.identifier()
    }
    
    //  argument = (
    //    ^ self variable
    //  )
    func argument() -> String {
        return self.variable()
    }
    
    //  identifier = (
    //    | s |
    //    s := text.
    //    (self accept: #primitive)
    //      ifFalse: [self expect: #identifier].
    //    ^ s
    //  )
    func identifier() -> String {
        if self.accept(.primitive) {
            self.expect(.identifier)
        }
        return text
    }

    //  keyword = (
    //    | s |
    //    s := text.
    //    self expect: #keyword.
    //    ^ s
    //  )
    func keyword() -> String {
        self.expect(.keyword)
        return text
    }

    //  string = (
    //    | s |
    //    s := text.
    //    self expect: #string.
    //    ^ s
    //  )
    func string() -> String {
        self.expect(.string)
        return text
    }
    //  selector = (
    //    (sym == #operatorSequence or: [self symIn: Parser singleOpSyms])
    //      ifTrue: [^ self binarySelector].
    //    (sym == #keyword or: [sym == #keywordSequence])
    //      ifTrue: [^ self keywordSelector].
    
    //    ^ self unarySelector
    //  )
    func selector() -> SSymbol {
        if (sym == .operatorSequence ||
            self.symIn(Parser.singleOpSyms)) {
            return self.binarySelector()
        } else {
            return self.keywordSelector()
        }
        return self.unarySelector()
    }
    //  keywordSelector = (
    //    | s |
    //    s := text.
    //    self expectOneOf: Parser keywordSelectorSyms.
    //    ^ universe symbolFor: s
    //  )
    func keywordSelector() -> SSymbol {
        let s = text
        self.expectOneOf(Parser.keywordSelectorSyms)
        return universe.symbolFor(s)
    }
    //  result: mgenc = (
    //    self expression: mgenc.
    
    //    mgenc isBlockMethod
    //      ifTrue: [bcGen emitReturnNonLocal: mgenc ]
    //      ifFalse: [bcGen emitReturnLocal: mgenc ].
    //    mgenc markAsFinished.
    
    //    self accept: #period
    //  )
    func result(_ mgenc: MethodGenerationContext) {
        self.expression(mgenc)
        if mgenc.isBlockMethod() {
            bcGen.emitReturnNonLocal(mgenc)
        } else {
            bcGen.emitReturnLocal(mgenc)
        }
        mgenc.markAsFinished()
        
        self.accept(.period)
    }
    
    //  expression: mgenc = (
    //    self peekForNextSymbolFromLexer.
    
    //    nextSym == #assign
    //      ifTrue: [self assignation: mgenc]
    //      ifFalse: [self evaluation: mgenc]
    //  )
    func expression(_ mgenc: MethodGenerationContext) {
        self.peekForNextSymbolFromLexer()
        if nextSym == .assign {
            self.assignation(mgenc)
        } else {
            self.evaluation(mgenc)
        }
    }
    
    //  assignation: mgenc = (
    //    | variables |
    //    variables := Vector new.
    
    //    self assignments: mgenc to: variables.
    //    self evaluation: mgenc.
    
    //    variables do: [:v | bcGen emitDup: mgenc ].
    //    variables do: [:v | self gen: mgenc popVariable: v ]
    //  )
    func assignation(_ mgenc: MethodGenerationContext) {
        var variables: [String] = []
        self.assignments(mgenc, to: variables)
        self.evaluation(mgenc)
        
        for v in variables {
            bcGen.emitDup(mgenc)
        }
        for v in variables {
            self.gen(mgenc, popVariable: v)
        }
    }
    
    //  assignments: mgenc to: variables = (
    //    sym == #identifier ifTrue: [
    //      variables append: (self assignment: mgenc).
    //      self peekForNextSymbolFromLexer.
    //      nextSym == #assign ifTrue: [
    //        self assignments: mgenc to: variables ] ]
    //  )
    func assignments(_ mgenc: MethodGenerationContext, to variables: [String]) {
        if sym == .identifier {
            variables.append(self.assignment(mgenc))
            self.peekForNextSymbolFromLexer()
            if nextSym == .assign {
                self.assignments(mgenc, to: variables)
            }
        }
    }
    
    //  assignment: mgenc = (
    //    | v |
    //    v := self variable.
    //    self expect: #assign.
    //    ^ v
    //  )
    func assignment(_ mgenc: MethodGenerationContext) -> String {
        let v = self.variable()
        self.expect(.assign)
        return v
    }
    
    //  evaluation: mgenc = (
    //    | superSend |
    //    superSend := self primary: mgenc.
    //    self symIsMethod ifTrue: [
    //      self messages: mgenc with: superSend ]
    //  )
    func evaluation(_ mgenc: MethodGenerationContext) {
        let superSend = self.primary(mgenc)
        if self.symIsMethod() {
            self.messages(mgenc, with: superSend)
        }
    }
    
    //  primary: mgenc = (
    //    | superSend |
    //    superSend := false.
    
    //    sym == #identifier ifTrue: [
    //      | v |
    //      v := self variable.
    //      v = 'super' ifTrue: [
    //        superSend := true.
    //        " sends to super, but pushes self as receiver"
    //        v := 'self' ].
    
    //      self gen: mgenc pushVariable: v.
    //      ^ superSend ].
    
    //    sym == #newTerm ifTrue: [
    //      self nestedTerm: mgenc.
    //      ^ superSend ].
    
    //    sym == #newBlock ifTrue: [
    //      | bgenc blockMethod |
    //      bgenc := MethodGenerationContext new: mgenc holder with: mgenc.
    //      bgenc markAsBlockMethod.
    
    //      self nestedBlock: bgenc.
    
    //      blockMethod := bgenc assembleMethod: universe.
    //      bcGen emit: mgenc pushBlock: blockMethod.
    //      ^ superSend ].
    
    //    self literal: mgenc.
    //    ^ superSend
    //  )
    func primary(_ mgenc: MethodGenerationContext) -> Bool {
        var superSend = false
        if sym == .identifier {
            var v = self.variable()
            if v == "super" {
                superSend = true
                v = "self"
            }
            self.gen(mgenc, pushVariable: v)
            return superSend
        }
        if sym == .newTerm {
            self.nestedTerm(mgenc)
            return superSend
        }
        if sym == .newBlock {
            var bgenc = MethodGenerationContext(mgenc.holder(), with: mgenc)
            bgenc.markAsBlockMethod()
            self.nestedBlock(bgenc)
            var blockMethod = bgenc.assembleMethod(universe)
            bcGen.emit(mgenc, pushBlock: blockMethod)
            return superSend
        }
        self.literal(mgenc)
        return superSend
    }
    
    //  messages: mgenc with: superSend = (
    //    sym == #identifier ifTrue: [
    //      "only the first message in a sequence can be a super send"
    //      self unaryMessage: mgenc with: superSend.
    
    //      [sym == #identifier] whileTrue: [
    //        self unaryMessage: mgenc with: false ].
    
    //      [sym == #operatorSequence or: [self symIn: Parser binaryOpSyms]] whileTrue: [
    //        self binaryMessage: mgenc with: false ].
    
    //      sym == #keyword ifTrue: [
    //        self keywordMessage: mgenc with: false ].
    
    //      ^ self ].
    
    //    (sym == #operatorSequence or: [self symIn: Parser binaryOpSyms]) ifTrue: [
    //      self binaryMessage: mgenc with: superSend.
    
    //      [sym == #operatorSequence or: [self symIn: Parser binaryOpSyms]] whileTrue: [
    //        self binaryMessage: mgenc with: false ].
    
    //      sym == #keyword ifTrue: [
    //        self keywordMessage: mgenc with: false ].
    
    //      ^ self ].
    
    //    self keywordMessage: mgenc with: superSend
    //  )
    func messages(_ mgenc: MethodGenerationContext, with superSend: Bool) {
        if sym == .identifier {
            self.unaryMessage(mgenc, with: superSend)
            while sym == .identifier {
                self.unaryMessage(mgenc, with: false)
            }
            while sym == .operatorSequence ||
                    self.symIn(Parser.binaryOpSyms) {
                self.binaryMessage(mgenc, with: false)
            }
            if sym == .keyword {
                self.keywordMessage(mgenc, with: false)
            }
            return self //??
        }
        if sym == .operatorSequence ||
            symIn(Parser.binaryOpSyms) {
            self.binaryMessage(mgenc, with: superSend)
            if sym == .operatorSequence ||
                symIn(Parser.binaryOpSyms) {
                self.binaryMessage(mgenc, with: false)
            }
            if sym == .keyword {
                self.keywordMessage(mgenc, with: false)
            }
        }
        self.keywordMessage(mgenc, with: superSend)
    }
    
    //  unaryMessage: mgenc with: superSend = (
    //    | msg |
    //    msg := self unarySelector.
    
    //    superSend ifTrue: [ bcGen emit: mgenc superSend: msg ]
    //              ifFalse: [ bcGen emit: mgenc send: msg ]
    //  )
    func unaryMessage(_ mgenc: MethodGenerationContext, with superSend: Bool) {
        let msg = self.unarySelector()
        if superSend {
            bcGen.emit(mgenc, superSend: msg)
        } else {
            bcGen.emit(mgenc, send: msg)
        }
    }
    //  binaryMessage: mgenc with: superSend = (
    //    | msg |
    //    msg := self binarySelector.
    //    self binaryOperand: mgenc.

    //    superSend ifTrue: [ bcGen emit: mgenc superSend: msg ]
    //              ifFalse: [ bcGen emit: mgenc send: msg ]
    //  )
    func binaryMessage(_ mgenc: MethodGenerationContext, with superSend: Bool) {
        let msg = self.binarySelector()
        self.binaryOperand(mgenc)
        
        if superSend {
            bcGen.emit(mgenc, superSend: msg)
        } else {
            bcGen.emit(mgenc, send: msg)
        }
    }

    //  binaryOperand: mgenc = (
    //    | superSend |
    //    superSend := self primary: mgenc.
    
    //    [sym == #identifier] whileTrue: [
    //      self unaryMessage: mgenc with: superSend.
    //      superSend := false ].
    
    //    ^ superSend
    //  )
    func binaryOperand(_ mgenc: MethodGenerationContext) -> Bool {
        var superSend = self.primary(mgenc)
        while sym == .identifier {
            self.unaryMessage(mgenc, with: superSend)
            superSend = false
        }
        return superSend
    }

    //  keywordMessage: mgenc with: superSend = (
    //    | kw msg |
    //    kw := self keyword.
    //    self formula: mgenc.
    
    //    [sym == #keyword] whileTrue: [
    //      kw := kw + self keyword.
    //      self formula: mgenc ].
    
    //    msg := universe symbolFor: kw.
    //    superSend ifTrue: [ bcGen emit: mgenc superSend: msg ]
    //              ifFalse: [ bcGen emit: mgenc send: msg ]
    //  )
    func keywordMessage(_ mgenc: MethodGenerationContext, with superSend: Bool) {
        var kw = self.keyword()
        self.formula(mgenc)
        
        while sym == .keyword {
            kw += self.keyword()
            self.fomula(mgenc)
        }
        msg = universe.symbolFor(kw)
        if superSend {
            bcGen.emit(mgenc, superSend: msg)
        } else {
            bcGen.emit(mgenc, send: msg)
        }
    }

    //  formula: mgenc = (
    //    | superSend |
    //    superSend := self binaryOperand: mgenc.
    
    //    "only the first message in a sequence can be a super send"
    //    [sym == #operatorSequence or: [self symIn: Parser binaryOpSyms]] whileTrue: [
    //        self binaryMessage: mgenc with: superSend.
    //        superSend := false ].
    //  )
    
    //  nestedTerm: mgenc = (
    //    self expect: #newTerm.
    //    self expression: mgenc.
    //    self expect: #endTerm.
    //  )
    
    //  nestedBlock: mgenc = (
    //    | blockSig argSize |
    //    mgenc addArgumentIfAbsent: '$block self'.
    
    //    self expect: #newBlock.
    
    //    sym == #colon ifTrue: [
    //      self blockPattern: mgenc ].
    
    //    "generate block signature"
    //    blockSig := '$block method'.
    //    argSize := mgenc numberOfArguments.
    //    (argSize - 1) timesRepeat: [
    //      blockSig := blockSig + ':' ].
    
    //    mgenc signature: (universe symbolFor: blockSig).
    
    //    self blockContents: mgenc.
    
    //    "if no return has been generated, we can be sure that the last expression
    //     in the block was not terminated by ., and can generate a return"
    //    mgenc isFinished ifFalse: [
    //      bcGen emitReturnLocal: mgenc.
    //      mgenc markAsFinished ].
    
    //    self expect: #endBlock
    //  )
    
    //  blockPattern: mgenc = (
    //    self blockArguments: mgenc.
    //    self expect: #or.
    //  )
    
    //  blockArguments: mgenc = (
    //    self expect: #colon.
    //    mgenc addArgumentIfAbsent: self argument.
    
    //    [sym == #colon] whileTrue: [
    //      self expect: #colon.
    //      mgenc addArgumentIfAbsent: self argument ]
    //  )
    
    //  literal: mgenc = (
    //    sym == #pound ifTrue: [
    //      self peekForNextSymbolFromLexerIfNecessary.
    //      nextSym == #newTerm
    //        ifTrue: [ self literalArray: mgenc ]
    //        ifFalse: [ self literalSymbol: mgenc ].
    //      ^ self ].
    
    //    sym == #string ifTrue: [
    //      self literalString: mgenc.
    //      ^ self ].
    
    //    self literalNumber: mgenc
    //  )
    
    //  literalArray: mgenc = (
    //    | arrayClassName arraySizePlaceholder
    //      newMessage atPutMessage arraySizeLiteralIndex i |
    //    self expect: #pound.
    //    self expect: #newTerm.
    
    //    arrayClassName := universe symbolFor: 'Array'.
    //    arraySizePlaceholder := universe symbolFor: 'ArraySizeLiteralPlaceholder'.
    //    newMessage := universe symbolFor: 'new:'.
    //    atPutMessage := universe symbolFor: 'at:put:'.
    
    //    "need the array size at a know idx so that we don't need a second pass
    //     over the array elements"
    //    arraySizeLiteralIndex := mgenc addLiteral: arraySizePlaceholder.
    
    //    "create empty array"
    //    bcGen emit: mgenc pushGlobal: arrayClassName.
    //    bcGen emit: mgenc pushConstantIdx: arraySizeLiteralIndex.
    //    bcGen emit: mgenc send: newMessage.
    
    //    i := 1.
    
    //    [sym == #endTerm] whileFalse: [
    //      | pushIndex |
    //      pushIndex := universe newInteger: i.
    //      bcGen emit: mgenc pushConstant: pushIndex.
    //      self literal: mgenc.
    //      bcGen emit: mgenc send: atPutMessage.
    //      i := i + 1 ].
    
    //    "replace the placeholder with the actual array size"
    //    mgenc updateLiteral: arraySizePlaceholder at: arraySizeLiteralIndex put: (universe newInteger: i - 1).
    //    self expect: #endTerm.
    //  )
    
    //  literalSymbol: mgenc = (
    //    | symb |
    //    self expect: #pound.
    //    sym == #string
    //      ifTrue: [
    //        | s |
    //        s := self string.
    //        symb := universe symbolFor: s ]
    //      ifFalse: [
    //        symb := self selector ].
    //    bcGen emit: mgenc pushConstant: symb
    //  )
    func literalString(_ mgenc: MethodGenerationContext) {
        self.expect(.pound)
        var symb: SSymbol
        if sym == .string {
            let s = self.string()
            symb = universe.symbolFor(s)
        } else {
            symb = self.selector()
        }
        bcGen.emit(mgenc, pushConstant: symb)
    }
    
    //  literalString: mgenc = (
    //    | s str |
    //    s := self string.
    //    str := universe newString: s.
    //    bcGen emit: mgenc pushConstant: str
    //  )
    func literalString(_ mgenc: MethodGenerationContext) {
        let s = self.string()
        let str = universe.newString(s: s)
        bcGen.emit(mgenc, pushConstant: str)
    }

    //  literalNumber: mgenc = (
    //    | lit |
    //    sym == #minus
    //      ifTrue: [lit := self negativeDecimal]
    //      ifFalse: [lit := self literalDecimal: false].
    
    //    bcGen emit: mgenc pushConstant: lit
    //  )
    func literalNumber(_ mgenc: MethodGenerationContext) {
        var lit: SObject
        if sym == .minus {
            lit = self.negativeDecimal()
        } else {
            lit = self.literalDecimal(false)
        }
        bcGen.emit(mgenc, pushConstant: lit)
    }
    
    //  negativeDecimal = (
    //    self expect: #minus.
    //    ^ self literalDecimal: true
    //  )
    func negativeDecimal() -> SObject {
        self.expect(.minus)
        return self.literalDecimal(true)
    }
    //  literalDecimal: isNegative = (
    //    sym == #integer
    //      ifTrue: [^ self literalInteger: isNegative]
    //      ifFalse: [^ self literalDouble: isNegative]
    //  )
    func literalDecimal(_ isNegative: Bool) -> SObject {
        if sym == .integer {
            return self.literalInteger(isNegative)
        } else {
            return self.literalDouble(isNegative)
        }
    }
    
    
    //  literalInteger: isNegative = (
    //    | i |
    //    i := Integer fromString: text.
    //    isNegative ifTrue: [
    //      i := i negated].
    //    self expect: #integer.
    //    ^ universe newInteger: i
    //  )
    func literalInteger(_ isNegative: Bool) -> SInteger {
        let d = Int(self.text) ?? 0
        self.expect(.integer)
        return universe.newInteger(i: d)
    }

    //  literalDouble: isNegative = (
    //    | d |
    //    d := Double fromString: text.
    //    isNegative ifTrue: [
    //      d := d negated ].
    //    self expect: #double.
    //    ^ universe newDouble: d
    //  )
    func literalDouble(_ isNegative: Bool) -> SDouble {
        let d = Float(self.text) ?? 0.0
        self.expect(.double)
        return universe.newDouble(d: d)
    }
    
    //  accept: s = (
    //    sym == s ifTrue: [
    //      self takeSymbolFromLexer.
    //      ^ true ].
    //    ^ false
    //  )
    func accept(_ s: LexerToken) -> Bool {
        if self.sym == s {
            self.takeSymbolFromLexer()
            return true
        }
        return false
    }
    
    
    //  acceptOneOf: ss = (
    //    (self symIn: ss) ifTrue: [
    //      self takeSymbolFromLexer.
    //      ^ true ].
    //    ^ false
    //  )
    func acceptOneOf(_ ss: [LexerToken]) -> Bool {
        if self.symIn(ss) {
            self.takeSymbolFromLexer()
            return true
        }
        return false
    }

    
    //  expect: s = (
    //    (self accept: s) ifTrue: [ ^ true ].
    func expect(_ s: LexerToken) -> Bool {
        if self.accept(s) {
            return true
        }
    //    self error: 'Parsing of ' + filename + ' failed, expected ' + s + ' but found ' + sym +
    //      ' (' + text + ').\nCurrent parser context: ' + lexer currentTextContext
    //  )
        print("Parsing of \(filename) failed, expected \(s) but found \(sym) \(text)).\nCurrent parser context: \(lexer.currentTextContext())")
        exit(2)
    }

    //  expectOneOf: ss = (
    //    | err |
    //    (self acceptOneOf: ss) ifTrue: [ ^ true ].
    
    //    err := 'Parsing of ' + filename + ' failed, expected one of '.
    
    //    ss do: [
    //      err := err + s + ', ' ].
    //    err := err + 'but found: ' + sym + ' (' + text + ').\nCurrent parser context: ' + lexer currentTextContext.
    
    //    self error: err
    //  )
    func expectOneOf(_ ss: [LexerToken]) -> Bool {
        if self.acceptOneOf(ss) {
            return true
        }
        print("Parsing of \(filename) failed, expected \(ss) but found \(sym) \(text)).\nCurrent parser context: \(lexer.currentTextContext())")
        exit(2)
    }

    //  symIn: ss = (
    //    ^ ss contains: sym
    //  )
    func symIn(_ ss: [LexerToken]) -> Bool {
        return ss.contains(self.sym)
    }
    
    //  symIsMethod = (
    //    sym == #identifier       ifTrue: [^ true].
    //    sym == #keyword          ifTrue: [^ true].
    //    sym == #operatorSequence ifTrue: [^ true].
    //    (self symIn: Parser binaryOpSyms) ifTrue: [^ true].
    //    ^ false
    //  )
    func symIsMethod() -> Bool {
        if sym == .identifier ||
            sym == .keyword ||
            sym == .operatorSequence {
            return true
        }
        if self.symIn(Parser.binaryOpSyms) {
            return true
        }
        return false
    }
    
    //  peekForNextSymbolFromLexer = (
    //    nextSym := lexer peek
    //  )
    func peekForNextSymbolFromLexer() {
        self.nextSym = lexer.peek()
    }
    
    //  peekForNextSymbolFromLexerIfNecessary = (
    //    lexer isPeekDone ifFalse: [
    //      self peekForNextSymbolFromLexer ]
    //  )
    func peekForNextSymbolFromLexerIfNecessary() {
        if lexer.isPeekDone() == false {
            self.peekForNextSymbolFromLexer()
        }
    }
    
    //  gen: mgenc popVariable: var = (
    //    | searchResult |
    //    "Needs to determine whether the variable that is to be popped off the stack
    //     is a local variable, argument, or object field.
    //     This is done by examining all available lexical contexts, starting with
    //     the innermost (i.e., the one represented by mgenc)."
    
    //    "index, context, isArgument"
    //    searchResult := Array with: 0 with: 0 with: false. "TODO support: #(0 0 false)"
    
    //    (mgenc findVar: var with: searchResult)
    //      ifTrue: [
    //        (searchResult at: 3) "isArgument"
    //          ifTrue: [bcGen emit: mgenc popArgument: (searchResult at: 1) in: (searchResult at: 2)]
    //          ifFalse: [bcGen emit: mgenc popLocal: (searchResult at: 1) in: (searchResult at: 2)]
    //      ]
    //      ifFalse: [
    //        | varSym |
    //        varSym := universe symbolFor: var.
    //        (mgenc hasField: varSym) ifFalse: [
    //          ^ self error: 'Write to variable with the name ' + var + ', but there is no variable or field defined with this name' ].
    //        bcGen emit: mgenc popField: varSym ].
    //  )
    
    //  gen: mgenc pushVariable: var = (
    //    "Needs to determine whether the variable to be pushed on the stack
    //     is a local variable, argument, or object field.
    //     This is done by examining all available lexical contexts, starting with
    //     the innermost (i.e., the one represented by mgenc)."
    
    //    "index, context, isArgument"
    //    | searchResult |
    //    searchResult := Array with: 0 with: 0 with: false. "TODO support: #(0 0 false)"
    
    //    (mgenc findVar: var with: searchResult)
    //      ifTrue: [
    //        (searchResult at: 3) "isArgument"
    //          ifTrue: [
    //            bcGen emit: mgenc pushArgument: (searchResult at: 1) in: (searchResult at: 2) ]
    //          ifFalse: [
    //            bcGen emit: mgenc pushLocal: (searchResult at: 1) in: (searchResult at: 2) ] ]
    //      ifFalse: [
    //        | varSym |
    //        varSym := universe symbolFor: var.
    //        (mgenc hasField: varSym)
    //          ifTrue: [
    //            bcGen emit: mgenc pushField: varSym ]
    //          ifFalse: [
    //            bcGen emit: mgenc pushGlobal: varSym ] ]
    //  )
    
    
    //  ----
    //  | singleOpSyms binaryOpSyms keywordSelectorSyms |
    static let singleOpSyms: [LexerToken] = [.not, .and,
                                             .or, .star, .div, .mod,
                                             .plus, .equal, .more, .less,
                                             .comma, .at, .per, .none]
    //  singleOpSyms = (
    //    singleOpSyms == nil ifTrue: [
    //      singleOpSyms := #(#not #and #or #star #div #mod #plus #equal
    //                        #more #less #comma #at #per #minus #none) ].
    //    ^ singleOpSyms
    //  )
    
    //  binaryOpSyms = (
    //    binaryOpSyms == nil ifTrue: [
    //      binaryOpSyms := #(#or #comma #minus #equal #not #and #or #star
    //                        #div #mod #plus #equal #more #less #comma #at
    //                        #per #none) ].
    //    ^ binaryOpSyms
    //  )
    static let binaryOpSyms: [LexerToken] = [.or, .comma, .minus, .equal, .not,
                                             .and, .star, .div,
                                             .mod, .plus, .more, .less, .at,
                                             .per, .none]
    //  keywordSelectorSyms = (
    //    keywordSelectorSyms == nil ifTrue: [
    //      keywordSelectorSyms := #(#keyword #keywordSequence) ].
    //    ^ keywordSelectorSyms
    //  )
    static let keywordSelectorSyms: [LexerToken] = [.keyword, .keywordSequence]
    
    //  newWith: aString for: aFilename in: universe = (
    //    ^ self new initializeWith: aString for: aFilename in: universe
    //  )
    
    //  load: aFileName in: universe = (
    //    | fileContent |
    //    fileContent := system loadFile: aFileName.
    //    fileContent == nil ifTrue: [ ^ nil ].
    
    //    ^ self new initializeWith: fileContent for: aFileName in: universe
    //  )
    //)
} // PARSER
