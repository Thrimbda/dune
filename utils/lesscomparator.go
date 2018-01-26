package utils

import (
	"strings"

	. "github.com/Thrimbda/dune"
)

// // inspired by go std package sort, a comparator only return true if i < j.

/*
string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

*/

func LessComparator(i, j interface{}) bool {
	switch i.(type) {
	case string:
		return StringsLessComparator(i, j)
	case int:
		return IntsLessComparator(i, j)
	case int8:
		return Int8sLessComparator(i, j)
	case int16:
		return Int16sLessComparator(i, j)
	case int32:
		return Int32sLessComparator(i, j)
	case int64:
		return Int64sLessComparator(i, j)
	case uint:
		return UintsLessComparator(i, j)
	case uint8:
		return Uint8sLessComparator(i, j)
	case uint16:
		return Uint16sLessComparator(i, j)
	case uint32:
		return Uint32sLessComparator(i, j)
	case uint64:
		return Uint64sLessComparator(i, j)
	case float32:
		return Float32sLessComparator(i, j)
	case float64:
		return Float64sLessComparator(i, j)
	default:
		a := i.(Elem)
		b := j.(Elem)
		return a.LessComparator(b)
	}
}

func StringsLessComparator(i, j interface{}) bool {
	a := i.(string)
	b := j.(string)

	if strings.Compare(a, b) > 0 {
		return false
	}
	return true
}

// TODO: except string, numerical values should be comparable.
//       e.g. comparing 1 < 1.1 shouldn't trigger a panic but return true.

func IntsLessComparator(i, j interface{}) bool {
	a := i.(int)
	b := j.(int)

	return a < b
}

func Int8sLessComparator(i, j interface{}) bool {
	a := i.(int8)
	b := j.(int8)

	return a < b
}

func Int16sLessComparator(i, j interface{}) bool {
	a := i.(int16)
	b := j.(int16)

	return a < b
}

func Int32sLessComparator(i, j interface{}) bool {
	a := i.(int32)
	b := j.(int32)

	return a < b
}

func Int64sLessComparator(i, j interface{}) bool {
	a := i.(int64)
	b := j.(int64)

	return a < b
}

func UintsLessComparator(i, j interface{}) bool {
	a := i.(uint)
	b := j.(uint)

	return a < b
}

func Uint8sLessComparator(i, j interface{}) bool {
	a := i.(uint8)
	b := j.(uint8)

	return a < b
}

func Uint16sLessComparator(i, j interface{}) bool {
	a := i.(uint16)
	b := j.(uint16)

	return a < b
}

func Uint32sLessComparator(i, j interface{}) bool {
	a := i.(uint32)
	b := j.(uint32)

	return a < b
}

func Uint64sLessComparator(i, j interface{}) bool {
	a := i.(uint64)
	b := j.(uint64)

	return a < b
}

func Float32sLessComparator(i, j interface{}) bool {
	a := i.(float32)
	b := j.(float32)

	return a < b
}

func Float64sLessComparator(i, j interface{}) bool {
	a := i.(float64)
	b := j.(float64)

	return a < b
}
