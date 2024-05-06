package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
Time Complexity : O(n)
Space Complexity: O(n) (recursive stack)

function reverseList(head)
    if head is null
        return null

    if head.Next is null
        return head

    fwdRef = head.Next
    newHead = reverseList(head.Next)

    fwdRef.Next = head
    head.Next = null

    return newHead
end function
*/

// reverseList is a recursive function that flips backward a forward-reference
// pointer to the prevailing next pointer.
func reverseList(head *Node) *Node {
	if head == nil {
		return nil // empty list corner case
	}

	// Through Tail-Recursion, head eventually becomes the
	// new Head of the reversed list.
	if head.Next == nil {
		return head
	}

	// Create a local copy of the forward-reference.
	fwdRef := head.Next
	newHead := reverseList(head.Next) // recurse down the list

	// Flip forward-reference 'Next' ptr backward, i.e. reverse the link.
	fwdRef.Next = head

	// Terminate current 'Next' ptr to ensure that the subsequent
	// recursive loop's base case is satisfied.
	head.Next = nil

	// Tail-Recursion ensures that newHead is passed back up
	// the recursive chain during the unwind phase.
	return newHead
}

///

func main() {

	// Test Nodes Only Data Structure
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

	// Test Linked List Data Structure
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
