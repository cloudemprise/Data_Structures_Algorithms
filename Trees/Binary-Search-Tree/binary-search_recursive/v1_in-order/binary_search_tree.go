package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/cloudemprise/binary-search_recursive/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "In-Order"

///

type myTree struct {
	tree.Tree
}

// Recursive in-order transversal that prints to stdout.
func (t *myTree) printTree(w io.Writer) {

	// Recursive closure.
	var printNodesInOrder func(*tree.Node)
	printNodesInOrder = func(n *tree.Node) {
		if n == nil {
			return
		}

		printNodesInOrder(n.Left)
		fmt.Fprintf(w, " %v", n.Data)
		printNodesInOrder(n.Right)
	}
	printNodesInOrder(t.Root)
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
