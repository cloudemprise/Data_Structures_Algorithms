package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 3, 5, 7, 9}, 3, 1},
		//{[]int{1, 3, 5, 7, 9}, 2, -1},
		{[]int{1, 3, 5, 7, 9}, 9, 4},
		//{[]int{1, 3, 5, 7, 9}, 0, -1},
		//{[]int{}, 1, -1},
	}

	for _, tt := range tests {
		got := binarySearch(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("binarySearch(%v, %v) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
