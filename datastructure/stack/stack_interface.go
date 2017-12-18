package stack

import (
	. "github.com/Thrimbda/dune/datastructure"
)

type Stack interface {
	Setup(size int)
	Clear()
	Push(item Elem)
	Pop() Elem
	TopValue() Elem
	IsEmpty() bool
}
