package vm

// SimpleHello = (
// | name |
// setName: aString (
// name := aString
// )
// printGreeting (
// (’Hello, ’, name) print )
// )

type VMObjectType int

const (
	NilVMObject VMObjectType = iota
	IntVMObject
	DoubleVMObject
	StringVMObject
	FieldVMObject
	MethodVMObject
	PrimitiveVMObject
	BlockVMObject
	ObjectVMObject
	ClassVMObject
)

type VMObject struct {
	Type   VMObjectType
	Fields []int  // Indices to Field entities in the heap
	Class  *Class // Points to the class of the object
	Value  interface{}
	// IntVal int32
	// DblVal float64
	// StrVal string
	// Field  *Field
	// Method *Method
	// Block  *Block
	// Class  *Class
}

// NewVMObject creates a new VMObject with the given type and value.
func NewVMObject(t VMObjectType, v interface{}) *VMObject {
	return &VMObject{Type: t, Value: v}
}

// NewIntVMObject creates a new VMObject with the given int value.
func NewIntVMObject(v int32) *VMObject {
	return &VMObject{Type: IntVMObject, Value: v}
}

// NewDoubleVMObject creates a new VMObject with the given float64 value.
func NewDoubleVMObject(v float64) *VMObject {
	return &VMObject{Type: DoubleVMObject, Value: v}
}

// NewStringVMObject creates a new VMObject with the given string value.
func NewStringVMObject(v string) *VMObject {
	return &VMObject{Type: StringVMObject, Value: v}
}

// NewFieldVMObject creates a new VMObject with the given Field value.
func NewFieldVMObject(v *Field) *VMObject {
	return &VMObject{Type: FieldVMObject, Value: v}
}

// NewMethodVMObject creates a new VMObject with the given Method value.
func NewMethodVMObject(v *Method) *VMObject {
	return &VMObject{Type: MethodVMObject, Value: v}
}

// NewPrimitiveVMObject creates a new VMObject with the given Primitive value.
func NewPrimitiveVMObject(v PrimitiveInterface) *VMObject {
	return &VMObject{Type: PrimitiveVMObject, Value: v}
}

// NewBlockVMObject creates a new VMObject with the given Block value.
func NewBlockVMObject(v *Block) *VMObject {
	return &VMObject{Type: BlockVMObject, Value: v}
}

// NewObjectVMObject creates a new VMObject with the given Object value.
func NewObjectVMObject(v *VMObject) *VMObject {
	return &VMObject{Type: ObjectVMObject, Value: v}
}

// NewClassVMObject creates a new VMObject with the given Class value.
func NewClassVMObject(v *Class) *VMObject {
	return &VMObject{Type: ClassVMObject, Value: v}
}

// ToInt returns the int value if the type is IntVMObject, otherwise returns 0.
func (e VMObject) ToInt() (int32, bool) {
	if e.Type == IntVMObject {
		return e.Value.(int32), true
	}
	return 0, false
}

// ToFloat returns the float64 value if the type is DoubleVMObject, otherwise returns 0.0.
func (e VMObject) ToFloat() (float64, bool) {
	if e.Type == DoubleVMObject {
		return e.Value.(float64), true
	}
	return 0.0, false
}

// ToString returns the string value if the type is StringVMObject, otherwise returns "".
func (e VMObject) ToString() (string, bool) {
	if e.Type == StringVMObject {
		return e.Value.(string), true
	}
	return "", false
}

// ToField returns the Field value if the type is FieldVMObject, otherwise returns nil.
func (e VMObject) ToField() (*Field, bool) {
	if e.Type == FieldVMObject {
		return e.Value.(*Field), true
	}
	return nil, false
}

// ToMethod returns the Method value if the type is MethodVMObject, otherwise returns nil.
func (e VMObject) ToMethod() (*Method, bool) {
	if e.Type == MethodVMObject {
		return e.Value.(*Method), true
	}
	return nil, false
}

// ToBlock returns the Block value if the type is BlockVMObject, otherwise returns nil.
func (e VMObject) ToBlock() (*Block, bool) {
	if e.Type == BlockVMObject {
		return e.Value.(*Block), true
	}
	return nil, false
}

// ToClass returns the Class value if the type is ClassVMObject, otherwise returns nil.
func (e VMObject) ToClass() (*Class, bool) {
	if e.Type == ClassVMObject {
		return e.Value.(*Class), true
	}
	return nil, false
}

// ToObject returns the Object value if the type is ObjectVMObject, otherwise returns nil.
func (e VMObject) ToObject() (*VMObject, bool) {
	if e.Type == ObjectVMObject {
		return e.Value.(*VMObject), true
	}
	return nil, false
}

// ToNil returns true if the type is NilVMObject, otherwise returns false.
func (e VMObject) ToNil() bool {
	return e.Type == NilVMObject
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
	NameIx int
	Impl   func() // Implementation in the VM
}

type Block struct {
	Arity        int
	Instructions []byte
}

type Class struct {
	Name        string // The name of the class
	ArityField  int    // Number of fields in the class
	ArityMethod int    // Number of methods in the class
	Fields      []int  // Indices to Field entities in the heap
	//Methods     []*Method // Map of method names to their indices in the heap
	SuperClass *Class // Reference to the superclass (if any)
}

type Frame struct {
	Arguments      []*VMObject   // Arguments passed to the method
	Locals         []*VMObject   // Local variables within the frame
	Stack          *OperandStack // Operand stack for the frame
	InitialAddress CodeAddress   // Initial address of the code
	ReturnAddress  CodeAddress   // Address to return after executing the frame
}
type CallStack struct {
	Frames []*Frame
}

