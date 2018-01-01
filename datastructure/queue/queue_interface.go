package queue

//Queue interface.
type Queue interface {
	Clear()
	Enqueue(item interface{})
	Dequeue() interface{}
	Peek() interface{}
	IsEmpty() bool
}
