package linkedqueue

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/datastructure"
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
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name      string
		args      args
		want      *LinkedQueue
		testPanic error
	}{
		{"new", args{[]interface{}{1, 2, 3}}, &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2, 3)}, nil},
		{"huge", args{make([]interface{}, 99)}, &LinkedQueue{linkedlist.ConvertToLinkedList(make([]interface{}, 99)...)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToLinkedQueue(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToLinkedQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedQueue_Clear(t *testing.T) {
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

type City struct {
	id   int
	name string
}

func (c *City) LessComparator(b datastructure.Elem) bool {
	return c.id < b.(*City).id
}
func (c *City) String() string {
	return fmt.Sprintf("%v. %v", c.id, c.name)
}

func TestLinkedQueue_String(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedQueue
		want string
	}{
		{"string1", NewLinkedQueue(), "()"},
		{"string2", ConvertToLinkedQueue(2, 3), "(2, 3)"},
		{"string3", &LinkedQueue{linkedlist.ConvertToLinkedList(1, 2, 3)}, "(1, 2, 3)"},
		{"string4", ConvertToLinkedQueue([]interface{}{&City{1, "Beijing"}, &City{2, "Shanghai"}, &City{3, "Xi'an"}}...), "(1. Beijing, 2. Shanghai, 3. Xi'an)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("LinkedQueue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
