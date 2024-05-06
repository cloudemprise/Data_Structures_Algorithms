package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3, 2},
				k:    2,
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				nums: []int{99, 99},
				k:    2,
			},
			want: true,
		},
		{
			name: "test6",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
				k:    3,
			},
			want: true,
		},
		{
			name: "test7",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				k:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
