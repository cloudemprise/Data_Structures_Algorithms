package main

import (
	"math"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{1, 12, -5, -6, 50, 3},
				k:    4,
			},
			want: 12.75,
		},
		{
			name: "test2",
			args: args{
				nums: []int{5},
				k:    1,
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{
				nums: []int{-5, -3},
				k:    2,
			},
			want: -4,
		},
		{
			name: "test4",
			args: args{
				nums: []int{-1},
				k:    1,
			},
			want: -1,
		},
		{
			name: "test5",
			args: args{
				nums: []int{-1},
				k:    2,
			},
			want: math.Inf(-1),
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
