package utils

type Elem interface {
	LessComparator(b Elem) bool
	//return true if a < b
}
