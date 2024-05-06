package main

import "fmt"

/*
Time Complexity:  O(n + m) , where n & m are the lengths of the two input lists.
Space Complexity: O(n + m) , due to the recursive call stack.
*/

/*
Instead of modifying the existing nodes like in the v1a solution, this approach
creates new nodes for each smaller element and recursively merges it with the
rest of the two lists.

Key Differences:

    - Node Creation: This solution creates new nodes for each merged element
	 instead of modifying existing nodes.

    - Return Type: The function returns a pointer to a newly created node for
	 each merge step.
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

	// Select the smaller valued element and return a new node with this smaller
	// value, together with the rest of the two merged lists.
	if list1.Val <= list2.Val {
		return &Node{list1.Val, mergeTwoLists(list1.Next, list2)}
	} else {
		return &Node{list2.Val, mergeTwoLists(list1, list2.Next)}
	}
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
