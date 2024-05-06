package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	// Test case 1: empty list
	var head *Node = nil
	result := reverseList(head)
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	// Test case 2: list with one element
	head = &Node{Val: 1, Next: nil}
	result = reverseList(head)
	if result != head {
		t.Errorf("Expected %v, but got %v", head, result)
	}

	// Test case 3: list with multiple elements
	head = &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: nil}}}
	result = reverseList(head)
	expected := &Node{Val: 3, Next: &Node{Val: 2, Next: &Node{Val: 1, Next: nil}}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
