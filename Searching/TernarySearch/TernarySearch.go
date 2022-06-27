package ternarySearch

func TernarySearch(array []int, target int) int {
	return searchInArray(array, 0, len(array)-1, target)
}

func searchInArray(array []int, start int, end int, target int) int {
	if start > end ||
		len(array) == 0 {
		return -1
	}

	mid1 := start + (end-start)/3
	mid2 := end - (end-start)/3

	mid1Value := array[mid1]
	mid2Value := array[mid2]

	if target == mid1Value {
		return mid1
	} else if target == mid2Value {
		return mid2
	} else if target < mid1Value {
		return searchInArray(array, start, mid1-1, target)
	} else if target > mid2Value {
		return searchInArray(array, mid2+1, end, target)
	} else {
		return searchInArray(array, mid1+1, mid2-1, target)
	}
}
