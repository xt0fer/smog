package smog

import "log"

type IsMethod interface {
}

type Method struct {
	Array
	receiverClassTable                []*Class
	invokedMethods                    []Invokable
	receiverClassIndex                int
	invocationCount                   int64
	bytecode                          []byte
	numberOfLocalsIndex               int //Integer?
	maximumNumberOfStackElementsIndex int //Integer?
	signature                         *Symbol
	numberOfMethodFields              int //Integer?
	holder                            *Class
}

func NewMethod() *Method {
	a := &Method{}
	a.Object.ObjectInit(2)
	a.ArrayInit(0)
	a.MethodInit(0)
	return a
}

func (a *Method) MethodInit(n int) {
	a.IndexableFields = make([]interface{}, n)
	a.bytecode = make([]byte, 0)
}

//   //Private variables for holding the last receiver class and invoked method
//   private java.util.ArrayList<Class> receiverClassTable = new java.util.ArrayList<Class>();
//   private java.util.ArrayList<Invokable> invokedMethods = new java.util.ArrayList<Invokable>();
//   private int receiverClassIndex = 0;

//   // Private variable holding number of invocations and backedges
//   private long invocationCount;

//   // Private variable holding byte array of bytecodes
//   private byte[] bytecodes;

//   // Static field indices and number of method fields
//   static final int numberOfLocalsIndex                 = 1 + classIndex;
//   static final int maximumNumberOfStackElementsIndex = 1 + numberOfLocalsIndex;
//   static final int signatureIndex                        = 1 + maximumNumberOfStackElementsIndex;
//   static final int holderIndex                           = 1 + signatureIndex;
//   static final int numberOfMethodFields                = 1 + holderIndex;
// }

// public class Method extends Array implements Invokable
// {

func (m *Method) isPrimitive() bool {
	return false
}

// public int getNumberOfLocals()
//
//	{
//	  // Get the number of locals (converted to a Java integer)
//	  return ((Integer) getField(numberOfLocalsIndex)).getEmbeddedInteger();
//	}
func (m *Method) GetNumberOfLocals() int {
	return m.numberOfLocalsIndex
}

// public void setNumberOfLocals(int value)
//
//	{
//	  // Set the number of locals
//	  setField(numberOfLocalsIndex, Universe.newInteger(value));
//	}
func (m *Method) SetNumberOfLocals(value int) {
	m.numberOfLocalsIndex = value
}

// public int getMaximumNumberOfStackElements()
//
//	{
//	  // Get the maximum number of stack elements (converted to a Java integer)
//	  return ((Integer) getField(maximumNumberOfStackElementsIndex)).getEmbeddedInteger();
//	}
func (m *Method) getMaximumNumberOfStackElements() int {
	return m.maximumNumberOfStackElementsIndex
}

// public void setMaximumNumberOfStackElements(int value)
//
//	{
//	  // Set the maximum number of stack elements
//	  setField(maximumNumberOfStackElementsIndex, Universe.newInteger(value));
//	}
func (m *Method) SetMaximumNumberOfStackElements(value int) {
	m.maximumNumberOfStackElementsIndex = value
}

// public Symbol getSignature()
//
//	{
//	  // Get the signature of this method by reading the field with signature index
//	  return (Symbol) getField(signatureIndex);
//	}
func (m *Method) GetSignature() *Symbol {
	return m.signature
}

// public void setSignature(Symbol value)
//
//	{
//	  // Set the signature of this method by writing to the field with signature index
//	  setField(signatureIndex, value);
//	}
func (m *Method) SetSignature(value *Symbol) {
	m.signature = value
}

// public Class getHolder()
//
//	{
//	  // Get the holder of this method by reading the field with holder index
//	  return (Class) getField(holderIndex);
//	}
func (m *Method) GetHolder() *Class {
	return m.holder
}

//   public void setHolder(Class value)
//   {
//     // Set the holder of this method by writing to the field with holder index
//     setField(holderIndex, value);

//	  // Make sure all nested invokables have the same holder
//	  for (int i = 0; i < getNumberOfIndexableFields(); i++)
//	    if (getIndexableField(i) instanceof Invokable)
//	      ((Invokable) getIndexableField(i)).setHolder(value);
//	}
func (m *Method) SetHolder(value *Class) {
	m.holder = value
	// Make sure all nested invokables have the same holder
	for i := 0; i < m.GetNumberOfIndexableFields(); i++ {
		if invokable, ok := m.GetIndexableField(i).(Invokable); ok {
			invokable.SetHolder(value)
		}
	}
}

// public Object getConstant(int bytecodeIndex)
//
//	{
//	  // Get the constant associated to a given bytecode index
//	  return getIndexableField(getBytecode(bytecodeIndex + 1));
//	}
func (m *Method) GetConstant(bytecodeIndex int) byte {
	return m.bytecode[bytecodeIndex+1]
}

// public int getNumberOfArguments()
//
//	{
//	  // Get the number of arguments of this method
//	  return getSignature().getNumberOfSignatureArguments();
//	}
func (m *Method) GetNumberOfArguments() int {
	return m.signature.GetNumberOfSignatureArguments()
}

