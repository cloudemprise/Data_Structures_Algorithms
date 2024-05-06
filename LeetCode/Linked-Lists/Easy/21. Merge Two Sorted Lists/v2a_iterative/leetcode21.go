package main

import "fmt"

// mergeTwoLists iteratively merges two sorted lists into a single sorted list.
func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	var head, tail, index *Node
	for list1 != nil || list2 != nil { // Only when both lists are exhausted.
		switch {
		// Evaluation order is important.
		case (list1 != nil && list2 == nil):
			// Size(list1) > Size(list2)
			index = list1      // cpy subsequent node
			list1 = list1.Next // shift ptr forward
			index.Next = nil   // clear nextPtr
		case (list1 == nil && list2 != nil):
			// Size(list2) > Size(list1)
			index = list2      // cpy subsequent node
			list2 = list2.Next // shift ptr forward
			index.Next = nil   // clear nextPtr
		case (list1.Val <= list2.Val):
			index = list1      // cpy subsequent node
			list1 = list1.Next // shift ptr forward
			index.Next = nil   // clear nextPtr
		case (list2.Val < list1.Val):
			index = list2      // cpy subsequent node
			list2 = list2.Next // shift ptr forward
			index.Next = nil   // clear nextPtr
		}

		if head == nil { // 1st node is covered
			head = index
			tail = head
		} else {
			tail.Next = index // Merge subsequent node
			tail = tail.Next  // shift ptr forward
		}
	}
	return head // merged list
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
