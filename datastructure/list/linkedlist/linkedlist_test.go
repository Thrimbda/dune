package linkedlist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/datastructure"
	"github.com/Thrimbda/dune/datastructure/arrayutils"
	"github.com/Thrimbda/dune/datastructure/linkutils"
)

func TestNewLinkedList(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{"empty", args{0}, &LinkedList{0, linkutils.NewDoubleLinkNode(nil, nil, nil), linkutils.NewDoubleLinkNode(nil, nil, nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Clear(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedList
		want *LinkedList
	}{
		{"test1", ConvertToLinkedList(0, 1, 2, 3), NewLinkedList(0)},
		{"test2", ConvertToLinkedList(0, make([]interface{}, 100, 100)...), NewLinkedList(0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Clear()
			if got := tt.a; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("list cleared = %v, want %v", got, tt.want)
			}
			if got := tt.a.Length(); got != 0 {
				t.Errorf("expect %v, but got %v", 0, got)
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	type args struct {
		items []interface{}
		index int
	}
	type want struct {
		isEmpty bool
		length  int
		elem    interface{}
		index   int
	}
	tests := []struct {
		name      string
		a         *LinkedList
		args      args
		want      want
		testPanic error
	}{
		{"append_1", ConvertToLinkedList(3, 1, 3, 4), args{[]interface{}{1}, 2}, want{false, 4, 1, 2}, nil},
		{"append_2", NewLinkedList(100), args{[]interface{}{2, 3}, 0}, want{false, 2, 3, 1}, nil},
		{"panic", NewLinkedList(0), args{[]interface{}{1, 1, 1}, 3}, want{false, 3, 1, 0}, &linkutils.NullCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				tt.a.Insert(tt.args.index, tt.args.items...)
				if got := tt.a.IsEmpty(); got != tt.want.isEmpty {
					t.Errorf("expect %v, but got %v", tt.want.isEmpty, got)
				}
				if got := tt.a.Length(); got != tt.want.length {
					t.Errorf("expect %v, but got %v", tt.want.length, got)
				}
				if got := tt.a.Get(tt.want.index); !reflect.DeepEqual(got, tt.want.elem) {
					t.Errorf("expect %v, but got %v", tt.want.elem, got)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but get %v", tt.testPanic, r)
					}
				}()
				tt.a.Insert(tt.args.index, tt.args.items...)
			}
		})
	}
}

func TestLinkedList_Append(t *testing.T) {
	type args struct {
		items []interface{}
	}
	type want struct {
		isEmpty bool
		length  int
		elem    interface{}
		index   int
	}
	tests := []struct {
		name string
		a    *LinkedList
		args args
		want want
	}{
		{"append_1", NewLinkedList(100), args{[]interface{}{1}}, want{false, 1, 1, 0}},
		{"append_2", NewLinkedList(100), args{[]interface{}{2, 3}}, want{false, 2, 3, 1}},
		{"appended", ConvertToLinkedList(3, 1, 2, 4), args{[]interface{}{1, 1, 1}}, want{false, 6, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Append(tt.args.items...)
			if got := tt.a.IsEmpty(); got != tt.want.isEmpty {
				t.Errorf("expect %v, but got %v", tt.want.isEmpty, got)
			}
			if got := tt.a.Length(); got != tt.want.length {
				t.Errorf("expect %v, but got %v", tt.want.length, got)
			}
			if got := tt.a.Get(tt.want.index); !reflect.DeepEqual(got, tt.want.elem) {
				t.Errorf("expect %v, but got %v", tt.want.elem, got)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type args struct {
		indexes []int
	}
	type want struct {
		value  interface{}
		length int
		remain *LinkedList
	}
	type testPanic struct {
		panic error
		round int
	}
	tests := []struct {
		name      string
		a         *LinkedList
		args      args
		want      []want
		testPanic testPanic
	}{
		{"remove1", ConvertToLinkedList(3, 1, 2, 3), args{[]int{0}},
			[]want{
				want{1, 2, ConvertToLinkedList(2, 2, 3)},
			}, testPanic{nil, 0}},
		{"remove2", ConvertToLinkedList(3, 1, 2, 3), args{[]int{0, 1}},
			[]want{
				want{1, 2, ConvertToLinkedList(2, 2, 3)},
				want{3, 1, ConvertToLinkedList(1, 2)},
			}, testPanic{nil, 0}},
		{"panic1", ConvertToLinkedList(2, 1, 2), args{[]int{0, 1}},
			[]want{
				want{1, 1, ConvertToLinkedList(1, 2)},
				want{2, 0, NewLinkedList(0)},
			}, testPanic{&linkutils.NullCurrError{}, 1}},
		{"panic2", NewLinkedList(0), args{[]int{0}},
			[]want{want{0, 0, NewLinkedList(0)}}, testPanic{&arrayutils.EmptyListError{}, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for round, index := range tt.args.indexes {
				if tt.testPanic.panic == nil || tt.testPanic.round != round {
					value := tt.a.Remove(index)
					if got := value; !reflect.DeepEqual(got, tt.want[round].value) {
						t.Errorf("LinkedList.Remove() = %v, want %v", got, tt.want[round].value)
					}
					if got := tt.a.Length(); got != tt.want[round].length {
						t.Errorf("expect %v, but got %v", tt.want[round].length, got)
					}
					if got := tt.a; !reflect.DeepEqual(got, tt.want[round].remain) {
						t.Errorf("expect %v, but got %v", tt.want[round].remain, got)
					}
				} else {
					defer func() {
						if r := recover(); !reflect.DeepEqual(r, tt.testPanic.panic) {
							t.Errorf("expect %v, but get %v", tt.testPanic.panic, r)
						}
					}()
					tt.a.Remove(index)
				}
			}
		})
	}
}

func TestLinkedList_Length(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedList
		want int
	}{
		{"length1", NewLinkedList(0), 0},
		{"length2", ConvertToLinkedList(3, 1, 2, 3), 3},
		{"length5", ConvertToLinkedList(2, make([]interface{}, 2, 2)...), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Length(); got != tt.want {
				t.Errorf("LinkedList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_SetValue(t *testing.T) {
	type args struct {
		index int
		item  interface{}
	}
	tests := []struct {
		name      string
		a         *LinkedList
		args      args
		want      interface{}
		testPanic error
	}{
		{"set1", ConvertToLinkedList(3, 2, 3, 3), args{0, 3}, 3, nil},
		{"set2", ConvertToLinkedList(2, 3, 4), args{1, 2}, 2, nil},
		{"panic", NewLinkedList(100), args{50, 1}, 1, &linkutils.NullCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				tt.a.SetValue(tt.args.index, tt.args.item)
				if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.args.item) {
					t.Errorf("LinkedList.SetValue() = %v, want %v", got, tt.want)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but get %v", tt.testPanic, r)
					}
				}()
				tt.a.SetValue(tt.args.index, tt.args.item)
			}
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		a         *LinkedList
		args      args
		want      interface{}
		testPanic error
	}{
		{"set1", ConvertToLinkedList(3, 2, 3, 3), args{0}, 2, nil},
		{"set2", ConvertToLinkedList(2, 2, 3), args{1}, 3, nil},
		{"panic", NewLinkedList(0), args{50}, 0, &linkutils.NullCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("LinkedList.Get() = %v, want %v", got, tt.want)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but got %v", tt.testPanic, r)
					}
				}()
				tt.a.Get(tt.args.index)
			}
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		a    *LinkedList
		args args
		want int
	}{
		{"index_of1", ConvertToLinkedList(3, 1, 2, 3), args{3}, 2},
		{"index_of2", ConvertToLinkedList(2, 2, 3), args{3}, 1},
		{"not_found1", ConvertToLinkedList(2, 1, 2), args{4}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IndexOf(tt.args.item); got != tt.want {
				t.Errorf("LinkedList.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		a    *LinkedList
		args args
		want bool
	}{
		{"index_of1", ConvertToLinkedList(3, 1, 2, 3), args{2}, true},
		{"index_of2", ConvertToLinkedList(2, 2, 3), args{3}, true},
		{"not_found1", ConvertToLinkedList(2, 1, 2), args{4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Contains(tt.args.item); got != tt.want {
				t.Errorf("LinkedList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedList
		want bool
	}{
		{"empty3", NewLinkedList(0), true},
		{"not_empty", ConvertToLinkedList(3, 1, 2, 3), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsEmpty(); got != tt.want {
				t.Errorf("LinkedList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_isInList(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		a    *LinkedList
		args args
		want bool
	}{
		{"in_list1", ConvertToLinkedList(3, 1, 2, 3), args{2}, true},
		{"in_list2", ConvertToLinkedList(2, 2, 3), args{1}, true},
		{"not_in_list1", ConvertToLinkedList(2, 2, 3), args{2}, false},
		{"not_in_list2", NewLinkedList(100), args{0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.isInList(tt.args.index); got != tt.want {
				t.Errorf("LinkedList.isInList() = %v, want %v", got, tt.want)
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

func TestLinkedList_String(t *testing.T) {
	tests := []struct {
		name string
		a    *LinkedList
		want string
	}{
		{"string1", NewLinkedList(100), "()"},
		{"string2", ConvertToLinkedList(2, 2, 3), "(2, 3)"},
		{"string3", ConvertToLinkedList(3, 1, 2, 3), "(1, 2, 3)"},
		{"string3", ConvertToLinkedList(3, []interface{}{&City{1, "Beijing"}, &City{2, "Shanghai"}, &City{3, "Xi'an"}}...), "(1. Beijing, 2. Shanghai, 3. Xi'an)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("LinkedList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
