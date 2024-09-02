package vm

import "github.com/xtofer/smog/gsrc/internal/vmobj" // Replace "path/to/vm/package" with the actual import path of the package that defines the Symbol type.

type ClassTable struct {
	Map map[string]*vmobj.VMClass // Replace "package" with the actual package name.
}

func (ct *ClassTable) Lookup(s string) *vmobj.VMClass {
	return ct.Map[s]
}

func (ct *ClassTable) Insert(sym *vmobj.VMClass) {

	ct.Map[sym.ToString()] = sym
}

var classTable *ClassTable

func init() {
	classTable = &ClassTable{make(map[string]*vmobj.VMClass)}
	classTable.Insert(vmobj.NewVMClass("VMClass", 0, 0))
	classTable.Insert(vmobj.NewVMClass("VMObject", 0, 0))
	classTable.Insert(vmobj.NewVMClass("VMSymbol", 0, 0))
}
