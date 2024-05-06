package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test1",
			args: args{strs: []string{"a"}},
			want: [][]string{{"a"}},
		},
		{
			name: "test2",
			args: args{strs: []string{""}},
			want: [][]string{{""}},
		},
		{
			name: "test3",
			args: args{strs: []string{"a", "b", "c"}},
			want: [][]string{{"a"}, {"b"}, {"c"}},
		},
		{
			name: "test4",
			args: args{strs: []string{"a", "a", "a"}},
			want: [][]string{{"a", "a", "a"}},
		},
		{
			name: "test5",
			args: args{strs: []string{"a", "a", "b"}},
			want: [][]string{{"a", "a"}, {"b"}},
		},
		{
			name: "test6",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "test7",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"nat", "tan"}, {"bat"}, {"ate", "eat", "tea"}},
		},
		{
			name: "test8",
			args: args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			want: [][]string{{"nat", "tan"}, {"ate", "eat", "tea"}, {"bat"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !isEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Custom comparison function to handle nested objects with unordered inner objects
func isEqual(a, b [][]string) bool {
	// Sort the inner slices before comparison
	sortSlice := func(s [][]string) {
		for i := range s {
			sort.Strings(s[i])
		}
		sort.Slice(s, func(i, j int) bool {
			return s[i][0] < s[j][0]
		})
	}
	sortSlice(a)
	sortSlice(b)
	return reflect.DeepEqual(a, b)
}
