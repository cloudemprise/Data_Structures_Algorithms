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
Explanation:

- The `isSymmetric` function takes the root of the binary tree as input and
	returns a boolean value indicating whether the tree is symmetric or not.
- If the root is nil, it returns true because an empty tree is always symmetric.
- Otherwise, it calls the `isMirror` function with the left and right subtrees
	of the root.
- The `isMirror` function takes two nodes as input and returns a boolean value
	indicating whether the two nodes are mirrors of each other or not.
- If both nodes are nil, it returns true because two empty nodes are always
	mirrors of each other.
- If one of the nodes is nil and the other is not, it returns false because a
	node and a nil value cannot be mirrors of each other.
- If the values of the two nodes are not equal, it returns false because two
	nodes with different values cannot be mirrors of each other.
- Otherwise, it recursively calls the `isMirror` function with the left and
	right children of the two nodes, checking if they are mirrors of each other.
- The `isSymmetric` function returns the result of calling `isMirror` on the
	left and right subtrees of the root.

This solution has a time complexity of O(n) and a space complexity of O(h),
where n is the number of nodes in the tree and h is the height of the tree.
*/

// isSymmetric checks if a binary-tree is symmetric wrt itself.
func isSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two trees are symmetrical, recursively.
func isMirror(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil { // structural check
		return true
	}
	if p == nil || q == nil { // uncorrelated nodes
		return false
	}
	if p.Val != q.Val { // values check
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
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

	/* fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.PrintTree(os.Stdout)
	fmt.Println(" ]") */

	fmt.Println()
	fmt.Printf("Is the tree symmetric? : \t %v\n", isSymmetric(aTree.Root))

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

	/* fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	bTree.PrintTree(os.Stdout)
	fmt.Println(" ]") */

	fmt.Println()
	fmt.Printf("Is the tree symmetric? : \t %v\n", isSymmetric(bTree.Root))
}
