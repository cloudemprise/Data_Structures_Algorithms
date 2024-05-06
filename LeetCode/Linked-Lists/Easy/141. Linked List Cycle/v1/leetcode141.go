package main

import "fmt"

/*
141. Linked List Cycle - Easy

Given head, the head of a linked list, determine if the linked list has a cycle
in it.

There is a cycle in a linked list if there is some node in the list that can be
reached again by continuously following the next pointer. Internally, pos is
used to denote the index of the node that tail's next pointer is connected to.
Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.


Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the
1st node (0-indexed).

Example 2:

Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the
0th node.

Example 3:

Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.


Constraints:

    The number of the nodes in the list is in the range [0, 104].
    -105 <= Node.val <= 105
    pos is -1 or a valid index in the linked-list.


Follow up: Can you solve it using O(1) (i.e. constant) memory?
*/

/*
Time complexity of O(n)
Space complexity of O(n)
*/

func hasCycle(head *Node) bool {

	visited := make(map[*Node]bool)

	current := head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}

///

func main() {

	var (
		input []int
		pos   int
		head  *Node
	)

	fmt.Print("Test1:\n")
	input = []int{3, 2, 0, -4}
	pos = 1
	fmt.Printf("input:\t%v\n", input)
	fmt.Printf("pos:\t%v\n", pos)
	head = createCyclicList(input, pos)
	fmt.Printf("Want: true \t Got: %v\n", hasCycle(head))
	//head.printNodes()
	fmt.Println()

	//*-*

	fmt.Print("Test2:\n")
	input = []int{1}
	pos = -1
	fmt.Printf("input:\t%v\n", input)
	fmt.Printf("pos:\t%v\n", pos)
	head = createCyclicList(input, pos)
	fmt.Printf("Want: false \t Got: %v\n", hasCycle(head))
	//fmt.Printf("Print:\t")
	//head.printNodes()
	fmt.Println()

	//*-*

}

// createCyclicList builds a linked-list with a cyclic pointer.
func createCyclicList(input []int, pos int) *Node {

	dummy := &Node{} // head pointer.
	// Mapping between link index & its memory address.
	nodeIndex := make(map[int]*Node)
	// Insert next link.
	current := dummy
	for i, val := range input {
		current.Next = &Node{Val: val}
		nodeIndex[i] = current.Next
		current = current.Next
	}
	// Create cyclic pointer.
	if pos > 0 {
		current.Next = nodeIndex[pos]
	}

	return dummy.Next
}

//*-*