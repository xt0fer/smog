package vm

import "github.com/xt0fer/smog/vmobjects"

type Universe struct {
	Globals map[vmobjects.Symbol]vmobjects.Object
	//
	PathSep string
	FileSep string
}