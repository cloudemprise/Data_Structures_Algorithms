package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode572/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

/*
The isSubtree function takes two pointers to TreeNode structs as input, 'root'
and 'subRoot', and returns a boolean indicating whether 'subRoot' is a subtree
of 'root'. The function first checks if 'root' is nil, in which case it returns
false. If 'root' and 'subRoot' are the same tree, it returns true. Otherwise, it
recursively checks if 'subRoot' is a subtree of the left or right subtree of
'root'. The isSameTree function is a helper function that checks if two trees
are the same. It returns true if both trees are nil, false if only one of them
is nil, and recursively checks if the values and subtrees of the two trees are
the same.

The time complexity of this solution is O(m*n) in the worst case, where m and n
are the number of nodes in 'root' and 'subRoot', respectively. This occurs when
'root' is a skewed tree and 'subRoot' is a tree with only one node.

The space complexity is O(n), where n is the number of nodes in 'root', due to the
recursive calls to isSubtree.
*/

// isSubtree is a nested recursive function that indicates whether tree
// 'subRoot' is a subtree of tree 'root'.
func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {
	// No need to check on subRoot being null as it
	// is invalid that null is subtree of null.
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// isSameTree is a tail-recursive check if two trees are the same.
func isSameTree(s *tree.TreeNode, t *tree.TreeNode) bool {
	if s == nil && t == nil { // structural check
		return true
	}
	if s == nil || t == nil { // uncorrelated nodes
		return false
	}
	if s.Val != t.Val { // values check
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

///

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
	fmt.Printf("\n//---\n\n")

	///

	var bTree = tree.New()
	//nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	//nums2 := []int{2, 4, 5, 7}
	nums2 := []int{2, 4, 5}

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
	fmt.Printf("\n//---\n\n")

	fmt.Println("Does the Secondary tree S exists within the Primary tree T =",
		isSubtree(aTree.Root, bTree.Root))
}

//*-*
