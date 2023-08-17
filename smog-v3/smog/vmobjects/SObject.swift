//
//  sample.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SObject: AbstractObject {
    var fields: [AbstractObject] = []
    var clazz: SClass
    
    init(nArgs: int, clazz: SClass) {
        self.fields = []
        self.clazz = clazz
    }
    
//    SObject = SAbstractObject (
//      | fields clazz |
//
//      initialize: numberOfFields with: nilObject = (
//        fields := Array new: numberOfFields withAll: nilObject
//      )
//
//      somClass = (
//        ^ clazz
//      )
//
//      somClass: aSClass = (
//        clazz := aSClass
//      )
//
//      somClassIn: universe = (
//        ^ clazz
//      )
//
//      fieldName: index = (
//        "Get the name of the field with the given index"
//        ^ clazz instanceFieldName: index
//      )
//
//      fieldIndex: name = (
//        "Get the index for the field with the given name"
//        ^ clazz lookupFieldIndex: name
//      )
//
//      numberOfFields = (
//        "Get the number of fields in this object"
//        ^ fields length
//      )
//
//      field: index = (
//        "Get the field with the given index"
//        ^ fields at: index
//      )
//
//      field: index put: value = (
//        "Set the field with the given index to the given value"
//        fields at: index put: value
//      )
//
//      "For using in debugging tools such as the Diassembler"
//      debugString = ( ^ 'SObject(' + clazz name string + ')' )
//
//      ----
//
//      new: numberOfFields with: nilObject = (
//        ^ self new initialize: numberOfFields with: nilObject
//      )
//    )

}
