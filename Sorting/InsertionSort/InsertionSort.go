package InsertionSort

func InsertionSort(slice []int) []int {
	sliceSorted := make([]int, 0, len(slice))

	for _, value := range slice {
		if len(sliceSorted) == 0 {
			sliceSorted = append(sliceSorted, value)
		} else {
			inserted := false
			for idx2, value2 := range sliceSorted {
				if value < value2 {
					sliceSorted = insert(sliceSorted, idx2, value)
					inserted = true
					break
				}
			}
			if !inserted {
				sliceSorted = append(sliceSorted, value)
			}
		}

	}

	return sliceSorted
}

func insert(slice []int, index int, value int) []int {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}
