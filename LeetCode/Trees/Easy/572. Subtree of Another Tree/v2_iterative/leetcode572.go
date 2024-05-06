package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode572/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

// This didn't perform very well within the leetcode engine.

func isSubtree(root *tree.TreeNode, subRoot *tree.TreeNode) bool {
	if root == nil {
		return false
	}
	if subRoot == nil {
		return true
	}

	// Push root on stack.
	stack := []*tree.TreeNode{root}

	for len(stack) > 0 {

		// Pop node off stack.
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		if node.Val == subRoot.Val && isSameTree(node, subRoot) {
			return true
		}

		stack = append(stack, node.Left, node.Right)
	}

	return false
}

func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {

	pQ := []*tree.TreeNode{p}
	qQ := []*tree.TreeNode{q}

	for len(pQ) > 0 && len(qQ) > 0 {

		pCurrent := pQ[0]
		pQ = pQ[1:]

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

		pQ = append(pQ, pCurrent.Left, pCurrent.Right)
		qQ = append(qQ, qCurrent.Left, qCurrent.Right)

	}
	return true
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
