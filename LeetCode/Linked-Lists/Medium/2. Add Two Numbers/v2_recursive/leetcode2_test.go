package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	// Define the test cases as a table
	testCases := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			l1:       &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
		{
			l1:       &ListNode{Val: 0},
			l2:       &ListNode{Val: 0},
			expected: &ListNode{Val: 0},
		},
		{
			l1:       &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}},
			l2:       &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			expected: &ListNode{Val: 6, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}},
		},
		{
			l1:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}},
			l2:       &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}},
			expected: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}},
		},
		// Add more test cases here if needed
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		// Call the addTwoNumbers function
		result := addTwoNumbers(tc.l1, tc.l2)

		// Compare the result with the expected result
		for result != nil && tc.expected != nil {
			if result.Val != tc.expected.Val {
				t.Errorf("Expected %d, but got %d", tc.expected.Val, result.Val)
			}

			result = result.Next
			tc.expected = tc.expected.Next
		}

		// Check if the result and expected result have different lengths
		if result != nil || tc.expected != nil {
			t.Errorf("Result and expected result have different lengths")
		}
	}
}
