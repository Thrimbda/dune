package queue

import (
	. "github.com/Thrimbda/dune/datastructure/arrayutils"
	"github.com/Thrimbda/dune/datastructure/list/arraylist"
)

type ArrayQueue struct {
	front int
	rear  int
	list  *arraylist.ArrayList
}

func NewArrayQueue(size int) *ArrayQueue {
	list := arraylist.NewArrayList(size)
	return &ArrayQueue{0, 0, list}
}

func (a ArrayQueue) clear() {
	a.list.Clear()
}

func (a ArrayQueue) Enqueue(item interface{}) {
	if a.rear+1%a.list.Length() == a.front {
		panic(&FullListError{})
	}
	a.rear = (a.rear + 1) % a.list.Length()
	a.list.SetValue(a.rear, item)
}

func (a ArrayQueue) Dequeue() interface{} {
	if a.IsEmpty() {
		panic(&EmptyListError{})
	}
	a.front = (a.front + 1) % a.list.Length()
	return a.list.Get(a.front)
}

func (a ArrayQueue) Peek() interface{} {
	if a.IsEmpty() {
		panic(&EmptyListError{})
	}
	return a.list.Get((a.front + 1) % a.list.Length())
}

func (a ArrayQueue) IsEmpty() bool {
	return a.front == a.rear
}
