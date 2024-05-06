package main

import "testing"

func TestCopyRandomList(t *testing.T) {
	// Create the original linked list
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node1.Random = node3
	node2.Random = node1
	node3.Random = node2

	// Call the function to copy the random list
	copy := copyRandomList(node1)

	// Check if the copied list is correct
	if copy.Val != 1 {
		t.Errorf("Expected copy.Val to be 1, got %d", copy.Val)
	}
	if copy.Next.Val != 2 {
		t.Errorf("Expected copy.Next.Val to be 2, got %d", copy.Next.Val)
	}
	if copy.Next.Next.Val != 3 {
		t.Errorf("Expected copy.Next.Next.Val to be 3, got %d", copy.Next.Next.Val)
	}

	// Create a set to store nodes of the original list.
	originalNodes := make(map[*Node]bool)

	// Iterate through the original list and store nodes in the set.
	current := node1
	for current != nil {
		originalNodes[current] = true
		current = current.Next
	}

	// Iterate through the copied list and check that the random pointers don't in fact point to the original list.
	current = copy
	for current != nil {
		if originalNodes[current.Random] {
			t.Errorf("Random pointer is pointing to a node in the original list: %v", current.Random)
		}
		current = current.Next
	}
}
