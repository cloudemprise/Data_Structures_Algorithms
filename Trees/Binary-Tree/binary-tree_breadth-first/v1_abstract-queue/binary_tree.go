package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/cloudemprise/binary-tree/queue"
	"github.com/cloudemprise/binary-tree/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Breadth-First"

///

type myTree struct {
	tree.Tree
}

// printTree prints a binary tree in breadth-first fashion.
func (t *myTree) printTree(w io.Writer) {

	q := queue.New()  // declare a new queue.
	q.EnQueue(t.Root) // Enqueue root node.

	var current *tree.Node
	for !q.IsEmpty() {
		// Dequeue element.
		current = q.DeQueue().(*tree.Node) // type assertion

		// Process Node.
		fmt.Fprintf(w, " %v", current.Data)

		// Move left.
		if current.Left != nil {
			q.EnQueue(current.Left)
		}
		// Move right.
		if current.Right != nil {
			q.EnQueue(current.Right)
		}
	}
}

///

func main() {

	aTree := myTree{}

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
	aTree.printTree(os.Stdout)
	fmt.Println(" ]")

	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

	aTree.Reset()

	fmt.Println("A 2D printout of a random", treeType, "using a", insertType)
	fmt.Println("insertion algo is as follows: ")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, num := range r.Perm(7) {
		fmt.Printf(">%v ", num+1)
		aTree.Insert(num + 1)
	}
	fmt.Printf("\n")

	aTree.Print2dTree(os.Stdout)

	fmt.Println()
	fmt.Println("With a", transversalType, "transversal as follows:")
	fmt.Printf("[")
	aTree.printTree(os.Stdout)
	fmt.Println(" ]")
}
