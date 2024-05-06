package main

import "fmt"

/*
Time Complexity:  O(n + m) , where n & m are the lengths of the two input lists.
Space Complexity: O(n + m) , due to the recursive call stack.
*/

// mergeTwoLists recursively merges two sorted lists into a single sorted list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	// If either of the lists are empty then return the other list
	// since merging with an empty list results in the same list.
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	// Select the smaller valued element and set its 'Next' pointer to the
	// result of merging the rest of the two lists.
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

///

func main() {

	// Create Input Node 1.
	head1 := &Node{Val: 1}
	values1 := []int{2, 4}
	for _, val1 := range values1 {
		head1.insertTail(val1)
	}
	fmt.Printf("Node 1:\t\t")
	head1.printNodes()

	// Create Input Node 2.
	head2 := &Node{Val: 1}
	values2 := []int{3, 4}
	for _, val2 := range values2 {
		head2.insertTail(val2)
	}
	fmt.Printf("Node 2:\t\t")
	head2.printNodes()

	newList := mergeTwoLists(head1, head2)
	fmt.Printf("Merged Lists:\t")
	newList.printNodes()
}
