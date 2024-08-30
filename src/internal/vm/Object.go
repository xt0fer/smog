package vm

// SimpleHello = (
// | name |
// setName: aString (
// name := aString
// )
// printGreeting (
// (’Hello, ’, name) print )
// )

type VMClass struct {
	Name        string    // The name of the class
	ArityField  int       // Number of fields in the class
	ArityMethod int       // Number of methods in the class
	Fields      []int     // Indices to Field entities in the constant pool
	Methods     []*Method // Map of method names to their indices in the constant pool
	SuperClass  *VMClass  // Reference to the superclass (if any)
}

type VMObject struct {
	// Properties of the object
	Fields []int    // Indices to Field entities in the constant pool
	Class  *VMClass // Points to the class of the object
}

type Frame struct {
	Arguments      []VMObject  // Arguments passed to the method
	Locals         []VMObject  // Local variables within the frame
	Stack          []VMObject  // Operand stack for the frame
	InitialAddress CodeAddress // Initial address of the code
	ReturnAddress  CodeAddress // Address to return after executing the frame
}
type ExecutionStack struct {
	Frames []*Frame
}

func (es *ExecutionStack) pushFrame(f *Frame) {
	es.Frames = append(es.Frames, f)
}

func (es *ExecutionStack) popFrame() *Frame {
	if len(es.Frames) == 0 {
		return nil
	}
	f := es.Frames[len(es.Frames)-1]
	es.Frames = es.Frames[:len(es.Frames)-1]
	return f
}

func (es *ExecutionStack) push(obj VMObject) {
	es.Frames[len(es.Frames)-1].Stack = append(es.Frames[len(es.Frames)-1].Stack, obj)
}

func (es *ExecutionStack) pop() VMObject {
	Stack := es.Frames[len(es.Frames)-1].Stack
	obj := Stack[len(Stack)-1]
	es.Frames[len(es.Frames)-1].Stack = Stack[:len(Stack)-1]
	return obj
}

type Heap struct {
	objects []*VMObject
}

func (h *Heap) allocObject() (*VMObject, int) {
	obj := &VMObject{Fields: make([]int, 1)}
	h.objects = append(h.objects, obj)
	return obj, len(h.objects) - 1
}

// func (h *Heap) newClazz(cname string, fields []int, methods []int, superc *VMClass) *VMClass {

// 	// Assuming we are adding a class to the constant pool
// 	vmClass := &VMClass{
// 		Name: "MyClass",
// 		Fields: []int{fieldIndex1, fieldIndex2}, // Indices to Field entities
// 		Methods: map[string]int{
// 			"initialize": methodIndex1,
// 			"doSomething": methodIndex2,
// 		},
// 		SuperClass: superc, // or reference to another VMClass for inheritance
// 	}

// 	// Add the class entity to the constant pool
// 	vm.ConstantPool = append(vm.ConstantPool, Entity{
// 		Type:  ClassEntity,
// 		Class: vmClass,
// 	})
// }

func (h *Heap) NewInstanceOf(class *VMClass) *VMObject {
	obj := &VMObject{
		Fields: make([]int, 1),
		Class:  class,
	}
	h.objects = append(h.objects, obj)
	return obj
}

type GlobalContext struct {
	classes                   map[string]*VMClass // Stores all the classes available in the context
	trueObj, falseObj, nilObj *VMObject           // Predefined true, false, nil objects
}

type CodeAddress struct {
	instructionPointer int
}

type Program struct {
	bytecode []byte // Bytecode representation of the program
}

type Interpreter struct{}

type VM struct {
	ConstantPool []Entity
	ExecStack    ExecutionStack
	Heap         Heap
}

// readability functions for the VM
func (vm *VM) TopFrame() *Frame {
	return vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
}
func (vm *VM) TopFrameStack() *[]VMObject {
	return &vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack
}