// public int getDefaultNumberOfFields()
//
//	{
//	  // Return the default number of fields in a method
//	  return numberOfMethodFields;
//	}
func (m *Method) getDefaultNumberOfFields() int {
	return m.numberOfMethodFields
}

// public int getNumberOfBytecodes()
//
//	{
//	  // Get the number of bytecodes in this method
//	  return bytecodes.length;
//	}
func (m *Method) GetNumberOfBytecodes() int {
	return len(m.bytecode)
}

// public void setNumberOfBytecodes(int value)
//
//	{
//	  // Set the number of bytecodes in this method
//	  bytecodes = new byte[value];
//	}
func (m *Method) SetNumberOfBytecodes(value int) {
	m.bytecode = make([]byte, value)
}

// public byte getBytecode(int index)
//
//	{
//	  // Get the bytecode at the given index
//	  return bytecodes[index];
//	}
func (m *Method) GetBytecode(index int) byte {
	return m.bytecode[index]
}

// public void setBytecode(int index, byte value)
//
//	{
//	  // Set the bytecode at the given index to the given value
//	  bytecodes[index] = value;
//	}
func (m *Method) SetBytecode(index int, value byte) {
	m.bytecode[index] = value
}

//	public void increaseInvocationCounter() {
//	  invocationCount++;
//	}
func (m *Method) IncreaseInvocationCounter() {
	m.invocationCount++
}

//	public long getInvocationCount() {
//	  return invocationCount;
//	}
func (m *Method) GetInvocationCount() int64 {
	return m.invocationCount
}

// public void invoke(Frame frame)
//
//	{
//	  // Increase the invocation counter
//	  invocationCount++;
//	  // Allocate and push a new frame on the interpreter stack
//	  Frame newFrame = Interpreter.pushNewFrame(this);
//	  newFrame.copyArgumentsFrom(frame);
//	}
func (m *Method) Invoke(frame *Frame) {
	m.invocationCount++
	newFrame := GetInterpreter().PushNewFrame(m)
	newFrame.CopyArgumentsFrom(frame)
}

//   public void replaceBytecodes()
//   {
//     byte newbc[] = new byte[bytecodes.length];
//     int idx = 0;

//     for (int i = 0; i < bytecodes.length; ) {
//       byte bc1 = bytecodes[i];
//       int len1 = Bytecodes.getBytecodeLength(bc1);

//       if (i + len1 >= bytecodes.length) {
//         // we're over target, so just copy bc1
//         for (int j = i; j < i + len1; ++j) {
//           newbc[idx++] = bytecodes[j];
//         }
//         break;
//       }

//       newbc[idx++] = bc1;

//       // copy args to bc1
//       for (int j = i + 1; j < i + len1; ++j) {
//         newbc[idx++] = bytecodes[j];
//       }

//       i += len1; // update i to point on bc2

// }
func (m *Method) ReplaceBytecodes() {
	log.Println("Not sure of ReplaceBytes")
	newbc := make([]byte, len(m.bytecode))
	idx := 0

	for i := 0; i < len(m.bytecode); i++ {
		bc1 := m.bytecode[i]
		byteCodeLen := bytecodeLength[int(m.bytecode[0])]

		if i+byteCodeLen >= byteCodeLen {
			// we're over target, so just copy bc1
			for j := i; j < i+byteCodeLen; j++ {
				newbc[idx] = m.bytecode[j]
				idx++
			}
			break
		}

		newbc[idx] = bc1
		idx++

		// copy args to bc1
		for j := i + 1; j < i+byteCodeLen; j++ {
			newbc[idx] = m.bytecode[j]
			idx++
		}

		i += byteCodeLen // update i to point on bc2
	}
	// we copy the new array because it may be shorter, and we don't
	// want to upset whatever dependence there is on the length
	m.bytecode = make([]byte, idx)
	for i := 0; i < idx; i++ {
		m.bytecode[i] = newbc[i]
	}
}

// public Class getReceiverClass(byte index)
//
//	{
//	  return receiverClassTable.get(index);
//	}
func (m *Method) GetReceiverClass(index byte) *Class {
	return m.receiverClassTable[index]
}

// public Invokable getInvokedMethod(byte index)
//
//	{
//	  //return the last invoked method for a particular send
//	  return invokedMethods.get(index);
//	}
func (m *Method) GetInvokedMethod(index byte) Invokable {
	return m.invokedMethods[index]
}

// public byte addReceiverClassAndMethod(Class recClass, Invokable invokable)
//
//	{
//	  receiverClassTable.add(receiverClassIndex, recClass);
//	  invokedMethods.add(receiverClassIndex, invokable);
//	  receiverClassIndex++;
//	  return (byte) (receiverClassIndex - 1);
//	}
func (m *Method) AddReceiverClassAndMethod(recClass *Class, invokable Invokable) byte {
	m.receiverClassTable[m.receiverClassIndex] = recClass
	m.invokedMethods[m.receiverClassIndex] = invokable
	m.receiverClassIndex++
	return byte(m.receiverClassIndex - 1)
}

//	public boolean isReceiverClassTableFull() {
//	  return receiverClassIndex == 255;
//	}
func (m *Method) IsReceiverClassTableFull() bool {
	return m.receiverClassIndex == 255
}

func (m *Method) IsPrimitive() bool {
	return false
}
