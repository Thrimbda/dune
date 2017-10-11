package binarytree

type BinNode interface {
	//interface of BinNode.
	Element() Elem
	SetElement(element Elem)
	Left() BinNode
	SetLeft(node BinNode)
	Right() BinNode
	SetRight(node BinNode)
	IsLeaf() bool
}

type Elem interface {
	key() int
}
