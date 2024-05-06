package main

import (
	"fmt"
	"strings"
)

// Suboptimal Double Loop Strategy.
// Time Complexity : O(n^2)
// Space Complexity: O(n)

// copyRandomList creates a deep copy of an original linked list.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Build a new list w/o populating the random pointers.
	var headNew *Node = &Node{Val: head.Val}
	old, new := head, headNew
	for old.Next != nil {
		old = old.Next                 // Move old forward.
		newNode := &Node{Val: old.Val} // Create a new node.
		new.Next = newNode             // Build new link.
		new = new.Next                 // Move new forward.
	} // Time: O(n) / Space: O(n)

	// Construct stacks of each list.
	var stackNew, stackOld []*Node
	old, new = head, headNew
	for old != nil {
		stackOld = append(stackOld, old)
		stackNew = append(stackNew, new)
		old = old.Next
		new = new.Next
	} // Time: O(n) / Space: O(n)

	// Perform a double loop through the old stack to discover where the
	// random pointers point to, in relation to their stack indices.
	for i, node := range stackOld { // loop old stack

		for j := 0; j < len(stackOld); j++ {

			if node.Random == nil {
				stackNew[i].Random = nil
				break
			} else if *node.Random == *stackOld[j] {
				// old random addresses match so...
				stackNew[i].Random = stackNew[j] // assign new random address
				break
			}
		}
	} // Time: O(n^2) / Space: O(1)

	return headNew
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
