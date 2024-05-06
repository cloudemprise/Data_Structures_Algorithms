package main

import "fmt"

// Using recursive closures

/*
Time Complexity: O(n)
---------------

Can be broken down as follows:

- Finding the first middle node : O(n/2)
- Splitting into two halves     : O(1)
- Reversing the second half     : O(n/2)
- Merging the two halves        : O(n/2)

Space Complexity : O(n) - recursive call stack
----------------
*/

// reorderList first finds the 1st middle node and then splits the list into
// two halves. It then reverses the 2nd list and finally merges both halves.
// This is achieved through the use of recursive closures.
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return // ignore < 2 nodes
	}

	// 1. Find the 1st middle node.
	var findMiddle func(*ListNode, *ListNode) *ListNode
	findMiddle = func(slow, fast *ListNode) *ListNode {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return slow
		}
		return findMiddle(slow.Next, fast.Next.Next)
	}
	middleNode := findMiddle(head, head) // closure call

	// 2. Split into two halves
	list1 := head
	list2 := middleNode.Next
	// Terminate list1 correctly
	middleNode.Next = nil

	// 3. Reverse list2
	var reverseList func(*ListNode) *ListNode
	reverseList = func(n *ListNode) *ListNode {
		if n == nil || n.Next == nil {
			return n
		}
		fwdPtr := n.Next               // Keep a forward-ptr copy
		newHead := reverseList(fwdPtr) // Recurse on forward-ptr
		fwdPtr.Next = n                // Flip forward-ptr ptr backward
		n.Next = nil                   // Terminating prevailing next-ptr
		return newHead                 // Tail-Recursion guarantee
	}
	list2 = reverseList(list2) // closure call

	// 4. Merge the two halves.
	var mergeLists func(*ListNode, *ListNode, bool) *ListNode
	mergeLists = func(l1 *ListNode, l2 *ListNode, isOdd bool) *ListNode {
		if l1 == nil {
			return l2
		}
		if l2 == nil {
			return l1
		}

		// Alternate between l1 & l2
		if isOdd {
			l1.Next = mergeLists(l1.Next, l2, false) // assemble from list1
			return l1
		} else {
			l2.Next = mergeLists(l1, l2.Next, true) // assemble from list2
			return l2
		}
	}
	mergeLists(list1, list2, true) // closure call
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
