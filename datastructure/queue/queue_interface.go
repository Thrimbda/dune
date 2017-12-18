package queue

import (
	. "github.com/Thrimbda/dune/datastructure"
)

type Queue interface {
	setup(size int)
	clear()
	enqueue(item Elem)
	dequeue() Elem
	firstValue() Elem
	isEmpty() bool
}
