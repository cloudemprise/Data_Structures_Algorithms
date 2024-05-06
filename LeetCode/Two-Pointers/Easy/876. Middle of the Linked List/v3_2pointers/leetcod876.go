package main

import (
	"fmt"
)

// Two Pointer Strategy. (Optimum)

// Time Complexity : O(n)/2 = O(n)
// Space Complexity: O(1)

// middleNode uses the two-pointer approach to return the 2nd middle node
// for an even number of nodes of a linked list.
func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

///

func main() {

	/* 	// A test list.
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
	   							Val:  5,
	   							Next: nil,
	   						},
	   					},
	   				},
	   			},
	   		},
	   	} */
	// A test list.
	l := &List{
		head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	// Print the test list.
	l.printList()

	middleNode := middleNode(l.head)
	fmt.Printf("Middle ListNode: %#v", middleNode.Val)
}
