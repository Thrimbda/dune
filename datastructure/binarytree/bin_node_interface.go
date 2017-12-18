package binarytree

type BinNode interface {
	//interface of BinNode.
	Element() interface{}
	SetElement(element interface{})
	Left() BinNode
	SetLeft(node BinNode)
	Right() BinNode
	SetRight(node BinNode)
	SetParent(node BinNode)
	Parent() BinNode
	IsLeaf() bool
}

// It seems like there is no need to create a BinNode interface.
// And I cannot figure out how to inherit interface in Golang and the point of it.
// But anyway, for congruity, I've done it in a folly way.
