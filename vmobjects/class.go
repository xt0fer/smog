package vmobjects

type Class struct {
	InvokablesTable map[*Symbol]*Invokable

}