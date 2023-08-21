//
//  SClass.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SClass: SObject {
    
    var universe: Universe
    var superClass: SClass
    var name: SSymbol
    var _instanceInvokables: SArray = SArray()
    var instanceInvokables: SArray {
        get{
            return _instanceInvokables
        }
        set {
            _instanceInvokables = newValue
            // do a loop to init to nil
            for (index, value) in _instanceInvokables.indexableFields.enumerated() {
                _instanceInvokables.indexableFields[index].holder(value: self)
                invokeCache[value.signature()] = value as! any Invokable
            }
        }
    }
    var instanceFields: [SObject] = []
    var invokeCache: [SSymbol:Invokable] = [:]

    init(_ u: Universe) {
        self.universe = u
        super.init(nArgs: 0, clazz: u.classClass)
    }

    init(_ numberOfFields: Int, u: Universe) {
        self.universe = u
        super.init(nArgs: numberOfFields, clazz: u.classClass)
    }


//      superClass = {
//        ^ superClass
//      }
//    func superClass() -> SClass {
//        return self.clazz!
//    }

//      superClass: aSClass = {
//        superClass := aSClass
//      }
    func superClass(put nc: SClass) {
        self.superClass = nc
    }

//      hasSuperClass = {
//        ^ superClass ~= universe nilObject
//      }
        func hasSuperClass() -> Bool {
            return self.superClass != Universe.shared.nilObject
        }

//      name = {
//        ^ name
//      }

//      name: aSSymbol = {
//        name := aSSymbol
//      }

//      instanceFields = {
//        ^ instanceFields
//      }

//      instanceFields: aSArray = {
//        instanceFields := aSArray
//      }

//      instanceInvokables = {
//        ^ instanceInvokables
//      }

//      instanceInvokables: aSArray = {
//        instanceInvokables := aSArray.

//        "Make sure this class is the holder of all invokables in the array"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          (instanceInvokables indexableField: i) holder: self ]
//      }
    // SEE the variable SET leg.

//      numberOfInstanceInvokables = {
//        ^ instanceInvokables numberOfIndexableFields
//      }
    func numberOfInstanceInvokables() -> Int {
        return self._instanceInvokables.fields.count
    }
//      instanceInvokable: idx = {
//        ^ instanceInvokables indexableField: idx
//      }
    func instanceInvokable(index: Int) -> Invokable {
        return self._instanceInvokables.indexableFields[index] as! Invokable
    }

//      instanceInvokable: idx put: aSInvokable = {
//        aSInvokable holder: self.
//        instanceInvokables indexableField: idx put: aSInvokable
//      }
    func instanceInvokable(index: Int, put ni: Invokable) {
        //ni.holder(value: self)
        instanceInvokables.indexableField(index, put: ni as! SObject)
    }

//      lookupInvokable: signature = {
//        | invokable |

//        "Lookup invokable with given signature in array of instance invokables"
//        1 to: instanceInvokables numberOfIndexableFields do: [:i |
//          "Get the next invokable in the instance invokable array"
//          invokable := instanceInvokables indexableField: i.

//          "Return the invokable if the signature matches"
//          invokable signature == signature ifTrue: [
//            ^ invokable ] ].

//        "Traverse the super class chain by calling lookup on the super class"
//        self hasSuperClass ifTrue: [
//          invokable := superClass lookupInvokable: signature.
//          invokable ~= nil ifTrue: [
//            ^ invokable ] ].

//        "Invokable not found"
//        ^ nil
//      }
    func lookupInvokable(signature: SSymbol) -> Invokable {
        //        "Lookup invokable with given signature in array of instance invokables"
        
    }

//      lookupFieldIndex: fieldName = {
//        "Lookup field with given name in array of instance fields"

//        self numberOfInstanceFields downTo: 1 do: [:i |
//          "Return the current index if the name matches"
//          fieldName == (self instanceFieldName: i}
//            ifTrue: [ ^ i ] ].

//        "Field not found"
//        ^ -1
//      }

//      addInstanceInvokable: value = {
//        "Add the given invokable to the array of instance invokables"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          "Get the next invokable in the instance invokable array"
//          | invokable |
//          invokable := self instanceInvokable: i.

//          "Replace the invokable with the given one if the signature matches"
//          invokable signature == value signature ifTrue: [
//            self instanceInvokable: i put: value.
//            ^ false ] ].

//        "Append the given method to the array of instance methods"
//        instanceInvokables := instanceInvokables copyAndExtendWith: value in: universe.
//        ^ true
//      }

//      addInstancePrimitive: value = {
//        self addInstancePrimitive: value dontWarn: false
//      }

//      addInstancePrimitive: value dontWarn: suppressWarning = {
//        value holder: self.
//        ((self addInstanceInvokable: value) and: [suppressWarning not]) ifTrue: [
//          Universe print: 'Warning: Primitive ' + value signature string.
//          Universe println: ' is not in class definition for class ' + name string ]
//      }

    func instanceFieldName(index: Int) -> SSymbol {
        
    }
//      instanceFieldName: index = {
//        "Get the name of the instance field with the given index"
//        index > self numberOfSuperInstanceFields
//          ifTrue: [
//            | idx |
//            "Adjust the index to account for fields defined in the super class"
//            idx := index - self numberOfSuperInstanceFields.

//            "Return the symbol representing the instance fields name"
//            ^ instanceFields indexableField: idx ]
//          ifFalse: [
//            "Ask the super class to return the name of the instance field"
//            ^ superClass instanceFieldName: index ]
//      }

//      numberOfInstanceFields = {
//        "Get the total number of instance fields in this class"
//        ^ instanceFields numberOfIndexableFields + self numberOfSuperInstanceFields
//      }
    func numberOfInstanceFields() -> Int {
        return self.instanceFields.count
    }

//      numberOfSuperInstanceFields = {
//        self hasSuperClass
//          ifTrue: [ ^ self superClass numberOfInstanceFields ]
//          ifFalse: [ ^ 0 ]
//      }
    func numberOfSuperInstanceFields() -> Int {
        if self.hasSuperClass() {
            return self.superClass!.numberOfInstanceFields()
        }
        return 0
    }

//      hasPrimitives = {
//        "Lookup invokable with given signature in array of instance invokables"
//        1 to: self numberOfInstanceInvokables do: [:i |
//          "Get the next invokable in the instance invokable array"
//          (self instanceInvokable: i) isPrimitive
//            ifTrue: [ ^ true ] ].
//        ^ false
//      }

//      loadPrimitives = {
//        | className primsClass |
//        className := (name string + 'Primitives') asSymbol.

//        "Try loading the primitives"

//        primsClass := system load: className.
//        primsClass ~= nil
//          ifTrue: [
//            (primsClass new: universe) installPrimitivesIn: self ]
//          ifFalse: [
//            Universe println: 'Primitives class ' + className + ' not found' ]
//      }


//      "For using in debugging tools such as the Diassembler"
//      debugString = ( ^ 'SClass(' + name string + ')' )

    override func debugString() -> String {
        return "SClass(\(String(describing: self.clazz?.name)))"
    }

    
}
