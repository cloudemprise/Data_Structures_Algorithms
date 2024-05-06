package main

import (
	"testing"

	"github.com/cloudemprise/leetcode104/tree"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root     *tree.TreeNode
		expected int
	}{
		{&tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 9}, Right: &tree.TreeNode{Val: 20, Left: &tree.TreeNode{Val: 15}, Right: &tree.TreeNode{Val: 7}}}, 3},
		{&tree.TreeNode{Val: 1, Left: &tree.TreeNode{Val: 2, Left: &tree.TreeNode{Val: 3, Left: &tree.TreeNode{Val: 4}, Right: &tree.TreeNode{Val: 4}}, Right: &tree.TreeNode{Val: 3}}, Right: &tree.TreeNode{Val: 2}}, 4},
		{nil, 0},
		{&tree.TreeNode{Val: 1}, 1},
	}

	for _, test := range tests {
		result := maxDepth(test.root)
		if result != test.expected {
			t.Errorf("maxDepth(%v) = %d; expected %d", test.root, result, test.expected)
		}
	}
}
