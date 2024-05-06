package main

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	// Create a linked list for testing
	head := &Node{1, nil}
	head.Next = &Node{2, nil}
	head.Next.Next = &Node{3, nil}
	head.Next.Next.Next = &Node{4, nil}
	head.Next.Next.Next.Next = &Node{5, nil}

	// Test case 1: Remove the second node from the end
	n := 2
	expected := []int{1, 2, 3, 5}

	result := removeNthFromEnd(head, n)

	// Validate the result
	if !compareList(result, expected) {
		t.Errorf("Test case 1 failed: Expected %v, got %v", expected, getListValues(result))
	}

	// Test case 2: Remove the first node from the end
	n = 1
	expected = []int{1, 2, 3}

	result = removeNthFromEnd(head, n)

	// Validate the result
	if !compareList(result, expected) {
		t.Errorf("Test case 2 failed: Expected %v, got %v", expected, getListValues(result))
	}

	// Test case 3: Remove the last node from the end
	n = 3
	expected = []int{2, 3}

	result = removeNthFromEnd(head, n)

	// Validate the result
	if !compareList(result, expected) {
		t.Errorf("Test case 3 failed: Expected %v, got %v", expected, getListValues(result))
	}
}

// Helper function to compare the values of a linked list with an array
func compareList(head *Node, values []int) bool {
	curr := head
	for i := 0; i < len(values); i++ {
		if curr == nil || curr.Val != values[i] {
			return false
		}
		curr = curr.Next
	}
	return curr == nil
}

// Helper function to get the values of a linked list as an array
func getListValues(head *Node) []int {
	values := []int{}
	curr := head
	for curr != nil {
		values = append(values, curr.Val)
		curr = curr.Next
	}
	return values
}
