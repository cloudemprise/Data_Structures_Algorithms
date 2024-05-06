package main

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		target int
		output bool
	}{
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			3,
			true,
		},
		{
			[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			13,
			false,
		},
		{
			[][]int{{1}, {3}},
			13,
			false,
		},
		{
			[][]int{{}},
			1,
			false,
		},
	}

	for _, test := range tests {
		result := searchMatrix(test.matrix, test.target)
		if result != test.output {
			t.Errorf("searchMatrix(%v, %d) = %t, expected %t", test.matrix, test.target, result, test.output)
		}
	}
}
