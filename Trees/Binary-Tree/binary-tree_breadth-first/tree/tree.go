package tree

import (
	"fmt"
	"io"
	"log"

	"github.com/cloudemprise/binary-tree/queue"
)

/*
These are the actual definitions of the Tree structures and the functions used
to build the test cases. In particular we have the insertNode() function that
uses a breath-first insertion algorithm to find the next available slot which
makes use of a FIFO queue to facilitate that process.
*/

// Tree object.
type Tree struct {
	Root *Node
}

// Tree constructor.
func New() *Tree {
	return &Tree{}
}

// Node object.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Node constructor.
func newNode(d int) *Node {
	return &Node{
		Data: d,
	}
}

// Reset Tree.
func (t *Tree) Reset() {
	t.Root = nil
}

// Insert Node into Tree.
func (t *Tree) Insert(d int) {
	if t.Root == nil {
		t.Root = newNode(d) // first Node.
	} else {
		t.Root.insertNode(d) // Insert into Node.
	}
}

// insertNode uses a breadth-first algorithm to find the next available slot
// which makes use of a FIFO queue to facilitate that process.
func (n *Node) insertNode(d int) {

	q := queue.New()
	q.EnQueue(n)

	for !q.IsEmpty() {

		index, ok := q.DeQueue().(*Node) // type assertion
		if !ok {
			log.Fatal("insertNode() : type assertion")
		}

		if index.Left == nil {
			index.Left = newNode(d)
			break
		} else {
			q.EnQueue(index.Left)
		}

		if index.Right == nil {
			index.Right = newNode(d)
			break
		} else {
			q.EnQueue(index.Right)
		}
	}
}

//-------------------------------------

func (t *Tree) PrintTree(w io.Writer) {
	t.Root.printNodeBreadthFirst(w)
}

///

// Breadth-First iterative print-out using a queue.
func (n *Node) printNodeBreadthFirst(w io.Writer) {

	q := queue.New()
	q.EnQueue(n)

	var current *Node

	for !q.IsEmpty() {
		current = q.DeQueue().(*Node) // type assertion
		fmt.Fprintf(w, " %v", current.Data)

		if current.Left != nil {
			q.EnQueue(current.Left)
		}
		if current.Right != nil {
			q.EnQueue(current.Right)
		}
	}
}

//-------------------------------------

/// SPECIAL 2D TREE PRINT-OUT ALGO:

// This code prints out a simple 2D visual representation of a binary tree;
// rotated on its side, anti-clockwise. A Node recursive (Right-leaning) helper
// function uses a reverse in-order (Right-Parent-Left) transversal to
// facilitate the print out.

func (t *Tree) Print2dTree(w io.Writer) {
	t.Root.inOrderReverse2D(w, 0)
}

//*-*

func (n *Node) inOrderReverse2D(w io.Writer, space int) {
	if n == nil { // base case
		return
	}

	const count = 10
	space += count

	n.Right.inOrderReverse2D(w, space)

	fmt.Fprintf(w, "\n")
	for i := count; i < space; i++ {
		fmt.Fprintf(w, " ")
	}
	fmt.Fprintf(w, "%v\n", n.Data)

	n.Left.inOrderReverse2D(w, space)
}

//-------------------------------------
