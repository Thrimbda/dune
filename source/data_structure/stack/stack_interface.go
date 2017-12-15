package stack

import (
	. "../../data_structure"
)

type Stack interface {
	setup(size int)
	clear()
	push(item Elem)
	pop() Elem
	topValue() Elem
	isEmpty() bool
}