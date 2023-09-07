//
//  SPrimitive.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SPrimitive: SObject, Invokable {
    func isNil() -> Bool {
        return false
    }
    
    
    var signatureSym: SSymbol
    var holderClass: SClass
    var isEmpty: Bool = false
    var operation: SBlock
    
    
    init(aSSymbol: SSymbol, block: SBlock) {
        self.signatureSym = aSSymbol
        self.isEmpty = false
        self.operation = block
        holderClass = Universe.shared.primClass
        super.init(nArgs: 0, clazz: Universe.shared.primClass)
        self.debugSDesc = "SPrimitive()"
    }
    //
    // this init looks like it puts lambda (block) of a failed message call into the opertion field.
    //
    //  initializeEmpty: aSSymbol in: universe = (
    //    signature := aSSymbol.
    //    isEmpty := true.
    //    operation := [:frame :interp |
    //      | receiver msg |
    //      signature numberOfSignatureArguments timesRepeat: [
    //        receiver := frame pop ].
    //      msg := 'Undefined primitive ' + (receiver somClassIn: universe) name string +
    //        '>>#' + signature string + ' called'.
    //      self send: 'error:' with: (Array with: receiver with: (universe newString: msg))
    //           in: universe using: interp ].
    //  )
        
    func isPrimitive() -> Bool { return false }
    
//    func invoke(frame: Frame, using interpreter: Interpreter) -> SObject {
//        print("invoke(frame:using:) not implemented")
//        return Universe.shared.nilObject //Block operation value: frame with: interp
//    }
    func invoke(frame: Frame,  using interpreter: Interpreter) {
        let newFrame = interpreter.pushNewFrame(invokable: self)
        newFrame.copyArgumentsFrom(frame: frame)
    }


    func invoke(frame: Frame) {
        self.invoke(frame: frame, using: Universe.shared.interpreter)
    }
    

//    func invoke(frame: Frame, using: Interpreter) {
//        <#code#>
//    }
    
    func signature() -> SSymbol {
        return self.signatureSym
    }
    
    func holder() -> SClass {
        return self.holderClass
    }
    
    func holder(value: SClass) {
        self.holderClass = value
    }

}
