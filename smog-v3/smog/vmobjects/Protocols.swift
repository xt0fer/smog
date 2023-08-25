//
//  AbstractObject.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

protocol Invokable {
    func isPrimitive() -> Bool
    func isNil() -> Bool
    func invoke(frame: Frame)
    func invoke(frame: Frame, using: Interpreter)
    func signature() -> SSymbol
    func holder() -> SClass
    mutating func holder(value: SClass)

}
protocol Debuggable{
    var debugSDesc: String { get }
    func debugString() -> String
}

protocol Sender {
    func error(_ e: String)
    func send(_ selectorString: String, withArguments: [SObject], in: Universe, using: Interpreter )
    func sendDoesNotUnderstand(_ selector: String, in: Universe, using: Interpreter)
    func sendUnknownGlobal(_ globalName: String, in: Universe, using: Interpreter)
    func sendEscapedBlock(_ block: SBlock, in: Universe, using: Interpreter)
}
