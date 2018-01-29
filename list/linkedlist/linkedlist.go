package linkedlist

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/Thrimbda/dune/arrayutils"
	"github.com/Thrimbda/dune/linkutils"
)

type LinkedList struct {
	numInList int
	head      linkutils.LinkNode
	tail      linkutils.LinkNode
}

func NewLinkedList() *LinkedList {
	head := linkutils.NewDoubleLinkNode(nil, nil, nil)
	tail := head
	return &LinkedList{0, head, tail}
}

func ConvertToLinkedList(listArray ...interface{}) *LinkedList {
	linkedList := NewLinkedList()
	linkedList.Append(listArray...)
	return linkedList
}

func (l *LinkedList) Clear() {
	l.numInList = 0
	l.head.SetNext(nil)
	l.tail = l.head
}

func (l *LinkedList) Insert(index int, items ...interface{}) {
	if !l.isInList(index) && index != l.numInList {
		panic(&linkutils.NullCurrError{})
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next()
	}
	for _, item := range items {
		curr.SetNext(linkutils.NewDoubleLinkNode(item, curr, curr.Next()))
		if curr.Next().Next() != nil {
			curr.Next().Next().SetPrev(curr.Next())
		}
		if l.tail == curr {
			l.tail = curr.Next()
		}
		curr = curr.Next()
		l.numInList++
	}
}

func (l *LinkedList) Append(items ...interface{}) {
	for _, item := range items {
		l.tail.SetNext(linkutils.NewDoubleLinkNode(item, l.tail, nil))
		l.tail = l.tail.Next()
		l.numInList++
	}
}

// TODO
func (l *LinkedList) Remove(index int) interface{} {
	if l.IsEmpty() {
		panic(&arrayutils.EmptyListError{})
	}
	if !l.isInList(index) {
		panic(&linkutils.NullCurrError{})
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next()
	}
	value := curr.Next().Element()
	if curr.Next().Next() != nil {
		curr.Next().Next().SetPrev(curr)
	} else {
		l.tail = curr
	}
	curr.SetNext(curr.Next().Next())
	l.numInList--
	curr = l.head
	return value
}

func (l *LinkedList) Length() int {
	return l.numInList
}

func (l *LinkedList) Set(index int, value interface{}) {
	if !l.isInList(index) {
		panic(&linkutils.NullCurrError{})
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next()
	}
	curr.Next().SetElement(value)
}

func (l *LinkedList) Get(index int) interface{} {
	if !l.isInList(index) {
		panic(&linkutils.NullCurrError{})
	}
	if !l.isInList(index) {
		panic(&linkutils.NullCurrError{})
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.Next()
	}
	return curr.Next().Element()
}

func (l *LinkedList) IndexOf(value interface{}) int {
	for i, item := 0, l.head.Next(); l.isInList(i); i, item = i+1, item.Next() {
		if reflect.DeepEqual(item.Element(), value) {
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

func (l *LinkedList) Values() []interface{} {
	values := make([]interface{}, l.numInList)
	for i, item := 0, l.head.Next(); l.isInList(i); i, item = i+1, item.Next() {
		values[i] = item.Element()
	}
	return values
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
