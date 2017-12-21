package list

import (
	"bytes"
	"fmt"
	"reflect"

	. "github.com/Thrimbda/dune/datastructure/arrayutils"
	. "github.com/Thrimbda/dune/datastructure/linkutils"
)

type LinkedList struct {
	numInList int
	curr      LinkNode
	head      LinkNode
	tail      LinkNode
}

func NewLinkedList(size int) *LinkedList {
	head := NewDoubleLinkNode(nil, nil, nil)
	tail := head
	curr := head
	return &LinkedList{0, head, tail, curr}
}

func ConvertToLinkedList(size int, listArray ...interface{}) *LinkedList {
	linkedList := NewLinkedList(0)
	linkedList.Append(listArray...)
	return linkedList
}

func (l *LinkedList) Clear() {
	l.numInList = 0
	l.head.SetNext(nil)
	l.tail = l.head
	l.curr = l.head
}

func (l *LinkedList) Insert(index int, items ...interface{}) {
	if !l.isInList(index) && index != l.numInList {
		panic(&NullCurrError{})
	}
	l.curr = l.head
	for i := 0; i < index; i++ {
		l.curr = l.curr.Next()
	}
	for _, item := range items {
		l.curr.SetNext(NewDoubleLinkNode(item, l.curr, l.curr.Next()))
		if l.curr.Next().Next() != nil {
			l.curr.Next().Next().SetPrev(l.curr.Next())
		}
		if l.tail == l.curr {
			l.tail = l.curr.Next()
		}
		l.curr = l.curr.Next()
		l.numInList++
	}
}

func (l *LinkedList) Append(items ...interface{}) {
	for _, item := range items {
		l.tail.SetNext(NewDoubleLinkNode(item, l.tail, nil))
		l.tail = l.tail.Next()
		l.numInList++
	}
}

// TODO
func (l *LinkedList) Remove(index int) interface{} {
	if l.IsEmpty() {
		panic(&EmptyListError{})
	}
	l.setPos(index)
	value := l.curr.Element()
	l.curr.SetNext(l.curr.Next().Next())
	if l.curr.Next() == l.tail {
		l.tail = l.curr
	}
	l.numInList--
	return value
}

func (l *LinkedList) Length() int {
	return l.numInList
}

func (l *LinkedList) setPos(index int) {
	if !l.isInList(index) {
		panic(&NullCurrError{})
	}
	l.curr = l.head
	for i := 0; i < index; i++ {
		l.curr = l.curr.Next()
	}
}

func (l *LinkedList) SetValue(index int, value interface{}) {
	l.setPos(index)
	l.curr.Next().SetElement(value)
}

func (l *LinkedList) Get(index int) interface{} {
	if !l.isInList(index) {
		// TODO:!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		panic(&NullCurrError{})
	}
	l.setPos(index)
	return l.curr.Next().Element()
}

func (l *LinkedList) IndexOf(value interface{}) int {
	for i, item := 0, l.head.Next(); l.isInList(i); i, item = i+1, item.Next() {
		if reflect.DeepEqual(item.Next().Element(), value) {
			return i
		}
	}
	return -1
}

func (l *LinkedList) Contains(value interface{}) bool {
	if l.IndexOf(value) == -1 {
		return false
	}
	return true
}

func (l *LinkedList) IsEmpty() bool {
	return l.head.Next() == nil
}

func (l *LinkedList) isInList(index int) bool {
	return index >= 0 && index < l.numInList
}

func (l *LinkedList) String() string {
	var buffer bytes.Buffer
	if l.IsEmpty() {
		return "()"
	}
	buffer.WriteString("(")
	for i, item := 0, l.head.Next(); l.isInList(i); i, item = i+1, item.Next() {
		value := item.Element()
		if i != l.numInList-1 {
			buffer.WriteString(fmt.Sprintf("%v, ", value))
		} else {
			buffer.WriteString(fmt.Sprintf("%v)", value))
		}
	}
	return buffer.String()
}
