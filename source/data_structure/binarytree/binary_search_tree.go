package binarytree

import "fmt"

type BSTimpl struct {
	root BinNode
}

func (b BSTimpl) clear() {
	b.root = nil
}

func (b BSTimpl) Insert(value Elem) {
	// b.root = InsertHelp(b.root, value)
	var y BinNode
	x := b.root
	z := BinNodePtr{value, nil, nil, nil}
	for x != nil {
		y = x
		if z.Element().key() < x.Element().key() {
			x = x.Left()
		} else {
			x = x.Right()
		}
	}
	z.SetParent(y)
	if y == nil {
		b.root = z
	} else if z.Element().key() < y.Element().key() {
		y.SetLeft(z)
	} else {
		y.SetRight(z)
	}
}

func (b BSTimpl) Delete(key int) {
	// b.root = DeleteHelp(b.root, key)
	node := SearchHelp(b.root, key)
	
	if node.Left() == nil {
		b.transplant(node, node.Right())
	} else if node.Right() == nil {
		b.transplant(node, node.Left())
	} else {
		helper := MinimumHelp(node)
		if helper.Parent() != node {
			b.transplant(helper, helper.Right())
			helper.SetRight(node.Right())
			helper.Right().SetParent(helper)
		}
		b.transplant(node, helper)
		helper.SetLeft(node.Left())
		helper.Left().SetParent(helper)
	}
}

func (b BSTimpl) Search(key int) BST {
	return BSTimpl{SearchHelp(b.root, key)}
}

func SearchHelp(root BinNode, key int) BinNode {
	node := root
	for node != nil && key != node.Element().key() {
		if key < node.Element().key() {
			node = node.Left()
		} else {
			node = node.Right()
		}
	}
	return node
}

func (b BSTimpl) transplant(u, v BinNode) {
	if u.Parent() == nil {
		b.root = v
	} else if u == u.Parent().Left() {
		u.Parent().SetLeft(v)
	} else {
		u.Parent().SetRight(v)
	}
	
	if v != nil {
		v.SetParent(u.Parent())
	}
}

func (b BSTimpl) Successor() BST {
	if b.root.Right() != nil {
		return BSTimpl{b.root.Right()}.Minimum()
	}
	helper := b.root
	successor := b.root.Parent()
	for successor != nil && helper == successor.Right() {
		helper = successor
		successor = successor.Parent()
	}
	return BSTimpl{successor}
}

func (b BSTimpl) Predecessor() BST {
	if b.root.Left() != nil {
		return BSTimpl{b.root.Left()}.Maximum()
	}
	helper := b.root
	predecessor := b.root.Parent()
	for predecessor != nil && helper == predecessor.Left() {
		helper = predecessor
		predecessor = predecessor.Parent()
	}
	return BSTimpl{predecessor}
}

func InorderWalkHelp(root BinNode) {
	if root != nil {
		InorderWalkHelp(root.Left())
		fmt.Printf("%d ", root.Element())
		InorderWalkHelp(root.Right())
	}
}

func (b BSTimpl) isEmpty() bool {
	return b.root == nil
}

func DeleteHelp(root BinNode, key int) BinNode {
	return nil
}

func (b BSTimpl) Minimum() BST {
	return BSTimpl{MinimumHelp(b.root)}
}

func MinimumHelp(root BinNode) BinNode {
	minTree := root
	for root.Left() != nil {
		minTree = root.Left()
	}
	return minTree
}

func (b BSTimpl) Maximum() BST {
	return BSTimpl{MaximumHelp(b.root)}
}

func MaximumHelp(root BinNode) BinNode {
	maxTree := root
	for root.Right() != nil {
		maxTree = root.Right()
	}
	return maxTree
}

func (b BSTimpl) InorderWalk() {
	InorderWalkHelp(b.root)
}
