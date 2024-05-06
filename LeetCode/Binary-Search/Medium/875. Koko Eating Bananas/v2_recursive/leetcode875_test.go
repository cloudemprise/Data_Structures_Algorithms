package main

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		piles []int
		h     int
		want  int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
	}

	for _, tt := range tests {
		got := minEatingSpeed(tt.piles, tt.h)
		if got != tt.want {
			t.Errorf("minEatingSpeed(%v, %v) = %v, want %v", tt.piles, tt.h, got, tt.want)
		}
	}
}
