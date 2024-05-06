package main

import (
	"reflect"
	"testing"
)

func TestGenerateAllSubsets(t *testing.T) {

	tests := []struct {
		name     string
		elements []any
		expected [][]any
	}{
		{
			name:     "Test1:",
			elements: []any{'a', 'b', 'c'},
			expected: [][]any{
				{},
				{'a'},
				{'b'},
				{'a', 'b'},
				{'c'},
				{'a', 'c'},
				{'b', 'c'},
				{'a', 'b', 'c'},
			},
		},

		// Add more test cases as needed
	}

	for _, test := range tests {
		result := generateAllSubsets(test.elements)

		// Check the length of the result
		if len(result) != len(test.expected) {
			t.Errorf("%s expected %q combinations, got %q", test.name, len(test.expected), len(result))
		}

		// Check the values of the combinations
		for i, combination := range result {
			if !reflect.DeepEqual(combination, test.expected[i]) {
				t.Errorf("%s expected %q, got %q", test.name, test.expected[i], combination)
			}
		}
	}
}
