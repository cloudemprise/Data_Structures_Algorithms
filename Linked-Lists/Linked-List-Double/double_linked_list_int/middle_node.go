package main

// middleNode1st will return the 1st middle node of a singly linked list.
func middleNode1st(head *node) *node {

	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

// middleNode2nd returns the 2nd middle node of a singly linked list.
func middleNode2nd(head *node) *node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
