package main

import (
	"testing"

	"github.com/cloudemprise/leetcode572/tree"
)

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name   string
		root   *tree.TreeNode
		sub    *tree.TreeNode
		expect bool
	}{
		{
			name: "subtree exists",
			root: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 1,
					},
					Right: &tree.TreeNode{
						Val: 2,
					},
				},
				Right: &tree.TreeNode{
					Val: 5,
				},
			},
			sub: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 2,
				},
			},
			expect: true,
		},
		{
			name: "subtree does not exist",
			root: &tree.TreeNode{
				Val: 3,
				Left: &tree.TreeNode{
					Val: 4,
					Left: &tree.TreeNode{
						Val: 1,
					},
					Right: &tree.TreeNode{
						Val: 2,
						Left: &tree.TreeNode{
							Val: 0,
						},
					},
				},
				Right: &tree.TreeNode{
					Val: 5,
				},
			},
			sub: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 1,
				},
				Right: &tree.TreeNode{
					Val: 2,
				},
			},
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSubtree(test.root, test.sub)
			if result != test.expect {
				t.Errorf("Test case %s failed: expected %v but got %v", test.name, test.expect, result)
			}
		})
	}
}
