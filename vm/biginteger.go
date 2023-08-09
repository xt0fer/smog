package vm

import (
	"fmt"
	"math/big"
)

type BigInteger struct {
	embeddedBigInt big.Int
}

func NewBigInteger(i big.Int) *BigInteger {
	return &BigInteger{embeddedBigInt: i}
}

func (i *BigInteger) getEmbeddedInteger() big.Int {
	return i.embeddedBigInt
}

func (i *BigInteger) setEmbeddedInteger(value big.Int) {
	i.embeddedBigInt = value
}

func (i *BigInteger) String() string {
	return fmt.Sprintf("%v", i.embeddedBigInt)
}
