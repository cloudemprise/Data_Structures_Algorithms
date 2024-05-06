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

type nodePath struct {
	node *TreeNode
	path []int // remember path to this node.
}

// binaryTreePaths returns all root-to-leaf paths of a binary tree in any order.
// The solution makes use of an iterative depth-first search algorithm, i.e
// a stack, to traverse the nodes of the binary tree as well as a nodePath
// object that records the history (path) that a particular permutation
// undertakes.
func binaryTreePaths(root *TreeNode) []string {

	// Record all the paths as routes.
	var routes []string

	// Create a stack of nodePath
	var stack []nodePath

	// Append root element
	element := nodePath{
		node: root,
		path: []int{}, // no history
	}
	stack = append(stack, element)

	// DFS while recording path to leaf node.
	for len(stack) > 0 {

		// Pop top off stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Check for leaf-node. Convert path to string.
		if current.node.Left == nil && current.node.Right == nil {
			var strPath string
			for _, p := range current.path {
				strPath += strconv.Itoa(p) + "->"
			}
			strPath += strconv.Itoa(current.node.Val)
			// Remember this route.
			routes = append(routes, strPath)
		}

		// Depth First Search while remembering path.
		if current.node.Right != nil {
			element := nodePath{
				node: current.node.Right,
				path: append(current.path, current.node.Val),
			}
			stack = append(stack, element)
		}
		if current.node.Left != nil {
			element := nodePath{
				node: current.node.Left,
				path: append(current.path, current.node.Val),
			}
			stack = append(stack, element)
		}
	}
	return routes
}
