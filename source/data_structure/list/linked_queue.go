package list

type LinkedQueue struct {
	front *LinkNode
	rear *LinkNode
}

func (l LinkedQueue) setup() {
	l.front = nil
	l.rear = nil
}

func (l LinkedQueue) clear() {
	l.setup()
}

func (l LinkedQueue) enqueue(item Elem) {
	if l.isEmpty() {
		l.rear = &LinkNode{item, l.rear, nil}
		l.front = l.rear
	} else {
		l.rear.next = &LinkNode{item, l.rear, nil}
		l.rear = l.rear.next	
	}
}

func (l LinkedQueue) dequeue() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	value := l.front.value
	l.front = l.front.next
	if l.front == nil {
		l.rear = nil
	} else {
		l.front.prev = nil	
	}
	return value, nil
}

func (l LinkedQueue) firstValue() (Elem, error) {
	if l.isEmpty() {
		return nil, EmptyListError{}
	}
	return l.front.value, nil
}

func (l LinkedQueue) isEmpty() bool {
	return l.front == l.rear
}