// one issue that needs to be resolved, is:
// when do we de-reference the fields of an object to an actual value for computation?
// and when do we keep it as a reference to a field in the constant pool?
// this is important for the VM to know when to de-reference a field and when to keep it as a reference
// for example, in the Lit instruction, we push a constant from the constant pool onto the stack
// we need to know if the constant is a primitive or a field in the constant pool
// if it is a field, we need to push the value of the field onto the stack
// if it is a primitive, we need to push the primitive onto the stack (?) or push a reference to the primitive/field?

func (vm *VM) Lit(i int) {
	entity := vm.ConstantPool[i]
	tstack := vm.TopFrameStack()
	*tstack = append(*tstack,
		VMObject{Fields: []int{entity}})
}
func (vm *VM) GetSlot(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	obj := frame.Stack[len(frame.Stack)-1]
	fieldName := vm.ConstantPool[i].StrVal
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	value := obj.Fields[fieldName]
	frame.Stack = append(frame.Stack, value.(VMObject)) // push the field value
}

func (vm *VM) SetSlot(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	obj := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	fieldName := vm.ConstantPool[i].StrVal
	obj.Fields[fieldName] = value
	frame.Stack = append(frame.Stack, obj)
}

func (vm *VM) Send(i, n int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	methodSelector := vm.ConstantPool[i].StrVal
	args := frame.Stack[len(frame.Stack)-n:]
	frame.Stack = frame.Stack[:len(frame.Stack)-n]

	// Assume `self` is on the Stack before args
	self := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]

	// Find the method to invoke
	class := self.Class
	method := vm.findMethod(class, methodSelector)

	// Create a new frame for method invocation
	newFrame := &Frame{
		Arguments:      append([]VMObject{self}, args...),
		Locals:         make([]VMObject, method.NumLocals),
		Stack:          []VMObject{},
		InitialAddress: CodeAddress{instructionPointer: vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].ReturnAddress.instructionPointer},
		ReturnAddress:  CodeAddress{instructionPointer: 0},
	}
	vm.ExecStack.pushFrame(newFrame)
	vm.runMethod(method)
}

func (vm *VM) GetLocal(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	local := frame.Locals[i]
	frame.Stack = append(frame.Stack, local)
}

func (vm *VM) SetLocal(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	frame.Locals[i] = value
}

func (vm *VM) GetSelf() {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	self := frame.arguments[0] // 'self' is always the first argument
	frame.Stack = append(frame.Stack, self)
}

func (vm *VM) GetArg(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	arg := frame.arguments[i]
	frame.Stack = append(frame.Stack, arg)
}

func (vm *VM) Block(i int) {
	block := vm.ConstantPool[i].Block
	blockObj, _ := vm.Heap.allocObject()
	blockObj.Fields["block"] = block
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, *blockObj)
}

func (vm *VM) Ret() {
	frame := vm.ExecStack.popFrame()
	value := frame.Stack[len(frame.Stack)-1]
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, value)
}

func (vm *VM) Retnl(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	// Logic to handle non-local return based on the argument i
	vm.ExecStack.Frames = vm.ExecStack.Frames[:len(vm.ExecStack.Frames)-i] // Adjust the Stack to the appropriate frame
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, value)
}

func (vm *VM) runMethod(m *Method) {
	ip := 0
	for ip < len(m.Instructions) {
		instruction := m.Instructions[ip]
		switch instruction {
		case 0x00: // LIT
			vm.Lit(int(m.Instructions[ip+1]))
			ip += 2
		case 0x01: // GET SLOT
			vm.GetSlot(int(m.Instructions[ip+1]))
			ip += 2
		case 0x02: // SET SLOT
			vm.SetSlot(int(m.Instructions[ip+1]))
			ip += 2
		case 0x03: // SEND
			vm.Send(int(m.Instructions[ip+1]), int(m.Instructions[ip+2]))
			ip += 3
		case 0x04: // GET LOCAL
			vm.GetLocal(int(m.Instructions[ip+1]))
			ip += 2
		case 0x05: // SET LOCAL
			vm.SetLocal(int(m.Instructions[ip+1]))
			ip += 2
		case 0x06: // GET SELF
			vm.GetSelf()
			ip++
		case 0x07: // GET ARG
			vm.GetArg(int(m.Instructions[ip+1]))
			ip += 2
		case 0x08: // BLOCK
			vm.Block(int(m.Instructions[ip+1]))
			ip += 2
		case 0x09: // RET
			vm.Ret()
			ip++
			return
		case 0x0A: // RETNL
			vm.Retnl(int(m.Instructions[ip+1]))
			ip += 2
			return
		}
	}
}

