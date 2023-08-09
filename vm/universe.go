package vm

import (
	"sync"

	"github.com/xt0fer/smog/vmobjects"
)

type universe struct {
	Globals map[vmobjects.Symbol]vmobjects.Object
	//
	PathSep string
	FileSep string
	//
	NillObject *vmobjects.Object
}

var instantiated *universe
var once sync.Once

func Universe() *universe {
	once.Do(func() {
		instantiated = &universe{}
	})
	return instantiated
}

func (u *universe) NewArray(size int) *vmobjects.Array {
	return vmobjects.NewArray(size)
}
