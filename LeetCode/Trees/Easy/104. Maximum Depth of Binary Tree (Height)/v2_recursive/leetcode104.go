package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode104/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

///

// Recursive solution without the helper function.
func maxDepth(n *tree.TreeNode) int {
	if n == nil {
		return 0
	}

	lhs := maxDepth(n.Left)
	rhs := maxDepth(n.Right)

	if lhs > rhs {
		return 1 + lhs
	}
	return 1 + rhs
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
