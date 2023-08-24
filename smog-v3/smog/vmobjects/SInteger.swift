//
//  SInteger.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SInteger: SObject {
    
    
    var wrappedValue: Int

    init(i: Int) {
        self.wrappedValue = i
        super.init(nArgs: 0, clazz: Universe.shared.integerClass)
        self.debugSDesc = "SInteger"
    }

    func integer() -> Int {
        return self.wrappedValue
    }
    
    func newFor() -> SInteger {
        return SInteger(i: self.wrappedValue)
    }

    //)
    
}
