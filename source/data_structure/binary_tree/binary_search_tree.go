package binary_tree

type BST struct {
	root BinNode
}

func (b BST) clear() {
	b.root = nil
}

func (b BST) insert(value Elem) {
	b.root = b.insertHelp(b.root, value)
}

func (b BST) remove(key int) {
	b.root = b.removeHelp(b.root, key)
}

func (b BST) find(key int) Elem {
	return b.findHelp(b.root, key)
}

func (b BST) isEmpty() bool {
	return b.root == nil
}

func (b BST) findHelp(root BinNode, key int) Elem {
	if root == nil {
		return nil
	}
	item := root.Element()
	if item.key() > key {
		return b.findHelp(root.Left(), key)
	} else if item.key() == key {
		return item
	} else {
		return b.findHelp(root.Right(), key)
	}
}

func (b BST) insertHelp(root BinNode, value Elem) BinNode {
	if root == nil {
		return BinNodePtr{value, nil, nil}
	}
	item := root.Element()
	if item.key() > value.key() {
		root.SetLeft(b.insertHelp(root.Left(), value))
	} else {
		root.SetRight(b.insertHelp(root.Right(), value))
	}
	return root
}

func (b BST) removeHelp(root BinNode, key int) BinNode {
	//TODO
	return nil
}

func (b BST) getMin(root BinNode) Elem {
	if root.Left() == nil {
		return root.Element()
	} else {
		return b.getMin(root.Left())
	}
}

func (b BST) deleteMin(root BinNode) BinNode {
	if root.Left() == nil {
		return root.Right()
	} else {
		root.SetLeft(b.deleteMin(root.Left()))
		return root
	}
}