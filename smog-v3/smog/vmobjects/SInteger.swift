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

    func integer() -> Int {
        return self.wrappedValue
    }
    
    func debugString() -> String {
        return "SInteger()"
    }

    func newFor() -> SInteger {
        return SInteger(i: self.wrappedValue)
    }

    //)
    
}
