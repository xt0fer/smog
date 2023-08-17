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
