package vm

import "fmt"

type Double struct {
	embeddedDouble float64
}

func NewDouble(d float64) *Double {
	return &Double{embeddedDouble: d}
}

func (d *Double) getEmbeddedDouble() float64 {
	return d.embeddedDouble
}

func (d *Double) setEmbeddedDouble(value float64) {
	d.embeddedDouble = value
}

func (d *Double) String() string {
	return fmt.Sprintf("%v", d.embeddedDouble)
}
