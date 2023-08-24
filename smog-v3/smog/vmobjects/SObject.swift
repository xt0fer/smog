//
//  sample.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation


class SObject: Sender, Debuggable, Hashable {
    
    var fields: [SObject] = []
    var clazz: SClass
    var debugSDesc = "SObject()"

    convenience init() {
        self.init(nArgs: 0, clazz: Universe.shared.objectClass)
    }
    
    var identifier: String {
        return UUID().uuidString
    }
    
    public func hash(into hasher: inout Hasher) {
        return hasher.combine(identifier)
    }
    
    public static func == (lhs: SObject, rhs: SObject) -> Bool {
        return lhs.identifier == rhs.identifier
    }

    init(nArgs: Int, clazz: SClass) {
        self.fields = Array(repeating: Universe.shared.nilObject, count: nArgs)
        self.clazz = clazz
    }

    func send(_ selectorString: String, withArguments: [SObject], in: Universe, using: Interpreter) {
        
    }
    
    func sendDoesNotUnderstand(_ selector: String, in: Universe, using: Interpreter) {
        
    }
    
    func sendUnknownGlobal(_ globalName: String, in: Universe, using: Interpreter) {
        
    }
    
    func sendEscapedBlock(_ block: SBlock, in: Universe, using: Interpreter) {
    
    }
    
    func error(_ e: String) {
        print(e)
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
    
    func asString() -> String {
        return self.clazz.name.s
    }

    func fieldName(index: Int) -> SSymbol {
        return self.clazz.instanceFieldName(index: index)
    }
    func fieldIndex(name: SSymbol) -> Int {
        return self.clazz.lookupFieldIndex(fieldName: name)
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
        return self.debugSDesc + "\(self.identifier)"
    }
}
