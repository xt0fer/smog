//
//  SInteger.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SInteger: SAbstractObject {}

//SInteger = SAbstractObject (
//  | value |
//
//  initialize: anInteger = (
//    value := anInteger
//  )
//
//  integer = ( ^ value )
//
//  somClassIn: universe = (
//    ^ universe integerClass
//  )
//
//  "For using in debugging tools such as the Diassembler"
//  debugString = ( ^ 'SInteger(' + value asString + ')' )
//
//  ----
//
//  "TODO: see whether it makes sense to have a cache"
//  for: anInteger = (
//    ^ self new initialize: anInteger
//  )
//)
