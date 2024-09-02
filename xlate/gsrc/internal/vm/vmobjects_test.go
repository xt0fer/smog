package vm

import (
	"testing"

	"github.com/xtofer/smog/gsrc/internal/vmobj"
)

func TestVMObjects(t *testing.T) {

	vmo := vmobj.NewObject(2, classTable.Lookup("VMObject"))

	if vmo.NumberOfFields() != 2 {
		t.Error("Number of class vars not 2")
	}

	ts := "test-symbol"
	vms := vmobj.NewVMSymbol(ts)
	if vms == nil {
		t.Error("VMSymbol not found")

	}
	if vms.ToString() != ts {
		t.Error("VMSymbol ToString() failed")
	}

	vmi := classTable.Lookup("VMInteger")
	if vmi == nil {
		t.Error("VMInteger not found")

	}

}
