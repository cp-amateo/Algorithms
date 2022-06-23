package MergeSort

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	middle := len(list) / 2

	sortedListL := MergeSort(list[:middle])
	sortedListR := MergeSort(list[middle:])

	return merge(sortedListL, sortedListR)
}

func merge(leftList []int, rightList []int) []int {
	mergedList := make([]int, len(leftList)+len(rightList))

	leftIdx := 0
	rightIdx := 0
	mergeIdx := 0
	for isBoundIndex(leftIdx, leftList) || isBoundIndex(rightIdx, rightList) {
		if isBoundIndex(leftIdx, leftList) &&
			(!isBoundIndex(rightIdx, rightList) ||
				leftList[leftIdx] < rightList[rightIdx]) {
			mergedList[mergeIdx] = leftList[leftIdx]
			leftIdx += 1
		} else {
			mergedList[mergeIdx] = rightList[rightIdx]
			rightIdx += 1
		}

		mergeIdx += 1
	}

	return mergedList
}

func isBoundIndex(idx int, array []int) bool {
	return idx < len(array)
}
