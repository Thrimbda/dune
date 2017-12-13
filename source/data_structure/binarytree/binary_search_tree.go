package binarytree

type BSTimpl struct {
	root BinNode
}

func (b BSTimpl) clear() {
	b.root = nil
}

func (b BSTimpl) insert(value Elem) {
	b.root = insertHelp(b.root, value)
}

func (b BSTimpl) remove(key int) {
	b.root = removeHelp(b.root, key)
}

func (b BSTimpl) find(key int) Elem {
	return findHelp(b.root, key)
}

func (b BSTimpl) isEmpty() bool {
	return b.root == nil
}

func findHelp(root BinNode, key int) Elem {
	if root == nil {
		return nil
	}
	item := root.Element()
	if item.key() > key {
		return findHelp(root.Left(), key)
	} else if item.key() == key {
		return item
	}
	return findHelp(root.Right(), key)
}

func insertHelp(root BinNode, value Elem) BinNode {
	if root == nil {
		return BinNodePtr{value, nil, nil, nil}
	}
	item := root.Element()
	if item.key() > value.key() {
		root.SetLeft(insertHelp(root.Left(), value))
	} else {
		root.SetRight(insertHelp(root.Right(), value))
	}
	return root
}

func removeHelp(root BinNode, key int) BinNode {
	if root == nil {
		return nil
	}
	item := root.Element()
	if key > item.key() {
		root.SetRight(removeHelp(root.Right(), key))
	} else if key < item.key() {
		root.SetLeft(removeHelp(root.Left(), key))
	} else {
		if root.Left() == nil {
			root = root.Right()
		} else if root.Right() == nil {
			root = root.Left()
		} else {
			temp := getMin(root.Right())
			root.SetRight(deleteMin(root.Right()))
			root.SetElement(temp)
		}
	}
	return root
}

func getMin(root BinNode) Elem {
	if root.Left() == nil {
		return root.Element()
	}
	return getMin(root.Left())
}

func deleteMin(root BinNode) BinNode {
	if root.Left() == nil {
		return root.Right()
	}
	root.SetLeft(deleteMin(root.Left()))
	return root
}
