package list

type ArrayStack struct {
	size int
	top int
	listArray []Elem
}

func (a ArrayStack) setup(size int) {
	a.size = size
	a.top = 0
	a.listArray = make([] Elem, size)
}

func (a ArrayStack) clear() {
	a.top = 0
}

func (a ArrayStack) push(item Elem) error {
	if a.top >= a.size {
		return FullListError{}
	}
	a.listArray[a.top] = item
	a.top++
	return nil
}

func (a ArrayStack) pop() (Elem, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	a.top--
	value := a.listArray[a.top]
	return value, nil
}

func (a ArrayStack) topValue() (Elem, error) {
	if a.isEmpty() {
		return nil, EmptyListError{}
	}
	return a.listArray[a.top - 1], nil
}

func (a ArrayStack) isEmpty() bool {
	return a.top == 0
}