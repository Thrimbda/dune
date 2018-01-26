package arraylist

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/arrayutils"
	"github.com/Thrimbda/dune/utils"
)

func TestNewArrayList(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *ArrayList
	}{
		{"empty", args{0}, &ArrayList{0, 0, make([]interface{}, 0)}},
		{"100", args{100}, &ArrayList{100, 0, make([]interface{}, 100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayList(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Clear(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
		want *ArrayList
	}{
		{"test1", &ArrayList{3, 3, []interface{}{1, 2, 3}}, &ArrayList{3, 0, []interface{}{1, 2, 3}}},
		{"test2", &ArrayList{100, 3, make([]interface{}, 3, 100)}, &ArrayList{100, 0, make([]interface{}, 3, 100)}},
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

func TestArrayList_Insert(t *testing.T) {
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
		a         *ArrayList
		args      args
		want      want
		testPanic error
	}{
		{"append_1", &ArrayList{100, 3, make([]interface{}, 3, 100)}, args{[]interface{}{1}, 2}, want{false, 4, 1, 2}, nil},
		{"append_2", NewArrayList(100), args{[]interface{}{2, 3}, 0}, want{false, 2, 3, 1}, nil},
		{"panic", NewArrayList(0), args{[]interface{}{1, 1, 1}, 0}, want{false, 3, 1, 0}, &arrayutils.FullListError{}},
		{"panic", NewArrayList(100), args{[]interface{}{1, 1, 1}, 3}, want{false, 3, 1, 0}, &arrayutils.BadCurrError{}},
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

func TestArrayList_Append(t *testing.T) {
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
		name      string
		a         *ArrayList
		args      args
		want      want
		testPanic error
	}{
		{"append_1", NewArrayList(100), args{[]interface{}{1}}, want{false, 1, 1, 0}, nil},
		{"append_2", NewArrayList(100), args{[]interface{}{2, 3}}, want{false, 2, 3, 1}, nil},
		{"panic", NewArrayList(0), args{[]interface{}{1, 1, 1}}, want{false, 3, 1, 0}, &arrayutils.FullListError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
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
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but get %v", tt.testPanic, r)
					}
				}()
				tt.a.Append(tt.args.items...)
			}
		})
	}
}

func TestArrayList_Remove(t *testing.T) {
	type args struct {
		indexes []int
	}
	type want struct {
		value  interface{}
		length int
		remain *ArrayList
	}
	type testPanic struct {
		panic error
		round int
	}
	tests := []struct {
		name      string
		a         *ArrayList
		args      args
		want      []want
		testPanic testPanic
	}{
		{"remove1", &ArrayList{3, 3, []interface{}{1, 2, 3}}, args{[]int{0}},
			[]want{
				want{1, 2, &ArrayList{3, 2, []interface{}{2, 3, 3}}},
			}, testPanic{nil, 0}},
		{"remove2", &ArrayList{3, 3, []interface{}{1, 2, 3}}, args{[]int{0, 1}},
			[]want{
				want{1, 2, &ArrayList{3, 2, []interface{}{2, 3, 3}}},
				want{3, 1, &ArrayList{3, 1, []interface{}{2, 3, 3}}},
			}, testPanic{nil, 0}},
		{"panic1", &ArrayList{2, 2, []interface{}{1, 2}}, args{[]int{0, 1}},
			[]want{
				want{1, 1, &ArrayList{2, 1, []interface{}{2, 2}}},
				want{2, 0, &ArrayList{2, 0, []interface{}{2, 2}}},
			}, testPanic{&arrayutils.BadCurrError{}, 1}},
		{"panic2", NewArrayList(0), args{[]int{0}},
			[]want{want{0, 0, NewArrayList(0)}}, testPanic{&arrayutils.EmptyListError{}, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for round, index := range tt.args.indexes {
				if tt.testPanic.panic == nil || tt.testPanic.round != round {
					value := tt.a.Remove(index)
					if got := value; !reflect.DeepEqual(got, tt.want[round].value) {
						t.Errorf("ArrayList.Remove() = %v, want %v", got, tt.want[round].value)
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

func TestArrayList_Length(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
		want int
	}{
		{"length1", NewArrayList(100), 0},
		{"length2", NewArrayList(10000), 0},
		{"length3", NewArrayList(0), 0},
		{"length4", &ArrayList{3, 3, []interface{}{2, 3, 3}}, 3},
		{"length5", &ArrayList{3, 2, []interface{}{2, 3, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Length(); got != tt.want {
				t.Errorf("ArrayList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	type args struct {
		index int
		item  interface{}
	}
	tests := []struct {
		name      string
		a         *ArrayList
		args      args
		want      interface{}
		testPanic error
	}{
		{"set1", &ArrayList{3, 3, []interface{}{2, 3, 3}}, args{0, 3}, 3, nil},
		{"set2", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{1, 2}, 2, nil},
		{"panic", NewArrayList(100), args{50, 1}, 1, &arrayutils.BadCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				tt.a.Set(tt.args.index, tt.args.item)
				if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.args.item) {
					t.Errorf("ArrayList.Set() = %v, want %v", got, tt.want)
				}
			} else {
				defer func() {
					if r := recover(); !reflect.DeepEqual(r, tt.testPanic) {
						t.Errorf("expect %v, but get %v", tt.testPanic, r)
					}
				}()
				tt.a.Set(tt.args.index, tt.args.item)
			}
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name      string
		a         *ArrayList
		args      args
		want      interface{}
		testPanic error
	}{
		{"set1", &ArrayList{3, 3, []interface{}{2, 3, 3}}, args{0}, 2, nil},
		{"set2", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{1}, 3, nil},
		{"panic", NewArrayList(100), args{50}, 0, &arrayutils.BadCurrError{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testPanic == nil {
				if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ArrayList.Get() = %v, want %v", got, tt.want)
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

func TestArrayList_IndexOf(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
		want int
	}{
		{"index_of1", &ArrayList{3, 3, []interface{}{1, 2, 3}}, args{2}, 1},
		{"index_of12", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{3}, 1},
		{"not_found1", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{4}, -1},
		{"not_found2", &ArrayList{3, 2, []interface{}{2, 4, 3}}, args{3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IndexOf(tt.args.item); got != tt.want {
				t.Errorf("ArrayList.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_Contains(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
		want bool
	}{
		{"contains1", &ArrayList{3, 3, []interface{}{1, 2, 3}}, args{2}, true},
		{"contains2", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{3}, true},
		{"not_found1", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{4}, false},
		{"not_found2", &ArrayList{3, 2, []interface{}{2, 4, 3}}, args{3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Contains(tt.args.item); got != tt.want {
				t.Errorf("ArrayList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
		want bool
	}{
		{"empty1", &ArrayList{3, 0, []interface{}{2, 3, 3}}, true},
		{"empty2", NewArrayList(100), true},
		{"empty3", NewArrayList(0), true},
		{"not_empty", &ArrayList{3, 3, []interface{}{1, 2, 3}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsEmpty(); got != tt.want {
				t.Errorf("ArrayList.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_isInList(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
		want bool
	}{
		{"in_list1", &ArrayList{3, 3, []interface{}{1, 2, 3}}, args{2}, true},
		{"in_list2", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{1}, true},
		{"not_in_list1", &ArrayList{3, 2, []interface{}{2, 3, 3}}, args{2}, false},
		{"not_in_list2", NewArrayList(100), args{0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.isInList(tt.args.index); got != tt.want {
				t.Errorf("ArrayList.isInList() = %v, want %v", got, tt.want)
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

func TestArrayList_String(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
		want string
	}{
		{"string1", NewArrayList(100), "()"},
		{"string2", &ArrayList{3, 2, []interface{}{2, 3, 3}}, "(2, 3)"},
		{"string3", &ArrayList{3, 3, []interface{}{1, 2, 3}}, "(1, 2, 3)"},
		{"string3", &ArrayList{3, 3, []interface{}{&City{1, "Beijing"}, &City{2, "Shanghai"}, &City{3, "Xi'an"}}}, "(1. Beijing, 2. Shanghai, 3. Xi'an)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("ArrayList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
