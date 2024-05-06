package main

import (
	"fmt"
)

/*
Explanation:

The recursive function buildSumList takes two linked lists (l1 and l2) and a
carry value as input. It returns a new linked list representing the sum of the
input lists.

Base Case:
	If both l1 and l2 are nil and the carry is 0, it means there are no more
	digits to add, so the function returns nil.

Recursive Case:
	If either l1 or l2 is not nil or the carry is non-zero, there are still
	digits to add.
	The function calculates the sum by adding the carry to the values of l1
	and l2 (if they are not nil).
	If l1 is not nil, it moves to the next node by assigning l1 to l1.Next.
	If l2 is not nil, it moves to the next node by assigning l2 to l2.Next.
	The function creates a new node with the digit sum modulo 10 and sets its
	Next pointer to the recursive call of buildSumList with the updated l1, l2,
	and the carry divided by 10.
	The recursion continues until the base case is reached.


Time Complexity:
	O(max(m, n)), where m and n are the lengths of the input linked lists.
Space Complexity:
	O(max(m, n)), as the length of the result list can be at most max(m, n) + 1.
*/

// addTwoNumbers uses recursion to traverse the input linked lists and perform
// the addition digit by digit. It creates a new linked list with the sum digits
// and carries the remaining value to the next recursive call.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// Define a recursive closure.
	var buildSumList func(*ListNode, *ListNode, int) *ListNode
	buildSumList = func(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil
		}

		sum := carry // from previous.

		if l1 != nil { // Accumulate list1 value.
			sum += l1.Val
			l1 = l1.Next // Shift forward.
		}

		if l2 != nil { // Accumulate list2 value.
			sum += l2.Val
			l2 = l2.Next // Shift forward.
		}

		return &ListNode{ // Create a new link with modulo.
			Val:  sum % 10,
			Next: buildSumList(l1, l2, sum/10),
		}
	}
	// Call recursive closure.
	return buildSumList(l1, l2, 0)
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
	in1 = []int{9, 9, 9, 9, 9, 9, 9}
	l1 = newList()
	for _, num := range in1 {
		l1.insertTail(num)
	}
	fmt.Printf("List1: ")
	l1.printList()

	in2 = []int{9, 9, 9, 9}
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
