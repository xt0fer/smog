package vm

import "github.com/xtofer/smog/gsrc/internal/vmobj" // Replace "path/to/vm/package" with the actual import path of the package that defines the Symbol type.

type SymbolTable struct {
	Map map[string]*vmobj.VMSymbol // Replace "package" with the actual package name.
}

func (st *SymbolTable) Lookup(s string) *vmobj.VMSymbol {
	return st.Map[s]
}

func (st *SymbolTable) Insert(sym *vmobj.VMSymbol) {

	st.Map[sym.ToString()] = sym
}
