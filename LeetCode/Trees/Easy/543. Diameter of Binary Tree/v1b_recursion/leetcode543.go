package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode543/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

/*
NOTE: The definition of the diameter of a binary tree as defined by this problem
is incorrect. The correct definition for the diameter of a binary tree is
defined as the number of NODES on the longest path between two end nodes and
not the number of EDGES as in the case of this problem.

Reference : https://www.geeksforgeeks.org/diameter-of-a-binary-tree/
*/

/*
Intuitive Understanding:
------------------------

The diameter of a tree T is the largest of the following quantities:

* the diameter of T’s left subtree.
* the diameter of T’s right subtree.
* the longest path between leaves that goes through the root of T, which can be
	computed from the heights of the subtrees of T.
*/

/*
An Improved Solution
--------------------

An optimization can be made to v1a_recursion by calculating the depth of each
node in the same recursion as the diameter calculation. This can be done by
modifying the depth() function to determine both the depth and diameter of the
binary tree. This optimization reduces the time complexity to O(n) and the space
complexity to O(h).

Thus:

The time complexity of this solution is O(n), where n is the number of nodes in
the binary tree. This is because we visit each node exactly once.

The space complexity of this solution is O(h), where h is the height of the
binary tree. This is because the maximum number of function calls that can be on
the call stack at any given time is equal to the height of the binary tree.

However, the space complexity can be reduced to O(1) by using an iterative
approach instead of a recursive approach.
*/

// diameterOfBinaryTree calculates the diameter of a binary tree using a
// recursive closure.
func diameterOfBinaryTree(root *tree.TreeNode) int {
	var maxDiameter int

	// Recursive closure.
	var maxDepth func(*tree.TreeNode) int
	maxDepth = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}

		// Get maxDepth of left & right subtrees.
		leftDepth := maxDepth(node.Left)
		rightDepth := maxDepth(node.Right)

		// Determine maxDiameter at this level.
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		// Calculate max depth at this level.
		maxDepth := max(leftDepth, rightDepth) + 1
		return maxDepth
	}
	maxDepth(root)

	return maxDiameter
}

// max returns the larger of two integers.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

///

func main() {

	var aTree = tree.New()

	const amount = 6
	var nums []int
	for i := range [amount]int{} {
		nums = append(nums, i+1)
	}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums {
		fmt.Printf(">%v ", num)
		aTree.Insert(num)
	}
	fmt.Println()

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]")
	fmt.Println()

	fmt.Println("Max Diameter = ", diameterOfBinaryTree(aTree.Root))

	fmt.Printf("\n//-------------------------------------\n\n")
}
