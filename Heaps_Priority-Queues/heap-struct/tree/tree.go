package tree

import (
	"fmt"
)

// node unit.
type node struct {
	data  int
	left  *node
	right *node
}

///

// node constructor.
func newNode(d int) *node {
	return &node{
		data: d,
	}
}

// tree object
type tree struct {
	root *node
}

///

// tree constructor.
func NewTree() *tree {
	return &tree{}
}

///

// Helper method.
func (t *tree) Reset() {
	t.root = nil
}

///

func (t *tree) Insert(d int) {
	if t.root == nil {
		t.root = newNode(d)
	} else {
		t.root.insertNode(d) // node insertion.
	}
}

///*-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-*

// An breadth-first search insertion algo that searches for the first
// available slot by making use of a FIFO queue.
func (n *node) insertNode(data int) {

	var queue []*node
	queue = append(queue, n)

	for len(queue) > 0 {

		tmp := queue[0]
		queue = queue[1:]

		if tmp.left == nil {
			tmp.left = newNode(data)
			break
		} else {
			queue = append(queue, tmp.left)
		}

		if tmp.right == nil {
			tmp.right = newNode(data)
			break
		} else {
			queue = append(queue, tmp.right)
		}
	}
}

///*-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-**-*

/// SPECIAL 2D TREE PRINT-OUT ALGO:

// This code prints out a simple 2D visual representation of a binary tree;
// rotated on its side, anti-clockwise. A node recursive (right-leaning) helper
// function uses a reverse in-order (Right-Parent-Left) transversal to
// facilitate the print out.

func (t *tree) PrintTree2D() {
	t.root.inOrderReverse2D(0)
}

///

func (n *node) inOrderReverse2D(space int) {
	if n == nil { // base case
		return
	}

	const count = 10
	space += count

	n.right.inOrderReverse2D(space)

	fmt.Printf("\n")
	for i := count; i < space; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%v\n", n.data)

	n.left.inOrderReverse2D(space)
}
