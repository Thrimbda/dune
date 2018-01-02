package linkedqueue

import (
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/datastructure/arrayutils"
	"github.com/Thrimbda/dune/datastructure/linkutils"
	"github.com/Thrimbda/dune/datastructure/list/linkedlist"
)

func TestNewLinkedQueue(t *testing.T) {
	tests := []struct {
		name string
		want *LinkedQueue
	}{
		{"new_empty", &LinkedQueue{linkedlist.NewLinkedList()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToLinkedQueue(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedQueue
	}{
		{"empty1", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2, 3)}},
		{"empty2", &LinkedQueue{linkedlist.ConvertToLinkedList(make([]int, 100, 100))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Clear()
			if !tt.a.list.IsEmpty() {
				t.Errorf("Clear failed, got size %v", tt.a.list.Length())
			}
		})
	}
}

func TestLinkedQueue_Clear(t *testing.T) {
	tests := []struct {
		name string
		l    LinkedQueue
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Clear()
		})
	}
}

func TestLinkedQueue_Enqueue(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name    string
		a       *LinkedQueue
		args    args
		length  int
		element interface{}
	}{
		{"enqueue", NewLinkedQueue(), args{1}, 1, 1},
		{"enqueue2", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2, 3)}, args{4}, 4, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Enqueue(tt.args.item)
			if got := tt.a.Peek(); !reflect.DeepEqual(got, tt.element) {
				t.Errorf("expect %v, but got %v", tt.element, got)
			}
			if got := tt.a.list.Length(); got != tt.length {
				t.Errorf("expect front %v, but got front %v", tt.length, got)
			}
		})
	}
}

func TestLinkedQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name      string
		a         *LinkedQueue
		length    int
		want      interface{}
		testPanic error
	}{
		{"dequeue", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2)}, 1, 1, nil},
		{"panic", &LinkedQueue{linkedlist.NewLinkedList()}, 0, 1, &arrayutils.EmptyListError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Dequeue(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedQueue.Dequeue() = %v, want %v", got, tt.want)
				}
				if got := tt.a.list.Length(); got != tt.length {
					t.Errorf("expect length %v, but got length %v", tt.length, got)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but got %v", tt.testPanic, r)
					}
				}()
				tt.a.Dequeue()
			}
		})
	}
}

func TestLinkedQueue_Peek(t *testing.T) {
	tests := []struct {
		name      string
		a         *LinkedQueue
		want      interface{}
		testPanic error
	}{
		{"peek", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2)}, 1, nil},
		{"panic", &LinkedQueue{linkedlist.NewLinkedList()}, 1, &linkutils.NullCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Peek(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedQueue.Dequeue() = %v, want %v", got, tt.want)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but got %v", tt.testPanic, r)
					}
				}()
				tt.a.Peek()
			}
		})
	}
}

func TestLinkedQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedQueue
		want bool
	}{
		{"empty", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2)}, false},
		{"not_empty", &LinkedQueue{linkedlist.NewLinkedList()}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsEmpty(); got != tt.want {
				t.Errorf("LinkedQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
