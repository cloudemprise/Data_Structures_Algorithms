package main

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		temps    []int
		expected []int
	}{
		{
			temps:    []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temps:    []int{30, 40, 50, 60},
			expected: []int{1, 1, 1, 0},
		},
		{
			temps:    []int{30, 60, 90},
			expected: []int{1, 1, 0},
		},
	}

	for _, tc := range testCases {
		result := dailyTemperatures(tc.temps)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test failed for temps=%v, expected=%v, got=%v", tc.temps, tc.expected, result)
		}
	}
}
