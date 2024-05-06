package main

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {

	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Test1:",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected:   [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
		{
			name:       "Test2:",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected:   [][]int{{1, 2, 2}, {5}},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := combinationSum2(test.candidates, test.target)

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
