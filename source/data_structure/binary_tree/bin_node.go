package binary_tree

type BinNodePtr struct {
	element Elem
	left *BinNode
	right *BinNode
}

func (b BinNodePtr) Element() Elem {
	return b.element
}

func (b BinNodePtr) SetElement(element Elem) {
	b.element = element
}

func (b BinNodePtr) Left() *BinNode {
	return b.left
}

func (b BinNodePtr) SetLeft(node *BinNode) {
	b.left = node
}

func (b BinNodePtr) Right() *BinNode {
	return b.right
}

func (b BinNodePtr) SetRight(node *BinNode) {
	b.right = node
}

func (b BinNodePtr) isLeaf() bool {
	return b.left == nil && b.right == nil
}
