package main

import "testing"

func TestCarFleet(t *testing.T) {
	testCases := []struct {
		target   int
		position []int
		speed    []int
		expected int
	}{
		{12, []int{10, 8, 0, 5}, []int{2, 4, 1, 1}, 3},
		{20, []int{0, 10, 15, 5}, []int{5, 2, 1, 3}, 1},
		{30, []int{5, 10, 15, 20}, []int{1, 2, 3, 4}, 4},
	}

	for _, tc := range testCases {
		result := carFleet(tc.target, tc.position, tc.speed)
		if result != tc.expected {
			t.Errorf("carFleet(%d, %v, %v) = %d; expected %d", tc.target, tc.position, tc.speed, result, tc.expected)
		}
	}
}
