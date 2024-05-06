package main

import "fmt"

/*
Time Complexity: O(N)

Here's a breakdown of the time complexity for each step:

    Pushing all node pointers onto a stack takes O(N) time, as we need to
	traverse the entire linked list once.
    Reallocating nodes using two pointers takes O(N/2) time, which is O(N),
	as we need to traverse half of the stack.

Therefore, the overall time complexity is O(N) + O(N) = O(N).

Space Complexity: O(N)
*/

// reorderList solves the problem by making use of only a stack.
func reorderList(head *Node) {
	if head == nil {
		return
	}

	// Push all node pointer onto a stack
	var stack []*Node
	current := head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// Reallocate nodes using two pointers.
	lo, hi := 0, len(stack)-1
	for lo < hi {
		// Connect hi to lo
		stack[hi].Next = nil // terminate before allocation
		stack[lo].Next = stack[hi]
		lo++

		// Connect lo to hi
		stack[hi].Next = stack[lo]
		stack[lo].Next = nil // terminate after allocation
		hi--
	}
	// The termination strategy above is to preempt & ensure that odd/even
	// length lists eventually have their Tails terminated correctly.
}

///

func main() {

	var aList *List
	const amount = 5

	// Create a list and print it.
	aList = newList()
	for i := range [amount]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()

	reorderList(aList.head)
	fmt.Printf("Reordered List:\t")
	aList.printList()

}
