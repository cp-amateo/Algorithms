package binarySearch

func BinarySearch(array []int, target int) int {
	return searchInArray(array, 0, len(array)-1, target)
}

func SearchWithIndex(array []int, start int, end int, target int) int {
	return searchInArray(array, start, end, target)
}

func searchInArray(array []int, start int, end int, target int) int {
	if start > end ||
		len(array) == 0 {
		return -1
	}

	keyIdx := (end + start) / 2
	keyValue := array[keyIdx]

	if keyValue == target {
		return keyIdx
	} else if target > keyValue {
		return searchInArray(array, keyIdx+1, end, target)
	} else {
		return searchInArray(array, start, keyIdx-1, target)
	}

}
