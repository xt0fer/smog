package smog

import (
	"fmt"
)

type BigInteger struct {
	Object
	embeddedBigInt int64
}

func NewBigInteger(i int64) *BigInteger {
	return &BigInteger{embeddedBigInt: i}
}

func (i *BigInteger) getEmbeddedInteger() int64 {
	return i.embeddedBigInt
}

func (i *BigInteger) setEmbeddedInteger(value int64) {
	i.embeddedBigInt = value
}

func (i *BigInteger) String() string {
	return fmt.Sprintf("%v", i.embeddedBigInt)
}
