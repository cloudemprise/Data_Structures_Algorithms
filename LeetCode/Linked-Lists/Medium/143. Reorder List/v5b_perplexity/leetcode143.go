package main

import "fmt"

func reorderList(head *Node) {
	if head == nil || head.Next == nil {
		return
	}

	// Find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the linked list
	prev, curr := slow, slow.Next
	for curr != nil && curr.Next != nil {
		temp := curr.Next
		curr.Next = temp.Next
		temp.Next = prev.Next
		prev.Next = temp
	}

	// Merge the two halves of the linked list
	p1, p2 := head, prev.Next
	for p1 != prev {
		prev.Next = p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = p2.Next
		p2 = prev.Next
	}
}

///

func main() {

	var aList *List
	const amount = 5

	// Create a list and print it.
	aList = newList()
	for i := range [amount]int{} {
		aList.insertTail(i + 1)
	}
	fmt.Printf("Input List:\t")
	aList.printList()

	reorderList(aList.head)
	fmt.Printf("Reordered List:\t")
	aList.printList()

}
