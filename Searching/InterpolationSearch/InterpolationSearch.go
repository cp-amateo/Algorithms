package interpolationSearch

func InterpolationSearch(array []int, target int) int {
	return searchInArray(array, 0, len(array)-1, target)
}

func searchInArray(array []int, start int, end int, target int) int {
	if start <= end &&
		target >= array[start] &&
		target <= array[end] {

		partitionIdx := calculatePartitionIdx(array, start, end, target)
		partitionValue := array[partitionIdx]

		if partitionValue == target {
			return partitionIdx
		} else if target > partitionValue {
			return searchInArray(array, partitionIdx+1, end, target)
		} else {
			return searchInArray(array, start, partitionIdx-1, target)
		}
	}
	return -1
}

func calculatePartitionIdx(array []int, start int, end int, target int) int {
	partitionIdx := start
	if end > start {
		//calculate the partition position with linear interpolation method
		partitionIdx = start +
			((end-start)/(array[end]-array[start]))*
				(target-array[start])
	}
	return partitionIdx
}
