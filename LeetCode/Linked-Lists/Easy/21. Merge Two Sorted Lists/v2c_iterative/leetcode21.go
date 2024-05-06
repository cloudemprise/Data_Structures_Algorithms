package main

import "fmt"

// More readable refactored version of v2a

/*
Time Complexity:  O(n + m) , where n & m are the lengths of the two input lists.
Space Complexity: O(1)
*/

// mergeTwoLists iteratively merges two sorted lists into a single sorted list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	// If either of the lists are empty then return the other list
	// since merging with an empty list results in the same list.
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Create two pointers in order to manage the
	// construction of a new merged list.
	var head, temp *Node

	// Determine which element will be the new Head.
	if list1.Val <= list2.Val {
		head = list1
		temp = list1       // make a local copy
		list1 = list1.Next // break link
	} else {
		head = list2
		temp = list2       // make a local copy
		list2 = list2.Next // break link
	}

	// Process elements from either list until one list is empty.
	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			temp.Next = list1
			list1 = list1.Next // remove front element from list1
		} else {
			temp.Next = list2
			list2 = list2.Next // remove front element from list2
		}
		temp = temp.Next // shift pointer forward.
	}

	// Append the remaining elements to the end of the merged list.
	if list1 != nil {
		temp.Next = list1
	}
	if list2 != nil {
		temp.Next = list2
	}

	return head
}

///

func main() {

	// Create Input List 1.
	head1 := &Node{Val: 1}
	values1 := []int{2, 4}
	for _, val1 := range values1 {
		head1.insertTail(val1)
	}
	fmt.Printf("List 1:\t\t")
	head1.printNodes()

	// Create Input List 2.
	head2 := &Node{Val: 1}
	values2 := []int{3, 4}
	for _, val2 := range values2 {
		head2.insertTail(val2)
	}
	fmt.Printf("List 2:\t\t")
	head2.printNodes()

	newList := mergeTwoLists(head1, head2)
	fmt.Printf("Merged Lists:\t")
	newList.printNodes()
}
