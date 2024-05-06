package main

import (
	"reflect"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	tests := []struct {
		prices []int
		result []int
	}{
		{[]int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 1, 1, 6}, []int{9, 0, 1, 6}},
		{[]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1}, []int{1, 3, 2, 1, 7, 0, 0, 6, 9, 1}},
	}

	for _, tt := range tests {
		result := finalPrices(tt.prices)
		if !reflect.DeepEqual(result, tt.result) {
			t.Errorf("finalPrices(%v) = %v, want %v", tt.prices, result, tt.result)
		}
	}
}
