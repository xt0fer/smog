//
//  SPrimitive.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SPrimitive: SObject {
    
    var signature: SSymbol
    var holder: SObject
    var isEmpty: Bool = false
    var operation: SBlock

    
    init(aSSymbol: SSymbol, block: SBlock) {
        self.signature = aSSymbol
        self.isEmpty = false
        self.operation = block
        holder = Universe.shared.nilObject
        super.init(nArgs: 0, clazz: Universe.shared.primClass)
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

    func holder(_ newHolder: SObject) {
        self.holder = newHolder
    }
    
    override func debugString() -> String {
        return "SMethod(\(String(describing: self.clazz.name)))"
    }
    
    func isPrimitive() -> Bool { return false }

    func invoke(frame: Frame, using interpreter: Interpreter) -> SObject {
        print("invoke(frame:using:) not implemented")
        return Universe.shared.nilObject //Block operation value: frame with: interp
    }

//    var debugString: String {
//        return "SPrimitive(\(String(describing: self.somClass())))"
//    }

}
