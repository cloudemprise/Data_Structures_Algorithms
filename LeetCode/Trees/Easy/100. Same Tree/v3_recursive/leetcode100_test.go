package main

import (
	"testing"

	"github.com/cloudemprise/leetcode100/tree"
)

func TestIsSameTree(t *testing.T) {
	testCases := []struct {
		p        *tree.TreeNode
		q        *tree.TreeNode
		expected bool
	}{
		{
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
			},
			q: &tree.TreeNode{
				Val: 1,
			},
			expected: true,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
			},
			q: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: false,
		},
		{
			p: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			q: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
				},
				Right: &tree.TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
	}

	for _, tc := range testCases {
		actual := isSameTree(tc.p, tc.q)
		if actual != tc.expected {
			t.Errorf("isSameTree(%v, %v) = %v; expected %v", tc.p, tc.q, actual, tc.expected)
		}
	}
}
