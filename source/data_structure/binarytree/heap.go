package binarytree

//a max-heap implementation.

type Heap struct {
	Heap      []Elem
	size      int
	numInHeap int
}

type PosILLegalError struct {
	errorMsg string
}

func (e PosILLegalError) Error() string {
	return e.errorMsg
}

type HeapFullError struct{}

func (e HeapFullError) Error() string {
	return "heap is full."
}

type HeapEmptyError struct{}

func (e HeapEmptyError) Error() string {
	return "heap is empty."
}

func (h Heap) setup(heap []Elem, num int, max int) {
	h.Heap = heap
	h.numInHeap = num
	h.size = max
	h.buildHeap()
}

func (h Heap) heapSize() int {
	return h.size
}

func (h Heap) isLeaf(pos int) bool {
	return pos >= h.numInHeap/2 && pos < h.numInHeap
}

func (h Heap) buildHeap() {
	for i := h.numInHeap/2 - 1; i >= 0; i-- {
		h.shiftDown(i)
	}
}

func (h Heap) leftChild(pos int) (int, error) {
	if pos >= h.numInHeap/2 {
		return -1, PosILLegalError{"element of given position has no left child"}
	}
	return 2*pos + 1, nil
}

func (h Heap) rightChild(pos int) (int, error) {
	if pos >= (h.numInHeap-1)/2 {
		return -1, PosILLegalError{"element of given position has no right child"}
	}
	return 2*pos + 2, nil
}

func (h Heap) parent(pos int) (int, error) {
	if pos <= 0 {
		return -1, PosILLegalError{"element of given position has no parent"}
	}
	return (pos - 1) / 2, nil
}

func (h Heap) shiftDown(pos int) error {
	if pos < 0 || pos > h.numInHeap {
		return PosILLegalError{"illegal heap position"}
	}
	for !h.isLeaf(pos) {
		j, _ := h.leftChild(pos)
		if j < (h.numInHeap-1) && h.Heap[j].Key() < h.Heap[j+1].Key() {
			j++
		}
		if h.Heap[pos].Key() >= h.Heap[j].Key() {
			return nil
		}
		h.Heap[pos], h.Heap[j] = h.Heap[j], h.Heap[pos]
		pos = j
	}
	return nil
}

func (h Heap) insert(val Elem) error {
	if h.numInHeap >= h.size {
		return HeapFullError{}
	}
	curr := h.numInHeap
	h.numInHeap++
	h.Heap[curr] = val
	parent, _ := h.parent(curr)
	for curr != 0 && h.Heap[curr].Key() > h.Heap[parent].Key() {
		h.Heap[curr], h.Heap[parent] = h.Heap[parent], h.Heap[curr]
		parent, _ = h.parent(curr)
		curr = parent
	}
	return nil
}

func (h Heap) removeMax() (Elem, error) {
	if h.numInHeap <= 0 {
		return nil, HeapEmptyError{}
	}
	h.numInHeap--
	h.Heap[0], h.Heap[h.numInHeap] = h.Heap[h.numInHeap], h.Heap[0]
	if h.numInHeap != 0 {
		h.shiftDown(0)
	}
	return h.Heap[h.numInHeap], nil
}

func (h Heap) remove(pos int) (Elem, error) {
	if pos <= 0 || pos >= h.numInHeap {
		return nil, PosILLegalError{"illegal heap position"}
	}
	h.numInHeap--
	h.Heap[pos], h.Heap[h.numInHeap] = h.Heap[h.numInHeap], h.Heap[pos]
	if h.numInHeap != 0 {
		h.shiftDown(pos)
	}
	return h.Heap[h.numInHeap], nil
}
