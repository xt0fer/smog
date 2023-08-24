//
//  SDouble.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SDouble: SObject {
    
    var wrappedValue: Float
    var debugId = "SDouble"
    
    init(d: Float) {
        self.wrappedValue = d
        super.init(nArgs: 0, clazz: Universe.shared.doubleClass)
    }
    func double() -> Float {
        return self.wrappedValue
    }
    func newFor() -> SDouble {
        return SDouble(d: self.wrappedValue)
    }
    
}
