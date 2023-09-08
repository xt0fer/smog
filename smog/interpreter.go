package smog

// Bytecodes

const (
	numBytecodes = 16
)

const (
	halt byte = iota //0
	dup
	push_local
	push_argument
	push_field
	push_block //5
	push_constant
	push_global
	pop
	pop_local
	pop_argument //10
	pop_field
	send
	super_send
	return_local
	return_non_local //15
)

var bytecodeNames [numBytecodes]string = [numBytecodes]string{
	"HALT            ",
	"DUP             ",
	"PUSH_LOCAL      ",
	"PUSH_ARGUMENT   ",
	"PUSH_FIELD      ",
	"PUSH_BLOCK      ",
	"PUSH_CONSTANT   ",
	"PUSH_GLOBAL     ",
	"POP             ",
	"POP_LOCAL       ",
	"POP_ARGUMENT    ",
	"POP_FIELD       ",
	"SEND            ",
	"SUPER_SEND      ",
	"RETURN_LOCAL    ",
	"RETURN_NON_LOCAL",
}

var bytecodeLength [numBytecodes]int = [numBytecodes]int{1, 1, 3, 3, 2, 2, 2, 2, 1, 3, 3, 2, 2, 2, 1, 1}

// interpreter

type Interpreter struct {
	Frame *Frame
}

func (itp *Interpreter) DoDup() {
	// Handle the dup bytecode
	itp.Frame.Push(itp.Frame.GetStackElement(0))
}

func (itp *Interpreter) DoPushLocal(bytecodeIndex int) {
	// Handle the push local bytecode
	itp.Frame.Push(
		itp.Frame.GetLocalLevel(
			int(itp.GetMethod().GetBytecode(bytecodeIndex+1)),
			int(itp.GetMethod().GetBytecode(bytecodeIndex+2))))
}

func (itp *Interpreter) DoPushArgument(bytecodeIndex int) {
	// Handle the push argument bytecode
	itp.Frame.Push(itp.Frame.GetArgument(itp.GetMethod().GetBytecode(bytecodeIndex+1),
		int(itp.GetMethod().GetBytecode(bytecodeIndex+2))))
}

func (itp *Interpreter) DoPushField(bytecodeIndex int) {
	// Handle the push field bytecode
	fieldName := itp.GetMethod().GetConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.GetSelf().GetFieldIndex(fieldName)

	// Push the field with the computed index onto the stack
	itp.Frame.push(itp.GetSelf().GetField(fieldIndex))
}

func (itp *Interpreter) DoPushBlock(bytecodeIndex int) {
	// Handle the push block bytecode
	blockMethod := itp.GetMethod().GetConstant(bytecodeIndex)

	// Push a new block with the current Frame as context onto the stack
	itp.Frame.push(Universe().newBlock(blockMethod, itp.Frame,
		blockMethod.GetNumberOfArguments()))
}

func (itp *Interpreter) DoPushConstant(bytecodeIndex int) {
	// Handle the push constant bytecode
	itp.Frame.push(itp.GetMethod().GetConstant(bytecodeIndex))
}

func (itp *Interpreter) DoPushGlobal(bytecodeIndex int) {
	// Handle the push global bytecode
	globalName := itp.GetMethod().GetConstant(bytecodeIndex)

	// Get the global from the universe
	global := Universe().GetGlobal(globalName)

	if global != nil {
		// Push the global onto the stack
		itp.Frame.Push(global)
	} else {
		// Send 'unknownGlobal:' to self
		arguments := []Object{globalName}
		itp.GetSelf().send("unknownGlobal:", arguments)
	}
}

func (itp *Interpreter) DoPop() {
	// Handle the pop bytecode
	itp.Frame.Pop()
}

func (itp *Interpreter) DoPopLocal(bytecodeIndex int) {
	// Handle the pop local bytecode
	itp.Frame.SetLocal(itp.GetMethod().GetBytecode(bytecodeIndex+1),
		itp.GetMethod().GetBytecode(bytecodeIndex+2),
		itp.Frame.Pop())
}

