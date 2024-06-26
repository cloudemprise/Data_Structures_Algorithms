package main

import (
	"fmt"
	"strings"
)

/*
Explanation:

The code starts by checking if the input head is nil (empty). If it is, the code
 returns nil, indicating an empty list. Otherwise, it creates a mapping between
 the original nodes and the newly created nodes using a map called copyMap.

Next, the code defines a recursive closure function called copyList. This
 function takes a node as input and returns a copy of the linked list starting
 from that node. The function first checks if the node already exists in the
 copyMap. If it does, it returns the corresponding copy of the node from the
 map. If the node is not in the copyMap, the code creates a new node with the
 same value as the original node and adds it to the copyMap.

Then, it recursively calls copyList on the next node (n.Next) and assigns the
 returned copy to the Next pointer of the newly created node.

After copying all the nodes in the list, the code checks if the original node
 has a random pointer (n.Random). If it does, it assigns the corresponding copy
 of the random node from the copyMap to the Random pointer of the newly created
 node.

Finally, the code returns the copy of the head node, which is the entry point of
 the copied linked list.
*/

/*
Time Complexity: The time complexity of the recursive solution can be analysed
by considering the number of recursive calls made and the work done in each
call. In this case, the code makes a recursive call for each node in the
original linked list. Therefore, the time complexity can be expressed as O(n),
where n is the number of nodes in the linked list.

Space Complexity: The space complexity of the solution is determined by the
additional space used during the recursion. In this case, the code uses a map
called copyMap to store the mapping between the original nodes and the newly
created nodes. The size of this map is proportional to the number of nodes in
the linked list. Therefore, the space complexity can be expressed as O(n),
where n is the number of nodes in the linked list.
*/

// copyRandomList creates a deep copy of a linked list with random pointers
// using a recursive closure.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// Create a mapping between original nodes and newly created nodes.
	copyMap := make(map[*Node]*Node) // Stores [org node] = [new node]

	// Recursive closure.
	var copyList func(*Node) *Node
	copyList = func(n *Node) *Node {
		if n == nil {
			return n
		}

		// If node exists, return it.
		if nodeCopy, ok := copyMap[n]; ok {
			return nodeCopy
		}

		// Allocate new node & create mapping.
		newNode := &Node{Val: n.Val}
		copyMap[n] = newNode

		// Recurse on next (original) node & assign copy to (new) Next pointer.
		newNode.Next = copyList(n.Next)

		// At this stage all nodes are copied making it possible
		// to assign random pointers during recursive unwind.
		if n.Random != nil {
			newNode.Random = copyMap[n.Random]
		}

		return newNode // Tail recursion => newHead.
	}
	// Call recursive closure.
	return copyList(head)
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
