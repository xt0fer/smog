package vm

import (
	"fmt"
)

type Interpreter struct{}

func (i *Interpreter) doDup() {
	// Handle the dup bytecode
	frame := getFrame()
	frame.push(frame.getStackElement(0))
}

func (i *Interpreter) doPushLocal(bytecodeIndex int) {
	// Handle the push local bytecode
	frame := getFrame()
	frame.push(frame.getLocal(getMethod().getBytecode(bytecodeIndex+1), getMethod().getBytecode(bytecodeIndex+2)))
}

func (i *Interpreter) doPushArgument(bytecodeIndex int) {
	// Handle the push argument bytecode
	frame := getFrame()
	frame.push(frame.getArgument(getMethod().getBytecode(bytecodeIndex+1), getMethod().getBytecode(bytecodeIndex+2)))
}

func (i *Interpreter) doPushField(bytecodeIndex int) {
	// Handle the push field bytecode
	fieldName := getMethod().getConstant(bytecodeIndex).(Symbol)

	// Get the field index from the field name
	fieldIndex := getSelf().getFieldIndex(fieldName)

	// Push the field with the computed index onto the stack
	frame := getFrame()
	frame.push(getSelf().getField(fieldIndex))
}

func (i *Interpreter) doPushBlock(bytecodeIndex int) {
	// Handle the push block bytecode
	blockMethod := getMethod().getConstant(bytecodeIndex).(Method)

	// Push a new block with the current frame as context onto the stack
	frame := getFrame()
	frame.push(Universe.newBlock(blockMethod, frame, blockMethod.getNumberOfArguments()))
}

func (i *Interpreter) doPushConstant(bytecodeIndex int) {
	// Handle the push constant bytecode
	frame := getFrame()
	frame.push(getMethod().getConstant(bytecodeIndex))
}

func (i *Interpreter) doPushGlobal(bytecodeIndex int) {
	// Handle the push global bytecode
	globalName := getMethod().getConstant(bytecodeIndex).(Symbol)

	// Get the global from the universe
	global := Universe.getGlobal(globalName)

	if global != nil {
		// Push the global onto the stack
		frame := getFrame()
		frame.push(global)
	} else {
		// Send 'unknownGlobal:' to self
		arguments := []interface{}{globalName}
		getSelf().send("unknownGlobal:", arguments)
	}
}

func (i *Interpreter) doPop() {
	// Handle the pop bytecode
	frame := getFrame()
	frame.pop()
}

func (i *Interpreter) doPopLocal(bytecodeIndex int) {
	// Handle the pop local bytecode
	frame := getFrame()
	frame.setLocal(getMethod().getBytecode(bytecodeIndex+1), getMethod().getBytecode(bytecodeIndex+2), frame.pop())
}

func (i *Interpreter) doPopArgument(bytecodeIndex int) {
	// Handle the pop argument bytecode
	frame := getFrame()
	frame.setArgument(getMethod().getBytecode(bytecodeIndex+1), getMethod().getBytecode(bytecodeIndex+2), frame.pop())
}

func (i *Interpreter) doPopField(bytecodeIndex int) {
	// Handle the pop field bytecode
	fieldName := getMethod().getConstant(bytecodeIndex).(Symbol)

	// Get the field index from the field name
	fieldIndex := getSelf().getFieldIndex(fieldName)

	// Set the field with the computed index to the value popped from the stack
	getSelf().setField(fieldIndex, getFrame().pop())
}

func (i *Interpreter) doSuperSend(bytecodeIndex int) {
	// Handle the super send bytecode
	signature := getMethod().getConstant(bytecodeIndex).(Symbol)

	// Send the message
	// Lookup the invokable with the given signature
	invokable := getMethod().getHolder().getSuperClass().lookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.invoke(getFrame())

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.getNumberOfSignatureArguments()

		// Compute the receiver
		receiver := getFrame().getStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := Universe.newArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.setIndexableField(i, getFrame().pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []interface{}{signature, argumentsArray}
		receiver.send("doesNotUnderstand:arguments:", arguments)
	}
}

func (i *Interpreter) doReturnLocal() {
	// Handle the return local bytecode
	result := getFrame().pop()

	// Pop the top frame and push the result
	popFrameAndPushResult(result)
}

func (i *Interpreter) doReturnNonLocal() {
	// Handle the return non local bytecode
	result := getFrame().pop()

	// Compute the context for the non-local return
	context := getFrame().getOuterContext()

	// Make sure the block context is still on the stack
	if !context.hasPreviousFrame() {
		// Try to recover by sending 'escapedBlock:' to the sending object
		// this can get a bit nasty when using nested blocks. In this case
		// the "sender" will be the surrounding block and not the object that
		// actually sent the 'value' message.
		block := getFrame().getArgument(0, 0).(Block)
		sender := getFrame().getPreviousFrame().getOuterContext().getArgument(0, 0)
		arguments := []interface{}{block}

		// pop the frame of the currently executing block...
		popFrame()

		// ... and execute the escapedBlock message instead
		sender.send("escapedBlock:", arguments)

		return
	}

	// Unwind the frames
	for getFrame() != context {
		popFrame()
	}

	// Pop the top frame and push the result
	popFrameAndPushResult(result)
}

