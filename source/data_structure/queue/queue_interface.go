package queue

import (
	. "../../data_structure"
)

type Queue interface {
	setup(size int)
	clear()
	enqueue(item Elem)
	dequeue() Elem
	firstValue() Elem
	isEmpty() bool
}