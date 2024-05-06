package main

import "testing"

func Test_lomutoPartition(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{2, 3, 4, 1},
			want: 0, // [1, 3, 4, 2]
		},
		{
			name: "test2",
			nums: []int{4, 2, 1, 3},
			want: 2, // [2, 1, 3, 4]
		},
		{
			name: "test3",
			nums: []int{1, 3, 2, 4},
			want: 3, // [1, 3, 2, 4]
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := lomutoPartition(tt.nums, 0, len(tt.nums)-1); got != tt.want {
				t.Errorf("lomutoPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
