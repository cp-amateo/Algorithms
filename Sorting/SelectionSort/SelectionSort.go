package SelectionSort

func SelectionSort(slice []int) []int {
	sliceSorted := make([]int, len(slice))
	copy(sliceSorted, slice)

	for idx, value := range sliceSorted {
		smallestValueIdx, smallestValue := findSmallestValue(idx, sliceSorted)

		if idx != smallestValueIdx {
			sliceSorted[smallestValueIdx] = value
			sliceSorted[idx] = smallestValue
		}
	}

	return sliceSorted
}

func findSmallestValue(initialIdx int, slice []int) (int, int) {
	smallestValue := slice[initialIdx]
	smallestValueIdx := initialIdx

	for idx := initialIdx + 1; idx < len(slice); idx++ {
		value := slice[idx]
		if value < smallestValue {
			smallestValue = value
			smallestValueIdx = idx
		}
	}

	return smallestValueIdx, smallestValue
}
