package main

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		nums1  []int
		nums2  []int
		result []int
	}{
		{[]int{4, 1, 2}, []int{1, 3, 4, 2}, []int{-1, 3, -1}},
		{[]int{2, 4}, []int{1, 2, 3, 4}, []int{3, -1}},
		{[]int{1, 3, 5, 2, 4}, []int{6, 5, 4, 3, 2, 1, 7}, []int{7, 7, 7, 7, 7}},
	}

	for _, tt := range tests {
		result := nextGreaterElement(tt.nums1, tt.nums2)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("nextGreaterElement(%v, %v) = %v, want %v", tt.nums1, tt.nums2, result, tt.result)
		}
	}
}
