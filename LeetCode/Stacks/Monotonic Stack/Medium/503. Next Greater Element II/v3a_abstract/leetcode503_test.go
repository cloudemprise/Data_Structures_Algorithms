package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		nums   []int
		result []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{[]int{4, 3, 2, 1}, []int{-1, 4, 4, 4}},
	}

	for _, tt := range tests {
		result := nextGreaterElements(tt.nums)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("nextGreaterElements(%v) = %v, want %v", tt.nums, result, tt.result)
		}
	}
}
