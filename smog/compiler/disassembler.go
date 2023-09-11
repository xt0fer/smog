package compiler

import (
	"fmt"

	"github.com/xt0fer/smog"
)

func Dump(cl *smog.Class) {
	for i := 0; i < cl.GetNumberOfInstanceInvokables(); i++ {
		inv := cl.GetInstanceInvokable(i)
		// output header and skip if the Invokable is a Primitive
		fmt.Printf("%s>>%s = ", cl.GetName().String(), inv.GetSignature().String())
		if inv.IsPrimitive() {
			fmt.Println("<primitive>")
			continue
		}
		// output actual method
		dumpMethod(inv.(*smog.Method), "\t")
	}
}

func dumpMethod(m *smog.Method, indent string) {
	fmt.Println("(")
	// output stack information
	fmt.Printf("%s<%d locals, %d stack, %d bc_count>\n", indent, m.GetNumberOfLocals(), m.GetMaximumNumberOfStackElements(), m.GetNumberOfBytecodes())
	// output bytecodes
	for b := 0; b < m.GetNumberOfBytecodes(); b += smog.GetBytecodeLength(m.GetBytecode(b)) {
		fmt.Print(indent)
		// bytecode index
		if b < 10 {
			fmt.Print(" ")
		}
		if b < 100 {
			fmt.Print(" ")
		}
		fmt.Printf(" %d:", b)
		// mnemonic
		bytecode := m.GetBytecode(b)
		fmt.Printf("%s  ", smog.BytecodeNames[bytecode])
		// parameters (if any)
		if smog.GetBytecodeLength(bytecode) == 1 {
			fmt.Println()
			continue
		}
		switch bytecode {
		case smog.PUSHLOCAL:
			fmt.Printf("local: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case smog.PUSHARGUMENT:
			fmt.Printf("argument: %d, context %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case smog.PUSHFIELD:
			fmt.Printf("(index: %d) field: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*smog.Symbol).ToString())
		case smog.PUSHBLOCK:
			fmt.Printf("block: (index: %d) ", m.GetBytecode(b+1))
			dumpMethod(m.GetConstant(b).(*smog.Method), indent+"\t")
		case smog.PUSHCONSTANT:
			constant := m.GetConstant(b)
			fmt.Printf("(index: %d) value: (%s) %s\n", m.GetBytecode(b+1), constant.GetSOMClass().GetName().ToString(), constant.ToString())
		case smog.PUSHGLOBAL:
			fmt.Printf("(index: %d) value: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*smog.Symbol).ToString())
		case smog.POPLOCAL:
			fmt.Printf("local: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case smog.POPARGUMENT:
			fmt.Printf("argument: %d, context: %d\n", m.GetBytecode(b+1), m.GetBytecode(b+2))
		case smog.POPFIELD:
			fmt.Printf("(index: %d) field: %s\n", m.GetBytecode(b+1), m.GetConstant(int(b)).(*smog.Symbol).ToString())
		case smog.SEND:
			fmt.Printf("(index: %d) signature: %s\n", m.GetBytecode(b+1), m.GetConstant(b).(*smog.Symbol).ToString())
		case smog.SUPERSEND:
			fmt.Printf("(index: %d) signature: %s\n", m.GetBytecode(b+1), m.GetConstant(b).(*smog.Symbol).ToString())
		default:
			fmt.Println("<incorrect bytecode>")
		}
	}
	fmt.Println(indent + ")")
}
