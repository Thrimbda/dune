package linkedstack

import (
	"bytes"
	"fmt"

	"github.com/Thrimbda/dune/list/linkedlist"
)

type LinkedStack struct {
	list *linkedlist.LinkedList
}

func NewLinkedStack() *LinkedStack {
	return &LinkedStack{linkedlist.NewLinkedList()}
}

func ConvertToLinkedStack(items ...interface{}) *LinkedStack {
	stack := NewLinkedStack()
	for _, item := range items {
		stack.Push(item)
	}
	return stack
}

func (l *LinkedStack) Clear() {
	l.list.Clear()
}

func (l *LinkedStack) Push(item interface{}) {
	l.list.Insert(0, item)
}

func (l *LinkedStack) Pop() interface{} {
	return l.list.Remove(0)
}

func (l *LinkedStack) Peek() interface{} {
	return l.list.Get(0)
}

func (l *LinkedStack) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedStack) String() string {
	var buffer bytes.Buffer
	if l.IsEmpty() {
		return "()"
	}
	buffer.WriteString("(")
	values := l.list.Values()
	for i := len(values) - 1; i >= 0; i-- {
		value := values[i]
		if i != 0 {
			buffer.WriteString(fmt.Sprintf("%v, ", value))
		} else {
			buffer.WriteString(fmt.Sprintf("%v)", value))
		}
	}
	return buffer.String()
}
