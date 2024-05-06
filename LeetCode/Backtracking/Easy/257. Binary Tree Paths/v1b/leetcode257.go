package main

import (
	"fmt"
	"strings"
)

/*
257. Binary Tree Paths - Easy

Given the root of a binary tree, return all root-to-leaf paths in any order.

Note: A leaf is a node with no children.

Example 1:

Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]

Example 2:

Input: root = [1]
Output: ["1"]

Constraints:

    The number of nodes in the tree is in the range [1, 100].
    -100 <= Node.val <= 100
*/

// Binary tree node declaration.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binaryTreePaths returns all root-to-leaf paths of a binary tree in any order.
// The algorithm makes use of a recursive closure to traverse the nodes of a
// binary tree.
// Note: Using strings.Builder to build the ouput string is less time-efficient
// than using string concantenation and strconv.Iota.
func binaryTreePaths(root *TreeNode) []string {

	// Record all the paths as routes.
	var routes []string

	// Define a function variable.
	var traverse func(*TreeNode, []int)

	// Recusive closure to record path to leaf nodes.
	traverse = func(n *TreeNode, path []int) {
		if n == nil {
			return // base case
		}

		// Pre-order value inclusion.
		path = append(path, n.Val)

		// Pre-order path extraction.
		if n.Left == nil && n.Right == nil {
			//fmt.Println(path)
			var b strings.Builder
			fmt.Fprintf(&b, "%v", path[0])
			for _, p := range path[1:] {
				fmt.Fprintf(&b, "->%v", p)
			}
			routes = append(routes, b.String())
		}

		// Visit left & right trees.
		traverse(n.Left, path)
		traverse(n.Right, path)
	}

	// Execute closure here.
	var path []int
	traverse(root, path)

	return routes
}
