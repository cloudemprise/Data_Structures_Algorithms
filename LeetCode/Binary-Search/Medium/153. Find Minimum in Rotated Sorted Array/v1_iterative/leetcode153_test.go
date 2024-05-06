package main

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TEST1", // middle element
			args: args{
				nums: []int{4, 5, 1, 2, 3},
			},
			want: 1,
		},
		{
			name: "TEST2", // first element
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 1,
		},
		{
			name: "TEST3", // last element
			args: args{
				nums: []int{2, 3, 4, 5, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
