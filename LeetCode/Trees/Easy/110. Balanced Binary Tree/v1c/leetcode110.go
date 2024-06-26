package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode110/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

///

// isBalanced is a recursive closure that checks if the left and right subtrees
// of every node differ in height by no more than 1.
func isBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	heightOK := true
	var getHeight func(*tree.TreeNode) int
	getHeight = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := getHeight(node.Left)
		rightHeight := getHeight(node.Right)

		if abs(leftHeight-rightHeight) > 1 {
			heightOK = false
		}

		return max(leftHeight, rightHeight) + 1

	}
	getHeight(root)

	return heightOK && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	const amount = 9
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

	fmt.Println("Is Tree Balanced = ", isBalanced(aTree.Root))

}

//*-*
