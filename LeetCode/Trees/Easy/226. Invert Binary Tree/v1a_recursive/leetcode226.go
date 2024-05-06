package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode226/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

/*
The time complexity of this solution is O(n), where n is the number of nodes in
the binary tree, since we visit each node exactly once.

The space complexity is O(h), where h is the height of the binary tree, since
the maximum number of function calls on the call stack at any given time is
equal to the height of the tree.
*/

// invertTree recursively inverts a binary tree by first swapping the left &
// right subtrees and then inverting the left & right trees respectively.
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	// Swap left & right subtrees.
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

///

func main() {

	var aTree = tree.New()
	nums := []int{1, 2, 3, 4, 5, 6, 7}

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

	aTree.Root = invertTree(aTree.Root)

	fmt.Println("A 2D printout of the inverted tree above: ")

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]")
}
