package Heap

type MaxHeap struct {
	Heap []int
}

func NewMaxHeap(slice []int) *MaxHeap {
	copySlice := make([]int, len(slice))
	copy(copySlice, slice)
	return &MaxHeap{copySlice}
}

func (h *MaxHeap) BuildHeap() {
	for i := len(h.Heap) / 2; i >= 0; i-- {
		h.Heapify(i, len(h.Heap))
	}
}

func (h *MaxHeap) Heapify(rootIdx, length int) {
	var max = rootIdx
	l := (rootIdx * 2) + 1
	r := (rootIdx * 2) + 2

	if l < length && h.Heap[l] > h.Heap[max] {
		max = l
	}
	if r < length && h.Heap[r] > h.Heap[max] {
		max = r
	}

	if max != rootIdx {
		h.Heap[rootIdx], h.Heap[max] = h.Heap[max], h.Heap[rootIdx]
		h.Heapify(max, length)
	}
}

func (h *MaxHeap) RemoveTop(length int) {
	lastIndex := length - 1
	h.Heap[0], h.Heap[lastIndex] = h.Heap[lastIndex], h.Heap[0]
	h.Heapify(0, lastIndex)
}
