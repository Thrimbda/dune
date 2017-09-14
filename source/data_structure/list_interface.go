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
	setValue(value interface{})
	currValue() interface{}
	isEmpty() bool
	isInList() bool
	print()
}