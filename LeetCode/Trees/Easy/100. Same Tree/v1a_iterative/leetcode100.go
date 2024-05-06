package main

import (
	"fmt"
	"os"

	"github.com/cloudemprise/leetcode100/queue"
	"github.com/cloudemprise/leetcode100/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

// Iterative structural and quantitative comparison with queue module.
func isSameTree(p *tree.TreeNode, q *tree.TreeNode) bool {
	pQ := queue.New()
	qQ := queue.New()

	pQ.EnQueue(p)
	qQ.EnQueue(q)

	for pQ.Size() > 0 && qQ.Size() > 0 {

		pCurrent := pQ.DeQueue().(*tree.TreeNode)
		qCurrent := qQ.DeQueue().(*tree.TreeNode)

		if pCurrent == nil && qCurrent == nil { // structural check
			continue
		}
		if pCurrent == nil || qCurrent == nil { // uncorrelated nodes
			return false
		}
		if pCurrent.Val != qCurrent.Val { // values check
			return false
		}

		pQ.EnQueue(pCurrent.Left)
		pQ.EnQueue(pCurrent.Right)

		qQ.EnQueue(qCurrent.Left)
		qQ.EnQueue(qCurrent.Right)

	}

	// Check if anything is still in queues.
	if pQ.Size() == 0 && qQ.Size() == 0 {
		return true
	}
	return false
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
