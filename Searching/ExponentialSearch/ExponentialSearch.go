package exponentialSearch

import (
	binarySearch "algorithms/Searching/BinarySearch"
)

func ExponentialSearch(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	i := 1
	for i < len(array) && array[i] <= target {
		i = i * 2
	}

	return binarySearch.SearchWithIndex(array, i/2, min(i, len(array)-1), target)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
