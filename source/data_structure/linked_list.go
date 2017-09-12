package data_structure

import (
	"fmt"
)

type LinkedList struct {
	numInList int
	curr *LinkNode
	head *LinkNode
	tail *LinkNode
}

type LinkNode struct {
	value interface{}
	prev *LinkNode
	next *LinkNode
}

type NullCurrError struct{}

func (e NullCurrError) Error() string {
	return "curr is nil!"
}

func (l LinkedList) setup() {
	l.head = &LinkNode{nil, nil, nil}
	l.tail = l.head
	l.curr = l.head
}

func (l LinkedList) clear() {
	l.head.next = nil
	l.tail = l.head
	l.curr = l.head
}

func (l LinkedList) insert(value interface{}) error {
	if l.curr == nil {
		return NullCurrError{}
	}
	l.curr.next = &LinkNode{value, l.curr, l.curr.next}
	if l.curr.next != l.tail {
		l.curr.next.next.prev = l.curr.next
	}
	if l.tail == l.curr {
		l.tail = l.curr.next
	}
	return nil
}

func (l LinkedList) append(value interface{}) {
	l.tail.next = &LinkNode{value, l.tail, nil}
	l.tail = l.tail.next
}
// TODO
func (l LinkedList) remove() (interface{}, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	if l.isInList() {
		return nil, NullCurrError{}
	}
	value := l.curr.value
	l.curr.next = l.curr.next.next
	if l.curr.next == l.tail {
		l.tail = l.curr
	}
	return value, nil
}

func (l LinkedList) setFirst() {
	l.curr = l.head
}

func (l LinkedList) next() {
	if l.curr != nil {
		l.curr = l.curr.next
	}
}

func (l LinkedList) prev() {
	if l.curr != nil {
		l.curr = l.curr.prev
	}
}

func (l LinkedList) length() int {
	cnt := 0
	for counter := l.head.next; counter != nil; counter = counter.next {
		cnt++
	}
	return cnt
}

func (l LinkedList) setPos(pos int) {
	l.setFirst()
	for i := 0; i < pos; i++ {
		l.next()
	}
}

func (l LinkedList) setValue(value interface{}) error {
	if l.isInList() {
		l.curr.next.value = value
		return nil
	} else {
		return NullCurrError{}
	}
}

func (l LinkedList) currValue() (interface{}, error) {
	if l.isInList() {
		return l.curr.next.value, nil
	} else {
		return nil, NullCurrError{}
	}
}

func (l LinkedList) isEmpty() bool {
	return l.head.next == nil
}

func (l LinkedList) isInList() bool {
	return l.curr != nil && l.curr.next != nil
}

func (l LinkedList) print() {
	if l.isEmpty() {
		fmt.Println("()")
	} else {
		fmt.Print("(")
		for item := l.head.next; item != nil; item = item.next {
			value, _ := l.currValue()
			if str, ok := value.(string); ok {
				fmt.Print(string(str) + " ")
			}
		}
	}
}