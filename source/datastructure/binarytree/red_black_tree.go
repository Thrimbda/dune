package binarytree

import (
	"github.com/Trimbda/dune/datastructure/utils"
)

type RBNode struct {
	BinNodePtr
	color bool
}

const (
	black = true
	red   = false
)

// black for black, red for red

func (node RBNode) SetColor(color bool) {
	node.color = color
}

func (node RBNode) Color() bool {
	return node.color
}

type RBTree struct {
	root *RBNode
}

var RBNil = &RBNode{BinNodePtr{nil, nil, nil, nil}, black}

//should I set the nil as an attribute of a RB-Tree?

func (rbt RBTree) Insert(value interface{}) {
	father := RBNil
	brother := rbt.root
	node := &RBNode{BinNodePtr{value, RBNil, RBNil, nil}, red}
	for brother != RBNil {
		father = brother
		if utils.LessComparator(node.Element(), brother.Element()) {
			brother = brother.Left().(*RBNode)
		} else {
			brother = brother.Right().(*RBNode)
		}
	}
	node.SetParent(father)
	if father == RBNil {
		rbt.root = node
	} else if utils.LessComparator(node.Element(), father.Element()) {
		father.SetLeft(node)
	} else {
		father.SetRight(node)
	}
	rbt.RBInsertFixUp(node)
}

func (rbt RBTree) RBInsertFixUp(node *RBNode) {
	var uncle *RBNode
	for !node.Parent().(*RBNode).Color() {
		if node.Parent() == node.Parent().Parent().Left() {
			uncle = node.Parent().Parent().Right().(*RBNode)
			if !uncle.Color() {
				node.Parent().(*RBNode).SetColor(black)
				uncle.SetColor(black)
				node.Parent().Parent().(*RBNode).SetColor(red)
				node = node.Parent().Parent().(*RBNode)
			} else if node == node.Parent().Right().(*RBNode) {
				node = node.Parent().(*RBNode)
				rbt.leftRotate(node)
			}
			node.Parent().(*RBNode).SetColor(black)
			node.Parent().Parent().(*RBNode).SetColor(red)
			rbt.rightRotate(node.Parent().Parent().(*RBNode))
		} else {
			uncle = node.Parent().Parent().Left().(*RBNode)
			if !uncle.Color() {
				node.Parent().(*RBNode).SetColor(black)
				uncle.SetColor(black)
				node.Parent().Parent().(*RBNode).SetColor(red)
				node = node.Parent().Parent().(*RBNode)
			} else if node == node.Parent().Left().(*RBNode) {
				node = node.Parent().(*RBNode)
				rbt.rightRotate(node)
			}
			node.Parent().(*RBNode).SetColor(black)
			node.Parent().Parent().(*RBNode).SetColor(red)
			rbt.leftRotate(node.Parent().Parent().(*RBNode))
		}
		rbt.root.SetColor(black)
	}
}

func (rbt RBTree) Search(key int) BST {
	return &RBTree{SearchHelp(rbt.root, key).(*RBNode)}
}

func (rbt RBTree) Delete(key int) {
	node := SearchHelp(rbt.root, key).(*RBNode)
	var tracker *RBNode
	replacement := node
	originColor := replacement.Color()

	if node.Left() == RBNil {
		tracker = node.Right().(*RBNode)
		rbt.RBTransplant(node, tracker)
	} else if node.Right() == RBNil {
		tracker = node.Left().(*RBNode)
		rbt.RBTransplant(node, tracker)
	} else {
		replacement = MinimumHelp(node.Right()).(*RBNode)
		originColor = replacement.Color()
		tracker = replacement.Right().(*RBNode)
		if replacement.Parent() != node {
			rbt.RBTransplant(replacement, replacement.Right().(*RBNode))
			replacement.SetRight(node.Right())
			replacement.Right().SetParent(node.Parent())
		}
		rbt.RBTransplant(node, replacement)
		replacement.SetLeft(node.Left())
		replacement.Left().SetParent(replacement)
		replacement.SetColor(node.Color())
		if originColor {
			rbt.RBDeleteFixUp(tracker)
		}
	}
}

