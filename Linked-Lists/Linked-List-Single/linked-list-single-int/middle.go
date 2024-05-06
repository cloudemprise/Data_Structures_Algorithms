package list

// Example problem: 876. Middle of the Linked List - Easy

/*
When a linked list has an even number of nodes, there are two middle nodes. Both
functions below utilize the two-pointer approach to find the first and second
middle nodes of a linked list, respectively. One pointer moves at double the
pace than the other to identify the middle element.

The difference in the loop condition allows the second function to return the
2nd middle node for an even number of nodes of a linked list.
*/

// middleNode1st uses the two-pointer approach to return the 1st middle node for
// an even number of nodes of a linked list.
func middleNode1st(head *node) *node {

	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// Debug purposes:
	/* fmt.Printf("fast: %v\n", fast.val)
	fmt.Printf("slow: %v\n", slow.val)
	fmt.Println() */

	return slow
}

// middleNode1st uses the two-pointer approach to return the 2nd middle node for
// an even number of nodes of a linked list.
func middleNode2nd(head *node) *node {

	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// Debug purposes:
	// For Even-length lists fast.val deferences a nil pointer, thus
	// segmentation violation ensues.
	/* //fmt.Printf("fast: %v\n", fast.val)
	fmt.Printf("slow: %v\n", slow.val)
	fmt.Println() */

	return slow
}
