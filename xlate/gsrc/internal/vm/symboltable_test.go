package vm

import (
	"testing"

	"github.com/xtofer/smog/gsrc/internal/vmobj" // Replace "path/to/vm/package" with the actual import path of the package that defines the Symbol type.
)

func TestSymbolTable(t *testing.T) {

	st := &SymbolTable{make(map[string]*vmobj.VMSymbol)}

	sym := vmobj.NewVMSymbol("test")

	n := sym.NumberOfFields()

	if n != 0 {
		t.Error("Number of fields not 0")
	}

	st.Insert(sym)

	if st.Lookup("test") != sym {
		t.Error("Symbol not found")

	}

	if st.Lookup("test2") != nil {
		t.Error("Symbol found")
	}

}
