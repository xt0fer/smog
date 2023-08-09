package vm

import (
	"sync"
)

type Universe struct {
	Globals map[Symbol]Object
	//
	PathSep string
	FileSep string
	//
	NilObject *Object
}

var instantiated *Universe
var once sync.Once

func GetUniverse() *Universe {
	once.Do(func() {
		instantiated = &Universe{}
	})
	return instantiated
}

func (u *Universe) NewArray(size int) *Array {
	return NewArray(size)
}
