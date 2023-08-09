package interpreter

import "github.com/xt0fer/smog/vm"

type Interpreter struct {
	Frame Frame
}

func (itp *Interpreter) doDup() {
	// Handle the dup bytecode
	itp.getFrame().push(itp.getFrame().getStackElement(0))
}

func (itp *Interpreter) doPushLocal(bytecodeIndex int) {
	// Handle the push local bytecode
	itp.getFrame().push(itp.getFrame().getLocal(itp.getMethod().getBytecode(bytecodeIndex+1),
		itp.getMethod().getBytecode(bytecodeIndex+2)))
}

func (itp *Interpreter) doPushArgument(bytecodeIndex int) {
	// Handle the push argument bytecode
	itp.getFrame().push(itp.getFrame().getArgument(itp.getMethod().getBytecode(bytecodeIndex+1),
		itp.getMethod().getBytecode(bytecodeIndex+2)))
}

func (itp *Interpreter) doPushField(bytecodeIndex int) {
	// Handle the push field bytecode
	fieldName := itp.getMethod().getConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.getSelf().getFieldIndex(fieldName)

	// Push the field with the computed index onto the stack
	itp.getFrame().push(itp.getSelf().getField(fieldIndex))
}

func (itp *Interpreter) doPushBlock(bytecodeIndex int) {
	// Handle the push block bytecode
	blockMethod := itp.getMethod().getConstant(bytecodeIndex)

	// Push a new block with the current getFrame() as context onto the stack
	itp.getFrame().push(Universe().newBlock(blockMethod, itp.getFrame(),
		blockMethod.getNumberOfArguments()))
}

func (itp *Interpreter) doPushConstant(bytecodeIndex int) {
	// Handle the push constant bytecode
	itp.getFrame().push(itp.getMethod().getConstant(bytecodeIndex))
}

func (itp *Interpreter) doPushGlobal(bytecodeIndex int) {
	// Handle the push global bytecode
	globalName := itp.getMethod().getConstant(bytecodeIndex)

	// Get the global from the universe
	global := Universe().getGlobal(globalName)

	if global != nil {
		// Push the global onto the stack
		itp.getFrame().push(global)
	} else {
		// Send 'unknownGlobal:' to self
		arguments := []Object{globalName}
		itp.getSelf().send("unknownGlobal:", arguments)
	}
}

func (itp *Interpreter) doPop() {
	// Handle the pop bytecode
	itp.getFrame().pop()
}

func (itp *Interpreter) doPopLocal(bytecodeIndex int) {
	// Handle the pop local bytecode
	itp.getFrame().setLocal(itp.getMethod().getBytecode(bytecodeIndex+1),
		itp.getMethod().getBytecode(bytecodeIndex+2),
		itp.getFrame().pop())
}

func (itp *Interpreter) doPopArgument(bytecodeIndex int) {
	// Handle the pop argument bytecode
	itp.getFrame().setArgument(itp.getMethod().getBytecode(bytecodeIndex+1),
		itp.getMethod().getBytecode(bytecodeIndex+2),
		itp.getFrame().pop())
}

func (itp *Interpreter) doPopField(bytecodeIndex int) {
	// Handle the pop field bytecode
	fieldName := itp.getMethod().getConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.getSelf().getFieldIndex(fieldName)

	// Set the field with the computed index to the value popped from the stack
	itp.getSelf().setField(fieldIndex, itp.getFrame().pop())
}

func (itp *Interpreter) doSuperSend(bytecodeIndex int) {
	// Handle the super send bytecode
	signature := itp.getMethod().getConstant(bytecodeIndex)

	// Send the message
	// Lookup the invokable with the given signature
	invokable := itp.getMethod().getHolder().getSuperClass().lookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.invoke(itp.getFrame())

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.getNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.getFrame().getStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := Universe().newArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.setIndexableField(i, itp.getFrame().pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []Object{signature, argumentsArray}
		receiver.send("doesNotUnderstand:arguments:", arguments)
	}
}

func (itp *Interpreter) doReturnLocal() {
	// Handle the return local bytecode
	result := itp.getFrame().pop()

	// Pop the top frame and push the result
	itp.popFrameAndPushResult(result)
}

func (itp *Interpreter) doReturnNonLocal() {
	// Handle the return non local bytecode
	result := itp.getFrame().pop()

	// Compute the context for the non-local return
	context := itp.getFrame().getOuterContext()

	// Make sure the block context is still on the stack
	if !context.hasPreviousFrame() {
		// Try to recover by sending 'escapedBlock:' to the sending object
		// this can get a bit nasty when using nested blocks. In this case
		// the "sender" will be the surrounding block and not the object that
		// acutally sent the 'value' message.
		block := itp.getFrame().getArgument(0, 0)
		sender := itp.getFrame().getPreviousFrame().getOuterContext().getArgument(0, 0)
		arguments := []Object{block}

		// pop the frame of the currently executing block...
		itp.popFrame()

		// ... and execute the escapedBlock message instead
		sender.send("escapedBlock:", arguments)

		return
	}

	// Unwind the frames
	for itp.getFrame() != context {
		itp.popFrame()
	}

	// Pop the top frame and push the result
	itp.popFrameAndPushResult(result)
}

