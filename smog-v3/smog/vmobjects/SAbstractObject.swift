//
//  AbstractObject.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class SAbstractObject {

    func send(_ selectorString: String, withArguments: [SObject], in: Universe, using: Interpreter ) {
        
    }
//send: selectorString with: arguments in: universe using: interpreter = (
//  | selector invokable |
//  selector := universe symbolFor: selectorString.
//
//  interpreter frame push: self.
//
//  arguments do: [:arg |
//    interpreter frame push: arg ].
//
//  invokable := (self somClassIn: universe) lookupInvokable: selector.
//
//  invokable invoke: interpreter frame using: interpreter
//)
//
    func sendDoesNotUnderstand(_ selector: String, in: Universe, using: Interpreter) {}
//sendDoesNotUnderstand: selector in: universe using: interpreter = (
//  | numberOfArguments frame argumentsArray args |
//  numberOfArguments := selector numberOfSignatureArguments.
//
//  frame := interpreter frame.
//  frame printStackTrace.
//
//  "Allocate an array with enough room to hold all arguments
//   except for the receiver, which is passed implicitly, as receiver of #dnu."
//  argumentsArray := universe newArray: numberOfArguments - 1.
//
//  "Remove all arguments and put them in the freshly allocated array"
//  numberOfArguments - 1 downTo: 1 do: [:i |
//    argumentsArray indexableField: i put: frame pop ].
//
//  frame pop. "pop receiver"
//
//  args := Array with: selector with: argumentsArray.
//  self send: 'doesNotUnderstand:arguments:' with: args in: universe using: interpreter
//)
//
    func sendUnknownGlobal(_ globalName: String, in: Universe, using: Interpreter) {}

//sendUnknownGlobal: globalName in: universe using: interpreter = (
//  | arguments |
//  arguments := Array with: globalName.
//  self send: 'unknownGlobal:' with: arguments in: universe using: interpreter
//)
//
    func sendEscapedBlock(_ block: SBlock, in: Universe, using: Interpreter) {}

//sendEscapedBlock: block in: universe using: interpreter = (
//  | arguments |
//  arguments := Array with: block.
//  self send: 'escapedBlock:' with: arguments in: universe using: interpreter
//)
//)

}
