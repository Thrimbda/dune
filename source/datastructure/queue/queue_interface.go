package queue

import (
	. "github.com/Trimbda/dune/datastructure"
)

type Queue interface {
	setup(size int)
	clear()
	enqueue(item Elem)
	dequeue() Elem
	firstValue() Elem
	isEmpty() bool
}
