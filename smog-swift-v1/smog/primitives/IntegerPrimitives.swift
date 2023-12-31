//
//  IntegerPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

//IntegerPrimitives = Primitives (
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: 'asString' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newString: rcvr integer asString) ]).
//
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'sqrt' in: universe with: [:frame :interp |
//        | rcvr result |
//        rcvr := frame pop.
//        result := rcvr integer sqrt.
//        result class == Integer
//          ifTrue: [frame push: (universe newInteger: result)]
//          ifFalse: [frame push: (universe newDouble: result)] ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'atRandom' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr integer atRandom) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'asDouble' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr integer asDouble) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '+' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (argument class == SDouble
//          ifTrue: [universe newDouble: rcvr integer + argument double]
//          ifFalse: [universe newInteger: rcvr integer + argument integer]) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '-' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (argument class == SDouble
//          ifTrue: [universe newDouble: rcvr integer - argument double]
//          ifFalse: [universe newInteger: rcvr integer - argument integer]) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '*' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (argument class == SDouble
//          ifTrue: [universe newDouble: rcvr integer * argument double]
//          ifFalse: [universe newInteger: rcvr integer * argument integer]) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '//' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (universe newDouble:
//          (argument class == SDouble
//            ifTrue: [rcvr integer // argument double]
//            ifFalse: [rcvr integer // argument integer])) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '/' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (universe newInteger:
//          (argument class == SDouble
//            ifTrue: [rcvr integer / argument double]
//            ifFalse: [rcvr integer / argument integer])) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '%' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (argument class == SDouble
//            ifTrue: [universe newDouble: rcvr integer % argument double]
//            ifFalse: [universe newInteger: rcvr integer % argument integer]) ]).
//
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'rem:' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer rem: argument integer))]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '&' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer & argument integer)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '=' in: universe with: [:frame :interp |
//        | argument rcvr left |
//        argument := frame pop.
//        rcvr := frame pop.
//        left := rcvr integer.
//
//        frame push: (self somBool: (
//          (argument class == SDouble)
//            ifTrue: [left = argument double]
//            ifFalse: [
//              argument class == SInteger
//                ifTrue: [left = argument integer]
//                ifFalse: [ false ] ])) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '==' in: universe with: [:frame :interp |
//        | argument rcvr left |
//        argument := frame pop.
//        rcvr := frame pop.
//        left := rcvr integer.
//
//        frame push: (self somBool: (
//          argument class == SInteger
//            ifTrue: [left = argument integer]
//            ifFalse: [ false ] )) ]) dontWarn: true.
//
//    self installInstancePrimitive: (
//      SPrimitive new: '<<' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer << argument integer)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '<' in: universe with: [:frame :interp |
//        | argument rcvr left |
//        argument := frame pop.
//        rcvr := frame pop.
//        left := rcvr integer.
//
//        frame push: (self somBool: (
//          (argument class == SDouble)
//            ifTrue: [left < argument double]
//            ifFalse: [
//              argument class == SInteger
//                ifTrue: [left < argument integer]
//                ifFalse: [ false ] ])) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'bitXor:' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer bitXor: argument integer)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'as32BitSignedValue' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer as32BitSignedValue)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'as32BitUnsignedValue' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer as32BitUnsignedValue)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '>>>' in: universe with: [:frame :interp |
//        | argument rcvr |
//        argument := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newInteger: (rcvr integer >>> argument integer)) ]).
//
//    self installClassPrimitive: (
//      SPrimitive new: 'fromString:' in: universe with: [:frame :interp |
//        | argument |
//        argument := frame pop.
//        frame pop.
//        frame push: (universe newInteger: (Integer fromString: argument string)) ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
