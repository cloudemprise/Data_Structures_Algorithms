package main

import "fmt"

/*
This just illustrates what the solution would look like when splitting the
linked list at the 2nd middle node as opposed to the 1st middle node as well as
not necessarily terminating (with nil) the first half, i.e. list1. Also, a
slightly different logic for the Merge Step that basically amounts to the same
thing
*/

// reorderList first finds the (2nd) middle node and then splits the list into
// two halves. It then reverses the 2nd list and finally merges both halves.
func reorderList(head *Node) {
	if head == nil || head.Next == nil {
		return // only 1 or 2 nodes
	}

	// 1. Find the 2nd middle node
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNode := slow // 2nd middle node

	// 2. Split into two halves.
	list1 := head
	list2 := midNode
	// Note: There is no way of separating list2 from list1 with this
	// approach... but it's not actually required when you are careful on how
	// you operate on list2.

	// 3. Reverse list2.
	var back, current, forward *Node // initialise 3 pointers
	back = nil
	current = list2
	forward = current.Next
	// loop until end of list2
	for current != nil {
		forward = current.Next // keep a copy of forward link
		current.Next = back    // flip current 'Next' ptr backward
		back = current         // shuffle pointers onwards
		current = forward
	}
	list2 = back // new list2 Head

	// 4. Merge the two halves.
	var fwdPtr1, fwdPtr2 *Node
	for list2.Next != nil {

		// Save the next nodes in both halves
		fwdPtr1 = list1.Next
		fwdPtr2 = list2.Next

		// Connect current node of list1 to current node of list2
		list1.Next = list2
		// Connect current node of list2 to the next node of list1
		list2.Next = fwdPtr1

		// Shift pointers forward
		list1 = fwdPtr1
		list2 = fwdPtr2
	}
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
