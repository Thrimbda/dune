package sort

import (
	"fmt"
	"testing"
)

type num struct {
	key int
}
type nums []num

func (a nums) Len() int {
	return len(a)
}

func (a nums) Less(i, j int) bool {
	return a[i].key < a[j].key
}

func (a nums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		array nums
	}
	tests := []struct {
		name string
		args args
	}{
		{"sorted array", args{[]num{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}},
		{"anti sorted array", args{[]num{{9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}, {0}}}},
		{"random array", args{[]num{{6}, {2}, {8}, {3}, {9}, {7}, {6}, {1}, {5}, {0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.array)
			fmt.Println(tt.args.array)
		})
	}
}
