package main

import "testing"

func Test_longestConsecutive(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "test2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "test3",
			nums: []int{},
			want: 0,
		},
		{
			name: "test4",
			nums: []int{0},
			want: 1,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
