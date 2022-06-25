package HeapSort

func HeapSort(slice []int) []int {
	heap := NewMaxHeap(slice)
	heap.BuildHeap()

	for length := len(slice); length > 1; length-- {
		heap.RemoveTop(length)
	}

	return heap.heap
}
