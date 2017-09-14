package data_structure

type List interface {
	clear()
	insert()
	append()
	remove()
	setFirst()
	next()
	prev()
	length()
	setPos(pos int)
	setValue(value Elem)
	currValue() Elem
	isEmpty() bool
	isInList() bool
	print()
}

type Elem interface {
	key() int
}