
//  Universe.swift
//  smog

//  Created by Kristofer Younger on 8/17/23.


import Foundation

class Universe {
    static let shared: Universe = {
        let instance = Universe()
        // setup code
        return instance
    }()
    
    init(){
        //      symbolTable = Dictionary new.
        //      globals = Dictionary new.
        //      interpreter = Interpreter new: self.
        self.dumpBytecodes = false
        self.avoidExit = false
        
    }
    
    
    
    
    //    symbolTable
    var symbolTable: [String : SSymbol] = [:]
    //    globals
    var globals: [SSymbol : SObject] = [:]
    var    classPath = [".", "classes", "src"]
    var    dumpBytecodes = false
    var    interpreter: Interpreter
    
    var avoidExit = false
    var lastExitCode = 0
    var exitBlock: SBlock
    
    var nilObject: SObject
    var trueObject: SObject
    var falseObject: SObject
    
    var objectClass: SClass
    var classClass: SClass
    var metaclassClass: SClass
    
    var nilClass: SClass
    var integerClass: SClass
    var arrayClass: SClass
    var methodClass: SClass
    var symbolClass: SClass
    var primClass: SClass
    var stringClass: SClass
    var systemClass: SClass
    var blockClass: SClass
    var doubleClass: SClass
    
    var trueSymbol: SSymbol
    var falseSymbol: SSymbol
    var systemObject: SObject
    var trueClass: SClass
    var falseClass: SClass
    
    let defaultClassPath = "."
    
    func interpret(_ args: ArraySlice<String>) {
        
    }
    
    //    initialize = (
    //      symbolTable = Dictionary new.
    //      globals = Dictionary new.
    //      interpreter = Interpreter new: self.
    //      dumpBytecodes = false.
    //      avoidExit = false
    //    )
    
    //    initialize: aBool = (
    //      self initialize.
    //      avoidExit = aBool
    //    )
    
    func exit(_ errorCode: Int) {
        if self.avoidExit {
            self.lastExitCode = errorCode
            ///exitBlock.value(errorCode)
        } else {
            self.errorExit("Failed with code \(errorCode)")
        }
    }
    
    func errorExit(_ msg: String) {
        print("Runtime Error \(msg)")
        exit(-1)
    }
    
    //    setupClassPath: cp = (
    //      | paths cps |
    //      "Create a new tokenizer to split up the string of directories"
    //      paths = cp split: ':'.
    
    //      cps = Vector new.
    //      cps appendAll: self defaultClassPath.
    //      cps appendAll: paths.
    
    //      classPath = cps asArray
    //    )
    
    //    handleArguments: args = (
    //      | gotClasspath remainingArgs cnt i sawOthers |
    //      gotClasspath = false.
    //      remainingArgs = Vector new.
    
    //      "read dash arguments only while we haven't seen other kind of arguments"
    //      sawOthers = false.
    
    //      i = 1.
    
    //      [i <= args length] whileTrue: [
    //        ((args at: i) = '-cp' and: sawOthers not)
    //          ifTrue: [
    //            i + 1 > args length ifTrue: [
    //              self printUsageAndExit ].
    //            self setupClassPath: (args at: i + 1).
    //            i = i + 1.
    //            gotClasspath = true ]
    //          ifFalse: [
    //            ((args at: i) = '-d' and: sawOthers not)
    //              ifTrue: [ dumpBytecodes = true ]
    //              ifFalse: [
    //                sawOthers = true.
    //                remainingArgs append: (args at: i) ] ].
    //          i = i + 1 ].
    
    //      gotClasspath ifFalse: [
    //        classPath = self defaultClassPath ].
    
    //      remainingArgs isEmpty ifFalse: [
    //        | split |
    //        split = self pathClassExtension: (remainingArgs at: 1).
    //        (split at: 1) = '' ifFalse: [
    //          classPath = classPath prependedWith: (split at: 1) ].
    //        remainingArgs at: 1 put: (split at: 2) ].
    
    //      ^ remainingArgs asArray
    //    )
    func handleArguments(_ args: [String]) -> [String]{
        return args
    }
    
