package stack

import (
	. "github.com/Thrimbda/dune/datastructure"
	. "github.com/Thrimbda/dune/datastructure/arrayutils"
	. "github.com/Thrimbda/dune/datastructure/linkutils"
)

type LinkedStack struct {
	top LinkNode
}

func (l LinkedStack) setup() {
	l.top = nil
}

func (l LinkedStack) clear() {
	l.top = nil
}

func (l LinkedStack) push(value Elem) {
	l.top = NewBaseLinkNode(value, l.top, nil)
}

func (l LinkedStack) pop() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	value := l.top.Element()
	l.top = l.top.Prev()
	l.top.SetNext(nil)
	return value, nil
}

func (l LinkedStack) topValue() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	return l.top.Element(), nil
}

func (l LinkedStack) isEmpty() bool {
	return l.top == nil
}
