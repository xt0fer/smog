package smog

type SymbolTable struct {
	symmap map[string]*Symbol
}

func NewSymbolTable() *SymbolTable {
	nst := &SymbolTable{}
	nst.symmap = make(map[string]*Symbol)
	return nst
}

func (st *SymbolTable) lookup(name string) *Symbol {
	return st.symmap[name]
}

func (st *SymbolTable) insert(sym *Symbol) {
	st.symmap[sym.Name] = sym
}
