package arrayqueue

import (
	"bytes"
	"fmt"

	"github.com/Thrimbda/dune/arrayutils"
	"github.com/Thrimbda/dune/list/arraylist"
)

type ArrayQueue struct {
	size  int
	front int
	rear  int
	list  *arraylist.ArrayList
}

func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{size, 0, 0, arraylist.ConvertToArrayList(size, make([]interface{}, size)...)}
}

func ConvertToArrayQueue(size int, items ...interface{}) *ArrayQueue {
	if size <= len(items) {
		panic(&arrayutils.FullListError{})
	}
	return &ArrayQueue{size, 0, len(items), arraylist.ConvertToArrayList(size, items...)}
}

func (a *ArrayQueue) Clear() {
	a.front = 0
	a.rear = 0
}

func (a *ArrayQueue) Enqueue(item interface{}) {
	if a.size == 0 || (a.rear+1)%a.size == a.front {
		panic(&arrayutils.FullListError{})
	}
	a.list.Set(a.rear, item)
	a.rear = (a.rear + 1) % a.size
}

func (a *ArrayQueue) Dequeue() interface{} {
	if a.IsEmpty() {
		panic(&arrayutils.EmptyListError{})
	}
	value := a.list.Get(a.front)
	a.front = (a.front + 1) % a.size
	return value
}

func (a *ArrayQueue) Peek() interface{} {
	if a.IsEmpty() {
		panic(&arrayutils.EmptyListError{})
	}
	return a.list.Get(a.front)
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.front == a.rear
}

func (a *ArrayQueue) String() string {
	var buffer bytes.Buffer
	if a.IsEmpty() {
		return "()"
	}
	buffer.WriteString("(")
	for i := a.front; i%a.size != a.rear; i++ {
		value := a.list.Get(i)
		if (i+1)%a.size != a.rear {
			buffer.WriteString(fmt.Sprintf("%v, ", value))
		} else {
			buffer.WriteString(fmt.Sprintf("%v)", value))
		}
	}
	return buffer.String()
}
