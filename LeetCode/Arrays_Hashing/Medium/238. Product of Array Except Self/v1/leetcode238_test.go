package main

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// Add test cases:
		{
			name: "test1",
			args: args{[]int{1, 2, 3, 4}},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "test2",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "test3",
			args: args{[]int{1}},
			want: []int{1},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
