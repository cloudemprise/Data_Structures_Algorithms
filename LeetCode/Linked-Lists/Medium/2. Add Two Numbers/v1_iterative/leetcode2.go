package main

import (
	"fmt"
)

/*
Explanation:

Traverse both linked lists simultaneously, starting from the head nodes.
Keep track of the carry while adding the corresponding digits from both lists.
Create a new linked list to store the sum digits.
Move to the next nodes in both lists until both lists are exhausted.
If there is a remaining carry, add a new node with the carry to the result list.
Return the head of the result list.

Time Complexity:
	O(max(m, n)), where m and n are the lengths of the input linked lists.
Space Complexity:
	O(max(m, n)), as the length of the result list can be at most max(m, n) + 1.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{} // dummy (new) head pointer.
	carry := 0           // addition carry.

	// Final carry could potentially add another link.
	current := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry // from previous iteration.

		if l1 != nil { // Accumulate list1 value.
			sum += l1.Val
			l1 = l1.Next // Shift forward.
		}

		if l2 != nil { // Accumulate list1 value.
			sum += l2.Val
			l2 = l2.Next // Shift forward.
		}

		// Create a new link with modulo.
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next // Shift forward.
		carry = sum / 10       // Is there a carry?
	}

	return dummy.Next // Return first node in link.
}

///

func main() {

	var in1, in2 []int
	var l1, l2, l3 *LinkedList

	//*-*

	fmt.Println("Test1:")
	in1 = []int{2, 4, 3}
	l1 = newList()
	for _, num := range in1 {
		l1.insertTail(num)
	}
	fmt.Printf("List1: ")
	l1.printList()

	in2 = []int{5, 6, 4}
	l2 = newList()
	for _, num := range in2 {
		l2.insertTail(num)
	}
	fmt.Printf("List2: ")
	l2.printList()

	l3 = newList()
	l3.head = addTwoNumbers(l1.head, l2.head)
	fmt.Printf("List3: ")
	l3.printList()

	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	in1 = []int{0}
	l1 = newList()
	for _, num := range in1 {
		l1.insertTail(num)
	}
	fmt.Printf("List1: ")
	l1.printList()

	in2 = []int{0}
	l2 = newList()
	for _, num := range in2 {
		l2.insertTail(num)
	}
	fmt.Printf("List2: ")
	l2.printList()

	l3 = newList()
	l3.head = addTwoNumbers(l1.head, l2.head)
	fmt.Printf("List3: ")
	l3.printList()

	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	in1 = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	l1 = newList()
	for _, num := range in1 {
		l1.insertTail(num)
	}
	fmt.Printf("List1: ")
	l1.printList()

	in2 = []int{5, 6, 4}
	l2 = newList()
	for _, num := range in2 {
		l2.insertTail(num)
	}
	fmt.Printf("List2: ")
	l2.printList()

	l3 = newList()
	l3.head = addTwoNumbers(l1.head, l2.head)
	fmt.Printf("List3: ")
	l3.printList()

	fmt.Println()

	//*-*

	fmt.Println("Test4:")
	//in1 = []int{9, 9, 9, 9, 9, 9, 9}
	in1 = []int{9, 9, 9, 9, 9}
	l1 = newList()
	for _, num := range in1 {
		l1.insertTail(num)
	}
	fmt.Printf("List1: ")
	l1.printList()

	in2 = []int{9, 9, 9}
	//in2 = []int{9, 9, 9, 9}
	l2 = newList()
	for _, num := range in2 {
		l2.insertTail(num)
	}
	fmt.Printf("List2: ")
	l2.printList()

	l3 = newList()
	l3.head = addTwoNumbers(l1.head, l2.head)
	fmt.Printf("List3: ")
	l3.printList()

	fmt.Println()

	//*-*

}

//*-*
