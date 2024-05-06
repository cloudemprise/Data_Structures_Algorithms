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
The time complexity of this solution is O(m*n) in the worst case, where m and n
are the number of nodes in the trees 'root' and 'subRoot', respectively. This is
because the function isSameTree() is called for each node in the tree 'root',
and each call to isSameTree() has a time complexity of O(n), where n is the
number of nodes in the tree 'subRoot'.

The space complexity of this solution is O(h), where h is the height of the tree
'root'. This is because the function isSameTree() is called recursively for each
node in the tree 'root', and the maximum number of recursive calls on the call
stack is equal to the height of the tree 'root'.
*/

// isSubtree is a recursive closure pre-order traversal function that indicates
// whether tree 'subRoot' is a subtree of tree 'root'.
func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {

	var isFound bool = false

	var traverse func(*tree.TreeNode)
	traverse = func(node *tree.TreeNode) {

		if node != nil && !isFound {

			isFound = isSameTree(node, subRoot)

			traverse(node.Left)
			traverse(node.Right)
		}
	}
	traverse(root)

	return isFound
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