func (es *CallStack) pushFrame(f *Frame) {
	es.Frames = append(es.Frames, f)
}

func (es *CallStack) popFrame() *Frame {
	if len(es.Frames) == 0 {
		return nil
	}
	f := es.Frames[len(es.Frames)-1]
	es.Frames = es.Frames[:len(es.Frames)-1]
	return f
}

func NewFrame() *Frame {
	return &Frame{
		Arguments:      []*VMObject{},
		Locals:         []*VMObject{},
		Stack:          NewOperandStack(),
		InitialAddress: CodeAddress{instructionPointer: 0},
		ReturnAddress:  CodeAddress{instructionPointer: 0},
	}
}

// func (es *CallStack) push(obj VMObject) {
// 	es.Frames[len(es.Frames)-1].Stack = append(es.Frames[len(es.Frames)-1].Stack, obj)
// }

// func (es *CallStack) pop() VMObject {
// 	Stack := es.Frames[len(es.Frames)-1].Stack
// 	obj := Stack[len(Stack)-1]
// 	es.Frames[len(es.Frames)-1].Stack = Stack[:len(Stack)-1]
// 	return obj
// }

type Heap struct {
	objects []*VMObject
	invmap  map[*VMObject]int
}

func NewHeap() *Heap {
	return &Heap{
		objects: []*VMObject{},
		invmap:  make(map[*VMObject]int),
	}
}

func (h *Heap) allocObject() *VMObject {
	obj := &VMObject{Fields: make([]int, 1)}
	h.objects = append(h.objects, obj)
	h.invmap[obj] = len(h.objects) - 1
	return obj
}

func (h *Heap) get(index int) *VMObject {
	return h.objects[index]
}

func (h *Heap) set(index int, obj *VMObject) {
	h.objects[index] = obj
}

func (h *Heap) getInv(obj *VMObject) int {
	return h.invmap[obj]
}

func (h *Heap) setInv(obj *VMObject, index int) {
	h.invmap[obj] = index
}

// func (h *Heap) newClazz(cname string, fields []int, methods []int, superc *VMClass) *VMClass {

// 	// Assuming we are adding a class to the heap
// 	vmClass := &VMClass{
// 		Name: "MyClass",
// 		Fields: []int{fieldIndex1, fieldIndex2}, // Indices to Field entities
// 		Methods: map[string]int{
// 			"initialize": methodIndex1,
// 			"doSomething": methodIndex2,
// 		},
// 		SuperClass: superc, // or reference to another VMClass for inheritance
// 	}

// 	// Add the class entity to the heap
// 	vm.ConstantPool = append(vm.ConstantPool, VMObject{
// 		Type:  ClassVMObject,
// 		Class: vmClass,
// 	})

// 	return vmClass
// }

func (h *Heap) NewInstanceOf(class *Class) *VMObject {
	obj := &VMObject{
		Fields: make([]int, class.ArityField),
		Class:  class,
	}
	h.objects = append(h.objects, obj)
	return obj
}

type GlobalContext struct {
	classTable map[string]*Class // Stores all the classes available in the context
	trueObj,
	falseObj,
	nilObj *VMObject // Predefined true, false, nil objects
}

type CodeAddress struct {
	instructionPointer int
}

type Program struct {
	bytecode []byte // Bytecode representation of the program
}

type VM struct {
	CallStack CallStack
	Heap      Heap
}

// readability functions for the VM
func (vm *VM) TopFrame() *Frame {
	return vm.CallStack.Frames[len(vm.CallStack.Frames)-1]
}
func (vm *VM) TopFrameOpStack() *OperandStack {
	return vm.CallStack.Frames[len(vm.CallStack.Frames)-1].Stack
}

// OperandStack is a generic stack data structure that holds Entities.
type OperandStack struct {
	items []VMObject
}

// NewOperandStack creates a new OperandStack.
func NewOperandStack() *OperandStack {
	return &OperandStack{items: []VMObject{}}
}

// Push adds an item to the top of the stack.
func (s *OperandStack) Push(item VMObject) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack.
// It returns false if the stack is empty.
func (s *OperandStack) Pop() (VMObject, bool) {
	if len(s.items) == 0 {
		return VMObject{}, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the item from the top of the stack without removing it.
// It returns false if the stack is empty.
func (s *OperandStack) Peek() (VMObject, bool) {
	if len(s.items) == 0 {
		return VMObject{}, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty.
func (s *OperandStack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *OperandStack) Size() int {
	return len(s.items)
}

// one issue that needs to be resolved, is:
// when do we de-reference the fields of an object to an actual value for computation?
// and when do we keep it as a reference to a field in the heap?
// this is important for the VM to know when to de-reference a field and when to keep it as a reference
// for example, in the Lit instruction, we push a constant from the heap onto the stack
// we need to know if the constant is a primitive or a field in the heap
// if it is a field, we need to push the value of the field onto the stack
// if it is a primitive, we need to push the primitive onto the stack (?) or push a reference to the primitive/field?

func (i *Interpreter) Interpret(p *Program, gc *GlobalContext) {
	// Execution loop to interpret bytecode
	// Fetch, decode, and execute instructions
	// Manage the stack, heap, and Frames
}

func main() {
	// Initialize global context with basic objects and classes
	globalContext := &GlobalContext{
		classTable: make(map[string]*Class),
		trueObj:    &VMObject{Fields: make([]int, 0)},
		falseObj:   &VMObject{Fields: make([]int, 0)},
		nilObj:     &VMObject{Fields: make([]int, 0)},
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
