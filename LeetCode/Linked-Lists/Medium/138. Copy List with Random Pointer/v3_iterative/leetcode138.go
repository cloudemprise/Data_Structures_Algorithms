package main

import (
	"fmt"
	"strings"
)

/*
Explanation:

 Step 1: Create a mapping between the original nodes and their corresponding
 copied nodes. This mapping is stored in the copyMap map, where the key is the
 original node and the value is the copied node. Initially, only the values of
 the copied nodes are set, as the pointers might point to nodes that don't
 currently exist.

 Step 2: Traverse the original linked list again and allocate the appropriate
 Next and Random pointers for the copied nodes using the mapping created in
 Step 1. For each node in the original list, we set the Next pointer of the
 copied node to the copied node corresponding to the Next pointer of the
 original node. Similarly, we set the Random pointer of the copied node to the
 copied node corresponding to the Random pointer of the original node.

 Finally, we return the head of the copied list, which is the value stored in
 the copyMap map for the original head node.

This solution effectively creates a deep copy of the original linked list,
ensuring that the Next and Random pointers of the copied nodes correctly point
to the corresponding copied nodes.
*/

// copyRandomList creates a deep copy of the original linked list.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Step 1: Create a mapping between old and newly created
	// value-only copied nodes.
	copyMap := make(map[*Node]*Node) // Stores [org node] = [new node]
	current := head
	for current != nil {
		copyNode := &Node{Val: current.Val}
		copyMap[current] = copyNode
		current = current.Next
	}

	// Step 2: Allocate appropriately .Next & .Random pointers of the newly
	// created copied list within the mapping.
	current = head
	for current != nil {
		copyMap[current].Next = copyMap[current.Next]
		copyMap[current].Random = copyMap[current.Random]
		current = current.Next
	}
	return copyMap[head]
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
