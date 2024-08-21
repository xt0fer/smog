package smog

import "fmt"

type Frame struct {
	Array
	stackPointer        int
	bytecodeIndex       int
	localOffset         int
	PreviousFrameIndex  int
	ContextIndex        int
	MethodIndex         int
	NumberOfFrameFields int
}

//    // Private variables holding the stack pointer and the bytecode index
//    private int stackPointer;
//    private int bytecodeIndex;

//    // the offset at which local variables start
//    private int localOffset;

//    // Static field indices and number of frame fields
//    static final int previousFrameIndex   = 1 + classIndex;
//    static final int contextIndex          = 1 + previousFrameIndex;
//    static final int methodIndex           = 1 + contextIndex;
//    static final int numberOfFrameFields = 1 + methodIndex;

func NewFrame() *Frame {
	f := &Frame{}
	f.PreviousFrameIndex = 0
	f.ContextIndex = 0
	f.MethodIndex = 0
	f.NumberOfFrameFields = 0
	return f
}

// frame is a subtype of array

/**
 * Frame layout:
 *
 * +-----------------+
 * | Arguments       | 0
 * +-----------------+
 * | Local Variables | <-- localOffset
 * +-----------------+
 * | Stack           | <-- stackPointer
 * | ...             |
 * +-----------------+
 */

func (f *Frame) GetPreviousFrame() *Frame {
	return f.GetField(f.PreviousFrameIndex).(*Frame)
}
func (f *Frame) SetPreviousFrame(value *Frame) {
	f.SetField(f.PreviousFrameIndex, value)
}

func (f *Frame) ClearPreviousFrame() {
	f.SetField(f.PreviousFrameIndex, GetUniverse().NilObject)
}

func (f *Frame) HasPreviousFrame() bool {
	return f.GetField(f.PreviousFrameIndex) != GetUniverse().NilObject
}

func (f *Frame) IsBootstrapFrame() bool {
	return !f.HasPreviousFrame()
}

//	   public Frame getContext()
//	   {
//		 // Get the context by reading the field with context index
//		 return (Frame) getField(contextIndex);
//	   }
func (f *Frame) GetContext() *Frame {
	return f.GetField(f.ContextIndex).(*Frame)
}

//	   public void setContext(Frame value)
//	   {
//		 // Set the context by writing to the field with context index
//		 setField(contextIndex, value);
//	   }
func (f *Frame) SetContext(value *Frame) {
	f.SetField(f.ContextIndex, value)
}

//	   public boolean hasContext()
//	   {
//		 return getField(contextIndex) != Universe.nilObject;
//	   }
func (f *Frame) HasContext() bool {
	return f.GetField(f.ContextIndex) != GetUniverse().NilObject
}

//	   public Frame getContext(int level)
//	   {
//		 // Get the context frame at the given level
//		 Frame frame = this;
//		 // Iterate through the context chain until the given level is reached
//		 while (level > 0) {
//		   // Get the context of the current frame
//		   frame = frame.getContext();
//		   // Go to the next level
//		   level = level - 1;
//		 }
//		 // Return the found context
//		 return frame;
//	   }
func (f *Frame) GetContextLevel(level int) *Frame {
	frame := f
	for level > 0 {
		frame = frame.GetContext()
		level = level - 1
	}
	return frame
}

//	   public Frame getOuterContext()
//	   {
//		 // Compute the outer context of this frame
//		 Frame frame = this;
//		 // Iterate through the context chain until null is reached
//		 while (frame.hasContext()) frame = frame.getContext();
//		 // Return the outer context
//		 return frame;
//	   }
func (f *Frame) GetOuterContext() *Frame {
	frame := f
	for frame.HasContext() {
		frame = frame.GetContext()
	}
	return frame
}

//	   public Method getMethod()
//	   {
//		 // Get the method by reading the field with method index
//		 return (Method) getField(methodIndex);
//	   }
func (f *Frame) GetMethod() *Method {
	return f.GetField(f.MethodIndex).(*Method)
}

//	   public void setMethod(Method value)
//	   {
//		 // Set the method by writing to the field with method index
//		 setField(methodIndex, value);
//	   }
func (f *Frame) SetMethod(value *Method) {
	f.SetField(f.MethodIndex, value)
}

//	   public int getDefaultNumberOfFields()
//	   {
//		 // Return the default number of fields in a frame
//		 return numberOfFrameFields;
//	   }
func (f *Frame) GetDefaultNumberOfFields() int {
	return f.NumberOfFrameFields
}

//	   public Object pop()
//	   {
//		 // Pop an object from the expression stack and return it
//		 int stackPointer = getStackPointer();
//		 setStackPointer(stackPointer - 1);
//		 return getIndexableField(stackPointer);
//	   }
func (f *Frame) Pop() interface{} {
	stackPointer := f.GetStackPointer()
	f.SetStackPointer(stackPointer - 1)
	return f.GetIndexableField(stackPointer)
}

//	   public void push(Object value)
//	   {
//		 // Push an object onto the expression stack
//		 int stackPointer = getStackPointer() + 1;
//		 setIndexableField(stackPointer, value);
//		 setStackPointer(stackPointer);
//	   }
func (f *Frame) Push(value interface{}) {
	stackPointer := f.GetStackPointer() + 1
	f.SetIndexableField(stackPointer, value)
	f.SetStackPointer(stackPointer)
}

