package InsertionSort

func InsertionSort(slice []int) []int {
	sliceSorted := make([]int, 0, len(slice))

	for _, value := range slice {
		idxToInsert := findSortedPosition(sliceSorted, value)
		sliceSorted = insert(sliceSorted, idxToInsert, value)
	}

	return sliceSorted
}

func findSortedPosition(slice []int, valueToInsert int) int {
	if len(slice) == 0 {
		return 0
	}

	idxToInsert := len(slice)
	for idx, value := range slice {
		if valueToInsert < value {
			idxToInsert = idx
			break
		}
	}
	return idxToInsert
}

func insert(slice []int, index int, value int) []int {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}
