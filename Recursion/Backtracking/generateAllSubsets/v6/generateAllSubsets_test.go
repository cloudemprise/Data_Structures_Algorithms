package main

import (
	"reflect"
	"testing"
)

func TestGenerateAllSubsets(t *testing.T) {

	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Test1:",
			nums: []int{1, 2, 3},
			expected: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := generateAllSubsets(test.nums)

		// Check the length of the result
		if len(result) != len(test.expected) {
			t.Errorf("%s expected %d combinations, got %d", test.name, len(test.expected), len(result))
		}

		// Check the values of the combinations
		for i, combination := range result {
			if !reflect.DeepEqual(combination, test.expected[i]) {
				t.Errorf("%s expected %d, got %d", test.name, test.expected[i], combination)
			}
		}
	}
}
