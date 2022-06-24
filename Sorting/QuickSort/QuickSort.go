package QuickSort

func QuickSort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	pivotIdx := len(slice) - 1
	idx := 0
	for idx2, value := range slice {
		if value < slice[pivotIdx] {
			swapElements(slice, idx, idx2)
			idx++
		}
	}
	swapElements(slice, idx, pivotIdx)

	QuickSort(slice[:idx])
	QuickSort(slice[idx+1:])

	return slice
}

func swapElements(slice []int, idx int, idx2 int) {
	value := slice[idx]
	slice[idx] = slice[idx2]
	slice[idx2] = value
}
