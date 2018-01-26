package arraystack

import (
	"github.com/Thrimbda/dune/list/arraylist"
)

type ArrayStack struct {
	list *arraylist.ArrayList
}

func NewArrayStack(size int) *ArrayStack {
	return &ArrayStack{arraylist.NewArrayList(size)}
}

func ConvertToArrayStack(size int, items ...interface{}) *ArrayStack {
	return &ArrayStack{arraylist.ConvertToArrayList(size, items)}
}

func (a *ArrayStack) Clear() {
	a.list.Clear()
}

func (a *ArrayStack) Push(item interface{}) {
	a.list.Append(item)
}

func (a *ArrayStack) Pop() interface{} {
	return a.list.Remove(a.list.Length() - 1)
}

func (a *ArrayStack) Peek() interface{} {
	return a.list.Get(a.list.Length() - 1)
}

func (a *ArrayStack) IsEmpty() bool {
	return a.list.IsEmpty()
}
