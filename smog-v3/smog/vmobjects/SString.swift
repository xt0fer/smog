//
//  SString.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SString: SObject {
    var s: String = ""
    
//    SString = SAbstractObject (
//        | string |
//
    
//          ----
//
//          new: aString = (
//            ^ self new initializeWith: aString
//          )
//    )

    init(s: String) {
        self.s = s
        super.init(nArgs: 0, clazz: Universe.shared.stringClass)
    }
//
//          string = ( ^ string )
    func string() -> String {
        return self.s
    }
//
//          somClassIn: universe = (
//            ^ universe stringClass
//          )
    func somClassInt(_ u: Universe) -> SClass {
        return Universe.shared.stringClass
    }
//
//          "For using in debugging tools such as the Diassembler"
//          debugString = ( ^ 'SString(' + string + ')' )

    override func debugString() -> String {
        return "SString(\(String(describing: self.clazz?.name)))"
    }
    
}
