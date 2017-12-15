package list

import (
	. "../../datastructure"
)

type List interface {
	clear()
	insert(item Elem)
	append(item Elem)
	remove() Elem
	setFirst()
	next()
	prev()
	length() int
	setPos(pos int)
	setValue(value Elem)
	currValue() Elem
	isEmpty() bool
	isInList() bool
	print()
}