func (i *Interpreter) doSend(bytecodeIndex int) {
	// Handle the send bytecode
	signature := getMethod().getConstant(bytecodeIndex).(Symbol)

	// Get the number of arguments from the signature
	numberOfArguments := signature.getNumberOfSignatureArguments()

	// Get the receiver from the stack
	receiver := getFrame().getStackElement(numberOfArguments - 1)

	// Send the message
	send(signature, receiver.getSOMClass(), bytecodeIndex)
}

func (i *Interpreter) start() {
	// Iterate through the bytecodes
	for {
		// Get the current bytecode index
		bytecodeIndex := getFrame().getBytecodeIndex()

		// Get the current bytecode
		bytecode := getMethod().getBytecode(bytecodeIndex)

		// Get the length of the current bytecode
		bytecodeLength := getBytecodeLength(bytecode)

		// Compute the next bytecode index
		nextBytecodeIndex := bytecodeIndex + bytecodeLength

		// Update the bytecode index of the frame
		getFrame().setBytecodeIndex(nextBytecodeIndex)

		// Handle the current bytecode
		switch bytecode {

		case HALT:
			{
				// Handle the halt bytecode
				return
			}

		case DUP:
			{
				i.doDup()
				break
			}

		case PUSH_LOCAL:
			{
				i.doPushLocal(bytecodeIndex)
				break
			}

		case PUSH_ARGUMENT:
			{
				i.doPushArgument(bytecodeIndex)
				break
			}

		case PUSH_FIELD:
			{
				i.doPushField(bytecodeIndex)
				break
			}

		case PUSH_BLOCK:
			{
				i.doPushBlock(bytecodeIndex)
				break
			}

		case PUSH_CONSTANT:
			{
				i.doPushConstant(bytecodeIndex)
				break
			}

		case PUSH_GLOBAL:
			{
				i.doPushGlobal(bytecodeIndex)
				break
			}

		case POP:
			{
				i.doPop()
				break
			}

		case POP_LOCAL:
			{
				i.doPopLocal(bytecodeIndex)
				break
			}

		case POP_ARGUMENT:
			{
				i.doPopArgument(bytecodeIndex)
				break
			}

		case POP_FIELD:
			{
				i.doPopField(bytecodeIndex)
				break
			}

		case SEND:
			{
				i.doSend(bytecodeIndex)
				break
			}

		case SUPER_SEND:
			{
				i.doSuperSend(bytecodeIndex)
				break
			}

		case RETURN_LOCAL:
			{
				i.doReturnLocal()
				break
			}

		case RETURN_NON_LOCAL:
			{
				i.doReturnNonLocal()
				break
			}

		default:
			fmt.Println("Nasty bug in interpreter")
			break
		}
	}
}

func (i *Interpreter) pushNewFrame(method Method) Frame {
	// Allocate a new frame and make it the current one
	frame := Universe.newFrame(frame, method)

	// Return the freshly allocated and pushed frame
	return frame
}

func getFrame() Frame {
	// Get the frame from the interpreter
	return frame
}

func getMethod() Method {
	// Get the method from the interpreter
	return getFrame().getMethod()
}

func getSelf() Object {
	// Get the self object from the interpreter
	return getFrame().getOuterContext().getArgument(0, 0)
}

func send(signature Symbol, receiverClass Class, bytecodeIndex int) {
	// Lookup the invokable with the given signature
	invokable := receiverClass.lookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.invoke(getFrame())

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.getNumberOfSignatureArguments()

		// Compute the receiver
		receiver := getFrame().getStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := Universe.newArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.setIndexableField(i, getFrame().pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []interface{}{signature, argumentsArray}
		receiver.send("doesNotUnderstand:arguments:", arguments)
	}
}

func popFrame() Frame {
	// Save a reference to the top frame
	result := frame

	// Pop the top frame from the frame stack
	frame = frame.getPreviousFrame()

	// Destroy the previous pointer on the old top frame
	result.clearPreviousFrame()

	// Return the popped frame
	return result
}

func popFrameAndPushResult(result Object) {
	// Pop the top frame from the interpreter frame stack and compute the number of arguments
	numberOfArguments := popFrame().getMethod().getNumberOfArguments()

	// Pop the arguments
	for i := 0; i < numberOfArguments; i++ {
		getFrame().pop()
	}

	// Push the result
	getFrame().push(result)
}

var frame Frame

func main() {
	interpreter := &Interpreter{}
	interpreter.start()
}