func (itp *Interpreter) DoPopArgument(bytecodeIndex int) {
	// Handle the pop argument bytecode
	itp.Frame.SetArgument(itp.GetMethod().GetBytecode(bytecodeIndex+1),
		itp.GetMethod().GetBytecode(bytecodeIndex+2),
		itp.Frame.Pop())
}

func (itp *Interpreter) DoPopField(bytecodeIndex int) {
	// Handle the pop field bytecode
	fieldName := itp.GetMethod().GetConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.GetSelf().GetFieldIndex(fieldName)

	// Set the field with the computed index to the value popped from the stack
	itp.GetSelf().setField(fieldIndex, itp.Frame.Pop())
}

func (itp *Interpreter) DoSuperSend(bytecodeIndex int) {
	// Handle the super send bytecode
	signature := itp.GetMethod().GetConstant(bytecodeIndex)

	// Send the message
	// Lookup the invokable with the given signature
	invokable := itp.GetMethod().GetHolder().GetSuperClass().lookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.Invoke(itp.Frame)

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.GetNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.Frame.GetStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := GetUniverse().NewArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.SetIndexableField(i, itp.Frame.Pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []Object{signature, argumentsArray}
		receiver.Send("doesNotUnderstand:arguments:", arguments)
	}
}

func (itp *Interpreter) DoReturnLocal() {
	// Handle the return local bytecode
	result := itp.Frame.Pop()

	// Pop the top frame and push the result
	itp.PopFrameAndPushResult(result)
}

func (itp *Interpreter) DoReturnNonLocal() {
	// Handle the return non local bytecode
	result := itp.Frame.Pop()

	// Compute the context for the non-local return
	context := itp.Frame.GetOuterContext()

	// Make sure the block context is still on the stack
	if !context.HasPreviousFrame() {
		// Try to recover by sending 'escapedBlock:' to the sending object
		// this can get a bit nasty when using nested blocks. In this case
		// the "sender" will be the surrounding block and not the object that
		// acutally sent the 'value' message.
		block := itp.Frame.GetArgument(0, 0)
		sender := itp.Frame.GetPreviousFrame().GetOuterContext().GetArgument(0, 0)
		arguments := []Object{block}

		// pop the frame of the currently executing block...
		itp.PopFrame()

		// ... and execute the escapedBlock message instead
		sender.Send("escapedBlock:", arguments)

		return
	}

	// Unwind the frames
	for itp.Frame != context {
		itp.PopFrame()
	}

	// Pop the top frame and push the result
	itp.PopFrameAndPushResult(result)
}

func (itp *Interpreter) doSend(bytecodeIndex int) {
	// Handle the send bytecode
	signature := itp.GetMethod().GetConstant(bytecodeIndex)

	// Get the number of arguments from the signature
	numberOfArguments := signature.GetNumberOfSignatureArguments()

	// Get the receiver from the stack
	receiver := itp.Frame.GetStackElement(numberOfArguments - 1)

	// Send the message
	itp.Send(signature, receiver.GetSOMClass(), bytecodeIndex)
}

