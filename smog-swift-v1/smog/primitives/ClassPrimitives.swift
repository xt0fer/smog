//
//  ClassPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

//ClassPrimitives = Primitives (
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: 'new' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInstance: rcvr) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'name' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: rcvr name ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'superclass' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: rcvr superClass ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'fields' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: rcvr instanceFields ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'methods' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: rcvr instanceInvokables ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
