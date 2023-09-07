//
//  ClassGenerationContext.swift
//  smog
//
//  Created by Kristofer Younger on 8/28/23.
//

import Foundation

class ClassGenerationContext {
    static let nilGenc = ClassGenerationContext()
    
    var universe: Universe
    var name: SSymbol = SSymbol()
    var superName: SSymbol = SSymbol()
    var classSide: Bool = false
    var classFields: [String] = []
    var instanceFields: [String] = []
    var classMethods: [Invokable] = []
    var instanceMethods: [Invokable] = []
    
    //ClassGenerationContext = (
    //  | universe
    //    name superName
    //    classSide
    //    classFields instanceFields
    //    classMethods instanceMethods |
    
    //  initalize: aUniverse = (
    //    universe := aUniverse.
    //    classSide := false.
    //    classFields := Vector new.
    //    instanceFields := Vector new.
    //    classMethods := Vector new.
    //    instanceMethods := Vector new.
    //  )
    init(_ u :Universe) {
        self.universe = u
    }
    init() {
        self.universe = Universe.shared
    }
    
    //  name = (
    //    ^ name
    //  )
    
    //  name: aSSymbol = (
    //    name := aSSymbol
    //  )
    
    //  superName = (
    //    ^ superName
    //  )
    
    //  superName: aSymbol = (
    //    superName := aSymbol
    //  )
    
    //  instanceFieldsOfSuper: aSArrayOfFieldNames = (
    //    | numFields |
    //    numFields := aSArrayOfFieldNames numberOfIndexableFields.
    //    1 to: numFields do: [:i |
    //      instanceFields append: (aSArrayOfFieldNames indexableField: i) ]
    //  )
    func instanceFieldsOfSuper(_ aSArrayOfFieldNames: [String]) {
        for s in aSArrayOfFieldNames {
            instanceFields.append(s)
        }
    }
    
    //  classFieldsOfSuper: aSArrayOfFieldNames = (
    //    | numFields |
    //    numFields := aSArrayOfFieldNames numberOfIndexableFields.
    //    1 to: numFields do: [:i |
    //      classFields append: (aSArrayOfFieldNames indexableField: i) ]
    //  )
    func classFieldsOfSuper(_ aSArrayOfFieldNames: [String]) {
        for s in aSArrayOfFieldNames {
            classFields.append(s)
        }
    }

    //  addField: aSymbol = (
    //    classSide
    //      ifTrue: [classFields append: aSymbol]
    //      ifFalse: [instanceFields append: aSymbol]
    //  )
    func addField(_ s: String) {
        if classSide {
            classFields.append(s)
        } else {
            instanceFields.append(s)
        }
    }
    
    //  hasField: aSymbol = (
    //    ^ classSide
    //      ifTrue: [classFields contains: aSymbol]
    //      ifFalse: [instanceFields contains: aSymbol]
    //  )
    func hasField(_ s: String) -> Bool {
        if classSide {
            return classFields.contains(s)
        } else {
            return instanceFields.contains(s)
        }
    }

    //  fieldIndex: aSymbol = (
    //    ^ classSide
    //      ifTrue: [classFields indexOf: aSymbol]
    //      ifFalse: [instanceFields indexOf: aSymbol]
    //  )
    func fieldIndex(_ s: String) -> Int {
        if classSide {
            return classFields.firstIndex(of: s) ?? -1
        } else {
            return instanceFields.firstIndex(of: s) ?? -1
        }
    }

    //  addMethod: anInvokable = (
    //    classSide
    //      ifTrue: [classMethods append: anInvokable]
    //      ifFalse: [instanceMethods append: anInvokable]
    //  )
    func addMethod(_ m: Invokable) {
        if classSide {
            classMethods.append(m)
        } else {
            instanceMethods.append(m)
        }
    }

    //  startClassSide = (
    //    classSide := true
    //  )
    func startClassSide() {
        classSide = true
    }
    
    //  assemble = (
    //    | ccname superClass resultClass superMClass result |
    //    "build class class name"
    //    ccname := name string + ' class'.
    
    //    "Load the super class"
    //    superClass := universe loadClass: superName.
    
    //    "Allocate the class of the resulting class"
    //    resultClass := universe newClass: universe metaclassClass.
    
    //    "Initialize the class of the resulting class"
    //    resultClass instanceFields: (universe newArrayFromVector: classFields).
    //    resultClass instanceInvokables: (universe newArrayFromVector: classMethods).
    //    resultClass name: (universe symbolFor: ccname).
    
    //    superMClass := superClass somClass.
    //    resultClass superClass: superMClass.
    
    //    "Allocate the resulting class"
    //    result := universe newClass: resultClass.
    
    //    "Initialize the resulting class"
    //    result name: name.
    //    result superClass: superClass.
    //    result instanceFields: (universe newArrayFromVector: instanceFields).
    //    result instanceInvokables: (universe newArrayFromVector: instanceMethods).
    
    //    ^ result
    //  )
    
    //  assembleSystemClass: systemClass = (
    //    | superMClass |
    //    systemClass instanceInvokables: (universe newArrayFromVector: instanceMethods).
    //    systemClass instanceFields: (universe newArrayFromVector: instanceFields).
    
    //    "class-bound == class-instance-bound"
    //    superMClass := systemClass somClass.
    //    superMClass instanceInvokables: (universe newArrayFromVector: classMethods).
    //    superMClass instanceFields: (universe newArrayFromVector: classFields).
    
    //    ^ systemClass
    //  )
    func assembleSystemClass(_ systemClass: SClass) {
//        systemClass.instanceInvokables( universe.newArrayFrom(instanceMethods))
//        systemClass.instanceFields( universe.newArrayFrom(instanceFields))
        
        
    }
    
    //  ----
    
    //  new: aUniverse = (
    //    ^ self new initalize: aUniverse
    //  )
    //)
}
