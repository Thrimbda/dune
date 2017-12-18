package arraylist

import (
	"reflect"
	"testing"
	// . "../../arrayutils"
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

func TestArrayList(t *testing.T) {
	list := NewArrayList(100)
	list.Append(1)
	list.Append(2, 3)
	if got := list.IsEmpty(); got != false {
		t.Errorf("expect %v, but got %v", false, got)
	}
	if got := list.Length(); got != 3 {
		t.Errorf("expect %v, but got %v", false, got)
	}
	if got := list.Get(1); got != 2 {
		t.Errorf("expect %v, but got %v", false, got)
	}
}

func TestArrayList_Clear(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Clear()
		})
	}
}

func TestArrayList_Insert(t *testing.T) {
	type args struct {
		index int
		items []interface{}
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Insert(tt.args.index, tt.args.items...)
		})
	}
}

func TestArrayList_Append(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
	}{
		{"append_1", NewArrayList(100), args{[]interface{}{1}}},
		{"append_2", NewArrayList(100), args{[]interface{}{2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Append(tt.args.items...)
			list := NewArrayList(100)
			list.Append(1)
			list.Append(2, 3)
			if got := list.IsEmpty(); got != false {
				t.Errorf("expect %v, but got %v", false, got)
			}
			if got := list.Length(); got != 3 {
				t.Errorf("expect %v, but got %v", false, got)
			}
			if got := list.Get(1); got != 2 {
				t.Errorf("expect %v, but got %v", false, got)
			}
		})
	}
}

func TestArrayList_Remove(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Remove(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayList.Remove() = %v, want %v", got, tt.want)
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Length(); got != tt.want {
				t.Errorf("ArrayList.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_SetValue(t *testing.T) {
	type args struct {
		index int
		item  interface{}
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.SetValue(tt.args.index, tt.args.item)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		a    *ArrayList
		args args
		want interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Get(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayList.Get() = %v, want %v", got, tt.want)
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.isInList(tt.args.index); got != tt.want {
				t.Errorf("ArrayList.isInList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayList_String(t *testing.T) {
	tests := []struct {
		name string
		a    *ArrayList
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("ArrayList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
