package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TEST1", // test middle element
			args: args{
				nums:   []int{4, 5, 1, 2, 3},
				target: 1,
			},
			want: 2,
		},
		{
			name: "TEST2", // test first element
			args: args{
				nums:   []int{4, 5, 1, 2, 3},
				target: 4,
			},
			want: 0,
		},
		{
			name: "TEST3", // test last element
			args: args{
				nums:   []int{4, 5, 1, 2, 3},
				target: 3,
			},
			want: 4,
		},
		{
			name: "TEST4", // test absent middle element
			args: args{
				nums:   []int{8, 10, 2, 4, 6},
				target: 1,
			},
			want: -1,
		},
		{
			name: "TEST5", // test absent lower element
			args: args{
				nums:   []int{12, 15, 3, 6, 9},
				target: 11,
			},
			want: -1,
		},
		{
			name: "TEST6", // test absent upper element
			args: args{
				nums:   []int{12, 15, 3, 6, 9},
				target: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
