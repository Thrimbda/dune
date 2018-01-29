package list

//List interface which doesn't support resize yet.
type List interface {
	Insert(index int, items ...interface{})
	Append(items ...interface{})
	Remove(index int) interface{}
	Length() int
	Get(index int) interface{}
	Set(index int, value interface{})
	Contains(value interface{}) bool
	IndexOf(value interface{}) int
	Clear()
	IsEmpty() bool
	Values() []interface{}
}
