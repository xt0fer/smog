package vm

type SymbolTable struct {
	smap map[string]*Symbol
}

func NewSymbolTable() *SymbolTable {
	nst := &SymbolTable{}
	nst.smap = make(map[string]*Symbol)
	return nst
}

func (st *SymbolTable) lookup(name string) *Symbol {
	return st.smap[name]
}

func (st *SymbolTable) insert(sym *Symbol) {
	st.smap[sym.Name] = sym
}
