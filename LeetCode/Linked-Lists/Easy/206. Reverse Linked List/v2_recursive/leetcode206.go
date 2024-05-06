package main

import (
	"bytes"
	"fmt"
	"os"
)

// Time Complexity : O(n)
//Space Complexity: O(n) (recursive stack)

// Leetcode solution using a recursive helper function.
func reverseList(head *Node) *Node {
	return reverse(head, head)
}

func reverse(root, node *Node) *Node {
	if root == nil || node.Next == nil {
		return root
	}

	node.Next, node.Next.Next, root = node.Next.Next, root, node.Next
	return reverse(root, node)
}

///

func main() {

	//------ Test Nodes Only Data Structure -------
	fmt.Println("Test Node Code:")
	fmt.Println("---------------")
	// Create a list for testing.
	slice1 := []int{4, 3, 2, 1}
	nodePtr := &Node{Val: 5}
	for _, val := range slice1 {
		nodePtr.insertTail(val)
	}

	// Printout the list
	fmt.Printf("Input List:\t")
	nodePtr.printNodes()

	// Reverse and printout
	nodePtr = reverseList(nodePtr)
	fmt.Printf("Reversed List:\t")
	nodePtr.printNodes()
	fmt.Println()

	/* //empty := &Node{}
	var empty *Node
	empty.printNodes() */

	//----- Test Linked List Data Structure -----
	fmt.Println("Test List Code:")
	fmt.Printf("---------------")
	// Create a list for testing.
	slice2 := []int{5, 4, 3, 2, 1}
	myList := newList()
	for _, val := range slice2 {
		myList.insertTail(val)
	}

	// Printout the list
	var buf1 bytes.Buffer
	myList.scanList(&buf1)
	fmt.Printf("Input myList:\t")
	buf1.WriteTo(os.Stdout)

	// Reverse and printout
	transpose := reverseList(myList.head)
	var buf2 bytes.Buffer
	transpose.scanNodes(&buf2)
	fmt.Printf("Output myList:\t")
	buf2.WriteTo(os.Stdout)

}
