package MergeSort

func HeapSort(slice []int) []int {
	heap := buildHeap(slice)

	for length := len(slice); length > 1; length-- {
		slice = removeTop(heap, length)
	}

	return heap
}

func removeTop(heap []int, length int) []int {
	lastIndex := length - 1
	heap[0], heap[lastIndex] = heap[lastIndex], heap[0]
	return heapify(heap, 0, lastIndex)
}

func buildHeap(slice []int) []int {
	for i := len(slice) / 2; i >= 0; i-- {
		heapify(slice, i, len(slice))
	}

	return slice
}

func heapify(slice []int, rootIdx int, length int) []int {
	var max = rootIdx
	l := (rootIdx * 2) + 1
	r := (rootIdx * 2) + 2

	if l < length && slice[l] > slice[max] {
		max = l
	}

	if r < length && slice[r] > slice[max] {
		max = r
	}

	if max != rootIdx {
		slice[rootIdx], slice[max] = slice[max], slice[rootIdx]
		slice = heapify(slice, max, length)
	}

	return slice
}
