package queue

import (
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/datastructure/arrayutils"
	"github.com/Thrimbda/dune/datastructure/list/arraylist"
)

func TestNewArrayQueue(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *ArrayQueue
	}{
		{"new_empty", args{0}, &ArrayQueue{0, 0, 0, arraylist.NewArrayList(0)}},
		{"new1", args{3}, &ArrayQueue{3, 0, 0, arraylist.NewArrayList(3)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayQueue(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayQueue_clear(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayQueue
	}{
		{"empty1", &ArrayQueue{3, 0, 3, arraylist.ConvertToArrayList(3, 1, 2, 3)}},
		{"empty2", &ArrayQueue{3, 3, 100, arraylist.ConvertToArrayList(3, make([]int, 100, 100))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Clear()
			if tt.a.front != 0 && tt.a.rear != 0 {
				t.Errorf("Clear failed, got front pointer %v, rear pointer %v", tt.a.front, tt.a.rear)
			}
		})
	}
}

func TestArrayQueue_Enqueue(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name      string
		a         *ArrayQueue
		args      args
		length    int
		element   interface{}
		testPanic error
	}{
		{"enqueue", NewArrayQueue(3), args{1}, 1, 1, nil},
		{"panic", NewArrayQueue(0), args{2}, 1, 2, &arrayutils.FullListError{}},
		{"panic", &ArrayQueue{2, 0, 1, arraylist.ConvertToArrayList(2, 1, 2)}, args{2}, 1, 2, &arrayutils.FullListError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				tt.a.Enqueue(tt.args.item)
				if got := tt.a.Peek(); !reflect.DeepEqual(got, tt.element) {
					t.Errorf("expect %v, but got %v", tt.element, got)
				}
				if got := func() int {
					if tt.a.front > tt.a.rear {
						return tt.a.rear + tt.a.size - tt.a.front
					}
					return tt.a.rear - tt.a.front
				}(); got != tt.length {
					t.Errorf("expect front %v, but got front %v", tt.length, got)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but got %v", tt.testPanic, r)
					}
				}()
				tt.a.Enqueue(tt.args.item)
			}
		})
	}
}

func TestArrayQueue_Dequeue(t *testing.T) {
	tests := []struct {
		name      string
		a         *ArrayQueue
		length    int
		want      interface{}
		testPanic error
	}{
		{"dequeue", &ArrayQueue{2, 0, 1, arraylist.ConvertToArrayList(2, 1, 2)}, 0, 1, nil},
		{"panic", &ArrayQueue{2, 1, 1, arraylist.ConvertToArrayList(2, 1, 2)}, 0, 1, &arrayutils.EmptyListError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Dequeue(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ArrayQueue.Dequeue() = %v, want %v", got, tt.want)
				}
				if got := func() int {
					if tt.a.front > tt.a.rear {
						return tt.a.rear + tt.a.size - tt.a.front
					}
					return tt.a.rear - tt.a.front
				}(); got != tt.length {
					t.Errorf("expect front %v, but got front %v", tt.length, got)
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

func TestArrayQueue_Peek(t *testing.T) {
	tests := []struct {
		name      string
		a         *ArrayQueue
		want      interface{}
		testPanic error
	}{
		{"peek", &ArrayQueue{2, 0, 1, arraylist.ConvertToArrayList(2, 1, 2)}, 1, nil},
		{"panic", &ArrayQueue{2, 1, 1, arraylist.ConvertToArrayList(2, 1, 2)}, 1, &arrayutils.EmptyListError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Peek(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ArrayQueue.Dequeue() = %v, want %v", got, tt.want)
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

func TestArrayQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayQueue
		want bool
	}{
		{"empty", &ArrayQueue{2, 0, 1, arraylist.ConvertToArrayList(2, 1, 2)}, false},
		{"not_empty", &ArrayQueue{2, 1, 1, arraylist.ConvertToArrayList(2, 1, 2)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsEmpty(); got != tt.want {
				t.Errorf("ArrayQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
