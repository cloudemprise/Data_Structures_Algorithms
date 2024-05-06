package main

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{1, 2, 1, 2, 0}, int(^uint(0) >> 1)}, // intentional false positive
	}

	for _, test := range tests {
		result := majorityElement(test.nums)
		if result != test.expected {
			t.Errorf("For nums = %v, got %d but want %d", test.nums, result, test.expected)
		}
	}
}
