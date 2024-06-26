package main

import "testing"

func Test_removeDuplicates(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			nums: []int{1, 1, 2},
			want: 2,
		}, // [1,2,_]
		{
			name: "test2",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		}, //[0,1,2,3,4,_,_,_,_,_]
		{
			name: "test3",
			nums: []int{1, 1},
			want: 1,
		}, //[1,_]
		{
			name: "test4",
			nums: []int{1},
			want: 1,
		}, //[1]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