    //    pathClassExtension: str = (
    //      | pathElements fileName parentPath nameParts |
    //      pathElements = str split: '/'.
    //      fileName = pathElements last.
    
    //      parentPath = ''.
    //      1 to: pathElements length - 1 do: [:i |
    //        parentPath = '' ifFalse: [
    //          parentPath = parentPath + '/' ].
    //        parentPath = parentPath + (pathElements at: i) ].
    
    //      nameParts = fileName split: '.'.
    //      ^ Array with: parentPath with: (nameParts at: 1)
    //    )
    
    //    interpret: args = (
    //      | remainingArgs result |
    //      remainingArgs = self handleArguments: args.
    //      result = self initializeInterpreter: remainingArgs.
    //      result class == SInteger
    //        ifTrue: [ ^ result integer ]
    //        ifFalse: [ ^ 1 ]
    //    )
    
    func interpret(_ args: [String]) -> Int {
//        let remainingArgs = self.handleArguments(args)
//        let result = self.initializeInterpreter(remainingArgs)
//        if result.clazz.name == "SInteger" {
//            result.integer()
//        }
//        else {
//            return 1
//        }
    }
    
    //    interpret: className with: selector = (
    //      | clazz initialize |
    //      self initializeObjectSystem.
    
    //      clazz = self loadClass: (self symbolFor: className).
    
    //      "Lookup the initialize invokable on the system class"
    //      initialize = (clazz somClassIn: self) lookupInvokable: (self symbolFor: selector).
    
    //      initialize == nil ifTrue: [
    //        self error: 'Lookup of ' + className + '>>#' + selector + ' failed' ].
    
    //      ^ self interpret: initialize in: clazz with: nil
    //    )
    func interpret(_ className: String, with: String) {
        let syscls = self.initializeObjectSystem()
        var clazz = self.loadClass(clsname: self.symbolFor(className))
        let initialize = (clazz.somClassIn(self).lookupInvokable(signature: self.symbolFor(with)))
        if initialize.isNil() {
            self.errorExit("Lookup failed for \(with)")
        }
    }
    
    //    initializeInterpreter: arguments = (
    //      | systemObject initialize argumentsArray |
    //      systemObject = self initializeObjectSystem.
    
    //      "Start the shell if no filename is given"
    //      arguments length == 0 ifTrue: [
    //        | shell |
    //        shell = Shell for: self using: interpreter.
    //        shell bootstrapMethod: self createBootstrapMethod.
    //        ^ shell start ].
    
    //      "Lookup the initialize invokable on the system class"
    //      initialize = systemClass lookupInvokable: (self symbolFor: 'initialize:').
    
    //      "Convert the arguments into an array"
    //      argumentsArray = self newArrayFromStrings: arguments.
    
    //      ^ self interpret: initialize in: systemObject with: argumentsArray
    //    )
    func initializeInterpreter(_ arguments: [String]) {
        
    }
    
    //    createBootstrapMethod = (
    //      | bootstrapMethod |
    //      "Create a fake bootstrap method to simplify later frame traversal"
    //      bootstrapMethod = self newMethod: (self symbolFor: 'bootstrap')
    //        bc: #(#halt) literals: #() numLocals: 0 maxStack: 2.
    
    //      bootstrapMethod holder: systemClass.
    //      ^ bootstrapMethod
    //    )
    
    //    interpret: invokable in: receiver with: arguments = (
    //      | bootstrapMethod bootstrapFrame |
    //      exitBlock = [:errorCode | ^ errorCode ].
    
    //      bootstrapMethod = self createBootstrapMethod.
    
    //      "Create a fake bootstrap frame with the system object on the stack"
    //      bootstrapFrame = interpreter pushNewFrame: bootstrapMethod.
    //      bootstrapFrame push: receiver.
    
    //      arguments ~= nil ifTrue: [
    //        bootstrapFrame push: arguments ].
    
    //      "Invoke the initialize invokable"
    //      invokable invoke: bootstrapFrame using: interpreter.
    
    //      "Start the interpreter"
    //      ^ interpreter start
    //    )
    
