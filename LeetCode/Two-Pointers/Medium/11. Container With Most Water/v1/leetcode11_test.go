package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		h []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				h: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "test2",
			args: args{
				h: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				h: []int{4, 3, 2, 1, 4},
			},
			want: 16,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := maxArea(tt.args.h); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
