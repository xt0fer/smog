//
//  Sblock.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SBlock: SObject, Invokable {
    func isNil() -> Bool {
        return true
    }
    
    func isPrimitive() -> Bool {
        false
    }
    
    func invoke(frame: Frame) {
        
    }
    
    func invoke(frame: Frame, using: Interpreter) {
        
    }
    
    func signature() -> SSymbol {
        self.method.signatureSym
    }
    
    func holder() -> SClass {
        return self.holderClass
    }
    
    func holder(value: SClass) {
        self.holderClass = value
    }
    
    var method: SMethod
    var context: Frame
    var blockClass: SClass
    var holderClass: SClass
    
    init(aSMethod: SMethod, aContext: Frame, aBlockClass: SClass) {
        self.method = aSMethod
        self.context = aContext
        self.blockClass = aBlockClass
        self.holderClass = Universe.shared.objectClass
        super.init(nArgs: 0, clazz: Universe.shared.blockClass)
        self.debugSDesc = "SBlock()"
    }
    
    // TODO: evaluationPrimitive: numberOfArguments in: universe = (
    //    ^ SPrimitive new: (self computeSignatureString: numberOfArguments)
    //                  in: universe
    //                with: [:frame :interp |
    //        | rcvr context newFrame |
    //        "Get the block (the receiver) from the stack"
    //        rcvr := frame stackElement: numberOfArguments - 1.
    //
    //        "Get the context of the block"
    //        context := rcvr context.
    //
    //        "Push a new frame and set its context to be the one specified in
    //         the block"
    //        newFrame := interp pushNewFrame: rcvr method with: context.
    //        newFrame copyArgumentsFrom: frame ]
    //  )
    static func evaluationPrimitive(_ numberOfArguments: Int, universe: Universe) -> SPrimitive {
        //return SPrimitive(aSSymbol: <#SSymbol#>, block: <#SBlock#>)
        return SPrimitive(aSSymbol: universe.symbolFor("Nil"), block: universe.nilObject as! SBlock)
    }
    
    func computeSignatureString(numberOfArguments: Int) -> String {
        var signatureString = "value"
        if numberOfArguments > 1 {
            signatureString.append(":")
        }
        for _ in 2...(numberOfArguments - 1) {
            signatureString.append("with:")
        }
        return signatureString
    }
}

