package smog

const (
	clsDouble  = "Double"
	clsString  = "String"
	clsSymbol  = "Symbol"
	clsArray   = "Array"
	clsMethod  = "Method"
	clsBlock   = "Block"
	clsBlock1  = "Block1"
	clsBlock2  = "Block2"
	clsClass   = "Class"
)

type VMObject struct {
	Clazz  *VMObject
	Kind   string
	Fields []*VMObject // local vars (any object) index of field is same as index of Class.InstanceFields
	N      int32
}

type Clazz struct {
	ClassFields     []*VMObject
	ClassMethods    []*VMObject
	InstanceFields  []*VMObject
	InstanceMethods []*VMObject
}

// type ObjectInterface interface {
// 	Send(selectorString string, arguments []VMObject, interpreter *Interpreter)
// 	SendDoesNotUnderstand(selector string, interpreter *Interpreter)
// 	SendUnknownGlobal(globalName VMObject, interpreter *Interpreter)
// 	SendEscapedBlock(block VMObject, interpreter *Interpreter)
// }

/* Literals are int, double, string, symbol, array, method, block */
type VMInteger struct {
	VMObject
	Value int32
}

type VMDouble struct {
	VMObject
	Value float64
}

func (i *VMInteger) doPlus(recv *VMObject, arg *VMObject) *VMObject {
	if recv.Kind == "Double" || arg.Kind == "Double" {
		return NewVMDouble(float64(recv.Value) + arg.Value)
	} else {
		return NewVMInteger(recv.Value + arg.Value)
	}
	return
}

func NewVMInteger(value int32) *VMInteger {
	return &VMInteger{Value: value}
}
func (i *VMInteger) Send(selectorString string, arguments []*VMObject, interpreter *Interpreter) {
	// TODO
}
func (i *VMInteger) SendDoesNotUnderstand(selector string, interpreter *Interpreter) {
	// TODO
}
func (i *VMInteger) SendUnknownGlobal(globalName *VMObject, interpreter *Interpreter) {
	// TODO
}
func (i *VMInteger) SendEscapedBlock(block *VMObject, interpreter *Interpreter) {
	// TODO
}

/*
Need 5 things:
- heap
- global variables; a map of keys to values of objects kept
- literals (constant pool)
- execution (operand) stack
- call stack
*/

// create a HEAP of vmobjects
type Symbol string

func (h *Heap) NewObject(clazz *VMObject) *VMObject {
	obj := &VMObject{Clazz: clazz}
	h.Objects = append(h.Objects, obj)
	return obj
}

// create a constant pool
type ConstantPool struct {
	Constants []VMObject
}

// type VMObjectInterner struct {
// 	Map map[VMObject]int32
// 	Vec []VMObject
// 	Buf VMObject
// }
// func (si *StringInterner) Intern(value string) int32 {
// 	if index, ok := si.Map[value]; ok {
// 		return index
// 	}
// 	index := int32(len(si.Vec))
// 	si.Map[value] = index
// 	si.Vec = append(si.Vec, value)
// 	return index
// }
// func (si *StringInterner) Lookup(index int32) string {
// 	return si.Vec[index]
// }
// create an execution stack
