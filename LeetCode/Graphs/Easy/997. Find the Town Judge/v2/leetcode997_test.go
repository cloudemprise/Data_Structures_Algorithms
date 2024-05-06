package main

import "testing"

func Test_findJudge(t *testing.T) {
	type args struct {
		n     int
		trust [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TEST1",
			args: args{
				n:     2,
				trust: [][]int{{1, 2}},
			},
			want: 2,
		},
		{
			name: "TEST2",
			args: args{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
				},
			},
			want: 3,
		},
		{
			name: "TEST3",
			args: args{
				n: 3,
				trust: [][]int{
					{1, 3},
					{2, 3},
					{3, 1},
				},
			},
			want: -1,
		},
		{
			name: "TEST4",
			args: args{
				n: 4,
				trust: [][]int{
					{1, 3},
					{1, 4},
					{2, 3},
					{2, 4},
					{4, 3},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findJudge(tt.args.n, tt.args.trust); got != tt.want {
				t.Errorf("findJudge() = %v, want %v", got, tt.want)
			}
		})
	}
}
