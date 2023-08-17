//
//  sample.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SObject: AbstractObject {
    var fields: [AbstractObject?] = []
    var clazz: SClass
    
    init(nArgs: Int, clazz: SClass) {
        self.fields = Array(repeating: nil, count: nArgs)
        self.clazz = clazz
    }
    
    func somClass() -> SClass {
        return self.clazz
    }

    func somClass(aSClass: SClass) {
        self.clazz = aSClass
    }
    func somClassIn(_ u: Universe) -> SClass {
        return self.clazz
    }

    func fieldName(index: Int) -> SString {
        return self.clazz.instanceFieldName(index)
    }
    func fieldIndex(name: SString) -> Int {
        return self.clazz.lookupFieldIndex(name)
    }
    func numberOfFields() -> Int {
        return fields.count
    }
    func field(index: Int) -> SObject {
        return self.fields[index]
    }
    func fieldAt(_ index: Int, put: SObject) {
        self.fields[index] = put
    }
    func debugString() -> String {
        return "SObject(\(self.clazz.name))"
    }
}
