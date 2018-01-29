package linkedstack

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/list/linkedlist"
	"github.com/Thrimbda/dune/utils"
)

func TestNewLinkedStack(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedStack
	}{
		{"new_empty", &LinkedStack{linkedlist.NewLinkedList()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToLinkedStack(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want *LinkedStack
	}{
		{"new_convert", args{[]interface{}{1, 2, 3, 4, 2242}}, &LinkedStack{linkedlist.ConvertToLinkedList(2242, 4, 3, 2, 1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToLinkedStack(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToLinkedStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedStack_Clear(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedStack
		want bool
	}{
		{"empty1", NewLinkedStack(), true},
		{"empty2", ConvertToLinkedStack(1, 2, 3, 4, 5), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Clear()
			if got := tt.l.IsEmpty(); got != tt.want {
				t.Errorf("Clear failed, got: %v", got)
			}
		})
	}
}

func TestLinkedStack_Push(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		l    *LinkedStack
		args args
		want interface{}
	}{
		{"empty_push", NewLinkedStack(), args{3}, 3},
		{"push1", ConvertToLinkedStack(1, 2, 3, 4), args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Push(tt.args.item)
			if got := tt.l.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push failed got top value: %v", got)
			}
		})
	}
}

func TestLinkedStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedStack
		want interface{}
	}{
		{"pop1", ConvertToLinkedStack(1, 2, 3), 3},
		{"pop2", ConvertToLinkedStack(make([]interface{}, 100)...), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedStack_Peek(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedStack
		want interface{}
	}{
		{"peek1", ConvertToLinkedStack(1, 2, 3), 3},
		{"peek2", ConvertToLinkedStack(make([]interface{}, 100)...), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedStack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedStack
		want bool
	}{
		{"empty", NewLinkedStack(), true},
		{"not_empty", ConvertToLinkedStack(2, 1, 2), false},
		{"not_empty", ConvertToLinkedStack(2), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsEmpty(); got != tt.want {
				t.Errorf("LinkedStack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

type City struct {
	id   int
	name string
}

func (c *City) LessComparator(b utils.Elem) bool {
	return c.id < b.(*City).id
}
func (c *City) String() string {
	return fmt.Sprintf("%v. %v", c.id, c.name)
}

func TestLinkedStack_String(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedStack
		want string
	}{
		{"string1", NewLinkedStack(), "()"},
		{"string2", ConvertToLinkedStack(2, 3), "(2, 3)"},
		{"string3", &LinkedStack{linkedlist.ConvertToLinkedList(3, 2, 1)}, "(1, 2, 3)"},
		{"string4", ConvertToLinkedStack([]interface{}{&City{1, "Beijing"}, &City{2, "Shanghai"}, &City{3, "Xi'an"}}...), "(1. Beijing, 2. Shanghai, 3. Xi'an)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("LinkedQueue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
