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
    case halt = "#halt"
    case dup = "#dup"
    case pushLocal = "#pushLocal"
    case pushArgument = "#pushArgument"
    case pushField = "#pushField"
    case pushBlock = "#pushBlock"
    case pushConstant = "#pushConstant"
    case pushGlobal = "#pushGlobal"
    case pop = "#pop"
    case popLocal = "#popLocal"
    case popArgument = "#popArgument"
    case popField = "#popField"
    case send = "#send"
    case superSend = "#superSend"
    case returnLocal = "#returnLocal"
    case returnNonLocal = "#returnNonLocal"
}

let bytecodeArgs = [1,1,3,3,2,2,2,2,1,3,3,2,2,2,1,1]

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
