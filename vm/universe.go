package vm

import (
	"sync"
)

type Universe struct {
	Globals map[*Symbol]*Object
	symboltable *SymbolTable
	//
	PathSep string
	FileSep string
	//
	NilObject   *Object
	TrueObject  *Object
	FalseObject *Object

	ObjectClass    *Class
	ClassClass     *Class
	MetaclassClass *Class

	NilClass        *Class
	IntegerClass    *Class
	BigIntegerClass *Class
	ArrayClass      *Class
	MethodClass     *Class
	SymbolClass     *Class
	FrameClass      *Class
	PrimitiveClass  *Class
	StringClass     *Class
	BlockClass      *Class
	DoubleClass     *Class
}

var instantiated *Universe
var once sync.Once

func GetUniverse() *Universe {
	once.Do(func() {
		instantiated = &Universe{}
		instantiated.initUniverse()
	})
	return instantiated
}

func (u *Universe) initUniverse() {
	u.symboltable = NewSymbolTable()
}

func (u *Universe) NewArray(size int) *Array {
	return NewArray(size)
}

func (u *Universe) symbolFor(sym string) *Symbol {
	result := u.symboltable.lookup(sym)
	if result != nil {
		return result
	}
	result = u.newSymbol(sym)
	return result
}

func (u *Universe) newSymbol(sym string) *Symbol {
	result := NewSymbol(sym, 1)
	result.setClass(*u.SymbolClass)
	u.symboltable.insert(result)
	return result
}
