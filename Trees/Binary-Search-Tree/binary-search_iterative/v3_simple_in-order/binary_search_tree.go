package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/cloudemprise/binary_search_tree/tree"
)

var treeType string = "binary-tree"
var insertType string = "breadth-first"
var transversalType string = "In-Order"

/*
Explanation

Use a while loop to iterate through the binary tree. The loop continues until
all nodes have been processed in the tree. Inside the loop, we use another while
loop to traverse the left subtree of the current node. We keep pushing the nodes
onto the stack until we reach the leftmost node. Once we reach the leftmost
node, we pop a node from the stack and add process it value. We then set the
current node to the right child of the popped node. We repeat this process until
we have processed all the nodes in the tree.

Complexity Analysis

The time complexity of this algorithm is O(n), where n is the number of nodes in
the binary tree. This is because we visit each node exactly once. The space
complexity of this algorithm is O(h), where h is the height of the binary tree.
This is because the maximum number of nodes that can be stored in the stack at
any given time is equal to the height of the binary tree. In the worst case, the
binary tree can be skewed, and the height can be equal to the number of nodes in
the tree, making the space complexity O(n).
*/

///

type myTree struct {
	tree.Tree
}

// printTree prints a binary tree in pre-order fashion.
// (loop within a loop)
func (t *myTree) printTree(w io.Writer) {
	if t.Root == nil {
		return // empty tree.
	}

	var s []*tree.Node // declare a new stack.
	current := t.Root  // Copy root node.

	// Loop while Stack is not empty.
	for len(s) > 0 || current != nil {

		// While node is not a leaf
		for current != nil {
			// Push on stack.
			s = append(s, current)
			// Move left.
			current = current.Left
		}

		// we're at the leaf node now.
		// Pop off the Stack.
		current = s[len(s)-1]
		s = s[:len(s)-1]

		// Process Node.
		fmt.Fprintf(w, " %v", current.Data)

		// Move right.
		current = current.Right
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
