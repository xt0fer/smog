package vm

import "github.com/xt0fer/smog/vmobjects"

type Universe struct {
	Globals map[vmobjects.Symbol]vmobjects.Object
	//
	PathSep string
	FileSep string
}

func NewUniverse() *Universe {
	return &Universe{}
}

var universe = NewUniverse()

func (u *Universe) NewArray(size int) *vmobjects.Array {
	return vmobjects.NewArray(size)
}