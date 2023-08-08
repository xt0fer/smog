package vmobjects

type Invokable interface {
	IsPrimitive() bool;
	Invoke(frame *Frame);
	GetSignature() *Symbol;
	GetHolder() *Class;
	SetHolder(value *Class);
}
