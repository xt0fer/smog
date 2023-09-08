package smog

type Block struct {
	*Object
	Method              *Method
	Context             *Frame
	numberOfBlockFields int
}

// public Method getMethod()
//
//	{
//	  // Get the method of this block by reading the field with method index
//	  return (Method) getField(methodIndex);
//	}
func (b *Block) GetMethod() *Method {
	return b.Method
}

// public void setMethod(Method value)
//
//	{
//	  // Set the method of this block by writing to the field with method index
//	  setField(methodIndex, value);
//	}
func (b *Block) SetMethod(value *Method) {
	b.Method = value
}

// public Frame getContext()
//
//	{
//	  // Get the context of this block by reading the field with context index
//	  return (Frame) getField(contextIndex);
//	}
func (b *Block) GetContext() *Frame {
	return b.Context
}

// public void setContext(Frame value)
//
//	{
//	  // Set the context of this block by writing to the field with context index
//	  setField(contextIndex, value);
//	}
func (b *Block) SetContext(value *Frame) {
	b.Context = value
}

// public int getDefaultNumberOfFields()
//
//	{
//	  // Return the default number of fields for a block
//	  return numberOfBlockFields;
//	}
func (b *Block) getDefaultNumberOfFields() int {
	return b.numberOfBlockFields
}

// public static Primitive getEvaluationPrimitive(int numberOfArguments)
//
//	{
//	  return new Evaluation(numberOfArguments);
//	}
func GetEvaluationPrimitive(numberOfArguments int) *Primitive {
	return NewPrimitive(computeSignatureString(numberOfArguments))
}

type Evaluation struct {
	*Primitive
	numberOfArguments int
}

// public static class Evaluation extends Primitive
//
//	{
//	  public Evaluation(int numberOfArguments)
//	  {
//		super(computeSignatureString(numberOfArguments));
//		this.numberOfArguments = numberOfArguments;
//	  }
func NewEvaluation(numberOfArguments int) *Evaluation {
	ne := &Evaluation{}
	ne.PrimitiveInit(computeSignatureString(numberOfArguments))
	ne.numberOfArguments = numberOfArguments
	return ne
}

//	  public void invoke(Frame frame)
//	  {
//		// Get the block (the receiver) from the stack
//		Block self = (Block) frame.getStackElement(numberOfArguments - 1);
//		// Get the context of the block...
//		Frame context = self.getContext();
//		// Push a new frame and set its context to be the one specified in the block
//		Frame newFrame = Interpreter.pushNewFrame(self.getMethod());
//		newFrame.copyArgumentsFrom(frame);
//		newFrame.setContext(context);
//	  }
func (e *Evaluation) Invoke(frame *Frame) {
	// Get the block (the receiver) from the stack
	self := frame.GetStackElement(e.numberOfArguments - 1).(*Block)
	// Get the context of the block...
	context := self.GetContext()
	// Push a new frame and set its context to be the one specified in the block
	newFrame := Interpreter.PushNewFrame(self.GetMethod())
	newFrame.copyArgumentsFrom(frame)
	newFrame.SetContext(context)
}

//	  private static java.lang.String computeSignatureString(int numberOfArguments)
//	  {
//		// Compute the signature string
//		java.lang.String signatureString = "value";
//		if (numberOfArguments > 1) signatureString += ":";
//		// Add extra value: selector elements if necessary
//		for (int i = 2; i < numberOfArguments; i++) signatureString += "with:";
//		// Return the signature string
//		return signatureString;
//	  }
func computeSignatureString(numberOfArguments int) string {
	signatureString := "value"
	if numberOfArguments > 1 {
		signatureString += ":"
	}
	for i := 2; i < numberOfArguments; i++ {
		signatureString += "with:"
	}
	return signatureString
}

//   private int numberOfArguments;
// }

// // Static field indices and number of block fields
// static final int methodIndex           = 1 + classIndex;
// static final int contextIndex          = 1 + methodIndex;
// static final int numberOfBlockFields = 1 + contextIndex;
// }
