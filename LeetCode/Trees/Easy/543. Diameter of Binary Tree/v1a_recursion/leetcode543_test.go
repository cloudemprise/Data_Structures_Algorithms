package main

import (
	"testing"

	"github.com/cloudemprise/leetcode543/tree"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *tree.TreeNode
		expected int
	}{
		{
			name:     "Example 1",
			root:     &tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 4}, Right: &tree.TreeNode{Val: 5}}, Right: &tree.TreeNode{Val: 3}},
			expected: 3,
		},
		{
			name: "Example 2",
			root: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 4,
						Left: &tree.TreeNode{
							Val: 8,
						},
					},
					Right: &tree.TreeNode{
						Val: 5,
					},
				},
			},
			expected: 3,
		},
		{
			name:     "Example 3",
			root:     &tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 2}},
			expected: 1,
		},
		{
			name:     "Empty Tree",
			root:     nil,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := diameterOfBinaryTree(tc.root)
			if result != tc.expected {
				t.Errorf("Expected %d but got %d", tc.expected, result)
			}
		})
	}
}
