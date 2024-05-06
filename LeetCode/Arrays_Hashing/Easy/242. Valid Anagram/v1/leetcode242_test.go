package main

import "testing"

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"ab", "a", false},
	}

	for _, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		if result != tc.expected {
			t.Errorf("Expected %t, but got %t for input (%s, %s)", tc.expected, result, tc.s, tc.t)
		}
	}
}
