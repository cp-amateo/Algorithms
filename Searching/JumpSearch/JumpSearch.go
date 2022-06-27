package jumpSearch

import (
	"algorithms/Searching/SequentialSearch"
	"math"
)

func JumpSearch(array []int, target int) int {
	blockSize := int(math.Sqrt(float64(len(array))))

	for i := 0; i < len(array); i += blockSize {
		value := array[i]
		if value == target {
			return i
		} else if value > target {
			return SequentialSearch.SequentialSearch(array[i-blockSize:i], target)
		}
	}

	return -1
}
