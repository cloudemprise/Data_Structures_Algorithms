package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode101/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"

//var transversalType string = "Breadth-First"

/*
This solution uses an iterative approach that uses two queues to check if the
binary tree is symmetric or not. The solution checks if the left and right nodes
are mirrors of each other by comparing their values and enqueuing their children
in a specific order.

The time and space complexity of the solution are both O(n), where n is the
number of nodes in the tree.
*/

// isSymmetric iteratively checks if a binary-tree is symmetric wrt itself.
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}

	// Left & Right queues.
	lhsQ := []*tree.TreeNode{root.Left}
	rhsQ := []*tree.TreeNode{root.Right}

	for len(lhsQ) > 0 && len(rhsQ) > 0 {

		lhsCurr := lhsQ[0]
		rhsCurr := rhsQ[0]

		lhsQ = lhsQ[1:]
		rhsQ = rhsQ[1:]

		if lhsCurr == nil && rhsCurr == nil { // structural check
			continue
		}
		if lhsCurr == nil || rhsCurr == nil { // uncorrelated nodes
			return false
		}
		if lhsCurr.Val != rhsCurr.Val { // values check
			return false
		}

		lhsQ = append(lhsQ, lhsCurr.Left, lhsCurr.Right)
		rhsQ = append(rhsQ, rhsCurr.Right, rhsCurr.Left)

	}
	return true
}

///

func main() {

	var aTree = tree.New()

	nums1 := []int{1, 3, 3, 5, 7, 7, 5}

	fmt.Println("2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums1 {
		fmt.Printf(">%v ", num)
		aTree.Insert(num)
	}
	fmt.Printf("\n")

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("Is the tree symmetric?")
	fmt.Printf("Want: true \t Got: %v\n", isSymmetric(aTree.Root))

	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

	var bTree = tree.New()

	nums2 := []int{1, 3, 3, 7, 5, 7, 5}

	fmt.Println("2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums2 {
		fmt.Printf(">%v ", num)
		bTree.Insert(num)
	}
	fmt.Printf("\n")

	bTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("Is the tree symmetric?")
	fmt.Printf("Want: false \t Got: %v\n", isSymmetric(bTree.Root))
}
