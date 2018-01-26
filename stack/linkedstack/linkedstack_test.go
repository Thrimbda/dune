package linkedstack

import (
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/list/linkedlist"
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

func TestConverToLinkedStack(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want *LinkedStack
	}{
		{"new_convert", args{[]interface{}{1, 2, 3, 4, 2242}}, &LinkedStack{linkedlist.ConvertToLinkedList(1, 2, 3, 4, 2242)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConverToLinkedStack(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConverToLinkedStack() = %v, want %v", got, tt.want)
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
		{"empty2", ConverToLinkedStack(1, 2, 3, 4, 5), true},
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
		{"push1", ConverToLinkedStack(1, 2, 3, 4), args{5}, 5},
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
		{"pop1", ConverToLinkedStack(1, 2, 3), 3},
		{"pop2", ConverToLinkedStack(make([]interface{}, 100)...), nil},
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
		{"peek1", ConverToLinkedStack(1, 2, 3), 3},
		{"peek2", ConverToLinkedStack(make([]interface{}, 100)...), nil},
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
		{"not_empty", ConverToLinkedStack(2, 1, 2), false},
		{"not_empty", ConverToLinkedStack(2), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsEmpty(); got != tt.want {
				t.Errorf("LinkedStack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
