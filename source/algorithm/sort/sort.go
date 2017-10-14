//Package sort is a sort collection implement by Thrimbda.
package sort

//Interface is that Any type implement this interface can be sorted.
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func insertionSort(array Interface) {
	for i := 1; i < array.Len(); i++ {
		for j := i; j > 0 && array.Less(j, j-1); j-- {
			array.Swap(j, j-1)
		}
	}
}
