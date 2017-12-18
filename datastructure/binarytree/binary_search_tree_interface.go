package binarytree

type BST interface {
	Insert(value interface{})
	Search(key int) BST
	Delete(key int)
	Predecessor() BST
	Successor() BST
	Minimum() BST
	Maximum() BST
	InorderWalk()
}
