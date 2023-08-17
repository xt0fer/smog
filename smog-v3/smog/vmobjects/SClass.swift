//
//  SClass.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SClass: SObject {
    
//    SClass = SObject (
//      | universe
//        superClass
//        name
//        instanceInvokables
//        instanceFields
//      |
//
    
//      ----
//
//      new: universe = (
//        ^ self new initialize: universe
//      )
//
//      new: numberOfFields in: universe = (
//        ^ self new initialize: numberOfFields in: universe
//      )
//    )

//      initialize: aUniverse = (
//        universe := aUniverse
//      )
//
//      initialize: numberOfFields in: aUniverse = (
//        super initialize: numberOfFields with: aUniverse nilObject.
//        universe := aUniverse
//      )
//
//      superClass = (
//        ^ superClass
//      )
//
//      superClass: aSClass = (
//        superClass := aSClass
//      )
//
//      hasSuperClass = (
//        ^ superClass ~= universe nilObject
//      )
//
//      name = (
//        ^ name
//      )
//
//      name: aSSymbol = (
//        name := aSSymbol
//      )
//
//      instanceFields = (
//        ^ instanceFields
//      )
//
//      instanceFields: aSArray = (
//        instanceFields := aSArray
//      )
//
//      instanceInvokables = (
//        ^ instanceInvokables
//      )
//
//      instanceInvokables: aSArray = (
//        instanceInvokables := aSArray.
//
//        "Make sure this class is the holder of all invokables in the array"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          (instanceInvokables indexableField: i) holder: self ]
//      )
//
//      numberOfInstanceInvokables = (
//        ^ instanceInvokables numberOfIndexableFields
//      )
//
//      instanceInvokable: idx = (
//        ^ instanceInvokables indexableField: idx
//      )
//
//      instanceInvokable: idx put: aSInvokable = (
//        aSInvokable holder: self.
//        instanceInvokables indexableField: idx put: aSInvokable
//      )
//
//      lookupInvokable: signature = (
//        | invokable |
//
//        "Lookup invokable with given signature in array of instance invokables"
//        1 to: instanceInvokables numberOfIndexableFields do: [:i |
//          "Get the next invokable in the instance invokable array"
//          invokable := instanceInvokables indexableField: i.
//
//          "Return the invokable if the signature matches"
//          invokable signature == signature ifTrue: [
//            ^ invokable ] ].
//
//        "Traverse the super class chain by calling lookup on the super class"
//        self hasSuperClass ifTrue: [
//          invokable := superClass lookupInvokable: signature.
//          invokable ~= nil ifTrue: [
//            ^ invokable ] ].
//
//        "Invokable not found"
//        ^ nil
//      )
//
//      lookupFieldIndex: fieldName = (
//        "Lookup field with given name in array of instance fields"
//
//        self numberOfInstanceFields downTo: 1 do: [:i |
//          "Return the current index if the name matches"
//          fieldName == (self instanceFieldName: i)
//            ifTrue: [ ^ i ] ].
//
//        "Field not found"
//        ^ -1
//      )
//
//      addInstanceInvokable: value = (
//        "Add the given invokable to the array of instance invokables"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          "Get the next invokable in the instance invokable array"
//          | invokable |
//          invokable := self instanceInvokable: i.
//
//          "Replace the invokable with the given one if the signature matches"
//          invokable signature == value signature ifTrue: [
//            self instanceInvokable: i put: value.
//            ^ false ] ].
//
//        "Append the given method to the array of instance methods"
//        instanceInvokables := instanceInvokables copyAndExtendWith: value in: universe.
//        ^ true
//      )
//
//      addInstancePrimitive: value = (
//        self addInstancePrimitive: value dontWarn: false
//      )
//
//      addInstancePrimitive: value dontWarn: suppressWarning = (
//        value holder: self.
//        ((self addInstanceInvokable: value) and: [suppressWarning not]) ifTrue: [
//          Universe print: 'Warning: Primitive ' + value signature string.
//          Universe println: ' is not in class definition for class ' + name string ]
//      )
//
//      instanceFieldName: index = (
//        "Get the name of the instance field with the given index"
//        index > self numberOfSuperInstanceFields
//          ifTrue: [
//            | idx |
//            "Adjust the index to account for fields defined in the super class"
//            idx := index - self numberOfSuperInstanceFields.
//
//            "Return the symbol representing the instance fields name"
//            ^ instanceFields indexableField: idx ]
//          ifFalse: [
//            "Ask the super class to return the name of the instance field"
//            ^ superClass instanceFieldName: index ]
//      )
//
//      numberOfInstanceFields = (
//        "Get the total number of instance fields in this class"
//        ^ instanceFields numberOfIndexableFields + self numberOfSuperInstanceFields
//      )
//
//      numberOfSuperInstanceFields = (
//        self hasSuperClass
//          ifTrue: [ ^ self superClass numberOfInstanceFields ]
//          ifFalse: [ ^ 0 ]
//      )
//
//      hasPrimitives = (
//        "Lookup invokable with given signature in array of instance invokables"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          "Get the next invokable in the instance invokable array"
//          (self instanceInvokable: i) isPrimitive
//            ifTrue: [ ^ true ] ].
//        ^ false
//      )
//
//      loadPrimitives = (
//        | className primsClass |
//        className := (name string + 'Primitives') asSymbol.
//
//        "Try loading the primitives"
//
//        primsClass := system load: className.
//        primsClass ~= nil
//          ifTrue: [
//            (primsClass new: universe) installPrimitivesIn: self ]
//          ifFalse: [
//            Universe println: 'Primitives class ' + className + ' not found' ]
//      )
//
//
//      "For using in debugging tools such as the Diassembler"
//      debugString = ( ^ 'SClass(' + name string + ')' )
//

    
}
