package smog

import "log"

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
	fieldName := itp.GetMethod().signature //.GetConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.GetSelf().GetFieldIndex(fieldName)

	// Push the field with the computed index onto the stack
	itp.Frame.Push(itp.GetSelf().GetField(fieldIndex))
}

func (itp *Interpreter) DoPushBlock(bytecodeIndex int) {
	// Handle the push block bytecode
	blockMethod := itp.GetMethod() //.signature //.GetConstant(bytecodeIndex)

	// Push a new block with the current Frame as context onto the stack
	itp.Frame.Push(GetUniverse().NewBlock(blockMethod, itp.Frame,
		blockMethod.GetNumberOfArguments()))
}

func (itp *Interpreter) DoPushConstant(bytecodeIndex int) {
	// Handle the push constant bytecode
	itp.Frame.Push(itp.GetMethod().GetConstant(bytecodeIndex))
}

func (itp *Interpreter) DoPushGlobal(bytecodeIndex int) {
	// Handle the push global bytecode
	globalName := itp.GetMethod().signature
	//.GetConstant(bytecodeIndex)

	// Get the global from the universe
	global := GetUniverse().GetGlobal(globalName).(*Object)

	if global != nil {
		// Push the global onto the stack
		itp.Frame.Push(global)
	} else {
		// Send 'unknownGlobal:' to self
		//arguments := []*Symbol{globalName}

		itp.GetSelf().Send("unknownGlobal:", nil) //arguments)
	}
}

func (itp *Interpreter) DoPop() {
	// Handle the pop bytecode
	itp.Frame.Pop()
}

func (itp *Interpreter) DoPopLocal(bytecodeIndex int) {
	// Handle the pop local bytecode
	itp.Frame.SetLocalLevel(int(itp.GetMethod().GetBytecode(bytecodeIndex+1)),
		int(itp.GetMethod().GetBytecode(bytecodeIndex+2)),
		itp.Frame.Pop())
}

func (itp *Interpreter) DoPopArgument(bytecodeIndex int) {
	// Handle the pop argument bytecode
	itp.Frame.SetArgument(int(itp.GetMethod().GetBytecode(bytecodeIndex+1)),
		int(itp.GetMethod().GetBytecode(bytecodeIndex+2)),
		itp.Frame.Pop())
}

func (itp *Interpreter) DoPopField(bytecodeIndex int) {
	// Handle the pop field bytecode
	fieldName := itp.GetMethod().signature
	//.GetConstant(bytecodeIndex)

	// Get the field index from the field name
	fieldIndex := itp.GetSelf().GetFieldIndex(fieldName)

	// Set the field with the computed index to the value popped from the stack
	itp.GetSelf().SetField(fieldIndex, itp.Frame.Pop())
}

func (itp *Interpreter) DoSuperSend(bytecodeIndex int) {
	// Handle the super send bytecode
	signature := itp.GetMethod().signature
	//.GetConstant(bytecodeIndex)

	// Send the message
	// Lookup the invokable with the given signature
	invokable := itp.GetMethod().GetHolder().GetSuperClass().LookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.Invoke(itp.Frame)

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.GetNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.Frame.GetStackElement(numberOfArguments - 1).(*Object)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := GetUniverse().NewArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.SetIndexableField(i, itp.Frame.Pop())
		}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		//arguments := []Object{signature, argumentsArray}
		receiver.Send("doesNotUnderstand:arguments:", nil) // arguments)
	}
}

func (itp *Interpreter) DoReturnLocal() {
	// Handle the return local bytecode
	result := itp.Frame.Pop().(*Object)

	// Pop the top frame and push the result
	itp.PopFrameAndPushResult(result)
}

func (itp *Interpreter) DoReturnNonLocal() {
	// Handle the return non local bytecode
	result := itp.Frame.Pop().(*Object)

	// Compute the context for the non-local return
	context := itp.Frame.GetOuterContext()

	// Make sure the block context is still on the stack
	if !context.HasPreviousFrame() {
		// Try to recover by sending 'escapedBlock:' to the sending object
		// this can get a bit nasty when using nested blocks. In this case
		// the "sender" will be the surrounding block and not the object that
		// actually sent the 'value' message.

		//block := itp.Frame.GetArgument(0, 0).(*Block)                                        // needs to be a Block

		sender := itp.Frame.GetPreviousFrame().GetOuterContext().GetArgument(0, 0).(*Object) // Object
		//arguments := []Object{block.(*Block)}

		// pop the frame of the currently executing block...
		itp.PopFrame()

		// ... and execute the escapedBlock message instead
		sender.Send("escapedBlock:", nil) // arguments)

		return
	}

	// Unwind the frames
	for itp.Frame != context {
		itp.PopFrame()
	}

	// Pop the top frame and push the result
	itp.PopFrameAndPushResult(result)
}

func (itp *Interpreter) DoSend(bytecodeIndex int) {
	// Handle the send bytecode
	signature := itp.GetMethod().signature
	//GetConstant(bytecodeIndex)

	// Get the number of arguments from the signature
	numberOfArguments := signature.GetNumberOfSignatureArguments()

	// Get the receiver from the stack
	receiver := itp.Frame.GetStackElement(numberOfArguments - 1).(*Object)

	// Send the message
	itp.Send(signature, receiver.GetSOMClass(), bytecodeIndex)
}

func (itp *Interpreter) Start() {
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
				itp.DoPushArgument(bytecodeIndex)
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
			log.Println("Nasty bug in interpreter")
			break
		}
	}
}

func (itp *Interpreter) PushNewFrame(method *Method) *Frame {
	// Allocate a new frame and make it the current one
	Frame := GetUniverse().NewFrame(itp.Frame, method)

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
	return itp.Frame.GetOuterContext().GetArgument(0, 0).(*Object)
}

func (itp *Interpreter) Send(signature *Symbol, receiverClass *Class, bytecodeIndex int) {
	// Lookup the invokable with the given signature
	invokable := receiverClass.LookupInvokable(signature)

	if invokable != nil {
		// Invoke the invokable in the current frame
		invokable.Invoke(itp.Frame)

	} else {
		// Compute the number of arguments
		numberOfArguments := signature.GetNumberOfSignatureArguments()

		// Compute the receiver
		receiver := itp.Frame.GetStackElement(numberOfArguments - 1).(*Object)

		// Allocate an array with enough room to hold all arguments
		argumentsArray := GetUniverse().NewArray(numberOfArguments)

		// Remove all arguments and put them in the freshly allocated array
		for i := numberOfArguments - 1; i >= 0; i-- {
			argumentsArray.SetIndexableField(i, itp.Frame.Pop())
		}

		log.Println("Send 'doesNotUnderstand:arguments:' to the receiver object, NEED TO ADD argumentsArray")
		//arguments := []string{signature.String()}

		// Send 'doesNotUnderstand:arguments:' to the receiver object
		receiver.Send("doesNotUnderstand:arguments:", nil) //arguments)
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

func (itp *Interpreter) PopFrameAndPushResult(result *Object) {
	// Pop the top frame from the interpreter frame stack and compute the number of arguments
	numberOfArguments := itp.PopFrame().GetMethod().GetNumberOfArguments()

	// Pop the arguments
	for i := 0; i < numberOfArguments; i++ {
		itp.Frame.Pop()
	}

	// Push the result
	itp.Frame.Push(result)
}
