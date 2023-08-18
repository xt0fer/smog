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
    
    //SArray = SAbstractObject (
    //  | indexableFields |
    var indexableFields: [SAbstractObject] = []
    //
    //  initializeWith: length and: nilObject = (
    //    indexableFields := Array new: length withAll: nilObject.
    //  )
    func initializeWith(_ length: Int, and: SAbstractObject) {
        
    }
    //
    //  somClassIn: universe = (
    //    ^ universe arrayClass
    //  )
    //
    //  indexableField: idx = (
    //    ^ indexableFields at: idx
    //  )
    //
    //  indexableField: idx put: val = (
    //    ^ indexableFields at: idx put: val
    //  )
    func indexableField(_ idx: Int, put: SAbstractObject) {
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
    //
    //  copyIndexableFieldsTo: destination = (
    //    indexableFields doIndexes: [:i |
    //      destination indexableField: i put: (indexableFields at: i) ]
    //  )
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
