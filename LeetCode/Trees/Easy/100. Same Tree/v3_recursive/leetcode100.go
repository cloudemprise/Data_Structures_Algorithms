package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode100/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

///

/*
Explanation:

- We define a function `isSameTree` that takes in two pointers to `TreeNode`
	structs, `p` and `q`, and returns a boolean value.
- We first check if both `p` and `q` are nil, in which case we return true since
	they are both empty trees and therefore identical.
- We then check if either `p` or `q` is nil, in which case we return false since
	they cannot be identical if one of them is empty.
- We then check if the values of `p` and `q` are equal. If they are not, we
	return false since they are not identical.
- Finally, we recursively call `isSameTree` on the left and right subtrees of
	`p` and `q`, and return the logical AND of the results. This checks if the
	left and right subtrees of `p` and `q` are identical as well.

This solution has a time complexity of O(n), where n is the number of nodes in
the tree, since we need to visit each node once.
The space complexity is O(h), where h is the height of the tree, since we need
to store the recursive call stack.
*/

// Recursive structural and quantitative comparison w/o queue module.
func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil { // structural check
		return true
	}
	if p == nil || q == nil { // uncorrelated nodes
		return false
	}
	if p.Val != q.Val { // values check
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/* // Same as above but with short-hand.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
} */

//-------------------------------------

func main() {

	var aTree = tree.New()

	nums1 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums1 {
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

	var bTree = tree.New()

	nums2 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("A 2D printout of a", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	for _, num := range nums2 {
		fmt.Printf(">%v ", num)
		bTree.Insert(num)
	}
	fmt.Printf("\n")

	bTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	bTree.PrintTree(os.Stdout)
	fmt.Println(" ]")

	//myTree.Reset()
	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

	fmt.Println("Are the two trees the same = ", isSameTree(aTree.Root, bTree.Root))
}

//*-*
