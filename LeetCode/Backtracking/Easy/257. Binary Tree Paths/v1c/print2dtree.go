package main

import (
	"fmt"
	"io"
)

/// SPECIAL 2D TREE PRINT-OUT ALGO:

// This code prints out a simple 2D visual representation of a binary tree;
// rotated on its side, anti-clockwise. A node recursive (right-leaning) helper
// function uses a reverse in-order (Right-Parent-Left) transversal to
// facilitate the print out.

func (t *TreeNode) printTree2D(w io.Writer) {
	t.inOrderReverse2D(w, 0)
}

//*-*

func (n *TreeNode) inOrderReverse2D(w io.Writer, space int) {
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
	fmt.Fprintf(w, "%v\n", n.Val)

	n.Left.inOrderReverse2D(w, space)
}

//-------------------------------------
