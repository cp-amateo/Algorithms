package ternarySearch

import (
	"testing"
)

func Test_ternary_search(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{"empty array", []int{}, 5, -1},
		{"array with one element", []int{1}, 1, 0},
		{"array with two elements", []int{2, 5}, 5, 1},
		{"not found", []int{1, 2, 3}, 5, -1},
		{"find at the beginning", []int{1, 2, 3}, 3, 2},
		{"find at the end", []int{1, 2, 3}, 1, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TernarySearch(test.array, test.target)
			if result != test.want {
				t.Errorf("Search: got = %d want = %d", result, test.want)
			}
		})

	}
}
