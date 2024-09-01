package vm

import (
	"fmt"
)

type SymbolTable struct {
	  Map map[string]*Symbol
}

func (st *SymbolTable) Lookup(s string) *Symbol {
	return st.Map[s]
}

func (st *SymbolTable) Insert(sym *Symbol) {

	st.Map[sym.String()] = sym
}

