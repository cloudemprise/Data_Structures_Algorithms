package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{5, 4, 3, 2, 1}, 5, 0},
		{[]int{1, 3, 5, 7, 9}, 2, -1},
		{[]int{2, 4, 6, 8, 10}, 10, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, test := range tests {
		got := search(test.nums, test.target)
		if got != test.want {
			t.Errorf("search(%v, %d) = %d, want %d", test.nums, test.target, got, test.want)
		}
	}
}
