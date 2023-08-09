package vm

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
	f.PreviousFrameIndex = 1 + f.ClassIndex
	f.ContextIndex = 1 + f.PreviousFrameIndex
	f.MethodIndex = 1 + f.ContextIndex
	f.NumberOfFrameFields = 1 + f.MethodIndex
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

func (f *Frame) getPreviousFrame() *Object {
	return f.getField(f.PreviousFrameIndex)
}
func (f *Frame) setPreviousFrame(value *Object) {
	f.setField(f.PreviousFrameIndex, value)
}

func (f *Frame) clearPreviousFrame() {
	f.setField(f.PreviousFrameIndex, GetUniverse().NilObject)
}

func (f *Frame) hasPreviousFrame() bool {
	return f.getField(f.PreviousFrameIndex) != GetUniverse().NilObject
}

func (f *Frame) isBootstrapFrame() bool {
	return !f.hasPreviousFrame()
}

//    public Frame getContext()
//    {
// 	 // Get the context by reading the field with context index
// 	 return (Frame) getField(contextIndex);
//    }

//    public void setContext(Frame value)
//    {
// 	 // Set the context by writing to the field with context index
// 	 setField(contextIndex, value);
//    }

//    public boolean hasContext()
//    {
// 	 return getField(contextIndex) != Universe.nilObject;
//    }

//    public Frame getContext(int level)
//    {
// 	 // Get the context frame at the given level
// 	 Frame frame = this;

// 	 // Iterate through the context chain until the given level is reached
// 	 while (level > 0) {
// 	   // Get the context of the current frame
// 	   frame = frame.getContext();

// 	   // Go to the next level
// 	   level = level - 1;
// 	 }

// 	 // Return the found context
// 	 return frame;
//    }

//    public Frame getOuterContext()
//    {
// 	 // Compute the outer context of this frame
// 	 Frame frame = this;

// 	 // Iterate through the context chain until null is reached
// 	 while (frame.hasContext()) frame = frame.getContext();

// 	 // Return the outer context
// 	 return frame;
//    }

//    public Method getMethod()
//    {
// 	 // Get the method by reading the field with method index
// 	 return (Method) getField(methodIndex);
//    }

//    public void setMethod(Method value)
//    {
// 	 // Set the method by writing to the field with method index
// 	 setField(methodIndex, value);
//    }

//    public int getDefaultNumberOfFields()
//    {
// 	 // Return the default number of fields in a frame
// 	 return numberOfFrameFields;
//    }

//    public Object pop()
//    {
// 	 // Pop an object from the expression stack and return it
// 	 int stackPointer = getStackPointer();
// 	 setStackPointer(stackPointer - 1);
// 	 return getIndexableField(stackPointer);
//    }

//    public void push(Object value)
//    {
// 	 // Push an object onto the expression stack
// 	 int stackPointer = getStackPointer() + 1;
// 	 setIndexableField(stackPointer, value);
// 	 setStackPointer(stackPointer);
//    }

//    public int getStackPointer()
//    {
// 	 // Get the current stack pointer for this frame
// 	 return stackPointer;
//    }

//    public void setStackPointer(int value)
//    {
// 	 // Set the current stack pointer for this frame
// 	 stackPointer = value;
//    }

//    public void resetStackPointer()
//    {
// 	 // arguments are stored in front of local variables
// 	 localOffset = getMethod().getNumberOfArguments();

// 	 // Set the stack pointer to its initial value thereby clearing the stack
// 	 setStackPointer(localOffset + getMethod().getNumberOfLocals() - 1);
//    }

//    public int getBytecodeIndex()
//    {
// 	 // Get the current bytecode index for this frame
// 	 return bytecodeIndex;
//    }

//    public void setBytecodeIndex(int value)
//    {
// 	 // Set the current bytecode index for this frame
// 	 bytecodeIndex = value;
//    }

//    public Object getStackElement(int index)
//    {
// 	 // Get the stack element with the given index
// 	 // (an index of zero yields the top element)
// 	 return getIndexableField(getStackPointer() - index);
//    }

//    public void setStackElement(int index, Object value)
//    {
// 	 // Set the stack element with the given index to the given value
// 	 // (an index of zero yields the top element)
// 	 setIndexableField(getStackPointer() - index, value);
//    }

//    private Object getLocal(int index) {
// 	 return getIndexableField(localOffset+index);
//    }

//    private void setLocal(int index,Object value) {
// 	 setIndexableField(localOffset+index, value);
//    }

//    public Object getLocal(int index, int contextLevel)
//    {
// 	 // Get the local with the given index in the given context
// 	 return getContext(contextLevel).getLocal(index);
//    }

//    public void setLocal(int index, int contextLevel, Object value)
//    {
// 	 // Set the local with the given index in the given context to the given value
// 	 getContext(contextLevel).setLocal(index, value);
//    }

//    public Object getArgument(int index, int contextLevel)
//    {
// 	 // Get the context
// 	 Frame context = getContext(contextLevel);

// 	 // Get the argument with the given index
// 	 return context.getIndexableField(index);
//    }

//    public void setArgument(int index, int contextLevel, Object value)
//    {
// 	 // Get the context
// 	 Frame context = getContext(contextLevel);

// 	 // Set the argument with the given index to the given value
// 	 context.setIndexableField(index, value);
//    }

//    public void copyArgumentsFrom(Frame frame) {
// 	 // copy arguments from frame:
// 	 // - arguments are at the top of the stack of frame.
// 	 // - copy them into the argument area of the current frame
// 	 int numArgs = getMethod().getNumberOfArguments();
// 	 for (int i=0; i < numArgs; ++i) {
// 	   setIndexableField(i, frame.getStackElement(numArgs-1-i));
// 	 }
//    }

//    public void printStackTrace()
//    {
// 	 // Print a stack trace starting in this frame
// 	 System.out.print(getMethod().getHolder().getName().getString());
// 	 System.out.print(getBytecodeIndex() + "@" + getMethod().getSignature().getString());
// 	 if (hasPreviousFrame()) getPreviousFrame().printStackTrace();
//    }
