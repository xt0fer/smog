//
//  ObjectPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation


//ObjectPrimitives = Primitives (
//
//  installPrimitives = (
//    self installInstancePrimitive: (
//      SPrimitive new: '==' in: universe with: [:frame :interp |
//        | op1 op2 |
//        op1 := frame pop.
//        op2 := frame pop.
//
//        frame push: (self somBool: op1 == op2) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'hashcode' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (universe newInteger: rcvr hashcode) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'objectSize' in: universe with: [:frame :interp |
//        | rcvr size clazz |
//        rcvr := frame pop.
//
//        size := 1.
//        clazz := (rcvr somClassIn: universe).
//        clazz == SArray ifTrue: [
//          size := size + rcvr numberOfIndexableFields ].
//        clazz == SObject ifTrue: [
//          size := size + rcvr numberOfFields ].
//
//        frame push: (universe newInteger: size) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'perform:' in: universe with: [:frame :interp |
//        | selector rcvr invokable |
//        selector := frame pop.
//        rcvr := frame stackElement: 0.
//
//        invokable := (rcvr somClassIn: universe) lookupInvokable: selector.
//        invokable invoke: frame using: interp ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'perform:inSuperclass:' in: universe with: [:frame :interp |
//        | selector clazz invokable |
//        clazz := frame pop.
//        selector := frame pop.
//
//        invokable := clazz lookupInvokable: selector.
//        invokable invoke: frame using: interp ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'perform:withArguments:' in: universe with: [:frame :interp |
//        | args selector rcvr invokable |
//        args := frame pop.
//        selector := frame pop.
//        rcvr := frame stackElement: 0.
//
//        1 to: args numberOfIndexableFields do: [:i |
//          frame push: (args indexableField: i) ].
//
//        invokable := (rcvr somClassIn: universe) lookupInvokable: selector.
//        invokable invoke: frame using: interp ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'instVarAt:' in: universe with: [:frame :interp |
//        | idx rcvr invokable |
//        idx := frame pop.
//        rcvr := frame pop.
//
//        frame push: (rcvr field: idx integer) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'instVarAt:put:' in: universe with: [:frame :interp |
//        | idx rcvr invokable val |
//        val := frame pop.
//        idx := frame pop.
//        rcvr := frame stackElement: 0.
//
//        rcvr field: idx integer put: val ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'class' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame pop.
//        frame push: (rcvr somClassIn: universe) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'halt' in: universe with: [:frame :interp |
//        | rcvr |
//        rcvr := frame stackElement: 0.
//        rcvr halt ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
