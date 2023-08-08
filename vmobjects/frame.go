package vmobjects

type Frame struct {
	Array
}

func NewFrame() *Frame {
	return &Frame{}
}

// frame is a subtype of array
