//
//  Interpreter.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class Interpreter {
    
    var u: Universe
    var frame: Frame?
    
    init(_ u: Universe) {
        self.u = u
        self.frame = nil //Universe.shared.nilObject
    }
    
    //  doDup = (
    //    frame push: (frame stackElement: 0)
    //  )
    func doDup() {
        if let frame = self.frame {
            frame.push(obj: (frame.stackElement(idx: 0)))
        }
    }
    //  doPushLocal: bytecodeIndex = (
    //    frame push: (
    //        frame local: (frame method bytecode: bytecodeIndex + 1)
    //                 at: (frame method bytecode: bytecodeIndex + 2))
    //  )
    func doPushLocal(idx: Int) {
        if let frame = self.frame {
            frame.push(
                obj: (frame.local(idx: (frame.method.bytecode(at: idx + 1)),
                                  at: (frame.method.bytecode(at: idx + 2)))))
        }
    }
    //
    //  doPushArgument: bytecodeIndex = (
    //    frame push: (
    //        frame argument: (frame method bytecode: bytecodeIndex + 1)
    //                    at: (frame method bytecode: bytecodeIndex + 2))
    //  )
    func doPushArgument(idx: Int) {
        if let frame = self.frame {
            frame.push(
                obj: (frame.argument(idx: (frame.method.bytecode(at: idx + 1)),
                                     at: (frame.method.bytecode(at: idx + 2)))))
        }
    }
    //
    //  doPushField: bytecodeIndex = (
    //    | fieldIndex |
    //    fieldIndex := frame method bytecode: bytecodeIndex + 1.
    //
    //    "Push the field with the computed index onto the stack"
    //    frame push: (self getSelf field: fieldIndex)
    //  )
    func doPushField(idx: Int) {
        if let frame = self.frame {
            let fieldIndex = frame.method.bytecode(at: idx + 1)
            
            frame.push(obj: self.getSelf().field(index: fieldIndex))
        }
    }
    
    //
    //  doPushBlock: bytecodeIndex = (
    //    | blockMethod |
    //    blockMethod := frame method constant: bytecodeIndex.
    //
    //    "Push a new block with the current frame as context onto the stack"
    //    frame push: (
    //        universe newBlock: blockMethod
    //                     with: frame
    //                  numArgs: blockMethod numberOfArguments)
    //  )
    func doPushBlock(idx: Int) {
        if let frame = self.frame {
            let blockMethod = frame.method.constant(bcIndex: idx) as! SMethod
            let univ = Universe.shared
            // "Push a new block with the current frame as context onto the stack"
            
            frame.push(
                obj: univ.newBlock(method: blockMethod , with: frame, numArgs: blockMethod.numberOfArguments())
            )
        }
    }
    
    //
    //  doPushConstant: bytecodeIndex = (
    //    frame push: (frame method constant: bytecodeIndex)
    //  )
    func doPushConstant(idx: Int) {
        if let frame = self.frame {
            frame.push(obj: frame.method.constant(bcIndex: idx))
        }
    }
    
    //
    //  doPushGlobal: bytecodeIndex = (
    //    | globalName global |
    //    globalName := frame method constant: bytecodeIndex.
    //
    //    "Get the global from the universe"
    //    global := universe global: globalName.
    //
    //    global ~= nil
    //      ifTrue: [ frame push: global  ]
    //      ifFalse: [
    //        "Send 'unknownGlobal:' to self"
    //        self getSelf sendUnknownGlobal: globalName in: universe using: self ]
    //  )
    func doPushGlobal(idx: Int) {
        if let frame = self.frame {
            let globalName = frame.method.constant(bcIndex: idx)
            let global = Universe.shared.global(globalName as! SSymbol)
            if global != Universe.shared.nilObject {
                frame.push(obj: global)
            } else {
                //        "Send 'unknownGlobal:' to self"
                let ns = globalName as! SSymbol
                self.getSelf().sendUnknownGlobal(ns.s, in: Universe.shared, using: self)
            }
        }
    }
    
    //
    //  doPop = (
    //    frame pop
    //  )
    func doPop() {
        if let frame = self.frame {
            _ = frame.pop()
        }
    }
    //
    //  doPopLocal: bytecodeIndex = (
    //    frame local: (frame method bytecode: bytecodeIndex + 1)
    //             at: (frame method bytecode: bytecodeIndex + 2)
    //            put: frame pop
    //  )
    func doPopLocal(idx: Int) {
        if let frame = self.frame {
            frame.local(idx: frame.method.bytecode(at: idx + 1),
                        at: frame.method.bytecode(at: idx + 2),
                        put: frame.pop())
        }
    }
    
    //
    //  doPopArgument: bytecodeIndex = (
    //    frame argument: (frame method bytecode: bytecodeIndex + 1)
    //                at: (frame method bytecode: bytecodeIndex + 2)
    //               put: frame pop
    //  )
    func doPopArgument(idx: Int) {
        if let frame = self.frame {
            frame.argument(idx: frame.method.bytecode(at: idx + 1),
                           at: frame.method.bytecode(at: idx + 2),
                           put: frame.pop())
        }
    }
    
    //
    //  doPopField: bytecodeIndex = (
    //    | fieldIndex |
    //    fieldIndex := frame method bytecode: bytecodeIndex + 1.
    //
    //    "Set the field with the computed index to the value popped from the stack"
    //    self getSelf field: fieldIndex put: frame pop
    //  )
    func doPopField(idx: Int) {
        if let frame = self.frame {
            let fieldIndex = frame.method.bytecode(at: idx + 1)
            //    "Set the field with the computed index to the value popped from the stack"
            self.getSelf().fieldAt(idx, put: frame.pop())
        }
    }
    
    //
    //  doSend: bytecodeIndex = (
    //    | signature numberOfArguments receiver |
    //    signature := frame method constant: bytecodeIndex.
    //    numberOfArguments := signature numberOfSignatureArguments.
    //    receiver := frame stackElement: numberOfArguments - 1.
    //    self send: signature rcvrClass: (receiver somClassIn: universe)
    //  )
    func doSend(idx: Int) {
        if let frame = self.frame {
            let signature = frame.method.constant(bcIndex: idx) as! SSymbol
            let numArgs = signature.numSignatureArguments
            let receiver = frame.stackElement(idx: numArgs - 1)
            self.send(selector: signature, rcvrClass: receiver.somClass())
        }
    }
    //
    //  doSuperSend: bytecodeIndex = (
    //    | signature holderSuper invokable |
    //    signature := frame method constant: bytecodeIndex.
    //
    //    "Send the message
    //     Lookup the invokable with the given signature"
    //    holderSuper := frame method holder superClass.
    //    invokable := holderSuper lookupInvokable: signature.
    //
    //    self activate: invokable orDnu: signature
    //  )
    func doSuperSend(idx: Int) {
        if let frame = self.frame {
            let signature = frame.method.constant(bcIndex: idx) as! SSymbol
            //    "Send the message
            //     Lookup the invokable with the given signature"
            let holderSuper = frame.method.holderClass.clazz.superClass
            let invokable = holderSuper.lookupInvokable(signature: signature)
            
            self.activate(invokable: invokable, orDnu: signature)
        }
    }
    //
    //  doReturnLocal = (
    //    | result |
    //    result := frame pop.
    //
    //    "Pop the top frame and push the result"
    //    self popFrameAndPushResult: result
    //  )
    func doReturnLocal() {
        if let frame = self.frame {
            let result = frame.pop()
            self.popFrameAndPushResult(result: result)
        }
    }
    //
    //  doReturnNonLocal = (
    //    | result context |
    //    result := frame pop.
    //
    //    "Compute the context for the non-local return"
    //    context := frame outerContext.
    //
    //    "Make sure the block context is still on the stack"
    //    context hasPreviousFrame ifFalse: [
    //      | block sender method numArgs |
    //      "Try to recover by sending 'escapedBlock:' to the sending object
    //       this can get a bit nasty when using nested blocks. In this case
    //       the 'sender' will be the surrounding block and not the object
    //       that actually sent the 'value' message."
    //      block := frame argument: 1 at: 0.
    //      sender := frame previousFrame outerContext argument: 1 at: 0.
    //
    //      "pop the frame of the currently executing block..."
    //      self popFrame.
    //
    //      "pop old arguments from stack"
    //      method := frame method.
    //      numArgs := method numberOfArguments.
    //      numArgs timesRepeat: [ frame pop ].
    //
    //      "... and execute the escapedBlock message instead"
    //      sender sendEscapedBlock: block in: universe using: self.
    //      ^ self ].
    //
    //    "Unwind the frames"
    //    [frame ~= context] whileTrue: [
    //      self popFrame ].
    //
    //    self popFrameAndPushResult: result
    //  )
    func doReturnNonLocal() {
        if let frame = self.frame {
            let result = frame.pop()
            let context = frame.outerContext()
            
            if context.hasPreviousFrame() == false {
                // "Try to recover by sending 'escapedBlock:' to the sending object
                // this can get a bit nasty when using nested blocks. In this case
                // the 'sender' will be the surrounding block and not the object
                // that actually sent the 'value' message."
                let block = frame.argument(idx: 1, at: 0) as! SBlock
                let sender = frame.previousFrame?.outerContext().argument(idx: 1, at: 0)
                
                _ = self.popFrame()
                
                let method = frame.method
                let numArgs = method.numberOfArguments()
                for _ in 0...numArgs { _ = frame.pop() }
                //      "... and execute the escapedBlock message instead"
                //      sender sendEscapedBlock: block in: universe using: self.
                sender!.sendEscapedBlock(block, in: Universe.shared, using: self)
                //return self
            }
            //    "Unwind the frames"
            //    [frame ~= context] whileTrue: [
            //      self popFrame ].
            while !(frame == context) {
                _ = self.popFrame()
            }
            self.popFrameAndPushResult(result: result)
        }
    }
    
    //
    //  start = (
    //    [true] whileTrue: [
    //      | bytecodeIndex bytecode bytecodeLength nextBytecodeIndex result |
    //      bytecodeIndex := frame bytecodeIndex.
    //      bytecode := frame method bytecode: bytecodeIndex.
    //      bytecodeLength := Bytecodes length: bytecode.
    //      nextBytecodeIndex := bytecodeIndex + bytecodeLength.
    //      frame bytecodeIndex: nextBytecodeIndex.
    //
    //      result := self dispatch: bytecode idx: bytecodeIndex.
    //      result ~= nil
    //        ifTrue: [ ^ result ] ]
    //  )
    
    func start() -> SObject {
        while true {
            let bytecodeIndex = self.frame!.bytecodeIndex
            let bytecode = self.frame!.method.bytecode(at: bytecodeIndex)
            
            let nextBytecodeIndex = bytecodeIndex + bytecodeLength // len is global
            self.frame!.bytecodeIndex(idx: nextBytecodeIndex)
            let result = self.dispatch(bc: bytecode, bytecodeIndex: bytecodeIndex)
            if !(result == Universe.shared.nilObject) {
                return result
            }
            
        }
    }
    
    //  dispatch: bytecode idx: bytecodeIndex = (
    func dispatch(bc: Int, bytecodeIndex: Int) -> SObject {
        let nilObject = Universe.shared.nilObject
        let bytecode = Bc(rawValue: bc)
        if bytecode == Bc.halt { return (frame?.stackElement(idx: bytecodeIndex))! }
        if bytecode == Bc.dup {
            self.doDup()
            return nilObject
        }
        if bytecode == Bc.pushLocal {
            self.doPushLocal(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pushArgument {
            self.doPushArgument(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pushField {
            self.doPushField(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pushBlock {
            self.doPushBlock(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pushConstant {
            self.doPushConstant(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pushGlobal {
            self.doPushGlobal(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.pop {
            self.doPop()
            return nilObject
        }
        if bytecode == Bc.popLocal {
            self.doPopLocal(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.popArgument {
            self.doPopArgument(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.popField {
            self.doPopField(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.send {
            self.doSend(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.superSend {
            self.doSuperSend(idx: bytecodeIndex)
            return nilObject
        }
        if bytecode == Bc.returnLocal {
            self.doReturnLocal()
            return nilObject
        }
        if bytecode == Bc.returnNonLocal {
            self.doReturnNonLocal()
            return nilObject
        }
        //    self error: 'Unknown bytecode' + bytecode asString
        print("Unknown bytecode \(bytecode!.rawValue)")
        return Universe.shared.nilObject
    }
    //
    //  pushNewFrame: method with: contextFrame = (
    //    frame := universe newFrame: frame with: method with: contextFrame.
    //    ^ frame
    //  )
    func pushNewFrame(invokable: Invokable, withContextFrame: Frame?) -> Frame {
        let f = Universe.shared.newFrame(previousFrame: self.frame!, method: invokable, withContextFrame: withContextFrame)
        return f
    }
    //
    //  pushNewFrame: method = (
    //    ^ self pushNewFrame: method with: nil
    //  )
    func pushNewFrame(invokable: Invokable) -> Frame {
        self.pushNewFrame(invokable: invokable, withContextFrame: nil)
    }
    
    //
    //  frame = (
    //    ^ frame
    //  )
    //
    //  method = (
    //    ^ frame method
    //  )
    //
    //  getSelf = (
    //    "Get the self object from the interpreter"
    //    ^ frame outerContext argument: 1 at: 0
    //  )
    func getSelf() -> SObject {
        //    "Get the self object from the interpreter"
        if let frame = self.frame {
            
            return frame.outerContext().argument(idx: 1, at: 0)
        }
        return Universe.shared.nilObject
    }
    //
    //  send: selector rcvrClass: receiverClass = (
    //    | invokable |
    //    invokable := receiverClass lookupInvokable: selector.
    //    self activate: invokable orDnu: selector
    //  )
    func send(selector: SSymbol, rcvrClass: SClass) {
        let invokable = rcvrClass.lookupInvokable(signature: selector)
        self.activate(invokable: invokable, orDnu: selector)
    }
    //
    //  activate: invokable orDnu: signature = (
    //    invokable ~= nil
    //        ifTrue: [
    //          "Invoke the invokable in the current frame"
    //          invokable invoke: frame using: self ]
    //        ifFalse: [
    //          | numberOfArguments receiver |
    //          numberOfArguments := signature numberOfSignatureArguments.
    //          receiver := frame stackElement: numberOfArguments - 1.
    //          receiver sendDoesNotUnderstand: signature in: universe using: self ]
    //  )
    func activate(invokable: Invokable, orDnu signature: SSymbol) {
        if let frame = self.frame {
            if invokable != nil {
                invokable.invoke(frame: frame, using: self )
            } else {
                let numberOfArguments = signature.numSignatureArguments
                let receiver = frame.stackElement(idx: numberOfArguments - 1)
                receiver.sendDoesNotUnderstand(signature.string(), in: self.u, using: self)
            }
        }
    }
    //
    //  popFrame = (
    //    | result |
    //    "Save a reference to the top frame"
    //    result := frame.
    //
    //    "Pop the top frame from the frame stack"
    //    frame := frame previousFrame.
    //
    //    "Destroy the previous pointer on the old top frame"
    //    result clearPreviousFrame.
    //
    //    "Return the popped frame"
    //    ^ result
    //  )
    func popFrame() -> Frame {
        let result = self.frame
        self.frame = self.frame!.previousFrame
        result?.clearPreviousFrame()
        return result!
    }
    //
    //  popFrameAndPushResult: result = (
    //    | numberOfArguments |
    //    "Pop the top frame from the interpreter frame stack and
    //     get the number of arguments"
    //    numberOfArguments := self popFrame method numberOfArguments.
    //
    //    "Pop the arguments"
    //    numberOfArguments
    //      timesRepeat: [ frame pop ].
    //
    //    frame push: result
    //  )
    func popFrameAndPushResult(result: SObject) {
        //    "Pop the top frame from the interpreter frame stack and
        //     get the number of arguments and pop them"
        let numArgs = self.popFrame().method.numberOfArguments()
        for i in 0...numArgs {
            _ = self.frame!.pop()
        }
        self.frame?.push(obj: result)
    }
    //
}
