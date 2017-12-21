package arraystack

import (
	. "github.com/Thrimbda/dune/datastructure/arrayutils"
)

//stack should be a higher layer of list
type arrayStack struct {
	size      int
	top       int
	listArray []interface{}
}

func (a arrayStack) Setup(size int) {
	a.size = size
	a.top = 0
	a.listArray = make([]interface{}, size)
}

func (a arrayStack) Clear() {
	a.top = 0
}

func (a arrayStack) Push(item interface{}) error {
	if a.top >= a.size {
		return FullListError{}
	}
	a.listArray[a.top] = item
	a.top++
	return nil
}

func (a arrayStack) Pop() (interface{}, error) {
	if a.IsEmpty() {
		return nil, EmptyListError{}
	}
	a.top--
	value := a.listArray[a.top]
	return value, nil
}

func (a arrayStack) TopValue() (interface{}, error) {
	if a.IsEmpty() {
		return nil, EmptyListError{}
	}
	return a.listArray[a.top-1], nil
}

func (a arrayStack) IsEmpty() bool {
	return a.top == 0
}
