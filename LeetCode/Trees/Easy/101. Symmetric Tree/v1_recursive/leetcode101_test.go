package main

import (
	"testing"

	"github.com/cloudemprise/leetcode101/tree"
)

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		name   string
		input  *tree.TreeNode
		output bool
	}{
		{
			name:   "Empty Tree",
			input:  nil,
			output: true,
		},
		{
			name: "Single Node Tree",
			input: &tree.TreeNode{
				Val: 1,
			},
			output: true,
		},
		{
			name: "Symmetric Tree",
			input: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 4,
					},
				},
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 4,
					},
					Right: &tree.TreeNode{
						Val: 3,
					},
				},
			},
			output: true,
		},
		{
			name: "Non-Symmetric Tree",
			input: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Right: &tree.TreeNode{
						Val: 3,
					},
				},
				Right: &tree.TreeNode{
					Val: 2,
					Right: &tree.TreeNode{
						Val: 3,
					},
				},
			},
			output: false,
		},
		{
			name: "Non-Symmetric Tree with Same Value",
			input: &tree.TreeNode{
				Val: 1,
				Left: &tree.TreeNode{
					Val: 2,
					Right: &tree.TreeNode{
						Val: 2,
					},
				},
				Right: &tree.TreeNode{
					Val: 2,
					Right: &tree.TreeNode{
						Val: 2,
					},
				},
			},
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isSymmetric(tc.input)
			if got != tc.output {
				t.Errorf("isSymmetric(%v) = %v; want %v", tc.input, got, tc.output)
			}
		})
	}
}
