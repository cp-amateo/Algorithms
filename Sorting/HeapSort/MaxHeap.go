package HeapSort

type MaxHeap struct {
	heap []int
}

func NewMaxHeap(slice []int) *MaxHeap {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return &MaxHeap{copySlice}
}

func (h *MaxHeap) BuildHeap() {
	for i := len(h.heap) / 2; i >= 0; i-- {
		h.Heapify(i, len(h.heap))
	}
}

func (h *MaxHeap) Heapify(rootIdx, length int) {
	var max = rootIdx
	l := (rootIdx * 2) + 1
	r := (rootIdx * 2) + 2

	if l < length && h.heap[l] > h.heap[max] {
		max = l
	}
	if r < length && h.heap[r] > h.heap[max] {
		max = r
	}

	if max != rootIdx {
		h.heap[rootIdx], h.heap[max] = h.heap[max], h.heap[rootIdx]
		h.Heapify(max, length)
	}
}

func (h *MaxHeap) RemoveTop(length int) {
	lastIndex := length - 1
	h.heap[0], h.heap[lastIndex] = h.heap[lastIndex], h.heap[0]
	h.Heapify(0, lastIndex)
}
