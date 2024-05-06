package main

import (
	"reflect"
	"testing"

	"github.com/cloudemprise/leetcode226/tree"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		input    *tree.TreeNode
		expected *tree.TreeNode
	}{
		{
			name: "Example 1",
			input: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 1,
					},
					Right: &tree.TreeNode{
						Val: 3,
					},
				},
				Right: &tree.TreeNode{
					Val: 7,
					Left: &tree.TreeNode{
						Val: 6,
					},
					Right: &tree.TreeNode{
						Val: 9,
					},
				},
			},
			expected: &tree.TreeNode{
				Val: 4,
				Left: &tree.TreeNode{
					Val: 7,
					Left: &tree.TreeNode{
						Val: 9,
					},
					Right: &tree.TreeNode{
						Val: 6,
					},
				},
				Right: &tree.TreeNode{
					Val: 2,
					Left: &tree.TreeNode{
						Val: 3,
					},
					Right: &tree.TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			name: "Example 2",
			input: &tree.TreeNode{
				Val: 2,
				Left: &tree.TreeNode{
					Val: 1,
				},
			},
			expected: &tree.TreeNode{
				Val: 2,
				Right: &tree.TreeNode{
					Val: 1,
				},
			},
		},
		{
			name:     "Example 3",
			input:    nil,
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := invertTree(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("expected: %v, got: %v", tc.expected, got)
			}
		})
	}
}
