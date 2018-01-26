package utils

import (
	"testing"

	. "github.com/Thrimbda/dune"
)

func TestStringsLessComparator(t *testing.T) {
	data := [][]interface{}{
		{"a", "b", true},
		{"a", "A", false},
		{"0", "a", true},
		{" ", "9", true},
		{"asdasd", "v", true},
	}
	for i, testCase := range data {
		got := StringsLessComparator(testCase[0], testCase[1])
		expect := testCase[2]
		if got != expect {
			t.Errorf("in test case %v, expect %v, but got %v", i, expect, got)
		}
	}
}

func TestIntsLessComparator(t *testing.T) {
	data := [][]interface{}{
		{1, 2, true},
		{0, 1, true},
		{-1, 1, true},
		{0, 0, false},
		{12345, 54321, true},
	}
	for i, testCase := range data {
		got := IntsLessComparator(testCase[0], testCase[1])
		expect := testCase[2]
		if got != expect {
			t.Errorf("in test case %v, expect %v, but got %v", i, expect, got)
		}
	}
}

func TestFloat64sLessComparator(t *testing.T) {
	data := [][]interface{}{
		{1.1, 2.2, true},
		{0.0, 1.0, true},
		{-1.1, 1.2, true},
		{0.0, 0.0, false},
		{1.123123, 1.1231234, true},
	}
	for i, testCase := range data {
		got := Float64sLessComparator(testCase[0], testCase[1])
		expect := testCase[2]
		if got != expect {
			t.Errorf("in test case %v, expect %v, but got %v", i, expect, got)
		}
	}
}

type City struct {
	id int
}

func (c *City) LessComparator(b Elem) bool {
	return c.id < b.(*City).id
}
func (c *City) String() string {
	return string(c.id)
}

func TestLessComparator(t *testing.T) {
	data := [][]interface{}{
		{1.1, 2.2, true},
		{"s", "A", false},
		{&City{12}, &City{1}, false},
		{&City{0}, &City{15}, true},
		{&City{3}, &City{3}, false},
		{0, 0, false},
		{1.123123, 1.1231234, true},
	}
	for i, testCase := range data {
		got := LessComparator(testCase[0], testCase[1])
		expect := testCase[2]
		if got != expect {
			t.Errorf("in test case %v, expect %v, but got %v", i, expect, got)
		}
	}
}
