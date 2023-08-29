//
//  BytecodeGenerator.swift
//  smog
//
//  Created by Kristofer Younger on 8/28/23.
//

import Foundation

class BytecodeGenerator {
    // All these are static methods...
    
    
    //BytecodeGenerator = (
    //  ----
    //  emitPop: mgenc = (
    //    self emit: mgenc bc: #pop
    //  )
    static func emitPop(_ mgenc: MethodGenerationContext) {
        BytecodeGenerator.emit(mgenc, bc: .pop)
    }
    
    //  emit: mgenc pushArgument: idx in: ctx = (
    //    self emit: mgenc bc: #pushArgument with: idx and: ctx
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                             pushArgument idx: Int, in ctx: Int) {
        BytecodeGenerator.emit(mgenc, bc: .pushArgument, with: idx, and: ctx)
    }
    
    //  emitReturnLocal: mgenc = (
    //    self emit: mgenc bc: #returnLocal
    //  )
    static func emitReturnLocal(_ mgenc: MethodGenerationContext) {
        BytecodeGenerator.emit(mgenc, bc: .returnLocal)
    }

    //  emitReturnNonLocal: mgenc = (
    //    self emit: mgenc bc: #returnNonLocal
    //  )
    static func emitReturnNonLocal(_ mgenc: MethodGenerationContext) {
        BytecodeGenerator.emit(mgenc, bc: .returnNonLocal)
    }

    //  emitDup: mgenc = (
    //    self emit: mgenc bc: #dup
    //  )
    static func emitDup(_ mgenc: MethodGenerationContext) {
        BytecodeGenerator.emit(mgenc, bc: .dup)
    }

    //  emit: mgenc pushBlock: blockMethod = (
    //    self emit: mgenc bc: #pushBlock with: (mgenc addLiteralIfAbsent: blockMethod)
    //  )
    static func emit(_ mgenc: MethodGenerationContext, pushBlock blockMethod: String) {
        BytecodeGenerator.emit(mgenc, bc: .pushBlock,
                               with: mgenc.addLiteralIfAbsent(blockMethod))
    }

    //  emit: mgenc pushLocal: idx in: ctx = (
    //    idx negative ifTrue: [ self error: 'pushLocal: ' + idx asString].
    //    self emit: mgenc bc: #pushLocal with: idx and: ctx
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                 pushLocal idx: Int, in ctx: Int) {
        if idx < 0 {
            error("pushLocal error, \(idx)")
        }
        BytecodeGenerator.emit(mgenc, bc: .pushLocal, with: idx, and: ctx)
    }

    //  emit: mgenc pushField: aSymbol = (
    //    (mgenc hasField: aSymbol) ifFalse: [ self error: 'pushField: field unknown ' + aSymbol ].
    //    self emit: mgenc bc: #pushField with: (mgenc fieldIndex: aSymbol)
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                          pushField symbol: String) {
        if !mgenc.hasField(symbol) {
            error("pushField: field unknown, \(symbol)")
        }
        BytecodeGenerator.emit(mgenc, bc: .pushField, with: mgenc.fieldIndex(symbol))
    }

    //  emit: mgenc pushGlobal: aSymbol = (
    //    self emit: mgenc bc: #pushGlobal with: (mgenc addLiteralIfAbsent: aSymbol)
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     pushGlobal symbol: String) {
        BytecodeGenerator.emit(mgenc, bc: .pushGlobal, with: mgenc.addLiteralIfAbsent(symbol))
    }

    //  emit: mgenc popArgument: idx in: ctx = (
    //    idx negative ifTrue: [ self error: 'popArgument: ' + idx asString].
    //    self emit: mgenc bc: #popArgument with: idx and: ctx
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     popArgument idx: Int, in ctx: Int) {
        if idx < 0 {
            error("popArgument error, \(idx)")
        }
        BytecodeGenerator.emit(mgenc, bc: .popArgument, with: idx, and: ctx)
    }

    //  emit: mgenc popLocal: idx in: ctx = (
    //    idx negative ifTrue: [ self error: 'popLocal: ' + idx asString].
    //    self emit: mgenc bc: #popLocal with: idx and: ctx
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     popLocal idx: Int, in ctx: Int) {
        if idx < 0 {
            error("popLocal error, \(idx)")
        }
        BytecodeGenerator.emit(mgenc, bc: .popLocal, with: idx, and: ctx)
    }

    //  emit: mgenc popField: aSymbol = (
    //    (mgenc hasField: aSymbol) ifFalse: [ self error: 'popField: field unknown ' + aSymbol ].
    //    self emit: mgenc bc: #popField with: (mgenc fieldIndex: aSymbol)
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     popField symbol: Int) {
        if !mgenc.hasField(symbol) {
            error("popField: field unknown, \(symbol)")
        }
        BytecodeGenerator.emit(mgenc, bc: .popField, with: mgenc.fieldIndex(symbol))
    }

    //  emit: mgenc superSend: aSymbol = (
    //    self emit: mgenc bc: #superSend with: (mgenc addLiteralIfAbsent: aSymbol)
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     superSend symbol: Int) {
        if !mgenc.hasField(symbol) {
            error("popField: field unknown, \(symbol)")
        }
        BytecodeGenerator.emit(mgenc, bc: .superSend, with: mgenc.addLiteralIfAbsent(symbol))
    }

    
    //  emit: mgenc send: aSymbol = (
    //    self emit: mgenc bc: #send with: (mgenc addLiteralIfAbsent: aSymbol)
    //  )
    static func emit(_ mgenc: MethodGenerationContext,
                     send symbol: Int) {
        if !mgenc.hasField(symbol) {
            error("send: field unknown, \(symbol)")
        }
        BytecodeGenerator.emit(mgenc, bc: .send, with: mgenc.addLiteralIfAbsent(symbol))
    }

    //  emit: mgenc pushConstant: anAbstractObject = (
    //    self emit: mgenc bc: #pushConstant with: (mgenc addLiteralIfAbsent: anAbstractObject)
    //  )
    static func emit(_ mgenc: MethodGenerationContext, pushConstant: SObject) {
        emit(mgenc, bc: .pushConstant, with: mgenc.addLiteralIfAbsent(anAbstractObject))
    }

    //  emit: mgenc pushConstantIdx: anInteger = (
    //    self emit: mgenc bc: #pushConstant with: anInteger
    //  )
    static func emit(_ mgenc: MethodGenerationContext, pushConstantIdx: Int) {
        emit(mgenc, bc: .pushConstant, with: pushConstantIdx)
    }

    //  emit: mgenc bc: aSymbol = (
    //    mgenc addBytecode: aSymbol.
    //  )
    static func emit(_ mgenc: MethodGenerationContext, bc: Bytecode) {
        mgenc.addBytecode(BytecodeMap[bc]!)
    }
    
    //  emit: mgenc bc: aSymbol with: anInteger = (
    //    mgenc addBytecode: aSymbol.
    //    mgenc addBytecode: anInteger.
    //  )
    static func emit(_ mgenc: MethodGenerationContext, bc: Bytecode, with arg1: Int) {
        mgenc.addBytecode(BytecodeMap[bc]!)
        mgenc.addBytecode(arg1)
    }

    //  emit: mgenc bc: aSymbol with: anInteger and: otherInteger = (
    //    mgenc addBytecode: aSymbol.
    //    mgenc addBytecode: anInteger.
    //    mgenc addBytecode: otherInteger.
    //  )
    static func emit(_ mgenc: MethodGenerationContext, bc: Bytecode, with arg1: Int, and arg2: Int) {
        mgenc.addBytecode(BytecodeMap[bc]!)
        mgenc.addBytecode(arg1)
        mgenc.addBytecode(arg2)
    }
    static func error(_ s: String) {
        print(s)
    }
}
