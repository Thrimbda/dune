package sort

import (
	"testing"
)

type nums []int

func (a nums) Len() int {
	return len(a)
}

func (a nums) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a nums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		data nums
	}
	tests := []struct {
		name string
		args args
	}{
		{"sorted array", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{"anti sorted array", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}},
		{"random array", args{[]int{6, 2, 8, 3, 9, 7, 6, 1, 5, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.args.data
			insertionSort(data)
			if !IsSorted(data) {
				t.Errorf("sorted %v", tt.args.data)
				t.Errorf("   got %v", data)
			}
		})
	}
}

func Test_bubbleSort(t *testing.T) {
	type args struct {
		data nums
	}
	tests := []struct {
		name string
		args args
	}{
		{"sorted array", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{"anti sorted array", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}},
		{"random array", args{[]int{6, 2, 8, 3, 9, 7, 6, 1, 5, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.args.data
			bubbleSort(data)
			if !IsSorted(data) {
				t.Errorf("sorted %v", tt.args.data)
				t.Errorf("   got %v", data)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		data nums
	}
	tests := []struct {
		name string
		args args
	}{
		{"sorted array", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{"anti sorted array", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}}},
		{"random array", args{[]int{6, 2, 8, 3, 9, 7, 6, 1, 5, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.args.data
			selectionSort(data)
			if !IsSorted(data) {
				t.Errorf("sorted %v", tt.args.data)
				t.Errorf("   got %v", data)
			}
		})
	}
}
