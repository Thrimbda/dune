package linkutils

//LinkNode is an interface for
type LinkNode interface {
	Element() interface{}
	SetElement(value interface{})
	Prev() LinkNode
	SetPrev(node LinkNode)
	Next() LinkNode
	SetNext(node LinkNode)
}

type linkNodeImpl struct {
	value interface{}
	prev  LinkNode
	next  LinkNode
}

func (node linkNodeImpl) Element() interface{} {
	return node.value
}

func (node *linkNodeImpl) SetElement(value interface{}) {
	node.value = value
}

func (node *linkNodeImpl) Prev() LinkNode {
	return node.prev
}

func (node *linkNodeImpl) SetPrev(prev LinkNode) {
	node.prev = prev
}

func (node *linkNodeImpl) Next() LinkNode {
	return node.next
}

func (node *linkNodeImpl) SetNext(next LinkNode) {
	node.next = next
}

func NewDoubleLinkNode(value interface{}, prev, next LinkNode) LinkNode {
	return &linkNodeImpl{value, prev, next}
}
