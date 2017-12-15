package binarytree

import (
	. "../../datastructure"
)

type BST interface {
	Insert(value Elem)
	Search(key int) BST
	Delete(key int)
	Predecessor() BST
	Successor() BST
	Minimum() BST
	Maximum() BST
	InorderWalk()
}