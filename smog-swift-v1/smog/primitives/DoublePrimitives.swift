//
//  DoublePrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation


//DoublePrimitives = Primitives (
//
//  coerceToDouble: anSAbstractObject = (
//    anSAbstractObject class == SDouble ifTrue: [
//      ^ anSAbstractObject double ].
//    anSAbstractObject class == SInteger ifTrue: [
//      ^ anSAbstractObject integer asDouble ].
//    self error: 'Cannot coerce ' + anSAbstractObject debugString + ' to double'.
//  )
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: 'asString' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newString: rcvr double asString) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'asInteger' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr double asInteger) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'sqrt' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double sqrt) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '+' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double + arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '-' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double - arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '*' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double * arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '//' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double // arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '%' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double % arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '=' in: universe with: [:frame :interp |
//        | argument rcvr left |
//        argument := frame pop.
//        rcvr := frame pop.
//        left := rcvr double.
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
//      SPrimitive new: '<' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := self coerceToDouble: frame pop.
//        rcvr := frame pop.
//        frame push: (self somBool: rcvr double < arg) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'round' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr double round) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'sin' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double sin) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'cos' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newDouble: rcvr double cos) ]).
//
//    self installClassPrimitive: (
//      SPrimitive new: 'PositiveInfinity' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newDouble: Double PositiveInfinity) ]).
//
//    self installClassPrimitive: (
//      SPrimitive new: 'fromString:' in: universe with: [:frame :interp |
//        | rcvr arg |
//        arg  := frame pop.
//        rcvr := frame pop.
//        frame push: (universe newDouble: (Double fromString: arg string)) ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
