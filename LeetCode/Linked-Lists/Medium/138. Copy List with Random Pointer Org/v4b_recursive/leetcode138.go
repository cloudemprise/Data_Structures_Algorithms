package main

/*
Explanation:

See version 4a.
*/

// copyRandomList creates a deep copy of the original linked list.
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
		// to assign random pointers during 1st recursive unwind.
		newNode.Random = copyList(n.Random)

		return newNode // Tail recursion => newHead.
	}
	// Call recursive closure.
	return copyList(head)
}

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

	aList := buildList(&input)
	//aList.printList() // Original print method.
	aList.listPrint() // New print method with random pointer info.

	bList := newList()
	bList.head = copyRandomList(aList.head)
	//bList.printList()
	bList.listPrint()

}

// Object to help build lists for development tests.
type inNode struct {
	val  int
	rand int
}

func buildList(input *[]inNode) *List {
	list := newList()
	var stack []*Node

	// Create a list & stack, w/o populating random pointer.
	for _, n := range *input {

		new := newNode(n.val)      // Create a new node.
		list.insertNode(new)       // Insert into list.
		stack = append(stack, new) // Make a copy on a stack.
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

//*-*
