package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	// Test case 1: both lists are nil
	var list1 *Node = nil
	var list2 *Node = nil
	result := mergeTwoLists(list1, list2)
	if result != nil {
		t.Errorf("Expected nil, but got %v", result)
	}

	// Test case 2: one list is nil
	list1 = &Node{Val: 1, Next: nil}
	list2 = nil
	result = mergeTwoLists(list1, list2)
	if result != list1 {
		t.Errorf("Expected %v, but got %v", list1, result)
	}

	// Test case 3: both lists have one element
	list1 = &Node{Val: 1, Next: nil}
	list2 = &Node{Val: 2, Next: nil}
	result = mergeTwoLists(list1, list2)
	expected := &Node{Val: 1, Next: &Node{Val: 2, Next: nil}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// Test case 4: both lists have multiple elements
	list1 = &Node{Val: 1, Next: &Node{Val: 3, Next: nil}}
	list2 = &Node{Val: 2, Next: &Node{Val: 4, Next: nil}}
	result = mergeTwoLists(list1, list2)
	expected = &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4, Next: nil}}}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
