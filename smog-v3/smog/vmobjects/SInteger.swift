//
//  SInteger.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SInteger: SObject {
    
    
    var wrappedValue: Int
    var debugId = "SInteger"

    init(i: Int) {
        self.wrappedValue = i
        super.init(nArgs: 0, clazz: Universe.shared.integerClass)
    }

    func integer() -> Int {
        return self.wrappedValue
    }
    
    func newFor() -> SInteger {
        return SInteger(i: self.wrappedValue)
    }

    //)
    
}
