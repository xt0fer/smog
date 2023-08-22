//
//  SArray.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SArray: SObject {
    
    
    //  new: length with: nilObject = (
    //    ^ self new initializeWith: length and: nilObject
    //  )
    init(size: Int, with: SObject) {
        self.indexableFields = Array(repeating: with, count: size)
        super.init(nArgs: 0, clazz: Universe.shared.arrayClass)
    }
    
    convenience init() {
        self.init(size: 0, with:  Universe.shared.nilObject)
    }
    
    var indexableFields: [SObject] = []
    //
    //  initializeWith: length and: nilObject = (
    //    indexableFields := Array new: length withAll: nilObject.
    //  )
    func initializeWith(_ length: Int, and: SObject) {
        
    }
    //
    //  somClassIn: universe = (
    //    ^ universe arrayClass
    //  )
    //
    //  indexableField: idx = (
    //    ^ indexableFields at: idx
    //  )
    func indexableField(idx: Int) -> SObject{
        return self.indexableFields[idx]
    }
    //
    //  indexableField: idx put: val = (
    //    ^ indexableFields at: idx put: val
    //  )
    func indexableField(_ idx: Int, put: SObject) {
        self.indexableFields[idx] = put
    }
    //  numberOfIndexableFields = (
    //    ^ indexableFields length
    //  )
    func numberOfIndexableFields() -> Int {
        self.indexableFields.count
    }
    
    //  copyAndExtendWith: value in: universe = (
    //    | result newLength |
    //    newLength := indexableFields length + 1.
    //    result := universe newArray: newLength.
    //
    //    self copyIndexableFieldsTo: result.
    //
    //    result indexableField: newLength put: value.
    //    ^ result
    //  )
    func copyAndExtendWith(value: SObject, in u: Universe) {
        let newSize = self.indexableFields.count+1
        let result = u.newArray(size: newSize)
        self.copyIndexableFieldsTo(destination: result)
        result.indexableFields[newSize] = value
    }
    //
    //  copyIndexableFieldsTo: destination = (
    //    indexableFields doIndexes: [:i |
    //      destination indexableField: i put: (indexableFields at: i) ]
    //  )
    func copyIndexableFieldsTo(destination: SArray) {
        for (idx, field) in self.indexableFields.enumerated() {
            destination.indexableFields[idx] = field
        }
    }
    //
    //  "For using in debugging tools such as the Diassembler"
    //  debugString = (
    //    | elems |
    //    elems := ''.
    //    indexableFields do: [:e |
    //      elems = '' ifTrue: [elems := e debugString]
    //                 ifFalse: [ elems := elems + ', ' + e debugString] ].
    //     ^ 'SArray(' + indexableFields length + '; ' + elems + ')' )
    //
    //  ----
    //
    //)
}
