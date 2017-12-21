package list

import (
	"reflect"
	"testing"

	// . "github.com/Thrimbda/dune/datastructure/arrayutils"
	. "github.com/Thrimbda/dune/datastructure/linkutils"
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
		{"empty", args{0}, &LinkedList{0, NewDoubleLinkNode(nil, nil, nil), NewDoubleLinkNode(nil, nil, nil), NewDoubleLinkNode(nil, nil, nil)}},
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
		{"panic", NewLinkedList(0), args{[]interface{}{1, 1, 1}, 3}, want{false, 3, 1, 0}, &NullCurrError{}},
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
	tests := []struct {
		name string
		l    LinkedList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Append(tt.args.items...)
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Length(t *testing.T) {
	tests := []struct {
		name string
		l    LinkedList
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Length(); got != tt.want {
				t.Errorf("LinkedList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_setPos(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.setPos(tt.args.index)
		})
	}
}

func TestLinkedList_SetValue(t *testing.T) {
	type args struct {
		index int
		value interface{}
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.SetValue(tt.args.index, tt.args.value)
		})
	}
}

func TestLinkedList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IndexOf(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IndexOf(tt.args.value); got != tt.want {
				t.Errorf("LinkedList.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Contains(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		l    LinkedList
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Contains(tt.args.value); got != tt.want {
				t.Errorf("LinkedList.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		l    LinkedList
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsEmpty(); got != tt.want {
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
		l    LinkedList
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.isInList(tt.args.index); got != tt.want {
				t.Errorf("LinkedList.isInList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_String(t *testing.T) {
	tests := []struct {
		name string
		l    LinkedList
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("LinkedList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
