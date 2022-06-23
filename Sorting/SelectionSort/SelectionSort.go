package MergeSort

func SelectionSort(slice []int) []int {
	sliceSorted := make([]int, len(slice))
	copy(sliceSorted, slice)

	for idx, value := range sliceSorted {
		smallestValue := value
		smallestValueIdx := idx
		for idx2 := idx + 1; idx2 < len(sliceSorted); idx2++ {
			value2 := sliceSorted[idx2]
			if value2 < smallestValue {
				smallestValue = value2
				smallestValueIdx = idx2
			}
		}
		if idx != smallestValueIdx {
			sliceSorted[smallestValueIdx] = value
			sliceSorted[idx] = smallestValue
		}
	}

	return sliceSorted
}
