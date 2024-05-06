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
Intuitive Understanding:
------------------------

The diameter of a tree T is the largest of the following quantities:

* the diameter of T’s left subtree.
* the diameter of T’s right subtree.
* the longest path between leaves that goes through the root of T, which can be
	computed from the heights of the subtrees of T.
*/

/* // This is similar to the Closure solution below but rather passing
// the variable to a separate function.
func diameterOfBinaryTree(root *tree.TreeNode) int {

	var maxDiameter int

	heightOfBinaryTree(root, &maxDiameter)

	return maxDiameter

}

// *-*

func heightOfBinaryTree(root *tree.TreeNode, maxDiameter *int) int {
	if root == nil {
		return 0
	}

	lhs := heightOfBinaryTree(root.Left, maxDiameter)
	rhs := heightOfBinaryTree(root.Right, maxDiameter)

	// Post-Order mutation phase here:
	if lhs+rhs > *maxDiameter {
		*maxDiameter = lhs + rhs
	}

	if lhs > rhs {
		return lhs + 1
	} else {
		return rhs + 1
	}
} */

//-------------------------------------

/* // From Leetcode: Time Complexity = O(n)
// Use a Closure to perform a Post-Order transversal height
// calculation, then capture maxDiameter during mutation phase.
func diameterOfBinaryTree(root *tree.TreeNode) int {

	var maxDiameter int

	var heightOfBinaryTree func(*tree.TreeNode) int

	// For each recursion we want to calculate the height
	// of each side via post-order transversal algo.
	heightOfBinaryTree = func(n *tree.TreeNode) int {
		if n == nil {
			return 0
		}

		lhs := heightOfBinaryTree(n.Left)
		rhs := heightOfBinaryTree(n.Right)

		// Post-Order mutation phase here:
		if lhs+rhs > maxDiameter {
			maxDiameter = lhs + rhs
		}

		if lhs > rhs {
			return lhs + 1
		} else {
			return rhs + 1
		}
	}

	heightOfBinaryTree(root)
	return maxDiameter
} */

//-------------------------------------

// Caluclate the Diameter of a Binary Tree using Recursion.
// Time Complexity: O(N^2)
// Space Complexity: O(N) for call stack

// Calculate the Diameter of a Binary Tree.
func diameterOfBinaryTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	// get the diameter of left and right sub-trees.
	lhsDia := diameterOfBinaryTree(root.Left)
	rhsDia := diameterOfBinaryTree(root.Right)

	// get the height of left and right sub-trees.
	lhsHght := heightOfBinaryTree(root.Left)
	rhsHght := heightOfBinaryTree(root.Right)

	// return the max of 3 variables.
	return max(max(lhsDia, rhsDia), lhsHght+rhsHght)
}

// Calculate the height of a binary tree using
// Post-Order Transversal.
func heightOfBinaryTree(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	lhs := heightOfBinaryTree(root.Left)
	rhs := heightOfBinaryTree(root.Right)

	if lhs > rhs {
		return 1 + lhs
	}
	return 1 + rhs
}

//*-*

// Return the largest of two values.
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//*-*

//-------------------------------------

func main() {

	var aTree = tree.New()
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//nums := []int{1, 2, 3, 4, 5, 6}
	//nums := []int{1, 2, 3, 4, 5}
	//nums := []int{1, 2, 3, 4}
	//nums := []int{1, 2, 3}
	//nums := []int{1, 2}
	//nums := []int{1}
	//nums := []int{}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums {
		fmt.Printf(">%v ", num)
		aTree.Insert(num)
	}
	fmt.Printf("\n")

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]")

	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

	fmt.Println("Max Diameter = ", diameterOfBinaryTree(aTree.Root))

}
