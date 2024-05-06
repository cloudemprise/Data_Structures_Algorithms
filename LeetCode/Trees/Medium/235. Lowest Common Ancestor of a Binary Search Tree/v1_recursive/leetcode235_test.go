package main

import (
	"testing"

	"github.com/cloudemprise/leetcode235/tree"
)

/*
		    6
		  /   \
		 2     8
		/ \   / \
	   0   4 7   8
*/
func TestLowestCommonAncestor(t *testing.T) {
	root := &tree.TreeNode{Val: 6,
		Left: &tree.TreeNode{Val: 2,
			Left: &tree.TreeNode{Val: 0},
			Right: &tree.TreeNode{Val: 4,
				Left:  &tree.TreeNode{Val: 3},
				Right: &tree.TreeNode{Val: 5},
			},
		},
		Right: &tree.TreeNode{Val: 8,
			Left:  &tree.TreeNode{Val: 7},
			Right: &tree.TreeNode{Val: 9},
		},
	}

	tests := []struct {
		p, q     *tree.TreeNode
		expected *tree.TreeNode
	}{
		{root.Left, root.Right, root},
		{root.Left, root.Left.Right, root.Left},
		{root.Right, root.Right.Left, root.Right},
		{root.Left.Right, root.Right.Left, root},
	}

	for _, tt := range tests {
		result := lowestCommonAncestor(root, tt.p, tt.q)
		if result != tt.expected {
			t.Errorf("lowestCommonAncestor(%v, %v, %v) = %v, want %v", root, tt.p, tt.q, result, tt.expected)
		}
	}
}
