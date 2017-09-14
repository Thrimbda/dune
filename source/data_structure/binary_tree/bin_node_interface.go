package binary_tree

type BinNode interface {
	element() Elem
	SetElement(element Elem)
	Left() BinNode
	SetLeft(node BinNode)
	SetRight(node BinNode)
	Right() BinNode
	isLeaf() bool
}

type Elem interface {
	key() int
}