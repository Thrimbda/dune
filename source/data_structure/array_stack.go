package data_structure

import (
	// "fmt"
)

type ArrayStack struct {
	size int
	top int
	listArray []interface{}
}

func (a ArrayStack) setup(size int) {
	a.size = size
	a.top = 0
	a.listArray = make([] interface{}, size)
}

func (a ArrayStack) clear() {
	a.top = 0
}

func (a ArrayStack) push(item interface{}) error {
	if a.top >= a.size {
		return FullListError{}
	}
	a.listArray[a.top] = item
	a.top++
	return nil
}

func (a ArrayStack) pop() (interface{}, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	a.top--
	value := a.listArray[a.top]
	return value, nil
}

func (a ArrayStack) topValue() (interface{}, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	return a.listArray[a.top - 1], nil
}

func (a ArrayStack) isEmpty() bool {
	return a.top == 0
}