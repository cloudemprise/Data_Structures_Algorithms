package main

import (
	"fmt"
	"strings"
)

// Suboptimal Double Loop Strategy.
// Time Complexity : O(n^2)
// Space Complexity: O(n)

// copyRandomList creates a deep copy of the original linked list.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Build a new list w/o populating the random
	// pointers but also make a count of its length.
	length := 0
	var headNew *Node = &Node{Val: head.Val} // initialise
	length++
	old, new := head, headNew
	for old.Next != nil {
		old = old.Next                 // Move old forward.
		newNode := &Node{Val: old.Val} // Create a new node.
		new.Next = newNode             // Build new link.
		new = new.Next                 // Move new forward.
		length++
	} // Time: O(n) / Space: O(n)

	// Construct stacks of each list and of length.
	stackNew := make([]*Node, length)
	stackOld := make([]*Node, length)
	old, new = head, headNew
	for i := 0; i < length; i++ {
		stackOld[i] = old
		stackNew[i] = new
		old = old.Next
		new = new.Next
	} // Time: O(n) / Space: O(n)

	// Create a mapping between the indices of nodes in stackOld
	// and to what node index each random pointer goes to.
	mapping := make(map[int]int)
	for i, n := range stackOld {

		if n.Random != nil {

			current := n.Random
			idx := 0
			for current != nil {
				idx++
				current = current.Next
			}
			mapping[i] = (length - idx)
		}
	} // Time: O(n^2) / Space: O(n)

	// Allocate that mapping to the stackNew random pointers.
	for i, n := range stackOld {
		if n.Random == nil {
			stackNew[i].Random = nil
		} else {
			stackNew[i].Random = stackNew[mapping[i]]
		}
	} // Time: O(n) / Space: O(1)

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
