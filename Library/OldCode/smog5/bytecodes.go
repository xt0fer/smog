package smog

type Bytecode int

const (
	HALT             Bytecode = iota // 1byte
	DUP                              // 1 byte
	PUSH_LOCAL                       // 3 byte
	PUSH_ARGUMENT                    // 3 byte
	PUSH_FIELD                       // 2 byte
	PUSH_BLOCK                       // 2 byte
	PUSH_CONSTANT                    // 2 byte
	PUSH_GLOBAL                      // 2 byte
	POP                              // 1 byte
	POP_LOCAL                        // 3 byte
	POP_ARGUMENT                     // 3 byte
	POP_FIELD                        // 2 byte
	SEND                             // 2 byte
	SUPER_SEND                       // 2 byte
	RETURN_LOCAL                     // 1 byte
	RETURN_NON_LOCAL                 // 1 byte
)

type Bytecodes struct {
	Bytecodes []Bytecode
	Bidx 	int
}

func (bc *Bytecodes) doHalt() {

	// stop the interpreter
}

func (bc *Bytecodes) doDup() {
	// duplicate the top of the stack
}

func (bc *Bytecodes) doPushLocal() {
	// push a local variable onto the stack
}

func (bc *Bytecodes) doPushArgument() {
	// push an argument onto the stack
}

func (bc *Bytecodes) doPushField() {
	// push a field onto the stack
}

func (bc *Bytecodes) doPushBlock() {
	// push a block onto the stack
}

func (bc *Bytecodes) doPushConstant() {
	// push a constant onto the stack
}

func (bc *Bytecodes) doPushGlobal() {
	// push a global onto the stack
}

func (bc *Bytecodes) doPop() {
	// pop the top of the stack
}

func (bc *Bytecodes) doPopLocal() {
	// pop a local variable
}

func (bc *Bytecodes) doPopArgument() {
	// pop an argument
}

func (bc *Bytecodes) doPopField() {
	// pop a field
}

func (bc *Bytecodes) doSend() {
	// send a message
}

func (bc *Bytecodes) doSuperSend() {
	// send a message to the superclass
}

func (bc *Bytecodes) doReturnLocal() {
	// return a local variable
}

func (bc *Bytecodes) doReturnNonLocal() {
	// return a non-local variable
}

func (bc *Bytecodes) Interpret() {
	bc.Bidx = 0 // bytecode pointer

	// interpret the bytecodes
	for _, bytecode := range bc.Bytecodes {
		switch bytecode {
		case HALT:
			bc.doHalt()
		case DUP:
			bc.doDup()
		case PUSH_LOCAL:
			bc.doPushLocal()
		case PUSH_ARGUMENT:
			bc.doPushArgument()
		case PUSH_FIELD:
			bc.doPushField()
		case PUSH_BLOCK:
			bc.doPushBlock()
		case PUSH_CONSTANT:
			bc.doPushConstant()
		case PUSH_GLOBAL:
			bc.doPushGlobal()
		case POP:
			bc.doPop()
		case POP_LOCAL:
			bc.doPopLocal()
		case POP_ARGUMENT:
			bc.doPopArgument()
		case POP_FIELD:
			bc.doPopField()
		case SEND:
			bc.doSend()
		case SUPER_SEND:
			bc.doSuperSend()
		case RETURN_LOCAL:
			bc.doReturnLocal()
		case RETURN_NON_LOCAL:
			bc.doReturnNonLocal()
		}
	}
}
