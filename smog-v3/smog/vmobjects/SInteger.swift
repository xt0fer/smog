//
//  SInteger.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SInteger: SAbstractObject {
    
    
    var wrappedValue: Int
    init(i: Int) {
        self.wrappedValue = i
        super.init()
    }
    
    
    
    //SInteger = SAbstractObject (
    //  | value |
    //
    //  initialize: anInteger = (
    //    value := anInteger
    //  )
    //
    //  integer = ( ^ value )
    func integer() -> Int {
        return self.wrappedValue
    }
    //
    //  somClassIn: universe = (
    //    ^ universe integerClass
    //  )
    func somClassIn(_ u: Universe) -> SClass {
        return u.integerClass
    }
    //
    //  "For using in debugging tools such as the Diassembler"
    //  debugString = ( ^ 'SInteger(' + value asString + ')' )
    //
    //  ----
    //
    //  "TODO: see whether it makes sense to have a cache"
    //  for: anInteger = (
    //    ^ self new initialize: anInteger
    //  )
    func newFor() -> SInteger {
        return SInteger(i: self.wrappedValue)
    }

    //)
    
}
