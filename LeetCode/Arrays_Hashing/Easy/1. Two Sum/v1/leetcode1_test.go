package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{nums: []int{1, 2, 3}, target: 7, expected: []int{}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For %v and target %d, expected %v but got %v", test.nums, test.target, test.expected, result)
		}
	}
}
