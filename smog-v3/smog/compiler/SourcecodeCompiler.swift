//
//  SourcecodeCompiler.swift
//  smog
//
//  Created by Kristofer Younger on 8/27/23.
//

import Foundation

//SourcecodeCompiler = (
class SourcecodeCompiler {
    //  ----
    //  compileClass: path name: fileName into: systemClass in: universe = (
    //    | fname parser result cname |
    //    fname := path + '/' + fileName + '.som'.
    
    //    parser := Parser load: fname in: universe.
    //    parser ifNil: [ ^ nil ].
    
    //    result := self compile: parser into: systemClass.
    
    //    cname := result name string.
    
    //    fileName ~= cname ifTrue: [
    //      self error: 'File name ' + fname
    //          + ' does not match class name (' + cname + ') in it.' ].
    //    ^ result
    //  )
    func compileClass(_ path: String, name fileName: String, into systemClass: SClass, in universe: Universe) -> ClassGenerationContext? {
        let fname = path + "/" + fileName + ".som"
        let parser = Parser.load(fname, in: universe)
//        if parser.isNil() {
//            return nil
//        }
        let result = self.compile(parser, into: systemClass)
        
        let cname = result.name.asString()
        if fileName != cname {
            print("File name \(fname) does not match class name \(cname) in it.")
        }
        return result
    }
    
    //  compileClass: stmt into: systemClass in: universe = (
    //    | parser |
    //    parser := Parser newWith: stmt for: '$string$' in: universe.
    //    ^ self compile: parser into: systemClass.
    //  )
//    func compileClass(_ stmt: String, into systemclass: SClass, in universe: Universe) -> ClassGenerationContext? {
//        let parser = Parser.newWith(stmt, forString: "$string$", in: universe)
//        return self.compile(parser, into: systemclass)
//    }

    //  compile: parser into: systemClass = (
    //    | cgc |
    //    cgc := parser classdef.
    //    systemClass == nil
    //      ifTrue: [ ^ cgc assemble ]
    //      ifFalse: [ ^ cgc assembleSystemClass: systemClass ]
    //  )
    //)
    func compile(_ parser: Parser, into sysclass: SClass) -> ClassGenerationContext {
        var cgc = parser.classDef()
        
//        if sysclass.isNil() {
//            return cgc.assemble()
//        } else {
//            return cgc.assembleSystemClass(sysclass)
//        }
        
        return cgc
    }
}
