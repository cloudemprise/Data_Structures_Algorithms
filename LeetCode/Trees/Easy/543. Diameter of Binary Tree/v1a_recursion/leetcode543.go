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
A Suboptimal Solution
---------------------

The time complexity of this solution is O(n^2), where n is the number of nodes
in the binary tree. This is because for each node, we calculate the depth of its
left and right subtrees, which takes O(n) time. Therefore, the total time
complexity is O(n^2).

The space complexity of this solution is O(h), where h is the height of the
binary tree. This is because the maximum number of function calls that can be on
the call stack at any given time is equal to the height of the binary tree.
*/

// diameterOfBinaryTree calculates the diameter of a binary tree using recursion.
func diameterOfBinaryTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	// Get diameter of left & right subtrees.
	leftDiameter := diameterOfBinaryTree(root.Left)
	rightDiameter := diameterOfBinaryTree(root.Right)

	// Get maxDepth of left & right subtrees.
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	// return the max of 3 variables.
	return max(max(leftDiameter, rightDiameter), leftDepth+rightDepth)
}

// maxDepth returns the maximum depth of a binary tree.
func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
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
