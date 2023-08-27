
//  CompilerAll.swift
//  smog

//  Created by Kristofer Younger on 8/17/23.


import Foundation

//BytecodeGenerator = (
//  ----
//  emitPop: mgenc = (
//    self emit: mgenc bc: #pop
//  )

//  emit: mgenc pushArgument: idx in: ctx = (
//    self emit: mgenc bc: #pushArgument with: idx and: ctx
//  )

//  emitReturnLocal: mgenc = (
//    self emit: mgenc bc: #returnLocal
//  )

//  emitReturnNonLocal: mgenc = (
//    self emit: mgenc bc: #returnNonLocal
//  )

//  emitDup: mgenc = (
//    self emit: mgenc bc: #dup
//  )

//  emit: mgenc pushBlock: blockMethod = (
//    self emit: mgenc bc: #pushBlock with: (mgenc addLiteralIfAbsent: blockMethod)
//  )

//  emit: mgenc pushLocal: idx in: ctx = (
//    idx negative ifTrue: [ self error: 'pushLocal: ' + idx asString].
//    self emit: mgenc bc: #pushLocal with: idx and: ctx
//  )

//  emit: mgenc pushField: aSymbol = (
//    (mgenc hasField: aSymbol) ifFalse: [ self error: 'pushField: field unknown ' + aSymbol ].
//    self emit: mgenc bc: #pushField with: (mgenc fieldIndex: aSymbol)
//  )

//  emit: mgenc pushGlobal: aSymbol = (
//    self emit: mgenc bc: #pushGlobal with: (mgenc addLiteralIfAbsent: aSymbol)
//  )

//  emit: mgenc popArgument: idx in: ctx = (
//    idx negative ifTrue: [ self error: 'popArgument: ' + idx asString].
//    self emit: mgenc bc: #popArgument with: idx and: ctx
//  )

//  emit: mgenc popLocal: idx in: ctx = (
//    idx negative ifTrue: [ self error: 'popLocal: ' + idx asString].
//    self emit: mgenc bc: #popLocal with: idx and: ctx
//  )

//  emit: mgenc popField: aSymbol = (
//    (mgenc hasField: aSymbol) ifFalse: [ self error: 'popField: field unknown ' + aSymbol ].
//    self emit: mgenc bc: #popField with: (mgenc fieldIndex: aSymbol)
//  )

//  emit: mgenc superSend: aSymbol = (
//    self emit: mgenc bc: #superSend with: (mgenc addLiteralIfAbsent: aSymbol)
//  )

//  emit: mgenc send: aSymbol = (
//    self emit: mgenc bc: #send with: (mgenc addLiteralIfAbsent: aSymbol)
//  )

//  emit: mgenc pushConstant: anAbstractObject = (
//    self emit: mgenc bc: #pushConstant with: (mgenc addLiteralIfAbsent: anAbstractObject)
//  )

//  emit: mgenc pushConstantIdx: anInteger = (
//    self emit: mgenc bc: #pushConstant with: anInteger
//  )

//  emit: mgenc bc: aSymbol = (
//    mgenc addBytecode: aSymbol.
//  )

//  emit: mgenc bc: aSymbol with: anInteger = (
//    mgenc addBytecode: aSymbol.
//    mgenc addBytecode: anInteger.
//  )

//  emit: mgenc bc: aSymbol with: anInteger and: otherInteger = (
//    mgenc addBytecode: aSymbol.
//    mgenc addBytecode: anInteger.
//    mgenc addBytecode: otherInteger.
//  )
//)

//ClassGenerationContext = (
//  | universe
//    name superName
//    classSide
//    classFields instanceFields
//    classMethods instanceMethods |

//  initalize: aUniverse = (
//    universe := aUniverse.
//    classSide := false.
//    classFields := Vector new.
//    instanceFields := Vector new.
//    classMethods := Vector new.
//    instanceMethods := Vector new.
//  )

//  name = (
//    ^ name
//  )

//  name: aSSymbol = (
//    name := aSSymbol
//  )

//  superName = (
//    ^ superName
//  )

//  superName: aSymbol = (
//    superName := aSymbol
//  )

