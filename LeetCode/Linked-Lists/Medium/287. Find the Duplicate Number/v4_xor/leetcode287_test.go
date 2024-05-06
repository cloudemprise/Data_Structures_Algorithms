package main

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 3, 4, 2, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{3, 1, 3, 4, 2},
			expected: 3,
		},
		{
			name:     "Example 3",
			nums:     []int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 6, 18},
			expected: 18,
		},
		/* { // OUT OF SCOPE
			name:     "Empty Array",
			nums:     []int{},
			expected: -1,
		}, */
		/* { // OUT OF SCOPE
			name:     "Single Element Array",
			nums:     []int{1},
			expected: -1,
		}, */

	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Call the function being tested
			actual := findDuplicate(test.nums)

			// Check if the actual result matches the expected result
			if actual != test.expected {
				t.Errorf("Test case %s failed: expected %d but got %d", test.name, test.expected, actual)
			}
		})
	}
}
