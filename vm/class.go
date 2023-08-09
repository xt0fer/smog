package vm

type Class struct {
	InvokablesTable map[*Symbol]*Invokable
}
