package binarytree

import (
	"fmt"
	. "../../datastructure"
)

type BSTimpl struct {
	root BinNode
}

func (b BSTimpl) Insert(value Elem) {
	var father BinNode
	brother := b.root
	node := &BinNodePtr{value, nil, nil, nil}
	for brother != nil {
		father = brother
		if node.Element().Key() < brother.Element().Key() {
			brother = brother.Left()
		} else {
			brother = brother.Right()
		}
	}
	node.SetParent(father)
	if father == nil {
		b.root = node
	} else if node.Element().Key() < father.Element().Key() {
		father.SetLeft(node)
	} else {
		father.SetRight(node)
	}
}

func (b BSTimpl) Delete(Key int) {
	node := SearchHelp(b.root, Key)
	
	if node.Left() == nil {
		b.transplant(node, node.Right())
	} else if node.Right() == nil {
		b.transplant(node, node.Left())
	} else {
		replacement := MinimumHelp(node)
		if replacement.Parent() != node {
			b.transplant(replacement, replacement.Right())
			replacement.SetRight(node.Right())
			replacement.Right().SetParent(replacement)
		}
		b.transplant(node, replacement)
		replacement.SetLeft(node.Left())
		replacement.Left().SetParent(replacement)
	}
}

func (b BSTimpl) Search(key int) BST {
	return &BSTimpl{SearchHelp(b.root, key)}
}

func SearchHelp(root BinNode, key int) BinNode {
	node := root
	for node != nil && key != node.Element().Key() {
		if key < node.Element().Key() {
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
	return &BSTimpl{SuccessorHelp(b.root)}
}

func SuccessorHelp(root BinNode) BinNode {
	if root.Right() != nil {
		return MinimumHelp(root.Right())
	}
	helper := root
	successor := root.Parent()
	for successor != nil && helper == successor.Right() {
		helper = successor
		successor = successor.Parent()
	}
	return successor
}

func (b BSTimpl) Predecessor() BST {
	return &BSTimpl{PredecessorHelp(b.root)}
}

func PredecessorHelp(root BinNode) BinNode {
	if root.Left() != nil {
		return MaximumHelp(root.Left())
	}
	helper := root
	predecessor := root.Parent()
	for predecessor != nil && helper == predecessor.Left() {
		helper = predecessor
		predecessor = predecessor.Parent()
	}
	return predecessor
}

func (b BSTimpl) isEmpty() bool {
	return b.root == nil
}

func DeleteHelp(root BinNode, Key int) BinNode {
	return nil
}

func (b BSTimpl) Minimum() BST {
	return &BSTimpl{MinimumHelp(b.root)}
}

func MinimumHelp(root BinNode) BinNode {
	minTree := root
	for root.Left() != nil {
		minTree = root.Left()
	}
	return minTree
}

func (b BSTimpl) Maximum() BST {
	return &BSTimpl{MaximumHelp(b.root)}
}

func MaximumHelp(root BinNode) BinNode {
	maxTree := root
	for root.Right() != nil {
		maxTree = root.Right()
	}
	return maxTree
}

func (b BSTimpl) InorderWalk() {
	inorderWalkHelp(b.root)
}

func inorderWalkHelp(root BinNode) {
	if root != nil {
		inorderWalkHelp(root.Left())
		fmt.Printf("%s ", root.Element())
		inorderWalkHelp(root.Right())
	}
}
