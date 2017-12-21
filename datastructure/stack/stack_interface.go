package stack

type Stack interface {
	Setup(size int)
	Clear()
	Push(item interface{})
	Pop() interface{}
	TopValue() interface{}
	IsEmpty() bool
}
