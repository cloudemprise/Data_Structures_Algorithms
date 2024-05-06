package main

import "testing"

func Test_isPalindrome(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		// TODO: Add test cases.
		{"test1", "A man, a plan, a canal: Panama", true},
		{"test2", "race a car", false},
		{"test3", " ", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
