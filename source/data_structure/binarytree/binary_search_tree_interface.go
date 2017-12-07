package binarytree

type BST interface {
	insert(value Elem)
	find(key int) Elem
	remove(key int) Elem
}