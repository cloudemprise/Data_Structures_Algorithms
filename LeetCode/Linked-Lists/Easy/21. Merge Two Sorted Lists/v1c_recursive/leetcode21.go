package main

import "fmt"

/*
Time Complexity:  O(n + m) , where n & m are the lengths of the two input lists.
Space Complexity: O(n + m) , due to the recursive call stack.
*/

/*
Key Differences with the other solutions:

    - Node Initialization: This solution initializes a new node at the beginning
	 of each recursive loop and updates it based on comparisons, unlike creating
	 new nodes for each merged element.

    - Assignment of Merged List: The merged list is constructed by updating the
	 Next pointers of the new nodes, which are initially set based on the
	 comparison of the first nodes of the input lists.
*/

// mergeTwoLists recursively merges two sorted lists into a single sorted list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {

	// Initialize a new node to store the next element.
	newNode := new(Node)

	// If both lists have elements then...
	if list1 != nil && list2 != nil {
		// ...select the smaller-valued element and assign it to
		// the new node, then set its 'Next' pointer to the
		// result of merging the rest of the two lists.
		if list1.Val <= list2.Val {
			newNode = list1
			newNode.Next = mergeTwoLists(list1.Next, list2)
		} else {
			newNode = list2
			newNode.Next = mergeTwoLists(list1, list2.Next)
		}
	} else {
		// If one of the input lists is empty then directly assign
		// the non-empty list to the end of the merged list.
		if list1 == nil {
			newNode = list2
		} else {
			newNode = list1
		}
	}
	return newNode // return the new processed node.
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
