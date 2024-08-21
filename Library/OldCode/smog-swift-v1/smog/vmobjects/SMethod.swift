//
//  SMethod.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SMethod: SObject, Invokable {
    func isNil() -> Bool {
        return false
    }
    
    
    var signatureSym: SSymbol
    var holderClass: SClass
    var bytecodes: [Int]
    var literals: [SObject] // These should be SStrings or SSymbols??
    var numberOfLocals: Int
    var maximumNumberOfStackElements: Int
    
    init(aSSymbol: SSymbol, bc: [Int], literals: [SString], numLocals: Int, maxStack: Int) {
        self.signatureSym = aSSymbol
        self.bytecodes = bc
        self.literals = literals
        self.numberOfLocals = numLocals
        self.maximumNumberOfStackElements = maxStack
        self.holderClass = Universe.shared.nilClass
        super.init(nArgs: 0, clazz: Universe.shared.methodClass)
        self.debugSDesc = "SMethod()"
    }
    
    func isPrimitive() -> Bool { return false }
    
    
    //
    //  holder: value = (
    //    holder := value.
    //
    //    literals == nil ifTrue: [ ^ self ].
    //
    //    "Make sure all nested invokables have the same holder"
    //    literals do: [:l |
    //      (l class == SMethod or: [l class == SPrimitive]) ifTrue: [
    //        l holder: value ] ]
    //  )
    func holder(value: SClass) {
        self.holderClass = value
        for lit in literals {
            if let l = lit as? SMethod {
                l.holder(value: value)
            }
            if let l = lit as? SPrimitive {
                l.holder(value: value)
            }
        }
    }
    
    func invoke(frame: Frame,  using interpreter: Interpreter) {
        let newFrame = interpreter.pushNewFrame(invokable: self)
        newFrame.copyArgumentsFrom(frame: frame)
    }

    func invoke(frame: Frame) {
        self.invoke(frame: frame, using: Universe.shared.interpreter)
    }
    
    func signature() -> SSymbol {
        return self.signatureSym
    }
    
    func holder() -> SClass {
        return self.holderClass
    }
    
    
    //
    //  constant: bytecodeIndex = (
    //    ^ literals at: (bytecodes at: bytecodeIndex + 1)
    //  )

    //    "Get the constant/literal/string/symbol associated to a given bytecode index"
    func constant(bcIndex: Int) -> SObject {
        return self.literals[Int(self.bytecodes[bcIndex + 1] as Int)]
    }
    //
    //  numberOfArguments = (
    //    "Get the number of arguments of this method"
    //    ^ signature numberOfSignatureArguments
    //  )
    func numberOfArguments() -> Int { return self.signatureSym.numSignatureArguments }
    //
    //  numberOfBytecodes = (
    //    "Get the number of bytecodes in this method"
    //    ^ bytecodes length
    //  )
    func numberOfBytecodes() -> Int { return self.bytecodes.count }
    //
    //  bytecode: index = (
    //    "Get the bytecode at the given index"
    //    ^ bytecodes at: index
    //  )
    func bytecode(at: Int) -> Int { return self.bytecodes[at] }
    //
}
