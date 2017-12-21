package list

type List interface {
	Insert(index int, items ...interface{})
	Append(items ...interface{})
	Remove(index int) interface{}
	Length() int
	Get(index int) interface{}
	SetValue(index int, value interface{})
	Contains(value interface{}) bool
	IndexOf(value interface{}) int
	Clear()
	IsEmpty() bool
}
