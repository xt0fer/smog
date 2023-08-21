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
    func double() -> Float {
        return self.wrappedValue
    }
    func newFor() -> SDouble {
        return SDouble(d: self.wrappedValue)
    }
    func debugString() -> String {
        return "SDouble()"
    }

}
