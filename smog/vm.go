package smog

// virtual machine runtme object
type Universe struct {
	Heap         *Heap
	GlobalVars   *GlobalVars //classes or objects
	CallStack    *CallStack
	OperandStack *OperandStack
	Literals     *Literals
}

func NewUniverse() *Universe {
	return &Universe{
		Heap:         &Heap{Objects: make([]*VMObject, 0)},
		GlobalVars:   &GlobalVars{Variables: make(map[string]*VMObject)},
		CallStack:    &CallStack{Stack: make([]Frame, 0)},
		OperandStack: &OperandStack{Stack: make([]*VMObject, 0)},
		Literals:     &Literals{Map: make(map[string]int32), Constants: make([]*VMObject, 0)},
	}
}

type Heap struct {
	Objects []*VMObject
}

type GlobalVars struct {
	Variables map[string]*VMObject
}

func (gv *GlobalVars) Get(name string) *VMObject {
	return gv.Variables[name]
}
func (gv *GlobalVars) Set(name string, value *VMObject) {
	gv.Variables[name] = value
}

/*
 * OperandStack
 */
type OperandStack struct {
	Stack []*VMObject
}

func (es *OperandStack) Push(obj *VMObject) {
	es.Stack = append(es.Stack, obj)
}
func (es *OperandStack) Pop() *VMObject {
	if len(es.Stack) == 0 {
		return &VMObject{}
	}
	obj := es.Stack[len(es.Stack)-1]
	es.Stack = es.Stack[:len(es.Stack)-1]
	return obj
}

/*
 * CallStack
 */
type Frame struct {
	Receiver  *VMObject
	Method    *VMObject
	IP        int
	Locals    []*VMObject
	Stack     []*VMObject
	Arguments []*VMObject
}
type CallStack struct {
	Stack []Frame
}

func (cs *CallStack) Push(frame Frame) {
	cs.Stack = append(cs.Stack, frame)
}
func (cs *CallStack) Pop() Frame {
	if len(cs.Stack) == 0 {
		return Frame{}
	}
	frame := cs.Stack[len(cs.Stack)-1]
	cs.Stack = cs.Stack[:len(cs.Stack)-1]
	return frame
}

/*
 * Literals
 */
type Literals struct {
	Map       map[string]int32
	Constants []*VMObject
}
func (l *Literals) Intern(value string) int32 {
	if index, ok := l.Map[value]; ok {
		return index
	}
	index := int32(len(l.Constants))
	l.Map[value] = index
	l.Constants = append(l.Constants, &VMObject{})
	return index
}
func (l *Literals) Lookup(index int32) *VMObject {
	return l.Constants[index]
}