//  instanceFieldsOfSuper: aSArrayOfFieldNames = (
//    | numFields |
//    numFields := aSArrayOfFieldNames numberOfIndexableFields.
//    1 to: numFields do: [:i |
//      instanceFields append: (aSArrayOfFieldNames indexableField: i) ]
//  )

//  classFieldsOfSuper: aSArrayOfFieldNames = (
//    | numFields |
//    numFields := aSArrayOfFieldNames numberOfIndexableFields.
//    1 to: numFields do: [:i |
//      classFields append: (aSArrayOfFieldNames indexableField: i) ]
//  )

//  addField: aSymbol = (
//    classSide
//      ifTrue: [classFields append: aSymbol]
//      ifFalse: [instanceFields append: aSymbol]
//  )

//  hasField: aSymbol = (
//    ^ classSide
//      ifTrue: [classFields contains: aSymbol]
//      ifFalse: [instanceFields contains: aSymbol]
//  )

//  fieldIndex: aSymbol = (
//    ^ classSide
//      ifTrue: [classFields indexOf: aSymbol]
//      ifFalse: [instanceFields indexOf: aSymbol]
//  )

//  addMethod: anInvokable = (
//    classSide
//      ifTrue: [classMethods append: anInvokable]
//      ifFalse: [instanceMethods append: anInvokable]
//  )

//  startClassSide = (
//    classSide := true
//  )

//  assemble = (
//    | ccname superClass resultClass superMClass result |
//    "build class class name"
//    ccname := name string + ' class'.

//    "Load the super class"
//    superClass := universe loadClass: superName.

//    "Allocate the class of the resulting class"
//    resultClass := universe newClass: universe metaclassClass.

//    "Initialize the class of the resulting class"
//    resultClass instanceFields: (universe newArrayFromVector: classFields).
//    resultClass instanceInvokables: (universe newArrayFromVector: classMethods).
//    resultClass name: (universe symbolFor: ccname).

//    superMClass := superClass somClass.
//    resultClass superClass: superMClass.

//    "Allocate the resulting class"
//    result := universe newClass: resultClass.

//    "Initialize the resulting class"
//    result name: name.
//    result superClass: superClass.
//    result instanceFields: (universe newArrayFromVector: instanceFields).
//    result instanceInvokables: (universe newArrayFromVector: instanceMethods).

//    ^ result
//  )

//  assembleSystemClass: systemClass = (
//    | superMClass |
//    systemClass instanceInvokables: (universe newArrayFromVector: instanceMethods).
//    systemClass instanceFields: (universe newArrayFromVector: instanceFields).

//    "class-bound == class-instance-bound"
//    superMClass := systemClass somClass.
//    superMClass instanceInvokables: (universe newArrayFromVector: classMethods).
//    superMClass instanceFields: (universe newArrayFromVector: classFields).

//    ^ systemClass
//  )

//  ----

//  new: aUniverse = (
//    ^ self new initalize: aUniverse
//  )
//)

//Disassembler = (

//  ----

//  dump: cl in: universe = (
//    1 to: cl numberOfInstanceInvokables do: [:i |
//      | inv |
//      inv := cl instanceInvokable: i.

//      "output header and skip if the Invokable is a Primitive"
//      Universe errorPrint: (cl name string + '>>' + inv signature string + ' = ').

//      inv isPrimitive
//        ifTrue: [ Universe errorPrintln: '<primitive>' ]
//        ifFalse: [ self dumpMethod: inv indent: '\t' in: universe ] ]
//  )

//  dumpInvokable: inv in: universe = (
//    | holderName |
//    holderName := inv holder == nil
//      ifTrue: ['nil']
//      ifFalse: [inv holder name string].
//    Universe errorPrint: (holderName + '>>#' + inv signature string + ' = ').
//    inv isPrimitive
//        ifTrue: [
//          Universe errorPrint: '<primitive>: '.
//          Universe errorPrintln: inv debugString ]
//        ifFalse: [ self dumpMethod: inv indent: '\t' in: universe ]
//  )

//  dumpMethod: m indent: indent in: universe = (
//    | b |
//    Universe errorPrintln: '('.