    func initializeObjectSystem() -> SObject {
        
        //      "Allocate the nil object"
        self.nilObject = SObject()
        
        //      "Allocate the Metaclass classes"
        self.metaclassClass = self.newMetaclassClass()
        
        //      "Allocate the rest of the system classes"
        self.objectClass = self .newSystemClass()
        self.nilClass = self .newSystemClass()
        self.classClass = self .newSystemClass()
        self.arrayClass = self .newSystemClass()
        self.symbolClass = self .newSystemClass()
        self.methodClass = self .newSystemClass()
        self.integerClass = self .newSystemClass()
        self.primClass = self .newSystemClass()
        self.stringClass = self .newSystemClass()
        self.stringClass = self .newSystemClass()
        
        //      "Setup the class reference for the nil object"
        self.nilObject.somClass(aSClass: nilClass)
        
        //      "Initialize the system classes."
        self.initializeSystemClass(class: objectClass, superClass: nilClass, name: "Object")
        self.initializeSystemClass(class: classClass, superClass: objectClass, name: "Class")
        self.initializeSystemClass(class: metaclassClass, superClass: classClass, name: "Metaclass")
        self.initializeSystemClass(class: nilClass, superClass: objectClass, name: "Nil")
        self.initializeSystemClass(class: arrayClass, superClass: objectClass, name: "Array")
        self.initializeSystemClass(class: methodClass, superClass: arrayClass, name: "Method")
        self.initializeSystemClass(class: stringClass, superClass: objectClass, name: "String")
        self.initializeSystemClass(class: symbolClass, superClass: stringClass, name: "Symbol")
        self.initializeSystemClass(class: integerClass, superClass: objectClass, name: "Integer")
        self.initializeSystemClass(class: primClass, superClass: objectClass, name: "Primitive")
        self.initializeSystemClass(class: doubleClass, superClass: objectClass, name: "Double")
        
        //      "Load methods and fields into the system classes"
        self.loadSystemClass(cls: objectClass)
        self.loadSystemClass(cls: classClass)
        self.loadSystemClass(cls: metaclassClass)
        self.loadSystemClass(cls: nilClass)
        self.loadSystemClass(cls: arrayClass)
        self.loadSystemClass(cls: methodClass)
        self.loadSystemClass(cls: symbolClass)
        self.loadSystemClass(cls: integerClass)
        self.loadSystemClass(cls: primClass)
        self.loadSystemClass(cls: stringClass)
        self.loadSystemClass(cls: doubleClass)
        
        //      "Fix up objectClass"
        self.objectClass.superClass(put: nilObject.clazz)
        
        //      "Load the generic block class"
        self.blockClass = self.loadClass(clsname: self.symbolFor("Block"))
        
        //      "Setup the true and false objects"
        self.trueSymbol = self.symbolFor("True")
        self.trueClass = self.loadClass(clsname: trueSymbol)
        self.trueObject = self.newInstance(trueClass)
        
        self.falseSymbol = self.symbolFor("False")
        self.falseClass = self.loadClass(clsname: falseSymbol)
        self.falseObject = self.newInstance(falseClass)
        
        //      "Load the system class and create an instance of it"
        self.systemClass = self.loadClass(clsname: self.symbolFor("System"))
        self.systemObject = self.newInstance(systemClass)
        
        //      "Put special objects and classes into the dictionary of globals"
        self.global(self.symbolFor("nil"), put: nilObject)
        self.global(self.symbolFor("true"), put: trueObject)
        self.global(self.symbolFor("false"), put: falseObject)
        self.global(self.symbolFor("system"), put: systemObject)
        self.global(self.symbolFor("System"), put: systemClass)
        self.global(self.symbolFor("Block"), put: blockClass)
        self.global(trueSymbol,  put: trueClass)
        self.global(falseSymbol, put: falseClass)
        return systemObject
    }
    
    //    symbolFor: aString = (
    //      | result |
    //      result = symbolTable at: aString.
    //      result == nil ifFalse: [
    //        ^ result ].
    
    //      ^ self.newSymbol: aString
    //    )
    func symbolFor(_ s: String) -> SSymbol {
        if let result = self.symbolTable[s] {
            return result
        }
        return self.newSymbol(s)
    }
    
