package main

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		result []string
	}{
		{
			name:   "Test Case 1",
			n:      3,
			result: []string{"()()()", "()(())", "(())()", "(()())", "((()))"},
		},
		{
			name:   "Test Case 2",
			n:      1,
			result: []string{"()"},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if !reflect.DeepEqual(got, tt.result) {
				t.Errorf("Test case %s failed. Expected %v, but got %v", tt.name, tt.result, got)
			}
		})
	}
}
