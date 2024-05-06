package main

import "fmt"

/*
This version finds the 1st middle node then splits the list into two halves. It
then reverses the 2nd list and finally merges both halves.
*/

/*
Time Complexity: O(n)
---------------

Can be broken down as follows:

- Finding the first middle node : O(n/2)
- Splitting into two halves     : O(1)
- Reversing the second half     : O(n/2)
- Merging the two halves        : O(n/2)

Space Complexity : O(1)
----------------
*/

// reorderList first finds the (1st) middle node and then splits the list into
// two halves. It then reverses the 2nd list and finally merges both halves.
func reorderList(head *Node) {
	if head == nil || head.Next == nil {
		return // only 1 or 2 nodes
	}

	// 1. Find the 1st middle node
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNode := slow // 1st middle node

	// 2. Split into two halves.
	list1 := head
	list2 := midNode.Next
	midNode.Next = nil // terminate list1 correctly

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
	for list2 != nil {

		// Save the next nodes in both halves
		fwdPtr1 = list1.Next
		fwdPtr2 = list2.Next

		// Connect current node of list1 to current node of list2
		list1.Next = list2
		list1 = fwdPtr1 // Shift list1 pointer forward

		// Connect current node of list2 to current node of list1
		list2.Next = list1
		list2 = fwdPtr2 // Shift list2 pointer forward
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