    //    newArray: size = (
    //      ^ SArray new: size with: nilObject
    //    )
    func newArray(size: Int) -> SArray {
        return SArray(size: size, with: nilObject)
    }
    
    //    newArrayFromStrings: strArray = (
    //      | sArr |
    //      sArr = self newArray: strArray length.
    //      1 to: strArray length do: [:i |
    //        sArr indexableField: i put: (self newString: (strArray at: i))].
    //      ^ sArr
    //    )
    func newArrayFrom(_ vector: [String]) -> SArray {
        let result = newArray(size: vector.count)
        for (i, v) in vector.enumerated() {
            result.indexableFields[i] = SString(s: v)
        }
        return result
    }

    //    newArrayFromVector: vector = (
    //      | result |
    //      "Allocate a new array with the same length as the list"
    //      result = self newArray: vector size.
    
    //      "Copy all elements from the list into the array"
    //      vector doIndexes: [:i |
    //        result indexableField: i put: (vector at: i) ].
    
    //      "Return the allocated and initialized array"
    //      ^ result
    //    )
//    func newArrayFrom(_ vector: [String]) -> SArray {
//        let result = newArray(size: vector.count)
//        for (i, v) in vector.enumerated() {
//            result.indexableFields[i] = self.symbolFor(v)
//        }
//        return result
//    }
    
    //    newBlock: method with: context numArgs: arguments = (
    //      ^ SBlock new: method in: context with: (self blockClass: arguments)
    //    )
    func newBlock(method: SMethod, with: Frame, numArgs: Int ) -> SBlock {
        return SBlock(aSMethod: method, aContext: with, aBlockClass: self.blockClass(numOfArgs: numArgs))
    }
    
    //    newClass: classClass = (
    //      | result |
    //      "Allocate a new class and set its class to be the given class class"
    //      result = SClass new: classClass numberOfInstanceFields in: self.
    //      result somClass: classClass.
    
    //      "Return the freshly allocated class"
    //      ^ result
    //    )
    func newClass(someClass: SClass) -> SClass {
        let result = SClass(someClass.numberOfFields(), u: self)
        result.somClass(aSClass: someClass)
        return result
    }
    
    //    newFrame: previousFrame with: method with: contextFrame = (
    //      | length result |
    //      "Compute the maximum number of stack locations (including arguments,
    //       locals and extra buffer to support doesNotUnderstand) and set the number
    //       of indexable fields accordingly"
    //      length = method numberOfArguments
    //          + method numberOfLocals
    //          + method maximumNumberOfStackElements + 2.
    
    //      result = Frame new: nilObject previous: previousFrame context: contextFrame method: method maxStack: length.
    
    //      "Return the freshly allocated frame"
    //      ^ result
    //    )
    func newFrame(previousFrame: Frame, method: Invokable, withContextFrame: Frame?) -> Frame {
        //      "Compute the maximum number of stack locations (including arguments,
        //       locals and extra buffer to support doesNotUnderstand) and set the number
        //       of indexable fields accordingly"
        if let method = method as? SMethod {
            let length = method.numberOfArguments()
            + method.numberOfLocals
            + method.maximumNumberOfStackElements + 2
            let result = Frame(with: nilObject, previousFrame: previousFrame, contextFrame: withContextFrame, method: method, maxStack: length)
            return result
        }
    }
    
    //    newSymbol: aString = (
    //      | result |
    //      result = SSymbol new: aString.
    //      symbolTable at: aString put: result.
    //      ^ result
    //    )
    
    func newSymbol(_ s: String) -> SSymbol {
        let ns = SSymbol(s: s)
        self.symbolTable[s] = ns
        return ns
    }
    
    //    newInstance: instanceClass = (
    //      | result |
    //      result = SObject new: instanceClass numberOfInstanceFields with: nilObject.
    //      result somClass: instanceClass.
    
    //      ^ result
    //    )
    func newInstance(_ sc: SClass) -> SObject {
        let result = SObject(nArgs: sc.numberOfInstanceFields(), clazz: sc)
        result.somClass(aSClass: sc)
        return result
    }
    
