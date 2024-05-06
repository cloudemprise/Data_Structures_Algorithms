package main

import (
	"fmt"
)

// removeNthFromEnd removed the nth element from a linked list by using the
// two-pointer technique, where the first pointer moves ahead faster than
// the second pointer.
func removeNthFromEnd(head *Node, n int) *Node {
	first := head
	second := head

	// Move the first pointer n steps ahead. Ensures a gap
	// of n nodes between the first and second pointers.
	for i := 0; i < n; i++ {
		first = first.Next
	}

	// NOTE! Special Case when n = numOfNodes, first pointer
	// reaches the end. So, just remove head node.
	if first == nil {
		return head.Next
	}

	// Move both pointers simultaneously until the
	// first pointer reaches the last node.
	for first.Next != nil {
		first = first.Next
		second = second.Next
	} // second will now be just before nth node.

	// Remove the nth node from the end.
	second.Next = second.Next.Next

	return head
}

///

func main() {

	var aList *List
	var N int

	// Create a list and print it.
	aList = newList()
	for i := range [5]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()
	N = 2
	fmt.Printf("n = %v\n", N)

	aList.head = removeNthFromEnd(aList.head, N)
	fmt.Printf("Resultant List:\t")
	aList.printList()
	fmt.Println()

	//*-*

	// Create a list and print it.
	aList = newList()
	for i := range [1]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()
	N = 1
	fmt.Printf("n = %v\n", N)

	aList.head = removeNthFromEnd(aList.head, N)
	fmt.Printf("Resultant List:\t")
	aList.printList()
	fmt.Println()

	//*-*

}

//*-*
