package binarytree

type RBNode interface {
	BinNode
	SetParent(node BinNode)
	Parent() BinNode
	SetColor(red bool)
	Color() bool
}

type RBTree struct {
	root RBNode
}

