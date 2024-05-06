package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Test case 1",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "Test case 2",
			nums:   []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "Test case 3",
			nums:   []int{1, 2, 3, 4, 5},
			target: 0,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
