package arraystack

import (
	"reflect"
	"testing"

	"github.com/Thrimbda/dune/list/arraylist"
)

func TestNewArrayStack(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *ArrayStack
	}{
		{"new1", args{3}, &ArrayStack{arraylist.NewArrayList(3)}},
		{"new2", args{23}, &ArrayStack{arraylist.NewArrayList(23)}},
		{"new_empty", args{0}, &ArrayStack{arraylist.NewArrayList(0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewArrayStack(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewArrayStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToArrayStack(t *testing.T) {
	type args struct {
		size  int
		items []interface{}
	}
	tests := []struct {
		name string
		args args
		want *ArrayStack
	}{
		{"new", args{3, []interface{}{1, 2, 3}}, &ArrayStack{arraylist.ConvertToArrayList(3, 1, 2, 3)}},
		{"huge", args{100, make([]interface{}, 100)}, &ArrayStack{arraylist.ConvertToArrayList(100, make([]interface{}, 100)...)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToArrayStack(tt.args.size, tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToArrayStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Clear(t *testing.T) {
	type fields struct {
		list *arraylist.ArrayList
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"empty1", fields{arraylist.ConvertToArrayList(4, 1, 2, 3, 4)}, true},
		{"empty2", fields{arraylist.ConvertToArrayList(100, make([]interface{}, 100)...)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack{
				list: tt.fields.list,
			}
			a.Clear()
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf("Clear failed, got: %v", got)
			}
		})
	}
}

func TestArrayStack_Push(t *testing.T) {
	type fields struct {
		list *arraylist.ArrayList
	}
	type args struct {
		item interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"empty_push", fields{arraylist.NewArrayList(3)}, args{3}, 3},
		{"push1", fields{arraylist.ConvertToArrayList(4, 1, 2, 3, 4)}, args{5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack{
				list: tt.fields.list,
			}
			a.Push(tt.args.item)
			if got := a.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Push failed got top value: %v", got)
			}
		})
	}
}

func TestArrayStack_Pop(t *testing.T) {
	type fields struct {
		list *arraylist.ArrayList
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"pop1", fields{arraylist.ConvertToArrayList(3, 1, 2, 3)}, 3},
		{"pop2", fields{arraylist.ConvertToArrayList(100, make([]interface{}, 100)...)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack{
				list: tt.fields.list,
			}
			if got := a.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Peek(t *testing.T) {
	type fields struct {
		list *arraylist.ArrayList
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{"peek1", fields{arraylist.ConvertToArrayList(4, 1, 2, 3)}, 3},
		{"peek2", fields{arraylist.ConvertToArrayList(100, make([]interface{}, 100)...)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack{
				list: tt.fields.list,
			}
			if got := a.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayStack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_IsEmpty(t *testing.T) {
	type fields struct {
		list *arraylist.ArrayList
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"empty", fields{arraylist.NewArrayList(10)}, true},
		{"empty", fields{arraylist.NewArrayList(0)}, true},
		{"not_empty", fields{arraylist.ConvertToArrayList(2, 1, 2)}, false},
		{"not_empty", fields{arraylist.ConvertToArrayList(2)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack{
				list: tt.fields.list,
			}
			if got := a.IsEmpty(); got != tt.want {
				t.Errorf("ArrayStack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
