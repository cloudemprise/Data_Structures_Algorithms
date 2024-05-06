package main

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		target   int
		expected []int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			target:   5,
			expected: []int{1, 1},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			target:   10,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		result := binarySearch2D(tt.matrix, tt.target)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("binarySearch2D(%v, %d) = %v, expected %v", tt.matrix, tt.target, result, tt.expected)
		}
	}
}
