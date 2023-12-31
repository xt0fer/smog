//
//  StringPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation


//StringPrimitives = Primitives (
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: 'concatenate:' in: universe with: [:frame :interp |
//        | rcvr argument |
//        argument := frame pop.
//        rcvr := frame pop.
//
//        frame push: (universe newString: rcvr string + argument string) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'asSymbol' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe symbolFor: rcvr string) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'length' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr string length) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: '=' in: universe with: [:frame :interp |
//        | rcvr argument argCls |
//        argument := frame pop.
//        rcvr := frame pop.
//        
//        argCls := argument somClassIn: universe.
//
//        frame push: (self somBool:
//          ((argCls == universe stringClass or: [argCls == universe symbolClass])
//            and: [rcvr string = argument string])) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'primSubstringFrom:to:' in: universe with: [:frame :interp |
//        | rcvr from to |
//        to := frame pop.
//        from := frame pop.
//        rcvr := frame pop.
//
//        frame push: (universe newString: (rcvr string primSubstringFrom: from integer to: to integer)) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'hashcode' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr string hashcode) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'isWhiteSpace' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (self somBool: rcvr string isWhiteSpace) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'isLetters' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (self somBool: rcvr string isLetters) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'isDigits' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (self somBool: rcvr string isDigits) ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
