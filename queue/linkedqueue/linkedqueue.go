package linkedqueue

import (
	"github.com/Thrimbda/dune/list/linkedlist"
)

// LinkedQueue is a queue implement with linklist.
type LinkedQueue struct {
	list *linkedlist.LinkedList
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{linkedlist.NewLinkedList()}
}

func ConvertToLinkedQueue(items ...interface{}) *LinkedQueue {
	return &LinkedQueue{linkedlist.ConvertToLinkedList(items...)}
}

func (l *LinkedQueue) Clear() {
	l.list.Clear()
}

func (l *LinkedQueue) Enqueue(item interface{}) {
	l.list.Append(item)
}

func (l *LinkedQueue) Dequeue() interface{} {
	return l.list.Remove(0)
}

func (l *LinkedQueue) Peek() interface{} {
	return l.list.Get(0)
}

func (l *LinkedQueue) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedQueue) String() string {
	return l.list.String()
}
