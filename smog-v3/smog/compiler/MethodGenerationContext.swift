//
//  MethodGenerationContext.swift
//  smog
//
//  Created by Kristofer Younger on 8/28/23.
//

import Foundation

class MethodGenerationContext {
    
    //             | holderGenc outerGenc
    //               arguments locals literals
    //               signature
    //               finished prim blockMethod
    //               bytecode |
    var bytecode: [UInt8] = []
    var holderGenc: ClassGenerationContext
    var outerGenc: ClassGenerationContext
    let nilGenc = ClassGenerationContext.nilGenc
    var arguments: [String] = []
    var locals: [String] = []
    var literals: [String] = []
    var signature: String = ""
    var finished = false
    var prim = false
    var blockMethod = false
    
    init(aHolderGenc: ClassGenerationContext, aOuterGenc: ClassGenerationContext) {
        self.holderGenc = aHolderGenc
        self.outerGenc = aOuterGenc
    }
    
    init(aHolderGenc: ClassGenerationContext) {
        self.holderGenc = aHolderGenc
        self.outerGenc = nilGenc
    }
    
    func addBytecode(_ code: Int) {
        let b = UInt8(truncatingIfNeeded: code)
        self.bytecode.append(b)
    }
    
    
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
    func holder() -> ClassGenerationContext{
        return self.holderGenc
    }
    
    //             signature: aSymbol = (
    //               signature := aSymbol
    //             )
    func signature(_ symbol: String) {
        self.signature = symbol
    }

//             addArgument: aString = (
//               arguments append: aString
//             )
    func addArgument(_ s: String) {
        self.arguments.append(s)
    }

//             addArgumentIfAbsent: aString = (
//               (arguments contains: aString)
//                 ifTrue: [^ false].

//               arguments append: aString.
//               ^ true
//             )
    func addArgumentIfAbsent(_ s: String) -> Bool {
        if arguments.contains(s) {
            return false
        }
        arguments.append(s)
        return true
    }

//             numberOfArguments = (
//               ^ arguments size
//             )
    func numberOfArguments() -> Int {
        return arguments.count
    }

//             addLocalIfAbsent: aString = (
//               (locals contains: aString)
//                 ifTrue: [^ false].

//               locals append: aString.
//               ^ true
//             )
    func addLocalIfAbsent(_ s: String) -> Bool {
        if locals.contains(s) {
            return false
        }
        locals.append(s)
        return true
    }

//             addLiteralIfAbsent: anAbstractObject = (
//               | idx |
//               idx := literals indexOf: anAbstractObject.
//               idx <> -1 ifTrue: [
//                 ^ idx ].

//               ^ self addLiteral: anAbstractObject
//             )
    func addLiteralIfAbsent(_ anAbstractObject: String) -> Int {
        if literals.contains(anAbstractObject) {
            return literals.firstIndex(of: anAbstractObject)!
        }
        return addLiteral(anAbstractObject)
    }

//             addLiteral: anAbstractObject = (
//               literals append: anAbstractObject.
//               ^ literals size
//             )
    func addLiteral(_ anAbstractObject: String) -> Int {
        literals.append(anAbstractObject)
        return literals.count
    }

//             updateLiteral: oldVal at: idx put: newVal = (
//               (literals at: idx) == oldVal ifFalse: [
//                 self error: 'updateLiteral saw wrong oldVal, indicates bug in parser' ].
//               literals at: idx put: newVal
//             )
    func updateLiteral(_ oldVal: String, at idx: Int, put newVal: String) {
        if literals[idx] != oldVal {
            print("updateLiteral saw wrong oldVal, indicates bug in parser")
        }
        literals[idx] = newVal
    }
    
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
    func findVar(_ var: String, with searchResult: [Int]) -> Bool {
        var sr: [Int] = []
        sr.append(locals.firstIndex(of: var))
        
    }

//             markAsFinished = (
//               finished := true
//             )
    func markAsFinished() {self.finished = true}

//             isFinished = (
//               ^ finished
//             )
    func isFinished() -> Bool {return self.finished}

//             markAsPrimitive = (
//               prim := true
//             )
    func markAsPrimitive() {self.prim = true}

//             isBlockMethod = (
//               ^ blockMethod
//             )
    func isBlockMethod() -> Bool { return blockMethod }

//             markAsBlockMethod = (
//               blockMethod := true
//             )
    func markAsBlockMethod() {blockMethod = true}


//             removeLastBytecode = (
//               bytecode remove
//             )
    func removeLastBytecode() { bytecode.removeLast() }

//             hasBytecodes = (
//               ^ bytecode isEmpty not
//             )
    func hasBytecodes() -> Bool {
        return !self.bytecode.isEmpty
    }

//             hasField: aSymbol = (
//               ^ holderGenc hasField: aSymbol
//             )
    func hasField(_ symbol: String) -> Bool {
        return holderGenc.hasField(symbol)
    }

//             fieldIndex: aSymbol = (
//               ^ holderGenc fieldIndex: aSymbol
//             )
    func fieldIndex(_ symbol: String) -> Int {
        return holderGenc.fieldIndex(symbol)
    }

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
}