func (rbt RBTree) RBDeleteFixUp(node *RBNode) {
	var brother *RBNode
	for node == rbt.root && node.Color() {
		if node != node.Parent().Left() {
			brother = node.Parent().Right().(*RBNode)
			if !brother.Color() {
				brother.SetColor(black)
				node.Parent().(*RBNode).SetColor(red)
				rbt.leftRotate(node.Parent().(*RBNode))
				brother = node.Parent().Right().(*RBNode)
			}
			if brother.Left().(*RBNode).Color() && brother.Right().(*RBNode).Color() {
				brother.SetColor(red)
				node = node.Parent().(*RBNode)
			} else if !brother.Right().(*RBNode).Color() {
				brother.Right().(*RBNode).SetColor(black)
				brother.SetColor(red)
				rbt.rightRotate(brother)
				brother = node.Parent().Right().(*RBNode)
			}
			brother.SetColor(node.Parent().(*RBNode).Color())
			node.Parent().(*RBNode).SetColor(black)
			brother.Right().(*RBNode).SetColor(black)
			rbt.leftRotate(node.Parent().(*RBNode))
			node = rbt.root
		} else {
			brother = node.Parent().Left().(*RBNode)
			if !brother.Color() {
				brother.SetColor(black)
				node.Parent().(*RBNode).SetColor(red)
				rbt.rightRotate(node.Parent().(*RBNode))
				brother = node.Parent().Left().(*RBNode)
			}
			if brother.Right().(*RBNode).Color() && brother.Left().(*RBNode).Color() {
				brother.SetColor(red)
				node = node.Parent().(*RBNode)
			} else if !brother.Left().(*RBNode).Color() {
				brother.Left().(*RBNode).SetColor(black)
				brother.SetColor(red)
				rbt.leftRotate(brother)
				brother = node.Parent().Left().(*RBNode)
			}
			brother.SetColor(node.Parent().(*RBNode).Color())
			node.Parent().(*RBNode).SetColor(black)
			brother.Left().(*RBNode).SetColor(black)
			rbt.rightRotate(node.Parent().(*RBNode))
			node = rbt.root
		}
	}
	node.SetColor(black)
}

func (rbt RBTree) RBTransplant(u, v *RBNode) {
	if u.Parent() == RBNil {
		rbt.root = v
	} else if u == u.Parent().Left() {
		u.Parent().SetLeft(v)
	} else {
		u.Parent().SetRight(v)
	}
	v.SetParent(u.Parent())
}

func (rbt RBTree) Predecessor() BST {
	return &RBTree{PredecessorHelp(rbt.root).(*RBNode)}
}

func (rbt RBTree) Successor() BST {
	return &RBTree{SuccessorHelp(rbt.root).(*RBNode)}
}

func (rbt RBTree) Minimum() BST {
	return &RBTree{MinimumHelp(rbt.root).(*RBNode)}
}

func (rbt RBTree) Maximum() BST {
	return &RBTree{MaximumHelp(rbt.root).(*RBNode)}
}

func (rbt RBTree) InorderWalk() {
	inorderWalkHelp(rbt.root)
}

func (rbt RBTree) leftRotate(node *RBNode) {
	rChild := node.Right().(*RBNode)
	node.SetRight(rChild.Left())
	if rChild.Left() != RBNil {
		rChild.Left().SetRight(node)
	}
	rChild.SetParent(node.Parent())
	if node.Parent() == RBNil {
		rbt.root = rChild
	} else if node == node.Parent().Left() {
		node.Parent().SetLeft(rChild)
	} else {
		node.Parent().SetRight(rChild)
	}
	rChild.SetLeft(node)
	node.SetParent(rChild)
}

func (rbt RBTree) rightRotate(node *RBNode) {
	lChild := node.Right().(*RBNode)
	node.SetLeft(lChild.Right())
	if lChild.Right() != RBNil {
		lChild.Right().SetParent(node)
	}
	lChild.SetParent(node.Parent())
	if node.Parent() == RBNil {
		rbt.root = lChild
	} else if node == node.Parent().Left() {
		node.Parent().SetLeft(lChild)
	} else {
		node.Parent().SetRight(lChild)
	}
	lChild.SetRight(node)
	node.SetParent(lChild)
}
