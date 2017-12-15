package stack

import (
	. "../../datastructure"
)

type Stack interface {
	setup(size int)
	clear()
	push(item Elem)
	pop() Elem
	topValue() Elem
	isEmpty() bool
}