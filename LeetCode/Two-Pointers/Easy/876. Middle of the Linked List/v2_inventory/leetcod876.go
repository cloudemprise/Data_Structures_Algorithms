package main

import (
	"fmt"
)

// Inventory Strategy: (Not Optimum)
// Create a list of node pointers and then use the length of the list to
// calculate the mid point.

// Time : O(n)
// Space: O(n)

// FindMiddleNode returns the (2nd) middle node of a linked list.
func (l *List) FindMiddleNode() *Node {
	store := []*Node{}
	currentNode := l.head
	for currentNode != nil {
		store = append(store, currentNode)
		currentNode = currentNode.Next
	}
	return store[len(store)/2] // integer division
}

///

func main() {

	// A test list.
	l := &List{
		head: &Node{
			Val: 1,
			Next: &Node{
				Val: 2,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 4,
						Next: &Node{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}

	// Print the test list.
	l.printList()

	middleNode := l.FindMiddleNode()
	fmt.Printf("Middle Node: %#v", middleNode.Val)
}
