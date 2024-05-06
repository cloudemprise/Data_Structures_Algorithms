package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
Time Complexity : O(n)
Space Complexity: O(1)

function reverseList(head)
    back = null
    current = head
    forward = current

    while current is not null
        forward = current.Next
        current.Next = back
        back = current
        current = forward

    return back
end function
*/

// reverseList is an iterative function that reverses a linked list by way of
// a 3-pointer backflip strategy.
func reverseList(head *Node) *Node {

	// Initialise three pointers.
	var back, current, forward *Node
	back = nil
	current = head
	forward = current

	// Loop until the end of the list.
	for current != nil {
		forward = current.Next // keep a copy of forward link
		current.Next = back    // flip current 'Next' ptr backward
		back = current         // shuffle pointers onwards
		current = forward
	}
	return back // new Head
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
