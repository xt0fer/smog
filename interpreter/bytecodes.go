package interpreter


const (
	numBytecodes = 16
)

const (
	halt = iota //0
	dup
	push_local
	push_argument
	push_field
	push_block //5
	push_constant
	push_global
	pop
	pop_local
	pop_argument //10
	pop_field
	send
	super_send
	return_local
	return_non_local //15
)

var bytecodeNames [numBytecodes]string = [numBytecodes]string{
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

var bytecodeLength [numBytecodes]int = [numBytecodes]int{1,1,3,3,2,2,2,2,1,3,3,2,2,2,1,1}

