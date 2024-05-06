package main

import (
	"reflect"
	"slices"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{nums: []int{1, 1, 1, 2, 2, 3}, k: 2}, want: []int{1, 2}},
		{name: "test2", args: args{nums: []int{1}, k: 1}, want: []int{1}},
		{name: "test3", args: args{nums: []int{1, 2}, k: 2}, want: []int{1, 2}},
		{
			name: "test4",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
				k:    10,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !isEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Custom comparison function to handle unordered slices.
func isEqual(a, b []int) bool {
	slices.Sort(a)
	slices.Sort(b)
	return reflect.DeepEqual(a, b)
}
