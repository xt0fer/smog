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
    
    //  compileClass: stmt into: systemClass in: universe = (
    //    | parser |
    //    parser := Parser newWith: stmt for: '$string$' in: universe.
    //    ^ self compile: parser into: systemClass.
    //  )
    
    //  compile: parser into: systemClass = (
    //    | cgc |
    //    cgc := parser classdef.
    
    //    systemClass == nil
    //      ifTrue: [ ^ cgc assemble ]
    //      ifFalse: [ ^ cgc assembleSystemClass: systemClass ]
    //  )
    //)
}
