package queue

//Queue interface.
type Queue interface {
	Clear()
	Enqueue(item interface{})
	Dequeue() interface{}
	FirstValue() interface{}
	IsEmpty() bool
}
