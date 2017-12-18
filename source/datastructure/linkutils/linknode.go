package linkutils

import (
	. "github.com/Trimbda/dune/datastructure"
)

type LinkNode interface {
	Element() Elem
	SetElement(value Elem)
	Prev() LinkNode
	SetPrev(node LinkNode)
	Next() LinkNode
	SetNext(node LinkNode)
}

type linkNodeImpl struct {
	value Elem
	prev  LinkNode
	next  LinkNode
}

func (node linkNodeImpl) Element() Elem {
	return node.value
}

func (node linkNodeImpl) SetElement(value Elem) {
	node.value = value
}

func (node linkNodeImpl) Prev() LinkNode {
	return node.prev
}

func (node linkNodeImpl) SetPrev(prev LinkNode) {
	node.prev = prev
}

func (node linkNodeImpl) Next() LinkNode {
	return node.next
}

func (node linkNodeImpl) SetNext(next LinkNode) {
	node.next = next
}

func NewBaseLinkNode(value Elem, prev, next LinkNode) LinkNode {
	return &linkNodeImpl{value, prev, next}
}
