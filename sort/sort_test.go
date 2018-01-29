package sort

import (
	"math"
	"testing"
)

type ints []int

func (a ints) Length() int { return len(a) }

func (a ints) Less(i, j int) bool { return a[i] < a[j] }

func (a ints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type float64s []float64

func (a float64s) Length() int { return len(a) }

func (a float64s) Less(i, j int) bool { return a[i] < a[j] || math.IsNaN(a[i]) && !math.IsNaN(a[j]) }

func (a float64s) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type strings []string

func (a strings) Length() int { return len(a) }

func (a strings) Less(i, j int) bool { return a[i] < a[j] }

func (a strings) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

var intCase = [...]int{1, 0, 0, 0, 654, 12312453, -12312312312, 64, 1563, 65537, -65536}

var float64Case = [...]float64{74.3, 59.0, math.Inf(1), 238.2, -784.0, 2.3, math.NaN(), math.NaN(), math.Inf(-1), 9845.768, -959.7485, 905, 7.8, 7.8}

var stringCase = [...]string{"", "", "", "asd", "qwe", "^&*&^&*", "=-=xc ", " "}

func Test_insertionSort(t *testing.T) {
	data := intCase
	a := ints(data[0:])
	insertionSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", intCase)
		t.Errorf("   got %v", data)
	}
}

func Test_bubbleSort(t *testing.T) {
	data := intCase
	a := ints(data[0:])
	bubbleSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", intCase)
		t.Errorf("   got %v", data)
	}
}

func Test_selectionSort(t *testing.T) {
	data := intCase
	a := ints(data[0:])
	selectionSort(a)
	if !IsSorted(a) {
		t.Errorf("sorted %v", intCase)
		t.Errorf("   got %v", data)
	}
}