//    "output stack information"
//    Universe errorPrintln: indent + '<' + m numberOfLocals + ' locals, '
//        + m maximumNumberOfStackElements + ' stack, '
//        + m numberOfBytecodes + ' bc_count>'.

//    "output bytecodes"
//    b := 1.
//    [b <= m numberOfBytecodes] whileTrue: [
//      | bytecode |
//      Universe errorPrint: indent.

//      b < 10 ifTrue: [ Universe errorPrint: ' ' ].
//      b < 100 ifTrue: [ Universe errorPrint: ' ' ].

//      Universe errorPrint: ' ' + b + ':'.

//      "mnemonic"
//      bytecode := m bytecode: b.
//      Universe errorPrint: (Bytecodes paddedBytecodeName: bytecode) + '  '.

//      "parameters (if any)"
//      (Bytecodes length: bytecode) = 1
//        ifTrue: [ Universe errorPrintln ]
//        ifFalse: [ self dumpBytecode: bytecode idx: b method: m indent: indent in: universe ].

//      b := b + (Bytecodes length: (m bytecode: b)) ].

//    Universe errorPrintln: indent + ')'
//  )

//  dumpBytecode: bc idx: b method: m indent: indent in: universe = (
//    bc == #pushLocal ifTrue: [
//      Universe errorPrintln: 'local: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushArgument ifTrue: [
//      Universe errorPrintln: 'argument: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushField ifTrue: [
//      | idx fieldName |
//      idx := m bytecode: b + 1.
//      fieldName := (m holder instanceFields indexableField: idx) string.
//      Universe errorPrintln: '(index: ' + idx + ') field: ' + fieldName.
//      ^ self ].
//    bc == #pushBlock ifTrue: [
//      Universe errorPrint: '(block: (index: ' + (m bytecode: b + 1) + ') '.
//      self dumpMethod: (m constant: b) indent: indent + '\t' in: universe.
//      ^ self ].
//    bc == #pushConstant ifTrue: [
//      | constant |
//      constant := m constant: b.
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') value: '
//        + '(' + (constant somClassIn: universe) name string + ') ' + constant debugString.
//      ^ self ].
//    bc == #pushGlobal ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') value: #' + (m constant: b) string.
//      ^ self ].
//    bc == #popLocal ifTrue: [
//      Universe errorPrintln: 'local: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #popArgument ifTrue: [
//      Universe errorPrintln: 'argument: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushField ifTrue: [
//      | idx fieldName |
//      idx := m bytecode: b + 1.
//      fieldName := (m holder instanceFields indexableField: idx) string.
//      Universe errorPrintln: '(index: ' + idx + ') field: ' + fieldName.
//      ^ self ].
//    bc == #send ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') signature: #' + (m constant: b) string.
//      ^ self ].
//    bc == #superSend ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') signature: #' + (m constant: b) string.
//      ^ self ].

//    Universe errorPrintln: '<unknown bytecode>'
//  )
//)
// LEXER

//MethodGenerationContext = (
//             | holderGenc outerGenc
//               arguments locals literals
//               signature
//               finished prim blockMethod
//               bytecode |

//             initializeWith: aHolderGenc and: aOuterGenc = (
//               holderGenc := aHolderGenc.
//               outerGenc := aOuterGenc.
//               arguments := Vector new.
//               locals := Vector new.
//               literals := Vector new.
//               finished := false.
//               prim := false.
//               blockMethod := false.
//               bytecode := Vector new.
//             )

//             holder = (
//               ^ holderGenc
//             )

//             signature: aSymbol = (
//               signature := aSymbol
//             )

//             addArgument: aString = (
//               arguments append: aString
//             )

//             addArgumentIfAbsent: aString = (
//               (arguments contains: aString)
//                 ifTrue: [^ false].

//               arguments append: aString.
//               ^ true
//             )

//             numberOfArguments = (
//               ^ arguments size
//             )

//             addLocalIfAbsent: aString = (
//               (locals contains: aString)
//                 ifTrue: [^ false].

//               locals append: aString.
//               ^ true
//             )

//             addLiteralIfAbsent: anAbstractObject = (
//               | idx |
//               idx := literals indexOf: anAbstractObject.
//               idx <> -1 ifTrue: [
//                 ^ idx ].

