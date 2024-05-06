package main

import "fmt"

/*
The time and space complexity can be analysed as follows:

Time Complexity:
---------------

- Finding the middle node: The slow and fast pointer approach takes O(n/2) time,
which simplifies to O(n) in the worst case.

- Splitting the list into two halves: This operation takes O(1) time as it
involves updating the `Next` pointer of the middle node.

- Reversing the second half of the linked list: This operation iterates through
the second half once, taking O(n/2) time, which simplifies to O(n) in the worst
case.

- Merging the first half and the reversed second half: This operation iterates
through both halves once, taking O(n/2) time, which simplifies to O(n) in the
worst case.

Overall, the time complexity of the solution is O(n).

Space Complexity:
----------------

- The modified solution uses a constant amount of extra space, regardless of the
size of the input linked list.

- It does not require any additional data structures or recursive calls.

- Therefore, the space complexity of the solution is O(1).

* Please note that the time and space complexity analysis assumes that the basic
operations, such as accessing and updating nodes in the linked list, take
constant time.
*/

// reorderList first finds the (1st) middle node and then splits the list into
// two halves. It then reverses the 2nd list and finally merges both halves.
// This is achieved through iterative means as well as the use of a stack to
// reverse the second half list.
func reorderList(head *Node) {
	if head == nil || head.Next == nil {
		return // only 1 or 2 nodes
	}

	// 1. Find the 1st middle node
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	midNode := slow // 1st middle node

	// 2. Split into two halves.
	list1 := head
	list2 := midNode.Next
	midNode.Next = nil // terminate list1 correctly

	// 3. Reverse list2 using a stack.
	stack := []*Node{}
	current := list2
	for current != nil { // push all node pointers onto a stack
		stack = append(stack, current)
		current = current.Next
	}
	list2 = stack[len(stack)-1] // reversed Head

	// Process each and every 'Next' pointer
	current = list2
	for i := len(stack) - 1; i >= 0; i-- {
		current.Next = stack[i] // flip pointer backward
		current = current.Next
	}
	current.Next = nil // terminate new tail correctly

	// 4. Merge the two halves.
	var fwdPtr1, fwdPtr2 *Node
	for list2 != nil {

		// Save the next nodes in both halves
		fwdPtr1 = list1.Next
		fwdPtr2 = list2.Next

		// Connect current node of list1 to current node of list2
		list1.Next = list2
		list1 = fwdPtr1 // Shift list1 pointer forward

		// Connect current node of list2 to current node of list1
		list2.Next = list1
		list2 = fwdPtr2 // Shift list2 pointer forward
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
