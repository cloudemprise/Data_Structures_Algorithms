package main

import (
	"testing"

	"github.com/cloudemprise/leetcode110/tree"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		name   string
		root   *tree.TreeNode
		expect bool
	}{
		{
			name:   "Example 1",
			root:   &tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 9}, Right: &tree.TreeNode{Val: 20, Left: &tree.TreeNode{Val: 15}, Right: &tree.TreeNode{Val: 7}}},
			expect: true,
		},
		{
			name:   "Example 2",
			root:   &tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 4}, Right: &tree.TreeNode{Val: 4}}, Right: &tree.TreeNode{Val: 3}}, Right: &tree.TreeNode{Val: 2}},
			expect: false,
		},
		{
			name:   "Empty tree",
			root:   nil,
			expect: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isBalanced(tc.root)
			if result != tc.expect {
				t.Errorf("Expected %v but got %v", tc.expect, result)
			}
		})
	}
}
