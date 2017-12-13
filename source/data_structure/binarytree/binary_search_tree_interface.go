package binarytree

type BST interface {
	Insert(value Elem)
	Search(key int) Elem
	Delete(key int) Elem
	Predecessor() BST
	Successor() BST
	Minimum() BST
	Maximum() BST
}