//Package sort is a sort collection implement by Thrimbda.
package sort

//Sortable is that Any type implement this interface can be sorted.
type Sortable interface {
	Length() int
	Less(i, j int) bool
	Swap(i, j int)
}

func insertionSort(data Sortable) {
	for i := 1; i < data.Length(); i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func bubbleSort(data Sortable) {
	for i := 0; i < data.Length()-1; i++ {
		for j := data.Length() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

func selectionSort(data Sortable) {
	for i := 0; i < data.Length()-1; i++ {
		minIndex := i
		for j := data.Length() - 1; j > i; j-- {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}

//IsSorted is a function to check whether a sequence is sorted.
func IsSorted(data Sortable) bool {
	for i := 0; i < data.Length()-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

//TODO: bubble sort, selection sort, quick sork, shell sork, merge sort, heap sort, bin sort, radix sort
