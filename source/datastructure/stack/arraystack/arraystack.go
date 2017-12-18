package arraystack

import (
	. "../../../datastructure"
	. "../../arrayutils"
)

//stack should be a higher layer of list
type arrayStack struct {
	size int
	top int
	listArray []Elem
}

func (a arrayStack) Setup(size int) {
	a.size = size
	a.top = 0
	a.listArray = make([] Elem, size)
}

func (a arrayStack) Clear() {
	a.top = 0
}

func (a arrayStack) Push(item Elem) error {
	if a.top >= a.size {
		return FullListError{}
	}
	a.listArray[a.top] = item
	a.top++
	return nil
}

func (a arrayStack) Pop() (Elem, error) {
	if a.IsEmpty() {
		return nil, EmptyListError{}
	}
	a.top--
	value := a.listArray[a.top]
	return value, nil
}

func (a arrayStack) TopValue() (Elem, error) {
	if a.IsEmpty() {
		return nil, EmptyListError{}
	}
	return a.listArray[a.top - 1], nil
}

func (a arrayStack) IsEmpty() bool {
	return a.top == 0
}