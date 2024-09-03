package vm

const (
	HALT             byte = 0
	DUP              byte = 1
	PUSH_LOCAL       byte = 2
	PUSH_ARGUMENT    byte = 3
	PUSH_FIELD       byte = 4
	PUSH_BLOCK       byte = 5
	PUSH_CONSTANT    byte = 6
	PUSH_GLOBAL      byte = 7
	POP              byte = 8
	POP_LOCAL        byte = 9
	POP_ARGUMENT     byte = 10
	POP_FIELD        byte = 11
	SEND             byte = 12
	SUPER_SEND       byte = 13
	RETURN_LOCAL     byte = 14
	RETURN_NON_LOCAL byte = 15
)

var bytecodeNames = []string{
	"HALT            ",
	"DUP             ",
	"PUSH_LOCAL      ",
	"PUSH_ARGUMENT   ",
	"PUSH_FIELD      ",
	"PUSH_BLOCK      ",
	"PUSH_CONSTANT   ",
	"PUSH_GLOBAL     ",
	"POP             ",
	"POP_LOCAL       ",
	"POP_ARGUMENT    ",
	"POP_FIELD       ",
	"SEND            ",
	"SUPER_SEND      ",
	"RETURN_LOCAL    ",
	"RETURN_NON_LOCAL",
}

func getBytecodeLength(bytecode byte) int {
	// Return the length of the given bytecode
	return bytecodeLength[bytecode]
}

var bytecodeLength = []int{ // this is the Arity of the bytecode
	1, // halt
	1, // dup
	3, // push_local
	3, // push_argument
	2, // push_field
	2, // push_block
	2, // push_constant
	2, // push_global
	1, // pop
	3, // pop_local
	3, // pop_argument
	2, // pop_field
	2, // send
	2, // super_send
	1, // return_local
	1, // return_non_local
}
