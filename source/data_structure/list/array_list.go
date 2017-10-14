package list

import (
	"fmt"
)

type ArrayList struct {
	maxSize   int
	numInList int
	curr      int
	listArray []Elem
}

type FullListError struct{}

func (e FullListError) Error() string {
	return "list is full!"
}

type BadCurrError struct{}

func (e BadCurrError) Error() string {
	return "bad current pointer position."
}

type EmptyListError struct{}

func (e EmptyListError) Error() string {
	return "list is empty!"
}

func (a ArrayList) setup(size int) {
	a.maxSize = size
	a.numInList = 0
	a.curr = 0
	a.listArray = make([]Elem, size)
}

func (a ArrayList) clear() {
	a.numInList = 0
	a.curr = 0
}

func (a ArrayList) insert(item Elem) error {
	// for each function of Array List, we must judge if list is full,
	// I wonder if there is a way to bind this action to every function.
	if a.numInList >= a.maxSize {
		return FullListError{}
	}
	if !a.isInList() {
		return BadCurrError{}
	}
	for i := a.numInList; i > a.curr; i-- {
		a.listArray[i] = a.listArray[i-1]
	}
	a.listArray[a.curr] = item
	a.numInList++
	return nil
}

func (a ArrayList) append(item Elem) error {
	if a.numInList >= a.maxSize {
		return FullListError{}
	}
	a.listArray[a.numInList] = item
	a.numInList++
	return nil
}

func (a ArrayList) remove() (Elem, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	value := a.listArray[a.curr]
	for i := a.curr; i < a.numInList-1; i++ {
		a.listArray[i] = a.listArray[i+1]
	}
	a.numInList--
	return value, nil
}

func (a ArrayList) setFirst() {
	a.curr = 0
}

func (a ArrayList) next() {
	a.curr++
}

func (a ArrayList) prev() {
	a.curr--
}

func (a ArrayList) length() int {
	return a.numInList
}

func (a ArrayList) setPos(pos int) {
	a.curr = pos
}

func (a ArrayList) setValue(value Elem) error {
	if !a.isInList() {
		return BadCurrError{}
	}
	a.listArray[a.curr] = value
	return nil
}

func (a ArrayList) currValue() (Elem, error) {
	if a.curr < 0 || a.curr > a.numInList {
		return nil, BadCurrError{}
	}
	return a.listArray[a.curr], nil
}

func (a ArrayList) isEmpty() bool {
	return a.numInList == 0
}

func (a ArrayList) isInList() bool {
	return a.curr > 0 && a.curr < a.numInList
}

func (a ArrayList) print() {
	if a.isEmpty() {
		fmt.Println("()")
	} else {
		fmt.Print("(")
		for a.setFirst(); a.isInList(); a.next() {
			value, _ := a.currValue()
			if str, ok := value.(string); ok {
				fmt.Print(string(str) + " ")
			}
		}
		fmt.Println(")")
	}
}
