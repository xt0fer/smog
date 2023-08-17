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
    init(s: String) {
        self.s = s
        super.init(nArgs: 0, clazz: Universe().SStringClass)
    }
//
//          string = ( ^ string )
//
//          somClassIn: universe = (
//            ^ universe stringClass
//          )
//
//          "For using in debugging tools such as the Diassembler"
//          debugString = ( ^ 'SString(' + string + ')' )
//
//          ----
//
//          new: aString = (
//            ^ self new initializeWith: aString
//          )
//    )
    
}
