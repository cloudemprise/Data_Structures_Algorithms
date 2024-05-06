package main

import (
	"fmt"
)

// Recursive Strategy.

// Time Complexity : O(n)
// Space Complexity: O(n)

// middleNode uses a recursive function to find and return the 2nd middle
// node for an even number of nodes of a linked list.
func middleNode(head *ListNode) *ListNode {

	return middle(head, head.Next)
}

// middle returns the 2nd middle node for an even number of nodes of a linked list.
func middle(slow, fast *ListNode) *ListNode {
	if fast == nil {
		return slow
	}

	if fast.Next == nil {
		return slow.Next
	}

	return middle(slow.Next, fast.Next.Next)
}

///

func main() {

	// A test list.
	l := &List{
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val:  6,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	// Print the test list.
	l.printList()

	middleNode := middleNode(l.head)
	fmt.Printf("Middle ListNode: %#v", middleNode.Val)
}
