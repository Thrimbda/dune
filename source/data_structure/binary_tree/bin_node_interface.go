package binary_tree

type BinNode interface {
	Element() Elem
	SetElement(element Elem)
	Left() BinNode
	SetLeft(node BinNode)
	Right() BinNode
	SetRight(node BinNode)
	isLeaf() bool
}

type Elem interface {
	key() int
}