package SequentialSearch

func SequentialSearch(array []int, target int) int {
	for idx, value := range array {
		if value == target {
			return idx
		}
	}

	return -1
}
