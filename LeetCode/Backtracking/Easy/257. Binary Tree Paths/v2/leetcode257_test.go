package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func newNode(Val int) *TreeNode {
	return &TreeNode{
		Val: Val,
	}
}

// insertNode inserts a new node into a binary tree. Uses a breadth-first
// transversal algorithms (i.e. queue) to locate the position of the next node.
func (n *TreeNode) insertNode(Val int) {
	// Use a queue for BFS
	var queue []*TreeNode
	// push root node.
	queue = append(queue, n)

	for len(queue) > 0 {
		// dequeue
		current := queue[0]
		queue = queue[1:]

		// Insert 1st available slot using
		// Breadth-first Search
		if current.Left == nil {
			current.Left = newNode(Val)
			break
		} else {
			queue = append(queue, current.Left)
		}
		if current.Right == nil {
			current.Right = newNode(Val)
			break
		} else {
			queue = append(queue, current.Right)
		}
	}
}

func TestBinaryTreePaths(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("input: %#v\n", nums)
	fmt.Println()

	aTree := newNode(nums[0])
	nums = nums[1:]
	for _, num := range nums {
		aTree.insertNode(num)
	}
	fmt.Println("A 2D printout of using a BFS insertion algo is as follows: ")
	fmt.Println()
	aTree.printTree2D(os.Stdout)
	fmt.Println()

	want := []string{
		"1->2->4",
		"1->2->5",
		"1->3->6",
		"1->3->7",
	}

	got := binaryTreePaths(aTree)
	fmt.Printf("got: %q\n", got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q want %q", got, want)
	}
}
