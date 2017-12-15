package list

import (
	"fmt"
	. "../../datastructure"
	. "../linkutils"
	. "../arrayutils"
)

type LinkedList struct {
	numInList int
	curr LinkNode
	head LinkNode
	tail LinkNode
}

func (l LinkedList) setup() {
	l.head = NewBaseLinkNode(nil, nil, nil)
	l.tail = l.head
	l.curr = l.head
}

func (l LinkedList) clear() {
	l.head.SetNext(nil)
	l.tail = l.head
	l.curr = l.head
}

func (l LinkedList) insert(value Elem) error {
	if l.curr == nil {
		return NullCurrError{}
	}
	l.curr.SetNext(NewBaseLinkNode(value, l.curr, l.curr.Next()))
	if l.curr.Next() != l.tail {
		l.curr.Next().Next().SetPrev(l.curr.Next())
	}
	if l.tail == l.curr {
		l.tail = l.curr.Next()
	}
	return nil
}

func (l LinkedList) append(value Elem) {
	l.tail.SetNext(NewBaseLinkNode(value, l.tail, nil))
	l.tail = l.tail.Next()
}
// TODO
func (l LinkedList) remove() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	if l.isInList() {
		return nil, NullCurrError{}
	}
	value := l.curr.Element()
	l.curr.SetNext(l.curr.Next().Next())
	if l.curr.Next() == l.tail {
		l.tail = l.curr
	}
	return value, nil
}

func (l LinkedList) setFirst() {
	l.curr = l.head
}

func (l LinkedList) next() {
	if l.curr != nil {
		l.curr = l.curr.Next()
	}
}

func (l LinkedList) prev() {
	if l.curr != nil {
		l.curr = l.curr.Prev()
	}
}

func (l LinkedList) length() int {
	cnt := 0
	for counter := l.head.Next(); counter != nil; counter = counter.Next() {
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

func (l LinkedList) setValue(value Elem) error {
	if l.isInList() {
		l.curr.Next().SetElement(value)
		return nil
	} else {
		return NullCurrError{}
	}
}

func (l LinkedList) currValue() (Elem, error) {
	if l.isInList() {
		return l.curr.Next().Element(), nil
	} else {
		return nil, NullCurrError{}
	}
}

func (l LinkedList) isEmpty() bool {
	return l.head.Next() == nil
}

func (l LinkedList) isInList() bool {
	return l.curr != nil && l.curr.Next() != nil
}

func (l LinkedList) print() {
	if l.isEmpty() {
		fmt.Println("()")
	} else {
		fmt.Print("(")
		for item := l.head.Next(); item != nil; item = item.Next() {
			value, _ := l.currValue()
			fmt.Printf("%s ", value)
		}
	}
}