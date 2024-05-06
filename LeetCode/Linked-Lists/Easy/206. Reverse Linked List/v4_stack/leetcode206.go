package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
Time Complexity : O(n) + O(n) = O(n)
Space Complexity: O(n)!!!

function reverseList(head)
	// Base case
    if head is null
        return null

    stack = empty array
    current = head

    while current is not null
        push current to stack
        current = current.Next

    ln = length of stack - 1
    newHead = stack[ln] // reversed Head

    current = stack[ln]
    for i from ln - 1 down to 0
        current.Next = stack[i] // flip pointer backward
        current = current.Next

    current.Next = null // terminate new tail correctly

    return newHead
end function
*/

// reverseList is an iterative function that reverses a linked-list
// by using a stack.
func reverseList(head *Node) *Node {
	if head == nil {
		return nil // corner case
	}

	// Create a stack.
	stack := []*Node{}
	// Push all nodes onto a stack.
	current := head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	ln := len(stack) - 1
	newHead := stack[ln] // reversed Head

	// Process each and every 'Next' pointer.
	current = stack[ln]
	for i := ln - 1; i >= 0; i-- {
		current.Next = stack[i] // flip pointer backward
		current = current.Next
	}
	current.Next = nil // terminate new tail correctly

	return newHead
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
