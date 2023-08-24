//
//  Sblock.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SBlock: SObject {
    var method: SMethod
    var context: Frame
    var blockClass: SClass

    init(aSMethod: SMethod, aContext: Frame, aBlockClass: SClass) {
        self.method = aSMethod
        self.context = aContext
        self.blockClass = aBlockClass
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
    func evaluationPrimitive(_ numberOfArguments: Int, universe: Universe) -> SPrimitive {
        return SPrimitive(aSSymbol: <#SSymbol#>, block: <#SBlock#>)
    }
    
    // TODO: computeSignatureString: numberOfArguments = (
    //    | signatureString |
    //    signatureString := 'value'.
    //    numberOfArguments > 1 ifTrue: [
    //      signatureString := signatureString + ':' ].
    //
    //    "Add extra with: selector elements if necessary"
    //    2 to: numberOfArguments - 1 do: [:i |
    //        signatureString := signatureString + 'with:' ].
    //
    //    ^ signatureString
    //  )
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

//SBlock = SAbstractObject (
//  | method context blockClass |
//
//  initialize: aSMethod in: aContext with: aBlockClass = (
//    method := aSMethod.
//    context := aContext.
//    blockClass := aBlockClass.
//  )
//
//  method = (
//    ^ method
//  )
//
//  context = (
//    ^ context
//  )
//
//  somClassIn: universe = (
//    ^ blockClass
//  )
//
//  "For using in debugging tools such as the Diassembler"
//  debugString = ( ^ 'SBlock(' + method asString + ')' )
//
//  ----
//
//  new: aSMethod in: aContext with: aBlockClass = (
//    ^ self new initialize: aSMethod in: aContext with: aBlockClass
//  )
//
//  evaluationPrimitive: numberOfArguments in: universe = (
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
//
//  computeSignatureString: numberOfArguments = (
//    | signatureString |
//    signatureString := 'value'.
//    numberOfArguments > 1 ifTrue: [
//      signatureString := signatureString + ':' ].
//
//    "Add extra with: selector elements if necessary"
//    2 to: numberOfArguments - 1 do: [:i |
//        signatureString := signatureString + 'with:' ].
//
//    ^ signatureString
//  )
//)
