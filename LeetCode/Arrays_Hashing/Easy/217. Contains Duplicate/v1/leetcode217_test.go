package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		result := containsDuplicate(tt.input)
		if result != tt.expected {
			t.Errorf("Expected %t, but got %t for input %v", tt.expected, result, tt.input)
		}
	}
}