//	   public int getStackPointer()
//	   {
//		 // Get the current stack pointer for this frame
//		 return stackPointer;
//	   }
func (f *Frame) GetStackPointer() int {
	return f.stackPointer
}

//	   public void setStackPointer(int value)
//	   {
//		 // Set the current stack pointer for this frame
//		 stackPointer = value;
//	   }
func (f *Frame) SetStackPointer(value int) {
	f.stackPointer = value
}

//	   public void resetStackPointer()
//	   {
//		 // arguments are stored in front of local variables
//		 localOffset = getMethod().getNumberOfArguments();
//		 // Set the stack pointer to its initial value thereby clearing the stack
//		 setStackPointer(localOffset + getMethod().getNumberOfLocals() - 1);
//	   }
func (f *Frame) ResetStackPointer() {
	f.localOffset = f.GetMethod().GetNumberOfArguments()
	f.SetStackPointer(f.localOffset + f.GetMethod().GetNumberOfLocals() - 1)
}

//	   public int getBytecodeIndex()
//	   {
//		 // Get the current bytecode index for this frame
//		 return bytecodeIndex;
//	   }
func (f *Frame) GetBytecodeIndex() int {
	return f.bytecodeIndex
}

//	   public void setBytecodeIndex(int value)
//	   {
//		 // Set the current bytecode index for this frame
//		 bytecodeIndex = value;
//	   }
func (f *Frame) SetBytecodeIndex(value int) {
	f.bytecodeIndex = value
}

//	   public Object getStackElement(int index)
//	   {
//		 // Get the stack element with the given index
//		 // (an index of zero yields the top element)
//		 return getIndexableField(getStackPointer() - index);
//	   }
func (f *Frame) GetStackElement(index int) interface{} {
	return f.GetIndexableField(f.GetStackPointer() - index)
}

//	   public void setStackElement(int index, Object value)
//	   {
//		 // Set the stack element with the given index to the given value
//		 // (an index of zero yields the top element)
//		 setIndexableField(getStackPointer() - index, value);
//	   }
func (f *Frame) SetStackElement(index int, value interface{}) {
	f.SetIndexableField(f.GetStackPointer()-index, value)
}

//	   private Object getLocal(int index) {
//		 return getIndexableField(localOffset+index);
//	   }
func (f *Frame) GetLocal(index int) interface{} {
	return f.GetIndexableField(f.localOffset + index)
}

//	   private void setLocal(int index,Object value) {
//		 setIndexableField(localOffset+index, value);
//	   }
func (f *Frame) SetLocal(index int, value interface{}) {
	f.SetIndexableField(f.localOffset+index, value)
}

//	   public Object getLocal(int index, int contextLevel)
//	   {
//		 // Get the local with the given index in the given context
//		 return getContext(contextLevel).getLocal(index);
//	   }
func (f *Frame) GetLocalLevel(index int, contextLevel int) interface{} {
	return f.GetContextLevel(contextLevel).GetLocal(index)
}

//	   public void setLocal(int index, int contextLevel, Object value)
//	   {
//		 // Set the local with the given index in the given context to the given value
//		 getContext(contextLevel).setLocal(index, value);
//	   }
func (f *Frame) SetLocalLevel(index int, contextLevel int, value interface{}) {
	f.GetContextLevel(contextLevel).SetLocal(index, value)
}

//	   public Object getArgument(int index, int contextLevel)
//	   {
//		 // Get the context
//		 Frame context = getContext(contextLevel);
//		 // Get the argument with the given index
//		 return context.getIndexableField(index);
//	   }
func (f *Frame) GetArgument(index byte, contextLevel int) interface{} {
	context := f.GetContextLevel(contextLevel)
	return context.GetIndexableField(int(index))
}

//	   public void setArgument(int index, int contextLevel, Object value)
//	   {
//		 // Get the context
//		 Frame context = getContext(contextLevel);
//		 // Set the argument with the given index to the given value
//		 context.setIndexableField(index, value);
//	   }
func (f *Frame) SetArgument(index int, contextLevel int, value interface{}) {
	context := f.GetContextLevel(contextLevel)
	context.SetIndexableField(index, value)
}

//	   public void copyArgumentsFrom(Frame frame) {
//		 // copy arguments from frame:
//		 // - arguments are at the top of the stack of frame.
//		 // - copy them into the argument area of the current frame
//		 int numArgs = getMethod().getNumberOfArguments();
//		 for (int i=0; i < numArgs; ++i) {
//		   setIndexableField(i, frame.getStackElement(numArgs-1-i));
//		 }
//	   }
func (f *Frame) CopyArgumentsFrom(frame *Frame) {
	numArgs := f.GetMethod().GetNumberOfArguments()
	for i := 0; i < numArgs; i++ {
		f.SetIndexableField(i, frame.GetStackElement(numArgs-1-i))
	}
}

//	   public void printStackTrace()
//	   {
//		 // Print a stack trace starting in this frame
//		 System.out.print(getMethod().getHolder().getName().getString());
//		 System.out.print(getBytecodeIndex() + "@" + getMethod().getSignature().getString());
//		 if (hasPreviousFrame()) getPreviousFrame().printStackTrace();
//	   }
func (f *Frame) PrintStackTrace() {
	fmt.Println(f.GetMethod().GetHolder().GetName().ToString())
	fmt.Println(f.GetBytecodeIndex(), "@", f.GetMethod().GetSignature().ToString())
	if f.HasPreviousFrame() {
		f.GetPreviousFrame().PrintStackTrace()
	}
}
