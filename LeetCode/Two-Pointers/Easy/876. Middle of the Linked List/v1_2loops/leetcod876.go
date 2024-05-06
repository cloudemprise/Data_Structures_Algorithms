package main

import (
	"fmt"
)

// Two-Loop Strategy. (Not optimum)

// Time : O(n) + O(n) = O(n)
// Space: O(1)

// FindMiddleNode returns the (2nd) middle node of a linked list.
func (l *List) FindMiddleNode() *Node {

	// Determine length of list.
	length := 0
	currentNode := l.head
	for currentNode != nil {
		currentNode = currentNode.Next
		length++
	} // Time: O(n) / Space: O(1)

	// Search for middle node.
	midPoint := length / 2

	currentNode = l.head
	for midPoint > 0 {
		currentNode = currentNode.Next
		midPoint--
	} // Time: O(n) / Space: O(1)

	return currentNode
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
