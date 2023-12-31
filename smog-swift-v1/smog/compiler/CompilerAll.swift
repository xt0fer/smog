
//  CompilerAll.swift
//  smog

//  Created by Kristofer Younger on 8/17/23.


import Foundation


class Disassembler {}
//Disassembler = (

//  ----

//  dump: cl in: universe = (
//    1 to: cl numberOfInstanceInvokables do: [:i |
//      | inv |
//      inv := cl instanceInvokable: i.

//      "output header and skip if the Invokable is a Primitive"
//      Universe errorPrint: (cl name string + '>>' + inv signature string + ' = ').

//      inv isPrimitive
//        ifTrue: [ Universe errorPrintln: '<primitive>' ]
//        ifFalse: [ self dumpMethod: inv indent: '\t' in: universe ] ]
//  )

//  dumpInvokable: inv in: universe = (
//    | holderName |
//    holderName := inv holder == nil
//      ifTrue: ['nil']
//      ifFalse: [inv holder name string].
//    Universe errorPrint: (holderName + '>>#' + inv signature string + ' = ').
//    inv isPrimitive
//        ifTrue: [
//          Universe errorPrint: '<primitive>: '.
//          Universe errorPrintln: inv debugString ]
//        ifFalse: [ self dumpMethod: inv indent: '\t' in: universe ]
//  )

//  dumpMethod: m indent: indent in: universe = (
//    | b |
//    Universe errorPrintln: '('.

//    "output stack information"
//    Universe errorPrintln: indent + '<' + m numberOfLocals + ' locals, '
//        + m maximumNumberOfStackElements + ' stack, '
//        + m numberOfBytecodes + ' bc_count>'.

//    "output bytecodes"
//    b := 1.
//    [b <= m numberOfBytecodes] whileTrue: [
//      | bytecode |
//      Universe errorPrint: indent.

//      b < 10 ifTrue: [ Universe errorPrint: ' ' ].
//      b < 100 ifTrue: [ Universe errorPrint: ' ' ].

//      Universe errorPrint: ' ' + b + ':'.

//      "mnemonic"
//      bytecode := m bytecode: b.
//      Universe errorPrint: (Bytecodes paddedBytecodeName: bytecode) + '  '.

//      "parameters (if any)"
//      (Bytecodes length: bytecode) = 1
//        ifTrue: [ Universe errorPrintln ]
//        ifFalse: [ self dumpBytecode: bytecode idx: b method: m indent: indent in: universe ].

//      b := b + (Bytecodes length: (m bytecode: b)) ].

//    Universe errorPrintln: indent + ')'
//  )

//  dumpBytecode: bc idx: b method: m indent: indent in: universe = (
//    bc == #pushLocal ifTrue: [
//      Universe errorPrintln: 'local: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushArgument ifTrue: [
//      Universe errorPrintln: 'argument: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushField ifTrue: [
//      | idx fieldName |
//      idx := m bytecode: b + 1.
//      fieldName := (m holder instanceFields indexableField: idx) string.
//      Universe errorPrintln: '(index: ' + idx + ') field: ' + fieldName.
//      ^ self ].
//    bc == #pushBlock ifTrue: [
//      Universe errorPrint: '(block: (index: ' + (m bytecode: b + 1) + ') '.
//      self dumpMethod: (m constant: b) indent: indent + '\t' in: universe.
//      ^ self ].
//    bc == #pushConstant ifTrue: [
//      | constant |
//      constant := m constant: b.
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') value: '
//        + '(' + (constant somClassIn: universe) name string + ') ' + constant debugString.
//      ^ self ].
//    bc == #pushGlobal ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') value: #' + (m constant: b) string.
//      ^ self ].
//    bc == #popLocal ifTrue: [
//      Universe errorPrintln: 'local: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #popArgument ifTrue: [
//      Universe errorPrintln: 'argument: ' + (m bytecode: b + 1) + ', context: ' + (m bytecode: b + 2).
//      ^ self ].
//    bc == #pushField ifTrue: [
//      | idx fieldName |
//      idx := m bytecode: b + 1.
//      fieldName := (m holder instanceFields indexableField: idx) string.
//      Universe errorPrintln: '(index: ' + idx + ') field: ' + fieldName.
//      ^ self ].
//    bc == #send ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') signature: #' + (m constant: b) string.
//      ^ self ].
//    bc == #superSend ifTrue: [
//      Universe errorPrintln: '(index: ' + (m bytecode: b + 1) + ') signature: #' + (m constant: b) string.
//      ^ self ].

//    Universe errorPrintln: '<unknown bytecode>'
//  )
//)

// LEXER
// LEXER
// LEXER
// LEXER

// MethodGenerationContext
// MethodGenerationContext
// MethodGenerationContext
// MethodGenerationContext



//PARSER

//SourcecodeCompiler
//SourcecodeCompiler
//SourcecodeCompiler

