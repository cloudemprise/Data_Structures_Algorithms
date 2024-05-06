package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Test1:",
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},

		// Add more test cases as needed
	}

	for _, test := range tests {
		result := subsetsWithDup(test.nums)

		// Check the length of the result
		if len(result) != len(test.expected) {
			t.Errorf("%s expected %d combinations, got %d", test.name, len(test.expected), len(result))
		}

		// Check the values of the combinations
		for i, combination := range result {
			if !reflect.DeepEqual(combination, test.expected[i]) {
				t.Errorf("%s expected %v, got %v", test.name, test.expected[i], combination)
			}
		}
	}
}
