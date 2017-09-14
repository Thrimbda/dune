package data_structure

type LinkedStack struct {
	top *LinkNode
}

func (l LinkedStack) setup() {
	l.top = nil
}

func (l LinkedStack) clear() {
	l.top = nil
}

func (l LinkedStack) push(value interface{}) {
	l.top = &LinkNode{value, l.top, nil}
}

func (l LinkedStack) pop() (interface{}, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	value := l.top.value
	l.top = l.top.prev
	l.top.next = nil
	return value, nil
}

func (l LinkedStack) topValue() (interface{}, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	return l.top.value, nil
}

func (l LinkedStack) isEmpty() bool {
	return l.top == nil
}