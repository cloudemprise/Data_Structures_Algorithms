package main

import "fmt"

/*
Reallocation Approach: The main difference lies in the approach used to
reallocate the nodes in the desired order. In the previous version, two
pointers (lo and hi) were used to traverse the stack from both ends and
connect the nodes alternately. In the new solution, a single pointer
(current) is used to traverse the list from the beginning, and the nodes
from the back of the stack are inserted one by one.
*/

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
	stack := []*Node{}
	current := head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// A single pointer is used to traverse the list from the beginning
	//while the nodes from the back of the stack are inserted one by one.
	current = head
	var fwdPtr *Node
	for i := len(stack) - 1; i > (len(stack)-1)/2; i-- {
		fwdPtr = current.Next
		current.Next = stack[i]
		current.Next.Next = fwdPtr
		current = fwdPtr
	}
	current.Next = nil // terminate Tail appropriately
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
