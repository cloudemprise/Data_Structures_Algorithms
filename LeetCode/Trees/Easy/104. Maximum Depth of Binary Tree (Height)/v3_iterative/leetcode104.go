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

Space Complexity: O(w), where w is the maximum width of the binary tree. This is
because the maximum number of nodes in the queue at any given time is
proportional to the width of the tree.
*/

// maxDepth returns the maximum depth of a binary-tree using the iterative
// breadth-first traversal algo that makes use of a slice as a queue.
func maxDepth(root *tree.TreeNode) int {
	if root == nil { // corner case.
		return 0
	}

	// Enqueue root of tree.
	var queue []*tree.TreeNode
	queue = append(queue, root)

	var depth int

	for len(queue) > 0 {

		// Dynamic queue so record length before looping.
		size := len(queue)

		// Process the whole node level.
		for i := 0; i < size; i++ {

			// Dequeue
			tmp := queue[0]
			queue = queue[1:]

			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
		depth++ // move on to next level
	}

	return depth
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
