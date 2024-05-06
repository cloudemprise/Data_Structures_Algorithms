package main

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
