//
//  Bytescodes.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

//Bytecodes = (
//  ----
//
//  length: bytecode = (
//    bytecode == #halt           ifTrue: [ ^ 1 ].
//    bytecode == #dup            ifTrue: [ ^ 1 ].
//    bytecode == #pushLocal      ifTrue: [ ^ 3 ].
//    bytecode == #pushArgument   ifTrue: [ ^ 3 ].
//    bytecode == #pushField      ifTrue: [ ^ 2 ].
//    bytecode == #pushBlock      ifTrue: [ ^ 2 ].
//    bytecode == #pushConstant   ifTrue: [ ^ 2 ].
//    bytecode == #pushGlobal     ifTrue: [ ^ 2 ].
//    bytecode == #pop            ifTrue: [ ^ 1 ].
//    bytecode == #popLocal       ifTrue: [ ^ 3 ].
//    bytecode == #popArgument    ifTrue: [ ^ 3 ].
//    bytecode == #popField       ifTrue: [ ^ 2 ].
//    bytecode == #send           ifTrue: [ ^ 2 ].
//    bytecode == #superSend      ifTrue: [ ^ 2 ].
//    bytecode == #returnLocal    ifTrue: [ ^ 1 ].
//    bytecode == #returnNonLocal ifTrue: [ ^ 1 ].
enum Bytecode: String {
    case halt = "#halt" //0
    case dup = "#dup"
    case pushLocal = "#pushLocal"
    case pushArgument = "#pushArgument"
    case pushField = "#pushField"
    case pushBlock = "#pushBlock" //5
    case pushConstant = "#pushConstant"
    case pushGlobal = "#pushGlobal"
    case pop = "#pop"
    case popLocal = "#popLocal"
    case popArgument = "#popArgument" //10
    case popField = "#popField"
    case send = "#send"
    case superSend = "#superSend"
    case returnLocal = "#returnLocal"
    case returnNonLocal = "#returnNonLocal" //15
}

enum Bc: Int { // with the BytecodeMap below, dunno if I need this anymore.
    case halt = 0
    case dup
    case pushLocal
    case pushArgument
    case pushField
    case pushBlock
    case pushConstant
    case pushGlobal
    case pop
    case popLocal
    case popArgument
    case popField
    case send
    case superSend
    case returnLocal
    case returnNonLocal
}

let BytecodeMap: [Bytecode: Int] = [
    .halt:0,
    .dup: 1,
    .pushLocal: 2,
    .pushArgument: 3,
    .pushField: 4,
    .pushBlock: 5,
    .pushConstant: 6,
    .pushGlobal: 7,
    .pop: 8,
    .popLocal: 9,
    .popArgument: 10,
    .popField: 11,
    .send: 12,
    .superSend: 13,
    .returnLocal: 14,
    .returnNonLocal: 15
]
let bytecodeArgs = [1,1,3,3,2,2,2,2,1,3,3,2,2,2,1,1]
let bytecodeLength = 16
//
//    self error: 'Unknown bytecode' + bytecode asString
//  )
//
//  paddedBytecodeName: bytecodeSymbol = (
//    | max padded |
//    max := #returnNonLocal length.
//    padded := bytecodeSymbol asString.
//    [padded length < max] whileTrue: [
//      padded := padded + ' ' ].
//    ^ padded
//  )
//)
func paddedBytecodeName(bytecodeSymbol: SSymbol) -> String {
    let max = Bytecode.returnNonLocal.rawValue.count
    var padded = bytecodeSymbol.s
    while padded.count < max {
        padded = padded + " "
    }
    return padded
}
