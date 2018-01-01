package queue

import (
	. "github.com/Thrimbda/dune/datastructure/arrayutils"
	"github.com/Thrimbda/dune/datastructure/list/arraylist"
)

type ArrayQueue struct {
	size  int
	front int
	rear  int
	list  *arraylist.ArrayList
}

func NewArrayQueue(size int) *ArrayQueue {
	list := arraylist.NewArrayList(size)
	return &ArrayQueue{size, 0, 0, list}
}

func (a *ArrayQueue) Clear() {
	a.front = 0
	a.rear = 0
}

func (a *ArrayQueue) Enqueue(item interface{}) {
	if a.size == 0 || a.rear+1%a.size == a.front {
		panic(&FullListError{})
	}
	if a.rear+1 >= a.list.Length() {
		a.rear = (a.rear + 1) % a.size
		a.list.Append(item)
	} else {
		a.list.SetValue(a.rear, item)
		a.rear += 1
	}
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.IsEmpty() {
		panic(&EmptyListError{})
	}
	value := a.list.Get(a.front)
	a.front = (a.front + 1) % a.size
	return value
}

func (a *ArrayQueue) Peek() interface{} {
	if a.IsEmpty() {
		panic(&EmptyListError{})
	}
	return a.list.Get(a.front)
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.front == a.rear
}
