package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/cloudemprise/binary_search_tree/stack"
	"github.com/cloudemprise/binary_search_tree/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Pre-Order"

///

type myTree struct {
	tree.Tree
}

// printTree prints a binary tree in pre-order fashion.
func (t *myTree) printTree(w io.Writer) {
	if t.Root == nil {
		return // empty tree.
	}

	s := stack.New() // declare a new stack.
	s.Push(t.Root)   // Push Root on stack.

	for !s.IsEmpty() {

		current, ok := s.Pop().(*tree.Node)
		if !ok {
			log.Fatal("printTree type assertion")
		}

		// Process Node.
		fmt.Fprintf(w, " %v", current.Data)

		// Push RHS.
		if current.Right != nil {
			s.Push(current.Right)
		}

		// Push LHS.
		if current.Left != nil {
			s.Push(current.Left)
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
