package heapSort

import "algorithms/DataStructures/Heap"

func HeapSort(slice []int) []int {
	heap := Heap.NewMaxHeap(slice)
	heap.BuildHeap()

	for length := len(slice); length > 1; length-- {
		heap.RemoveTop(length)
	}

	return heap.Heap
}
