package SelectionSort

import (
	"reflect"
	"testing"
)

func Test_insertion_sort(t *testing.T) {
	tests := []struct {
		name     string
		arg      []int
		wantList []int
	}{
		{"empty array", []int{}, []int{}},
		{"array with one element", []int{1}, []int{1}},
		{"array with two elements", []int{2, 1}, []int{1, 2}},
		{"array with three elements", []int{3, 2, 1}, []int{1, 2, 3}},
		{"array with repeated elements",
			[]int{7, 3, 8, 2, 2, 5, 20, 5, 4, 1, 9, 6},
			[]int{1, 2, 2, 3, 4, 5, 5, 6, 7, 8, 9, 20}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortListResult := SelectionSort(test.arg)
			if !reflect.DeepEqual(test.wantList, sortListResult) {
				t.Errorf("MergeSort: got = %v want = %v", sortListResult, test.wantList)
			}
		})

	}
}
