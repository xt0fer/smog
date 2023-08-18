//
//  SDouble.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SDouble: SAbstractObject {
    
    var wrappedValue: Float
    init(d: Float) {
        self.wrappedValue = d
        super.init()
    }

//SDouble = SAbstractObject (
//  | value |
//
//  initialize: aDouble = (
//    value := aDouble
//  )
//
//  double = ( ^ value )
    func double() -> Float {
        return self.wrappedValue
    }
//
//  somClassIn: universe = (
//    ^ universe doubleClass
//  )
    func somClassIn(_ u: Universe) -> SClass {
        return u.doubleClass
    }
//
//  "For using in debugging tools such as the Diassembler"
//  debugString = ( ^ 'SDouble(' + value asString + ')' )
//
//  ----
//
//  for: aDouble = (
//    ^ self new initialize: aDouble
//  )
    func newFor() -> SDouble {
        return SDouble(d: self.wrappedValue)
    }
}
//)