    //    newInteger: anInteger = (
    //      ^ SInteger for: anInteger
    //    )
    func newInteger(i: Int) -> SInteger {
        return SInteger(i: i)
    }
    
    //    newDouble: aDouble = (
    //      ^ SDouble for: aDouble
    //    )
    func newDouble(d: Float) -> SDouble {
        return SDouble(d: d)
    }
    //    newMetaclassClass = (
    //      | result |
    //      "Allocate the metaclass classes"
    //      result = SClass new: self.
    //      result somClass: (SClass new: self).
    
    //      "Setup the metaclass hierarchy"
    //      result somClass somClass: result.
    
    //      "Return the freshly allocated metaclass class"
    //      ^ result
    //    )
    func newMetaclassClass() -> SClass {
        let result = SClass(self)
        result.somClass(aSClass: SClass(self))
        
        result.somClass().somClass(aSClass: result)
        return result
    }
    
    //    newMethod: aSSymbol bc: bcArray literals: literalsArray numLocals: numLocals maxStack: maxStack = (
    //      ^ SMethod new: aSSymbol bc: bcArray literals: literalsArray numLocals: numLocals maxStack: maxStack
    //    )
    func newMethod(aSSymbol: SSymbol, bc: [Int], literals: [SString], numLocals: Int, maxStack: Int) -> SMethod {
        return SMethod(aSSymbol: aSSymbol, bc: bc, literals: literals, numLocals: numLocals, maxStack: maxStack)
    }
    
    //    newString: aString = (
    //      ^ SString new: aString
    //    )
    func newString(s: String) -> SString {
        return SString(s: s)
    }
    //    newSystemClass = (
    //      | symbolClass |
    //      "Allocate the new system class"
    //      systemClass = SClass new: self.
    
    //      "Setup the metaclass hierarchy"
    //      systemClass somClass: (SClass new: self).
    //      systemClass somClass somClass: metaclassClass.
    
    //      "Return the freshly allocated system class"
    //      ^ systemClass
    //    )
    func newSystemClass() -> SClass {
        systemClass = SClass(self)
        systemClass.somClass(aSClass: SClass(self))
        systemClass.somClass().somClass(aSClass: self.metaclassClass)
        return systemClass
    }
    
    //    initializeSystemClass: systemClass superClass: superClass name: name = (
    func initializeSystemClass(class: SClass, superClass: SClass, name: String) {
        
        
        //      "Initialize the superclass hierarchy"
        //      superClass ~= nil
        //        ifTrue: [
        //          systemClass superClass: superClass.
        //          systemClass somClass superClass: (superClass somClass) ]
        //        ifFalse: [
        //          systemClass somClass superClass: classClass ].
        
        //      "Initialize the array of instance fields"
        //      systemClass instanceFields: (self newArray: 0).
        //      systemClass somClass instanceFields: (self newArray: 0).
        
        //      "Initialize the array of instance invokables"
        //      systemClass instanceInvokables: (self newArray: 0).
        //      systemClass somClass instanceInvokables: (self newArray: 0).
        
        //      "Initialize the name of the system class"
        //      systemClass name: (self symbolFor: name).
        //      systemClass somClass name: (self symbolFor: name + " class").
        
        //      "Insert the system class into the dictionary of globals"
        //      self global: systemClass name put: systemClass.
        //    )
    }
    
    //    global: aSSymbol = (
    //      "Return the global with the given name if it's in the dictionary of globals"
    //      (self hasGlobal: aSSymbol) ifTrue: [
    //        ^ globals at: aSSymbol ].
    
    //      "Global not found"
    //      ^ nil
    //    )
    func global(_ symbol: SSymbol) -> SObject {
        return self.globals[symbol] ?? nilObject
    }
    
    //    global: aSSymbol put: aSAbstractObject = (
    //      "Insert the given value into the dictionary of globals"
    //      globals at: aSSymbol put: aSAbstractObject
    //    )
    func global(_ symbol: SSymbol, put no: SObject ) {
        self.globals[symbol] = no
    }
    
    //    hasGlobal: aSSymbol = (
    //      "Returns if the universe has a value for the global of the given name"
    //      ^ globals containsKey: aSSymbol
    //    )
    func hasGlobal(_ symbol: SSymbol) -> Bool {
        return self.globals.index(forKey: symbol) != nil
    }
    
