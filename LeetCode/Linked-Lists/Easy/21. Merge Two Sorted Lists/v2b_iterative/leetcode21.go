package main

import "fmt"

// More readable refactored version of v2a

/*
Time Complexity:  O(n + m) , where n & m are the lengths of the two input lists.
Space Complexity: O(1)
*/

// mergeTwoLists iteratively merges two sorted lists into a single sorted list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	// Initialize three pointers in order to manage the
	// construction of a new merged list.
	var head, tail, temp *Node

	// If either list still has elements.
	for list1 != nil || list2 != nil {

		// Extract the next applicable element by "closing over" its link:
		if list1 == nil {
			temp = list2       // make a local copy of next element
			list2 = list2.Next // remove front element from list2
		} else if list2 == nil {
			temp = list1       // make a local copy of next element
			list1 = list1.Next // remove front element from list1
		} else {
			if list1.Val <= list2.Val {
				temp = list1       // make a local copy of next element
				list1 = list1.Next // remove front element from list1
			} else {
				temp = list2       // make a local copy of next element
				list2 = list2.Next // remove front element from list2
			}
		}
		// Completely "isolate" the temp node by
		// resetting its 'Next' pointer.
		temp.Next = nil

		// If necessary, at first initialise the new Head...
		if head == nil {
			head = temp
			tail = head
		} else {
			// ...or reconstitute the next link of the merged list.
			tail.Next = temp
			tail = tail.Next // shift pointer forward
		}
	}

	return head // The new merged Head.
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
