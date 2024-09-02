package vm

import (
	"testing"
)

func TestClassTable(t *testing.T) {

	vmo := classTable.Lookup("VMObject")
	if vmo == nil {
		t.Error("VMObject not found")

	}

	if vmo.NumberOfClassVars() != 0 {
		t.Error("Number of class vars not 0")
	}

	vmc := classTable.Lookup("VMClass")
	if vmc == nil {
		t.Error("VMClass not found")

	}

	vms := classTable.Lookup("VMSymbol")
	if vms == nil {
		t.Error("VMSymbol not found")

	}
	vmi := classTable.Lookup("VMInteger")
	if vmi == nil {
		t.Error("VMInteger not found")

	}

}
