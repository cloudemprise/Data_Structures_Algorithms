package main

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
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
				s: "ABAB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: 4,
		},
		{
			name: "test3",
			args: args{
				s: "ABAA",
				k: 0,
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				s: "ABBB",
				k: 2,
			},
			want: 4,
		},
		{
			name: "test5",
			args: args{
				s: "BAAAB",
				k: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