func (itp *Interpreter) doSend(bytecodeIndex int) {
	// Handle the send bytecode
	signature := itp.getMethod().getConstant(bytecodeIndex)

	// Get the number of arguments from the signature
	numberOfArguments := signature.getNumberOfSignatureArguments()

	// Get the receiver from the stack
	receiver := itp.getFrame().getStackElement(numberOfArguments - 1)

	// Send the message
	send(signature, receiver.getSOMClass(), bytecodeIndex)
}

func (itp *Interpreter) start() {
	// Iterate through the bytecodes
	for {

		// Get the current bytecode index
		bytecodeIndex := itp.getFrame().getBytecodeIndex()

		// Get the current bytecode
		bytecode := itp.getMethod().getBytecode(bytecodeIndex)

		// Get the length of the current bytecode
		bytecodeLength := Bytecodes.getBytecodeLength(bytecode)

		// Compute the next bytecode index
		nextBytecodeIndex := bytecodeIndex + bytecodeLength

		// Update the bytecode index of the frame
		itp.getFrame().setBytecodeIndex(nextBytecodeIndex)

		// Handle the current bytecode
		switch bytecode {

		case Bytecodes.halt:
			{
				// Handle the halt bytecode
				return
			}

		case Bytecodes.dup:
			{
				itp.doDup()
				break
			}

		case Bytecodes.push_local:
			{
				itp.doPushLocal(bytecodeIndex)
				break
			}

		case Bytecodes.push_argument:
			{
				itp.doPushArgument(bytecodeIndex)
				break
			}

		case Bytecodes.push_field:
			{
				itp.doPushField(bytecodeIndex)
				break
			}

		case Bytecodes.push_block:
			{
				itp.doPushBlock(bytecodeIndex)
				break
			}

		case Bytecodes.push_constant:
			{
				itp.doPushConstant(bytecodeIndex)
				break
			}

		case Bytecodes.push_global:
			{
				itp.doPushGlobal(bytecodeIndex)
				break
			}

		case Bytecodes.pop:
			{
				itp.doPop()
				break
			}

		case Bytecodes.pop_local:
			{
				itp.doPopLocal(bytecodeIndex)
				break
			}

		case Bytecodes.pop_argument:
			{
				itp.doPopArgument(bytecodeIndex)
				break
			}

		case Bytecodes.pop_field:
			{
				itp.doPopField(bytecodeIndex)
				break
			}

		case Bytecodes.send:
			{
				itp.doSend(bytecodeIndex)
				break
			}

		case Bytecodes.super_send:
			{
				itp.doSuperSend(bytecodeIndex)
				break
			}

		case Bytecodes.return_local:
			{
				itp.doReturnLocal()
				break
			}

		case Bytecodes.return_non_local:
			{
				itp.doReturnNonLocal()
				break
			}

		default:
			fmt.println("Nasty bug in interpreter")
			break
		}
	}
}

func (itp *Interpreter) pushNewFrame(method Method) Frame {
	// Allocate a new frame and make it the current one
	Frame := Universe.newFrame(itp.Frame, method)

	// Return the freshly allocated and pushed frame
	return Frame
}

func (itp *Interpreter) getFrame() Frame {
	// Get the frame from the interpreter
	return itp.Frame
}

func (itp *Interpreter) getMethod() Method {
	// Get the method from the interpreter
	return itp.getFrame().getMethod()
}

func (itp *Interpreter) getSelf() vm.Object {
	// Get the self object from the interpreter
	return itp.getFrame().getOuterContext().getArgument(0, 0)
}

func (itp *Interpreter) send(signature vm.Symbol, receiverClass vm.Class, bytecodeIndex int) {
	// Lookup the invokable with the given signature
	invokable := receiverClass.lookupInvokable(signature)

	if invokable != null {
		// Invoke the invokable in the current frame
		invokable.invoke(itp.getFrame())

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.getNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.getFrame().getStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := Universe().newArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.setIndexableField(i, itp.getFrame().pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []vm.Object{signature, argumentsArray}
		receiver.send("doesNotUnderstand:arguments:", arguments)
	}
}

func (itp *Interpreter) popFrame() *vm.Frame {
	// Save a reference to the top frame
	result := itp.Frame

	// Pop the top frame from the frame stack
	itp.Frame = itp.Frame.getPreviousFrame()

	// Destroy the previous pointer on the old top frame
	result.clearPreviousFrame()

	// Return the popped frame
	return result
}

func (itp *Interpreter) popFrameAndPushResult(result Object) {
	// Pop the top frame from the interpreter frame stack and compute the number of arguments
	numberOfArguments := popFrame().getMethod().getNumberOfArguments()

	// Pop the arguments
	for i = 0; i < numberOfArguments; i++ {
		getFrame().pop()
	}

	// Push the result
	getFrame().push(result)
}
