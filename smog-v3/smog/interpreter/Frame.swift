//
//  Frame.swift
//  smog
//
//  Created by Kristofer Younger on 8/17/23.
//

import Foundation

class Frame: SArray {
    
    //  "Points at the top element"
    var stackPointer = 0
    var bytecodeIndex = 0
    //
    //  "the offset at which local variables start"
    var localOffset = 0
    
    var method: SMethod
    var contextFrame: Frame?
    var previousFrame: Frame?
    var stack: SArray
    
    //|
    //  initialize: nilObject previous: prevFrame context: contextFrame method: aSMethod maxStack: stackElements = (
    //    previousFrame := prevFrame.
    //    context := contextFrame.
    //    method := aSMethod.
    //    stack := Array new: stackElements withAll: nilObject.
    //
    //    "Reset the stack pointer and the bytecode index"
    //    self resetStackPointer.
    //    bytecodeIndex := 1.
    //  )
    init(with: SObject, previousFrame: Frame, contextFrame: Frame, method: SMethod, maxStack: Int) {
        self.previousFrame = previousFrame
        self.contextFrame = contextFrame
        self.method = method
        self.stack = SArray(size: maxStack, with: with) // prob nilObject
        super.init(size: 0, with: Universe.shared.nilObject)
        
        self.resetStackPointer()
        self.bytecodeIndex = 1
    }

    // TODO: Lots to do in Frame

//Frame = (
//"
//Frame layout:
//+-----------------+
//| Arguments       | 1
//+-----------------+
//| Local Variables | <-- localOffset
//+-----------------+
//| Stack           | <-- stackPointer
//| ...             |
//+-----------------+
//"
//|
//
//  previousFrame = (
//    ^ previousFrame
//  )
//
//  clearPreviousFrame = (
//    previousFrame := nil
//  )
    func clearPreviousFrame() {
        self.previousFrame = nil
    }
//
//  hasPreviousFrame = (
//    ^ previousFrame ~= nil
//  )
    func hasPreviousFrame() -> Bool {
        return self.previousFrame != nil
    }
//
//  isBootstrapFrame = (
//    ^ self hasPreviousFrame not
//  )
    func isBootstrapFrame() -> Bool {
        return !self.hasPreviousFrame()
    }
//
//  context = (
//    ^ context
//  )
//
//  hasContext = (
//    ^ context ~= nil
//  )
    func hasContext() -> Bool {
        return self.contextFrame != nil
    }

//
//  context: level = (
//    | frame |
//    "Get the context frame at the given level"
//    frame := self.
//
//    "Iterate through the context chain until the given level is reached"
//    [level > 0] whileTrue: [
//      "Get the context of the current frame"
//      frame := frame context.
//
//      "Go to the next level"
//      level := level - 1 ].
//
//    ^ frame
//  )
    func context(level: Int) -> Frame {
        var f = self
        var l = level
        while level > 0 {
            f = f.contextFrame!
            l -= 1
        }
        return f
    }
//
//  outerContext = (
//    | frame |
//    "Compute the outer context of this frame"
//    frame := self.
//
//    "Iterate through the context chain until null is reached"
//    [frame hasContext] whileTrue: [
//      frame := frame context ].
//
//    ^ frame
//  )
    func outerContext() -> Frame {
        var f = self
        while f.hasContext() {
            f = f.contextFrame!
        }
        return f
    }
//
//  method = (
//    ^ method
//  )
//
//  pop = (
//    | sp |
//    "Pop an object from the expression stack and return it"
//    sp := stackPointer.
//    stackPointer := stackPointer - 1.
//    ^ stack at: sp.
//  )
    func pop() -> SObject {
        let sp = self.stackPointer
        self.stackPointer -= 1
        return self.stack.field(index: sp)
    }
//
//  push: aSAbstractObject = (
//    "Push an object onto the expression stack"
//    | sp |
//    sp := stackPointer + 1.
//    stack at: sp put: aSAbstractObject.
//    stackPointer := sp
//  )
    func push(obj: SObject) {
        let sp = stackPointer + 1
        self.stack.fieldAt(sp, put: obj)
        self.stackPointer = sp
    }
//
//  resetStackPointer = (
//    "arguments are stored in front of local variables"
//    localOffset := method numberOfArguments + 1.
//
//    "Set the stack pointer to its initial value thereby clearing the stack"
//    stackPointer := localOffset + method numberOfLocals - 1
//  )
func resetStackPointer() {
    self.localOffset = self.method.numberOfArguments() + 1
    self.stackPointer = self.localOffset + method.numberOfLocals - 1
}
//
//  bytecodeIndex = (
//    "Get the current bytecode index for this frame"
//    ^ bytecodeIndex
//  )
//    func bytecodeIndex() -> Int {
//        return self.bytecodeIndex
//    }
//
//  bytecodeIndex: value = (
//    "Set the current bytecode index for this frame"
//    bytecodeIndex := value
//  )
    func bytecodeIndex(idx: Int) {
        self.bytecodeIndex = idx
    }
//
//  stackElement: index = (
//    "Get the stack element with the given index
//     (an index of zero yields the top element)"
//    ^ stack at: stackPointer - index
//  )
    func stackElement(idx: Int) -> SObject {
        return self.stack.indexableFields[idx]
    }
//
//  stackElement: index put: value = (
//    "Set the stack element with the given index to the given value
//     (an index of zero yields the top element)"
//    stack at: stackPointer - index put: value
//  )
//
//  local: index = (
//    ^ stack at: localOffset + index - 1
//  )
//
//  local: index put: value = (
//    stack at: localOffset + index - 1 put: value
//  )
//
//  local: index at: contextLevel = (
//    "Get the local with the given index in the given context"
//    ^ (self context: contextLevel) local: index
//  )
//
//  local: index at: contextLevel put: value = (
//    "Set the local with the given index in the given context to the given value"
//    (self context: contextLevel) local: index put: value
//  )
//
//  argument: index = (
//    ^ stack at: index
//  )
//
//  argument: index put: value = (
//    ^ stack at: index put: value
//  )
//
//  argument: index at: contextLevel = (
//    | context |
//    "Get the context"
//    context := self context: contextLevel.
//
//    "Get the argument with the given index"
//    ^ context argument: index
//  )
//
//  argument: index at: contextLevel put: value = (
//    | context |
//    "Get the context"
//    context := self context: contextLevel.
//
//    "Set the argument with the given index to the given value"
//    context argument: index put: value
//  )
//
//  copyArgumentsFrom: frame = (
//    | numArgs |
//    "copy arguments from frame:
//     - arguments are at the top of the stack of frame.
//     - copy them into the argument area of the current frame"
//    numArgs := method numberOfArguments.
//    0 to: numArgs - 1 do: [:i |
//      stack at: i + 1 put: (frame stackElement: numArgs - 1 - i) ]
//  )
//
//  printStackTrace = (
//    | className methodName |
//    "Print a stack trace starting in this frame"
//    self hasPreviousFrame ifTrue: [
//      previousFrame printStackTrace ].
//
//    className := method holder name string.
//    methodName := method signature string.
//    Universe println: className + '>>#' + methodName + ' @bi: ' + bytecodeIndex
//  )
//
//  ----
//
//  new: nilObject previous: prevFrame context: contextFrame method: aSMethod maxStack: stackElements = (
//    ^ self new initialize: nilObject previous: prevFrame context: contextFrame method: aSMethod maxStack: stackElements
//  )
//)
}
