package queue

import (
	. "../../datastructure"
	. "../arrayutils"
	. "../linkutils"
)

type LinkedQueue struct {
	front LinkNode
	rear  LinkNode
}

func (l LinkedQueue) setup() {
	l.front = nil
	l.rear = nil
}

func (l LinkedQueue) clear() {
	l.setup()
}

func (l LinkedQueue) enqueue(item Elem) {
	if l.isEmpty() {
		l.rear = NewBaseLinkNode(item, l.rear, nil)
		l.front = l.rear
	} else {
		l.rear.SetNext(NewBaseLinkNode(item, l.rear, nil))
		l.rear = l.rear.Next()
	}
}

func (l LinkedQueue) dequeue() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	value := l.front.Element()
	l.front = l.front.Next()
	if l.front == nil {
		l.rear = nil
	} else {
		l.front.SetPrev(nil)
	}
	return value, nil
}

func (l LinkedQueue) firstValue() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	return l.front.Element(), nil
}

func (l LinkedQueue) isEmpty() bool {
	return l.front == l.rear
}
