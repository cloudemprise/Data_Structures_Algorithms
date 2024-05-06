package main

import (
	"fmt"
	"strings"
)

/*
Explanation:

 In the first step, we create a copy of each node and insert it next to the
 original node. This can be done by iterating through the original linked list
 and creating a new copy node with the same value. We then update the Next
 pointer of the original node to point to the copy node.

 In the second step, we assign the random pointers for the copied nodes. We
 iterate through the original linked list and update the Random pointer of the
 copy node to point to the corresponding copy node of the original node's random
 pointer.

 In the third step, we separate the original and copied nodes into two separate
 lists. We iterate through the modified linked list and update the Next pointers
 of the original nodes and the copied nodes to skip the copied nodes. We also
 keep track of the head of the copied list.

 Finally, we return the head of the copied list.

The time complexity of this solution is O(N) because we iterate through the
linked list three times, each time visiting each node once.
The space complexity is also O(N) because we create a copy of each node,
resulting in a new list with the same number of nodes as the original list.
*/

// copyRandomList creates a deep copy of the original linked list.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Create a copy of each node and slot it into the list
	// following next to the original node.
	current := head
	for current != nil {
		newNode := &Node{Val: current.Val, Next: current.Next}
		current.Next = newNode
		current = current.Next.Next
	}

	// Step 2: Assign random pointers for the copied nodes.
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// Step 3: Separate the original and copied nodes into two separate lists.
	copyHead := head.Next
	old := head // Two-pointer strategy.
	new := copyHead
	for new.Next != nil {
		old.Next = new.Next
		new.Next = old.Next.Next
		old = old.Next
		new = new.Next
	}
	// Terminate the original list correctly.
	old.Next = nil

	return copyHead
}

///

// To help build lists
type inNode struct {
	val  int
	rand int
}

// buildList creates test lists.
func buildList(input *[]inNode) *List {
	list := newList()
	var stack []*Node

	// Create a list & stack, w/o populating random pointer.
	list.head = newNode((*input)[0].val) // initialise list
	stack = append(stack, list.head)     // initialise stack
	index := list.head
	for i := 1; i < len(*input); i++ {
		index.Next = newNode((*input)[i].val) // insert into list
		stack = append(stack, index.Next)     // copy onto stack
		index = index.Next
	}

	// Populate random pointers by cross referencing between the input
	// data and their corresponding stack indices.
	for i, n := range *input {

		if n.rand < 0 {
			stack[i].Random = nil
		} else {
			stack[i].Random = stack[n.rand]
		}
	}

	return list
}

///

// printListSpecial prints the elements of a linked list along with the index
// of their corresponding random pointers.
func (l *List) printListSpecial() {

	var str strings.Builder

	// Push all nodes onto a stack.
	stack := []*Node{}
	current := l.head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// Ascertain the index of the random pointers by comparing their
	// memory addresses.
	str.WriteString("[")
	for _, n := range stack {

		str.WriteString(" [")
		fmt.Fprintf(&str, "%v", n.Val)

		for i := 0; i < len(stack); i++ {

			if n.Random == nil {
				fmt.Fprintf(&str, " nil")
				str.WriteString("]")
				break // found
			} else if *n.Random == *stack[i] { // Compare memory addresses.
				fmt.Fprintf(&str, " %v", i)
				str.WriteString("]")
				break // found
			}
		}
	}
	str.WriteString(" ]")

	fmt.Println(str.String())
}

//

func main() {

	//var input []inNode

	input := []inNode{
		{7, -1},
		{13, 0},
		{11, 4},
		{10, 2},
		{1, 0},
	}
	//fmt.Printf("%T %#v\n", input, input)

	orgList := buildList(&input)
	//orgList.printList() // Original print method.
	orgList.printListSpecial() // New print method with random pointer info.

	cpyList := newList()
	cpyList.head = copyRandomList(orgList.head)
	//cpyList.printList()
	cpyList.printListSpecial()

}
