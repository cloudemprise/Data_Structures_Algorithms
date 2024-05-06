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

- We define a function `isSameTree` that takes in two pointers to
	`tree.TreeNode` structs, `p` and `q`, and returns a boolean value.
- We create two queues of `tree.TreeNode` pointers, `pQ` and `qQ`, and push `p`
	and `q` onto their respective queues.
- We then enter a loop that continues as long as there are nodes in both queues.
- In each iteration of the loop, we dequeue the first node from both `pQ` and
	`qQ` and assign them to `pCurrent` and `qCurrent`, respectively.
- We then check if both `pCurrent` and `qCurrent` are nil. If they are, we
	continue to the next iteration of the loop since they are both empty trees
	and therefore identical.
- We then check if either `pCurrent` or `qCurrent` is nil. If one of them is
	nil, we return false since they cannot be identical if one of them is empty.
- We then check if the values of `pCurrent` and `qCurrent` are equal. If they
	are not, we return false since they are not identical.
- Finally, we enqueue the left and right subtrees of `pCurrent` and `qCurrent`
	onto their respective queues, in that order. This ensures that we compare
	the left subtrees before the right subtrees.
- If we exit the loop without returning false, we know that `p` and `q` are
	identical, so we return true.

This solution has a time complexity of O(n), where n is the number of nodes in
the tree, since we need to visit each node once.
The space complexity is O(n), since in the worst case the queues will contain
all the nodes in the tree.
*/

// Iterative structural and quantitative comparison w/o queue module.
func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {

	// Two queues.
	pQ := []*tree.TreeNode{p}
	qQ := []*tree.TreeNode{q}

	for len(pQ) > 0 && len(qQ) > 0 {

		// Dequeue
		pCurrent := pQ[0]
		pQ = pQ[1:]

		// Dequeue
		qCurrent := qQ[0]
		qQ = qQ[1:]

		if pCurrent == nil && qCurrent == nil { // structural check
			continue
		}
		if pCurrent == nil || qCurrent == nil { // uncorrelated nodes
			return false
		}
		if pCurrent.Val != qCurrent.Val { // values check
			return false
		}

		// Enqueue left & right nodes.
		pQ = append(pQ, pCurrent.Left, pCurrent.Right)
		qQ = append(qQ, qCurrent.Left, qCurrent.Right)

	}
	return true
}

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
