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

func (a ArrayQueue) Enqueue(item interface{}) error {
	if a.rear+1%a.list.Length() == a.front {
		return FullListError{}
	}
	a.rear = (a.rear + 1) % a.list.Length()
	a.list.SetValue(a.rear, item)
	return nil
}

func (a ArrayQueue) Dequeue() (interface{}, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	a.front = (a.front + 1) % a.list.Length()
	return a.list.Get(a.front), nil
}

func (a ArrayQueue) Peek() (interface{}, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	return a.list.Get((a.front + 1) % a.list.Length()), nil
}

func (a ArrayQueue) IsEmpty() bool {
	return a.front == a.rear
}
