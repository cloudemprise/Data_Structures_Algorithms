package main

import (
	"fmt"
)

func removeNthFromEnd(head *Node, n int) *Node {
	index := removeNode(head, n)
	if index == n {
		return head.Next
	}
	return head
}

func removeNode(node *Node, n int) int {
	if node == nil {
		return 0
	}
	index := removeNode(node.Next, n) + 1
	if index == n+1 {
		node.Next = node.Next.Next
	}
	return index
}

///

func main() {

	var aList *List
	var N int

	// Create a list and print it.
	aList = newList()
	for i := range [5]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()
	N = 2
	fmt.Printf("n = %v\n", N)

	aList.head = removeNthFromEnd(aList.head, N)
	fmt.Printf("Resultant List:\t")
	aList.printList()
	fmt.Println()

	//*-*

	// Create a list and print it.
	aList = newList()
	for i := range [1]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()
	N = 1
	fmt.Printf("n = %v\n", N)

	aList.head = removeNthFromEnd(aList.head, N)
	fmt.Printf("Resultant List:\t")
	aList.printList()
	fmt.Println()

	//*-*

}

//*-*
