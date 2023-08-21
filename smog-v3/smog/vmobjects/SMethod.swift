//
//  SMethod.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SMethod: SObject {
    
    var signature: SSymbol
    var holder: SObject
    var bytecodes: [Int]
    var literals: [SObject]
    var numberOfLocals: Int
    var maximumNumberOfStackElements: Int
    
    init(aSSymbol: SSymbol, bc: [Int], literals: [SString], numLocals: Int, maxStack: Int) {
        self.signature = aSSymbol
        self.bytecodes = bc
        self.literals = literals
        self.numberOfLocals = numLocals
        self.maximumNumberOfStackElements = maxStack
        self.holder = Universe.shared.nilObject
        super.init(nArgs: 0, clazz: Universe.shared.methodClass)
    }
    
    override func debugString() -> String {
        return "SMethod(\(String(describing: self.clazz.name)))"
    }
    
    override func isPrimitive() -> Bool { return false }
    
    
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
    func holder(_ newHolder: SObject) {
        self.holder = newHolder
        for lit in literals {
            if let l = lit as? SMethod {
                l.holder(newHolder)
            }
            if let l = lit as? SPrimitive {
                l.holder(newHolder)
            }
        }
    }
    
    //
    //  constant: bytecodeIndex = (
    //    "Get the constant associated to a given bytecode index"
    //    ^ literals at: (bytecodes at: bytecodeIndex + 1)
    //  )
    func constant(bcIndex: Int) -> SObject {
        return self.literals[Int(self.bytecodes[bcIndex + 1] as Int)]
    }
    //
    //  numberOfArguments = (
    //    "Get the number of arguments of this method"
    //    ^ signature numberOfSignatureArguments
    //  )
    func numberOfArguments() -> Int { return self.signature.numSignatureArguments }
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
    //  invoke: frame using: interpreter = (
    //    | newFrame |
    //    "Allocate and push a new frame on the interpreter stack"
    //    newFrame := interpreter pushNewFrame: self.
    //    newFrame copyArgumentsFrom: frame
    //  )
    //
    
    func invoke(frame: Frame,  using interpreter: Interpreter) {
        let newFrame = interpreter.pushNewFrame(self)
        newFrame.copyArgumentsFrom(frame)
    }
}
