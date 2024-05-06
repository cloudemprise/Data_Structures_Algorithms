package main

// copyRandomList creates a deep copy of the original linked list.
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var length int // length of the list

	// Build a new list w/o populating the random pointers.
	var newHead *Node = &Node{Val: head.Val}

	old := head
	new := newHead
	length++
	for old.Next != nil {
		old = old.Next                 // Move old forward.
		newNode := &Node{Val: old.Val} // Create a new node.
		new.Next = newNode             // Build new link.
		new = new.Next                 // Move new forward.
		length++
	}

	// Construct stacks of each list.
	newStack := make([]*Node, length)
	oldStack := make([]*Node, length)

	old = head
	new = newHead
	for i := 0; i < length; i++ {
		oldStack[i] = old
		newStack[i] = new
		old = old.Next
		new = new.Next
	}

	// Correlate the random pointers by discovering where they
	// point to in relation to their stack indices.
	for i, node := range oldStack {

		for j := 0; j < len(oldStack); j++ {

			if node.Random == nil {
				newStack[i].Random = nil
				break
			} else if *node.Random == *oldStack[j] {

				newStack[i].Random = newStack[j]
				break
			}
		}
	}
	return newHead
}

func main() {

	var input []inNode

	input = []inNode{
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

// Object to help build lists.
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