//               ^ self addLiteral: anAbstractObject
//             )

//             addLiteral: anAbstractObject = (
//               literals append: anAbstractObject.
//               ^ literals size
//             )

//             updateLiteral: oldVal at: idx put: newVal = (
//               (literals at: idx) == oldVal ifFalse: [
//                 self error: 'updateLiteral saw wrong oldVal, indicates bug in parser' ].
//               literals at: idx put: newVal
//             )

//             findVar: var with: searchResult = (
//               "searchResult: index, context, isArgument"
//               searchResult at: 1 put: (locals indexOf: var).
//               (searchResult at: 1) = -1 ifTrue: [
//                 searchResult at: 1 put: (arguments indexOf: var).
//                 (searchResult at: 1) = -1
//                   ifTrue: [
//                     outerGenc == nil
//                       ifTrue: [^ false]
//                       ifFalse: [
//                         searchResult at: 2 put: (searchResult at: 2) + 1.
//                         ^ outerGenc findVar: var with: searchResult ] ]
//                   ifFalse: [
//                     searchResult at: 3 put: true ] ].

//               ^ true
//             )

//             markAsFinished = (
//               finished := true
//             )

//             isFinished = (
//               ^ finished
//             )

//             markAsPrimitive = (
//               prim := true
//             )

//             isBlockMethod = (
//               ^ blockMethod
//             )

//             markAsBlockMethod = (
//               blockMethod := true
//             )

//             addBytecode: code = (
//               bytecode append: code
//             )

//             removeLastBytecode = (
//               bytecode remove
//             )

//             hasBytecodes = (
//               ^ bytecode isEmpty not
//             )

//             hasField: aSymbol = (
//               ^ holderGenc hasField: aSymbol
//             )

//             fieldIndex: aSymbol = (
//               ^ holderGenc fieldIndex: aSymbol
//             )

//             assemble: universe = (
//               prim
//                 ifTrue: [
//                   ^ SPrimitive emptyPrimitive: signature string in: universe ]
//                 ifFalse: [
//                   ^ self assembleMethod: universe ]
//             )

//             assembleMethod: universe = (
//               | numLocals meth i |
//               "create a method instance with the given number of bytecodes"
//               numLocals := locals size.

//               meth := universe newMethod: signature
//                   bc: bytecode asArray literals: literals asArray
//                   numLocals: numLocals maxStack: self computeStackDepth.

//               "return the method - the holder field is to be set later on!"
//               ^ meth
//             )

//             computeStackDepth = (
//               | depth maxDepth i |
//               depth := 0.
//               maxDepth := 0.
//               i := 1.

//               [i <= bytecode size] whileTrue: [
//                 | bc |
//                 bc := bytecode at: i.

//                 (bc == #dup           or: [
//                  bc == #pushLocal     or: [
//                  bc == #pushArgument  or: [
//                  bc == #pushField     or: [
//                  bc == #pushBlock     or: [
//                  bc == #pushConstant  or: [
//                  bc == #pushGlobal ] ] ] ] ] ])  ifTrue: [
//                   depth := depth + 1 ] ifFalse: [

//                 (bc == #pop         or: [
//                  bc == #popLocal    or: [
//                  bc == #popArgument or: [
//                  bc == #popField ] ] ]) ifTrue: [
//                   depth := depth - 1 ] ifFalse: [

//                 (bc == #send or: [bc == #superSend]) ifTrue: [
//                   | sig |
//                   "these are special: they need to look at the number of
//                    arguments (extractable from the signature)"
//                   sig := literals at: (bytecode at: i + 1).
//                   depth := depth - sig numberOfSignatureArguments.
//                   depth := depth + 1 "return value" ] ] ].

//                 i := i + (Bytecodes length: bc).

//                 depth > maxDepth ifTrue: [
//                   maxDepth := depth ] ].

//               ^ maxDepth
//             )

//             ----

//             new: holderGenc = (
//               ^ self new initializeWith: holderGenc and: nil
//             )

//             new: holderGenc with: outerGenc = (
//               ^ self new initializeWith: holderGenc and: outerGenc
//             )
//           )



//Parser = (
//  | lexer sym text nextSym filename cgenc universe bcGen |

