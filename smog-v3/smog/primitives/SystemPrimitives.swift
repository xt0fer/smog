//
//  SystemPrimitives.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

//SystemPrimitives = Primitives (
//  installPrimitives = (
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'load:' in: universe with: [:frame :interp |
//        | arg result |
//        arg := frame pop.
//        frame pop.
//
//        result := universe loadClass: arg.
//
//        frame push: (result == nil
//          ifTrue: [ universe nilObject ]
//          ifFalse: [ result ]) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'exit:' in: universe with: [:frame :interp |
//        | error |
//        frame printStackTrace.
//        error := frame pop.
//        universe exit: error integer ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'global:' in: universe with: [:frame :interp |
//        | argument result |
//        argument := frame pop.
//        frame pop.
//
//        result := universe global: argument.
//        frame push: (result == nil
//          ifTrue: [ universe nilObject ]
//          ifFalse: [ result ]) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'global:put:' in: universe with: [:frame :interp |
//        | value argument |
//        value := frame pop.
//        argument := frame pop.
//
//        universe global: argument put: value ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'printString:' in: universe with: [:frame :interp |
//        | arg |
//        arg := frame pop.
//        "Universe print: arg somClass asString."
//        Universe print: arg string ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'printNewline' in: universe with: [:frame :interp |
//        Universe println ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'errorPrint:' in: universe with: [:frame :interp |
//        | arg |
//        arg := frame pop.
//        Universe errorPrint: arg string ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'errorPrintln:' in: universe with: [:frame :interp |
//        | arg |
//        arg := frame pop.
//        Universe errorPrintln: arg string ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'time' in: universe with: [:frame :interp |
//        | time |
//        frame pop. "ignore"
//        time := system time.
//        frame push: (universe newInteger: time) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'ticks' in: universe with: [:frame :interp |
//        | ticks |
//        frame pop. "ignore"
//        ticks := system ticks.
//        frame push: (universe newInteger: ticks) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'gcStats' in: universe with: [:frame :interp |
//        | gcStats arr |
//        frame pop. "ignore"
//        gcStats := system gcStats.
//        arr := universe newArray: 3.
//        arr indexableField: 1 put: (universe newInteger: (gcStats at: 1)).
//        arr indexableField: 2 put: (universe newInteger: (gcStats at: 2)).
//        arr indexableField: 3 put: (universe newInteger: (gcStats at: 3)).
//
//        frame push: arr ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'totalCompilationTime' in: universe with: [:frame :interp |
//        | cTime |
//        frame pop. "ignore"
//        cTime := system totalCompilationTime.
//        frame push: (universe newInteger: cTime) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'fullGC' in: universe with: [:frame :interp |
//        frame pop. "ignore"
//        system fullGC.
//        frame push: (universe trueObject) ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'loadFile:' in: universe with: [:frame :interp |
//        | fileName content |
//        fileName := frame pop.
//        frame pop.
//
//        content := system loadFile: fileName string.
//        content == nil
//          ifTrue: [frame push: universe nilObject]
//          ifFalse: [frame push: (universe newString: content)] ]).
//
//    self installInstancePrimitive: (
//      SPrimitive new: 'printStackTrace' in: universe with: [:frame :interp |
//        frame pop. "ignore"
//        frame printStackTrace.
//        frame push: (universe trueObject) ]).
//  )
//
//  ----
//
//  new: universe = (
//    ^ self new initialize: universe
//  )
//)
