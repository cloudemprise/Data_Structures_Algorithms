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
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Test2:",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Test3:",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Test4:",
			candidates: []int{8, 2, 6, 3},
			target:     13,
			expected:   [][]int{{8, 2, 3}, {2, 2, 2, 2, 2, 3}, {2, 2, 6, 3}, {2, 2, 3, 3, 3}},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := combinationSum(test.candidates, test.target)

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