//  initializeWith: aString for: aFilename in: aUniverse = (
//    filename := aFilename.
//    lexer := Lexer new: aString.
//    universe := aUniverse.
//    cgenc := ClassGenerationContext new: universe.
//    self takeSymbolFromLexer.

//    "This is just a convenient abbreviation."
//    bcGen := BytecodeGenerator.
//  )

//  takeSymbolFromLexer = (
//    sym := lexer sym.
//    text := lexer text.
//    nextSym := #none.
//  )

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

//  classBody = (
//    self fields.
//    [self symIsMethod] whileTrue: [
//       | mgenc |
//       mgenc := MethodGenerationContext new: cgenc.
//       mgenc addArgument: 'self'.
//       self method: mgenc.
//       cgenc addMethod: (mgenc assemble: universe) ].
//  )

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

//  binarySelector = (
//    | s |
//    s := text.

//    (self accept: #operatorSequence) or: [
//    (self acceptOneOf: Parser singleOpSyms) or: [
//        self expect: #none ] ].

//    ^ universe symbolFor: s
//  )

//  variable = (
//    ^ self identifier
//  )

//  argument = (
//    ^ self variable
//  )

//  identifier = (
//    | s |
//    s := text.
//    (self accept: #primitive)
//      ifFalse: [self expect: #identifier].
//    ^ s
//  )

//  keyword = (
//    | s |
//    s := text.
//    self expect: #keyword.
//    ^ s
//  )

//  string = (
//    | s |
//    s := text.
//    self expect: #string.
//    ^ s
//  )

//  selector = (
//    (sym == #operatorSequence or: [self symIn: Parser singleOpSyms])
//      ifTrue: [^ self binarySelector].
//    (sym == #keyword or: [sym == #keywordSequence])
//      ifTrue: [^ self keywordSelector].

//    ^ self unarySelector
//  )

//  keywordSelector = (
//    | s |
//    s := text.
//    self expectOneOf: Parser keywordSelectorSyms.
//    ^ universe symbolFor: s
//  )

//  result: mgenc = (
//    self expression: mgenc.

//    mgenc isBlockMethod
//      ifTrue: [bcGen emitReturnNonLocal: mgenc ]
//      ifFalse: [bcGen emitReturnLocal: mgenc ].
//    mgenc markAsFinished.

//    self accept: #period
//  )

//  expression: mgenc = (
//    self peekForNextSymbolFromLexer.

//    nextSym == #assign
//      ifTrue: [self assignation: mgenc]
//      ifFalse: [self evaluation: mgenc]
//  )

//  assignation: mgenc = (
//    | variables |
//    variables := Vector new.

//    self assignments: mgenc to: variables.
//    self evaluation: mgenc.

//    variables do: [:v | bcGen emitDup: mgenc ].
//    variables do: [:v | self gen: mgenc popVariable: v ]
//  )

//  assignments: mgenc to: variables = (
//    sym == #identifier ifTrue: [
//      variables append: (self assignment: mgenc).
//      self peekForNextSymbolFromLexer.
//      nextSym == #assign ifTrue: [
//        self assignments: mgenc to: variables ] ]
//  )

//  assignment: mgenc = (
//    | v |
//    v := self variable.
//    self expect: #assign.
//    ^ v
//  )

//  evaluation: mgenc = (
//    | superSend |
//    superSend := self primary: mgenc.
//    self symIsMethod ifTrue: [
//      self messages: mgenc with: superSend ]
//  )

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

//  unaryMessage: mgenc with: superSend = (
//    | msg |
//    msg := self unarySelector.

//    superSend ifTrue: [ bcGen emit: mgenc superSend: msg ]
//              ifFalse: [ bcGen emit: mgenc send: msg ]
//  )

//  binaryMessage: mgenc with: superSend = (
//    | msg |
//    msg := self binarySelector.
//    self binaryOperand: mgenc.

//    superSend ifTrue: [ bcGen emit: mgenc superSend: msg ]
//              ifFalse: [ bcGen emit: mgenc send: msg ]
//  )

//  binaryOperand: mgenc = (
//    | superSend |
//    superSend := self primary: mgenc.