    //    blockClass: numberOfArguments = (
    //      | name result |
    //      "Determine the name of the block class with the given number of arguments"
    //      name = self symbolFor: 'Block' + numberOfArguments.
    
    //      "Lookup the block class in the dictionary of globals and return it"
    //      (self hasGlobal: name) ifTrue: [
    //        ^ self global: name ].
    
    //      result = self loadClass: name into: nil.
    
    //      "Add the appropriate value primitive to the block class"
    //      result addInstancePrimitive:
    //        (SBlock evaluationPrimitive: numberOfArguments in: self).
    
    //      self global: name put: result.
    //      ^ result
    //    )
    func blockClass(numOfArgs: Int) -> SClass {
        let name = self.symbolFor("Block\(numOfArgs)")
        if self.hasGlobal(name) {
            return self.global(name) as! SClass
        }
        let result = self.loadClass(clsname: name, into: nilClass)
        
        result.addInstancePrimitive(SBlock.evaluationPrimitive(numOfArgs, universe: self))
        
    }
    
    //    loadClass: name = (
    //      | result |
    //      "Check if the requested class is already in the dictionary of globals"
    //      (self hasGlobal: name) ifTrue: [
    //        ^ self global: name ].
    
    //      "Load the class"
    //      result = self loadClass: name into: nil.
    
    //      "Load primitives (if necessary) and return the resulting class"
    //      (result ~= nil and: [result hasPrimitives]) ifTrue: [
    //        result loadPrimitives ].
    
    //      self global: name put: result.
    //      ^ result
    //    )
    func loadClass(clsname: SSymbol) -> SClass {
        
    }
    
    //    loadSystemClass: systemClass = (
    //      | result |
    //      "Load the system class"
    //      result = self loadClass: systemClass name into: systemClass.
    
    //      "Load primitives if necessary"
    //      result hasPrimitives ifTrue: [
    //        result loadPrimitives ].
    //    )
    func loadSystemClass(cls: SClass) {
        let result = self.loadClass(clsname: systemClass.name, into: systemClass)
        if result.hasPrimitives() {
            result.loadPrimitives()
        }
    }
    
    //    loadClass: name into: systemClass = (
    //      "Try loading the class from all different paths"
    //      classPath do: [:cpEntry |
    //        | result |
    //        "Load the class from a file and return the loaded class"
    //        result = SourcecodeCompiler compileClass: cpEntry name: name string into: systemClass in: self.
    
    //        (result notNil and: dumpBytecodes) ifTrue: [
    //          Disassembler dump: result somClass in: self.
    //          Disassembler dump: result in: self ].
    
    //        result ifNotNil: [ ^ result ] ].
    
    //      "The class could not be found."
    //      ^ nil
    //    )
    func loadClass(clsname: SSymbol, into: SClass) -> SClass {
        print("OOPS loadClass clsname:into:() not implemented")
        return objectClass
    }
    
    //    loadShellClass: stmt = (
    //      | result |
    //      "Load the class from a stream and return the loaded class"
    //      result = SourcecodeCompiler compileClass: stmt into: nil in: self.
    //      dumpBytecodes ifTrue: [
    //        Disassembler dump: result in: self ].
    //      ^ result
    //    )
    func loadShellClass(stmt: String) -> SClass {
        print("OOPS loadShellClass() not implemented")
        return objectClass
    }
    
    //    ----
    
    //    new = (
    //      ^ super new initialize
    //    )
    
    //    new: avoidExit = (
    //      ^ super new initialize: avoidExit
    //    )
    
    //    errorPrint: msg = (
    //      system errorPrint: msg
    //    )
    
    //    errorPrintln: msg = (
    //      system errorPrintln: msg
    //    )
    
    //    errorPrintln = (
    //      system errorPrintln: ''
    //    )
    
    //    print: msg = (
    //      system errorPrint: msg
    //    )
    
    //    println: msg = (
    //      system errorPrintln: msg
    //    )
    
    //    println = (
    //      system errorPrintln
    //    )
    //  )
    
}
