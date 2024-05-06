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

The space complexity of this solution is O(w), where w is the maximum width of
the binary tree, since the maximum number of nodes in the queue at any given
time is equal to the maximum width of the tree.
*/

// invertTree iteratively inverts a binary tree by performing a breadth-first
// search (using a queue) and then swapping the left & right subtrees of each
// node.
func invertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil {
		return nil
	}

	var queue []*tree.TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {

		node := queue[0]
		queue = queue[1:] 

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

	}
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
