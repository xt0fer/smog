package vm

import "github.com/xt0fer/smog/vmobjects"

type SymbolTable struct {
	smap map[string]*vmobjects.Symbol
}

func NewSymbolTable() *SymbolTable {
	nst := &SymbolTable{}
	nst.smap = make(map[string]*vmobjects.Symbol)
	return nst
}

func (st *SymbolTable) lookup(name string) *vmobjects.Symbol {
	return st.smap[name];
}


func (st *SymbolTable) insert(sym *vmobjects.Symbol) {
	st.smap[sym.Name] = sym
}