func (vm *VM) findMethod(class *Class, selector string) *Method {
	for _, index := range class.Fields {
		if vm.ConstantPool[index].Type == MethodEntity && vm.ConstantPool[index].Method.Selector == selector {
			return vm.ConstantPool[index].Method
		}
	}
	return nil
}

func (i *Interpreter) Interpret(p *Program, gc *GlobalContext) {
	// Execution loop to interpret bytecode
	// Fetch, decode, and execute instructions
	// Manage the stack, heap, and Frames
}

func main() {
	// Initialize global context with basic objects and classes
	globalContext := &GlobalContext{
		classes:  make(map[string]*VMClass),
		trueObj:  &VMObject{fields: make(map[string]interface{})},
		falseObj: &VMObject{fields: make(map[string]interface{})},
		nilObj:   &VMObject{fields: make(map[string]interface{})},
	}

	// Load program bytecode (this could be loaded from a file, for example)
	program := &Program{
		bytecode: []byte{ /* Bytecode instructions */ },
	}

	// Create an interpreter instance
	interpreter := &Interpreter{}

	// Execute the program
	interpreter.Interpret(program, globalContext)
}

type EntityType int

const (
	NilEntity EntityType = iota
	IntEntity
	DoubleEntity
	StringEntity
	FieldEntity
	MethodEntity
	PrimitiveEntity
	BlockEntity
	ClassEntity
)

type Entity struct {
	Type   EntityType
	IntVal int32
	DblVal float64
	StrVal string
	Field  *Field
	Method *Method
	Block  *Block
	Class  *Class
}

type Field struct {
	Name string
}

type Method struct {
	Selector     string
	Arity        int
	NumLocals    int
	Instructions []byte
}

type Primitive struct {
	Name string
	Impl func() // Implementation in the VM
}

type Block struct {
	Arity        int
	Instructions []byte
}

// LIT i: Push a constant from the constant pool onto the stack
func (vm *VM) Lit(i int) {
	entity := vm.ConstantPool[i]
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, VMObject{Fields: map[string]interface{}{"value": entity}})
}

// GET SLOT i: Retrieve a field from an object and push it onto the stack
func (vm *VM) GetSlot(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	obj := frame.Stack[len(frame.Stack)-1]
	fieldName := vm.ConstantPool[i].StrVal
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	value := obj.Fields[fieldName]
	frame.Stack = append(frame.Stack, value.(VMObject)) // push the field value
}

// SET SLOT i: Set a field in an object with the value on the top of the stack
func (vm *VM) SetSlot(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	obj := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	fieldName := vm.ConstantPool[i].StrVal
	obj.Fields[fieldName] = value
	frame.Stack = append(frame.Stack, obj)
}

// SEND i n: Send a message to an object, creating a new frame
func (vm *VM) Send(i, n int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	methodSelector := vm.ConstantPool[i].StrVal
	args := frame.Stack[len(frame.Stack)-n:]
	frame.Stack = frame.Stack[:len(frame.Stack)-n]

	// Assume `self` is on the stack before args
	self := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]

	// Find the method to invoke
	class := self.Class
	method := vm.findMethod(class, methodSelector)

	// Create a new frame for method invocation
	newFrame := &Frame{
		Args:       append([]VMObject{self}, args...),
		Locals:     make([]VMObject, method.NumLocals),
		Stack:      []VMObject{},
		ReturnAddr: CodeAddress{InstructionPointer: vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].ReturnAddr.InstructionPointer},
	}
	vm.ExecStack.PushFrame(newFrame)
	vm.runMethod(method)
}

