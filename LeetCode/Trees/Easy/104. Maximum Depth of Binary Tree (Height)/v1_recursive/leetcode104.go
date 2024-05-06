package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode104/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

/*
Time Complexity: O(n), where n is the number of nodes in the binary tree. This
is because we visit each node exactly once.

Space Complexity: O(h), where h is the height of the binary tree. This is
because the maximum amount of space used by the function call stack is
proportional to the height of the tree.
*/

// Recursive solution with helper function.
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

	const amount = 8
	var nums []int
	for i := range [amount]int{} {
		nums = append(nums, i+1)
	}

	fmt.Println("A 2D printout [of a", treeType, "using a", insertType,
		"insertion algo] is as follows: ")

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

	fmt.Println("Max Depth = ", maxDepth(aTree.Root))

}
