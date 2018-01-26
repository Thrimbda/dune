package stack

type Stack interface {
	Clear()
	Push(item interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}