//    [sym == #identifier] whileTrue: [
//      self unaryMessage: mgenc with: superSend.
//      superSend := false ].

//    ^ superSend
//  )

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

//  literalString: mgenc = (
//    | s str |
//    s := self string.
//    str := universe newString: s.
//    bcGen emit: mgenc pushConstant: str
//  )

//  literalNumber: mgenc = (
//    | lit |
//    sym == #minus
//      ifTrue: [lit := self negativeDecimal]
//      ifFalse: [lit := self literalDecimal: false].

//    bcGen emit: mgenc pushConstant: lit
//  )

//  negativeDecimal = (
//    self expect: #minus.
//    ^ self literalDecimal: true
//  )

//  literalDecimal: isNegative = (
//    sym == #integer
//      ifTrue: [^ self literalInteger: isNegative]
//      ifFalse: [^ self literalDouble: isNegative]
//  )

//  literalInteger: isNegative = (
//    | i |
//    i := Integer fromString: text.
//    isNegative ifTrue: [
//      i := i negated].

//    self expect: #integer.
//    ^ universe newInteger: i
//  )

//  literalDouble: isNegative = (
//    | d |
//    d := Double fromString: text.
//    isNegative ifTrue: [
//      d := d negated ].

//    self expect: #double.
//    ^ universe newDouble: d
//  )

//  accept: s = (
//    sym == s ifTrue: [
//      self takeSymbolFromLexer.
//      ^ true ].
//    ^ false
//  )

//  acceptOneOf: ss = (
//    (self symIn: ss) ifTrue: [
//      self takeSymbolFromLexer.
//      ^ true ].
//    ^ false
//  )

//  expect: s = (
//    (self accept: s) ifTrue: [ ^ true ].

//    self error: 'Parsing of ' + filename + ' failed, expected ' + s + ' but found ' + sym +
//      ' (' + text + ').\nCurrent parser context: ' + lexer currentTextContext
//  )

//  expectOneOf: ss = (
//    | err |
//    (self acceptOneOf: ss) ifTrue: [ ^ true ].

//    err := 'Parsing of ' + filename + ' failed, expected one of '.

//    ss do: [
//      err := err + s + ', ' ].
//    err := err + 'but found: ' + sym + ' (' + text + ').\nCurrent parser context: ' + lexer currentTextContext.

//    self error: err
//  )

//  symIn: ss = (
//    ^ ss contains: sym
//  )

//  symIsMethod = (
//    sym == #identifier       ifTrue: [^ true].
//    sym == #keyword          ifTrue: [^ true].
//    sym == #operatorSequence ifTrue: [^ true].
//    (self symIn: Parser binaryOpSyms) ifTrue: [^ true].
//    ^ false
//  )

//  peekForNextSymbolFromLexer = (
//    nextSym := lexer peek
//  )

//  peekForNextSymbolFromLexerIfNecessary = (
//    lexer isPeekDone ifFalse: [
//      self peekForNextSymbolFromLexer ]
//  )

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

//  keywordSelectorSyms = (
//    keywordSelectorSyms == nil ifTrue: [
//      keywordSelectorSyms := #(#keyword #keywordSequence) ].
//    ^ keywordSelectorSyms
//  )

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

//SourcecodeCompiler = (
//  ----
//  compileClass: path name: fileName into: systemClass in: universe = (
//    | fname parser result cname |
//    fname := path + '/' + fileName + '.som'.

//    parser := Parser load: fname in: universe.
//    parser ifNil: [ ^ nil ].

//    result := self compile: parser into: systemClass.

//    cname := result name string.

//    fileName ~= cname ifTrue: [
//      self error: 'File name ' + fname
//          + ' does not match class name (' + cname + ') in it.' ].
//    ^ result
//  )

//  compileClass: stmt into: systemClass in: universe = (
//    | parser |
//    parser := Parser newWith: stmt for: '$string$' in: universe.
//    ^ self compile: parser into: systemClass.
//  )

//  compile: parser into: systemClass = (
//    | cgc |
//    cgc := parser classdef.

//    systemClass == nil
//      ifTrue: [ ^ cgc assemble ]
//      ifFalse: [ ^ cgc assembleSystemClass: systemClass ]
//  )
//)