// GET LOCAL i: Retrieve a local variable and push it onto the stack
func (vm *VM) GetLocal(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	local := frame.Locals[i]
	frame.Stack = append(frame.Stack, local)
}

// SET LOCAL i: Pop a value from the stack and store it as a local variable
func (vm *VM) SetLocal(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	frame.Locals[i] = value
}

// GET SELF: Push the current object (self) onto the stack
func (vm *VM) GetSelf() {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	self := frame.Args[0] // 'self' is always the first argument
	frame.Stack = append(frame.Stack, self)
}

// GET ARG i: Retrieve the i-th argument of the current message and push it on top of the stack
func (vm *VM) GetArg(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	arg := frame.Args[i]
	frame.Stack = append(frame.Stack, arg)
}

// BLOCK i: Create a code block object and push it onto the stack
func (vm *VM) Block(i int) {
	block := vm.ConstantPool[i].Block
	blockObj := vm.Heap.NewObject()
	blockObj.Fields["block"] = block
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, *blockObj)
}

// RET: Return from a method call
func (vm *VM) Ret() {
	frame := vm.ExecStack.PopFrame()
	value := frame.Stack[len(frame.Stack)-1]
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, value)
}

// RETNL i: Non-local return
func (vm *VM) Retnl(i int) {
	frame := vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1]
	value := frame.Stack[len(frame.Stack)-1]
	// Logic to handle non-local return based on the argument i
	vm.ExecStack.Frames = vm.ExecStack.Frames[:len(vm.ExecStack.Frames)-i] // Adjust the stack to the appropriate frame
	vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack = append(vm.ExecStack.Frames[len(vm.ExecStack.Frames)-1].Stack, value)
}

// Helper methods

// runMethod executes the instructions of a method
func (vm *VM) runMethod(m *Method) {
	ip := 0
	for ip < len(m.Instructions) {
		instruction := m.Instructions[ip]
		switch instruction {
		case 0x00: // LIT
			vm.Lit(int(m.Instructions[ip+1]))
			ip += 2
		case 0x01: // GET SLOT
			vm.GetSlot(int(m.Instructions[ip+1]))
			ip += 2
		case 0x02: // SET SLOT
			vm.SetSlot(int(m.Instructions[ip+1]))
			ip += 2
		case 0x03: // SEND
			vm.Send(int(m.Instructions[ip+1]), int(m.Instructions[ip+2]))
			ip += 3
		case 0x04: // GET LOCAL
			vm.GetLocal(int(m.Instructions[ip+1]))
			ip += 2
		case 0x05: // SET LOCAL
			vm.SetLocal(int(m.Instructions[ip+1]))
			ip += 2
		case 0x06: // GET SELF
			vm.GetSelf()
			ip++
		case 0x07: // GET ARG
			vm.GetArg(int(m.Instructions[ip+1]))
			ip += 2
		case 0x08: // BLOCK
			vm.Block(int(m.Instructions[ip+1]))
			ip += 2
		case 0x09: // RET
			vm.Ret()
			ip++
			return
		case 0x0A: // RETNL
			vm.Retnl(int(m.Instructions[ip+1]))
			ip += 2
			return
		}
	}
}

// findMethod is a helper function to locate a method in a class
func (vm *VM) findMethod(class *Class, selector string) *Method {
	for _, index := range class.Fields {
		if vm.ConstantPool[index].Type == MethodEntity && vm.ConstantPool[index].Method.Selector == selector {
			return vm.ConstantPool[index].Method
		}
	}
	return nil
}
