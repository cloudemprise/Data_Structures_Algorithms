package main

import "strconv"

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
// Note: Using string concantenation and strconv.Iota() is more time-efficient
// than using strings.Builder to build the output string.
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

		// Pre-order path extraction results in tail-recursion
		// which can be compiler optimised.

		// A leaf-node requires both left & right
		// ptrs to be equal to nill
		if n.Left == nil && n.Right == nil {
			var soln string
			soln = strconv.Itoa(path[0])
			for _, p := range path[1:] {
				soln = soln + "->" + strconv.Itoa(p)
			}
			routes = append(routes, soln)
		}

		// Visit left & right trees.
		traverse(n.Left, path)
		traverse(n.Right, path)
	}

	// Execute closure here.
	var p []int
	traverse(root, p)

	return routes
}