func (itp *Interpreter) start() {
	// Iterate through the bytecodes
	for {

		// Get the current bytecode index
		bytecodeIndex := itp.Frame.GetBytecodeIndex()

		// Get the current bytecode
		bytecode := itp.GetMethod().GetBytecode(bytecodeIndex)

		// Get the length of the current bytecode
		bytecodeLength := bytecodeLength[bytecode]

		// Compute the next bytecode index
		nextBytecodeIndex := bytecodeIndex + bytecodeLength

		// Update the bytecode index of the frame
		itp.Frame.SetBytecodeIndex(nextBytecodeIndex)

		// Handle the current bytecode
		switch bytecode {

		case halt:
			{
				// Handle the halt bytecode
				return
			}

		case dup:
			{
				itp.DoDup()
				break
			}

		case push_local:
			{
				itp.DoPushLocal(bytecodeIndex)
				break
			}

		case push_argument:
			{
				itp.SoPushArgument(bytecodeIndex)
				break
			}

		case push_field:
			{
				itp.DoPushField(bytecodeIndex)
				break
			}

		case push_block:
			{
				itp.DoPushBlock(bytecodeIndex)
				break
			}

		case push_constant:
			{
				itp.DoPushConstant(bytecodeIndex)
				break
			}

		case push_global:
			{
				itp.DoPushGlobal(bytecodeIndex)
				break
			}

		case pop:
			{
				itp.DoPop()
				break
			}

		case pop_local:
			{
				itp.DoPopLocal(bytecodeIndex)
				break
			}

		case pop_argument:
			{
				itp.DoPopArgument(bytecodeIndex)
				break
			}

		case pop_field:
			{
				itp.DoPopField(bytecodeIndex)
				break
			}

		case send:
			{
				itp.DoSend(bytecodeIndex)
				break
			}

		case super_send:
			{
				itp.DoSuperSend(bytecodeIndex)
				break
			}

		case return_local:
			{
				itp.DoReturnLocal()
				break
			}

		case return_non_local:
			{
				itp.DoReturnNonLocal()
				break
			}

		default:
			fmt.println("Nasty bug in interpreter")
			break
		}
	}
}

func (itp *Interpreter) pushNewFrame(method Method) *Frame {
	// Allocate a new frame and make it the current one
	Frame := Universe.newFrame(itp.Frame, method)

	// Return the freshly allocated and pushed frame
	return Frame
}

func (itp *Interpreter) GetFrame() *Frame {
	// Get the frame from the interpreter
	return itp.Frame
}

func (itp *Interpreter) GetMethod() *Method {
	// Get the method from the interpreter
	return itp.Frame.GetMethod()
}

func (itp *Interpreter) GetSelf() *Object {
	// Get the self object from the interpreter
	return itp.Frame.GetOuterContext().GetArgument(0, 0)
}

func (itp *Interpreter) send(signature Symbol, receiverClass Class, bytecodeIndex int) {
	// Lookup the invokable with the given signature
	invokable := receiverClass.lookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.invoke(itp.Frame)

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.GetNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.Frame.GetStackElement(numberOfArguments - 1)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := GetUniverse().NewArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.SetIndexableField(i, itp.Frame.Pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		arguments := []Object{signature, argumentsArray}
		receiver.send("doesNotUnderstand:arguments:", arguments)
	}
}

func (itp *Interpreter) PopFrame() *Frame {
	// Save a reference to the top frame
	result := itp.Frame

	// Pop the top frame from the frame stack
	itp.Frame = itp.Frame.GetPreviousFrame()

	// Destroy the previous pointer on the old top frame
	result.ClearPreviousFrame()

	// Return the popped frame
	return result
}

func (itp *Interpreter) PopFrameAndPushResult(result Object) {
	// Pop the top frame from the interpreter frame stack and compute the number of arguments
	numberOfArguments := PopFrame().GetMethod().GetNumberOfArguments()

	// Pop the arguments
	for i := 0; i < numberOfArguments; i++ {
		Frame.Pop()
	}

	// Push the result
	Frame.Push(result)
}

// public static Frame pushNewFrame(Method method)
// {
//   // Allocate a new frame and make it the current one
//   frame = Universe.newFrame(frame, method);

//	  // Return the freshly allocated and pushed frame
//	  return frame;
//	}
func (itp *Interpreter) PushNewFrame(method Method) *Frame {
	// Allocate a new frame and make it the current one
	Frame := Universe.NewFrame(itp.Frame, method)

	// Return the freshly allocated and pushed frame
	return Frame
}
