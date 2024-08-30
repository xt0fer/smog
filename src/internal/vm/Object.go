package vm

// SimpleHello = (
// | name |
// setName: aString (
// name := aString
// )
// printGreeting (
// (’Hello, ’, name) print )
// )

type Class struct {
	Name        string // The name of the class
	ArityField  int    // Number of fields in the class
	ArityMethod int    // Number of methods in the class
	Fields      []int  // Indices to Field entities in the constant pool
	//Methods     []*Method // Map of method names to their indices in the constant pool
	SuperClass *Class // Reference to the superclass (if any)
}

type VMObject struct {
	// Properties of the object
	Fields []int  // Indices to Field entities in the constant pool
	Class  *Class // Points to the class of the object
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

func (h *Heap) NewInstanceOf(class *Class) *VMObject {
	obj := &VMObject{
		Fields: make([]int, 1),
		Class:  class,
	}
	h.objects = append(h.objects, obj)
	return obj
}

type GlobalContext struct {
	classes                   map[string]*Class // Stores all the classes available in the context
	trueObj, falseObj, nilObj *VMObject         // Predefined true, false, nil objects
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

// EntityStack is a generic stack data structure that holds Entities.
type EntityStack struct {
	items []Entity
}

// Push adds an item to the top of the stack.
func (s *EntityStack) Push(item Entity) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// It returns false if the stack is empty.
func (s *EntityStack) Pop() (Entity, bool) {
	if len(s.items) == 0 {
		return Entity{}, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the item from the top of the stack without removing it.
// It returns false if the stack is empty.
func (s *EntityStack) Peek() (Entity, bool) {
	if len(s.items) == 0 {
		return Entity{}, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *EntityStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *EntityStack) Size() int {
	return len(s.items)
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
	frame := vm.TopFrame()
	obj := frame.Stack[len(frame.Stack)-1]
	fieldName := vm.ConstantPool[i].StrVal
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	value := obj.Fields[fieldName]
	frame.Stack = append(frame.Stack, value.(VMObject)) // push the field value
}

func (vm *VM) SetSlot(i int) {
	frame := vm.TopFrame()
	// pop the value from the stack
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	// pop the object from the stack
	obj := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1] // pop the object
	// set the field value name (name of the field) to the value
	fieldName, ok := vm.ConstantPool[i].ToString()
	if !ok {
		panic("SetSlot: expected a string")
	}
	obj.Fields[fieldName] = value
	frame.Stack = append(frame.Stack, obj)
}

func (vm *VM) Send(i, n int) {
	frame := vm.TopFrame()
	methodSelector, ok := vm.ConstantPool[i].ToString()
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
		InitialAddress: CodeAddress{instructionPointer: vm.TopFrame().ReturnAddress.instructionPointer},
		ReturnAddress:  CodeAddress{instructionPointer: 0},
	}
	vm.ExecStack.pushFrame(newFrame)
	vm.runMethod(method)
}

func (vm *VM) GetLocal(i int) {
	frame := vm.TopFrame()
	local := frame.Locals[i]
	frame.Stack = append(frame.Stack, local)
}

func (vm *VM) SetLocal(i int) {
	frame := vm.TopFrame()
	value := frame.Stack[len(frame.Stack)-1]
	frame.Stack = frame.Stack[:len(frame.Stack)-1]
	frame.Locals[i] = value
}

func (vm *VM) GetSelf() {
	frame := vm.TopFrame()
	self := frame.Arguments[0] // 'self' is always the first argument
	frame.Stack = append(frame.Stack, self)
}

func (vm *VM) GetArg(i int) {
	frame := vm.TopFrame()
	arg := frame.Arguments[i]
	frame.Stack = append(frame.Stack, arg)
}

func (vm *VM) Block(i int) {
	block := vm.ConstantPool[i].Block
	blockObj, _ := vm.Heap.allocObject()
	blockObj.Fields["block"] = block
	vm.TopFrame().Stack = append(vm.TopFrame().Stack, *blockObj)
}

func (vm *VM) Ret() {
	frame := vm.ExecStack.popFrame()
	value := frame.Stack[len(frame.Stack)-1]
	vm.TopFrame().Stack = append(vm.TopFrame().Stack, value)
}

func (vm *VM) Retnl(i int) {
	frame := vm.TopFrame()
	value := frame.Stack[len(frame.Stack)-1]
	// Logic to handle non-local return based on the argument i
	vm.ExecStack.Frames = vm.ExecStack.Frames[:len(vm.ExecStack.Frames)-i] // Adjust the Stack to the appropriate frame
	vm.TopFrame().Stack = append(vm.TopFrame().Stack, value)
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
		if vm.ConstantPool[index].Type == MethodEntity {
			m, _ := vm.ConstantPool[index].ToMethod()
			return m
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
	Type  EntityType
	Value interface{}
	// IntVal int32
	// DblVal float64
	// StrVal string
	// Field  *Field
	// Method *Method
	// Block  *Block
	// Class  *Class
}

// NewEntity creates a new Entity with the given type and value.
func NewEntity(t EntityType, v interface{}) *Entity {
	return &Entity{Type: t, Value: v}
}

// ToInt returns the int value if the type is IntEntity, otherwise returns 0.
func (e Entity) ToInt() (int32, bool) {
	if e.Type == IntEntity {
		return e.Value.(int32), true
	}
	return 0, false
}

// ToFloat returns the float64 value if the type is DoubleEntity, otherwise returns 0.0.
func (e Entity) ToFloat() (float64, bool) {
	if e.Type == DoubleEntity {
		return e.Value.(float64), true
	}
	return 0.0, false
}

// ToString returns the string value if the type is StringEntity, otherwise returns "".
func (e Entity) ToString() (string, bool) {
	if e.Type == StringEntity {
		return e.Value.(string), true
	}
	return "", false
}

// ToField returns the Field value if the type is FieldEntity, otherwise returns nil.
func (e Entity) ToField() (*Field, bool) {
	if e.Type == FieldEntity {
		return e.Value.(*Field), true
	}
	return nil, false
}

// ToMethod returns the Method value if the type is MethodEntity, otherwise returns nil.
func (e Entity) ToMethod() (*Method, bool) {
	if e.Type == MethodEntity {
		return e.Value.(*Method), true
	}
	return nil, false
}

// ToBlock returns the Block value if the type is BlockEntity, otherwise returns nil.
func (e Entity) ToBlock() (*Block, bool) {
	if e.Type == BlockEntity {
		return e.Value.(*Block), true
	}
	return nil, false
}

// ToClass returns the Class value if the type is ClassEntity, otherwise returns nil.
func (e Entity) ToClass() (*Class, bool) {
	if e.Type == ClassEntity {
		return e.Value.(*Class), true
	}
	return nil, false
}

// ToNil returns true if the type is NilEntity, otherwise returns false.
func (e Entity) ToNil() bool {
	return e.Type == NilEntity
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

// Interface for Primitive
type PrimitiveInterface interface {
	// Implementations for the primitive
}

type Primitive struct {
	Name string
	Impl func() // Implementation in the VM
}

type Block struct {
	Arity        int
	Instructions []byte
}
