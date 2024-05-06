package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/cloudemprise/binary-tree/stack"
	"github.com/cloudemprise/binary-tree/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "Post-Order"

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
	var result []int // For Node processing.

	for !s.IsEmpty() {

		// Pop off the Stack.
		current, ok := s.Pop().(*tree.Node)
		if !ok {
			log.Fatal("printTree type assertion")
		}

		// Keep a record for further processing.
		result = append(result, current.Data)

		// Push LHS.
		if current.Left != nil {
			s.Push(current.Left)
		}

		// Push RHS.
		if current.Right != nil {
			s.Push(current.Right)
		}
	}

	// Reverse captured result.
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Fprintf(w, " %v", result[i])
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
