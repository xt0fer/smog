//
//  ArrayPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

//ArrayPrimitives = Primitives (
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: 'at:' in: universe with: [:frame :interp |
//        | idx rcvr |
//        idx := frame pop.
//        rcvr := frame pop.
//        frame push: (rcvr indexableField: idx integer)]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'at:put:' in: universe with: [:frame :interp |
//        | rcvr idx value |
//        value := frame pop.
//        idx   := frame pop.
//        rcvr  := frame stackElement: 0.
//        rcvr indexableField: idx integer put: value ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'length' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//
//        frame push: (universe newInteger: rcvr numberOfIndexableFields) ]).
//
//    self installClassPrimitive: (
//      SPrimitive new: 'new:' in: universe with: [:frame :interp |
//        | arg |
//        arg := frame pop.
//        frame pop.
//
//        frame push: (universe newArray: arg integer) ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
//
