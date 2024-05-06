package list

/*
// To practise the iterative algorithm:
// -----------------------------------

func (l *List) ReverseList3Ptr() {

}
*/

/*
// To practise the recursive algorithm:
// -----------------------------------

func (l *List) ReverseListRecursive() {

}

func (n *node) reverseNodesRecursive() *node {

}
*/

// ReverseList3Ptr reverses a linked list iteratively by implementing a 3
// pointer backflip strategy.
func (l *List) ReverseList3Ptr() {
	// Initialise 3 pointers.
	current := l.head
	var back, front *node
	for current != nil {
		front = current.next // keep a copy of forward node.
		current.next = back  // flip current next ptr backward.
		back = current       // shift forward.
		current = front
	}
	l.head = back // new head
}

///

// ReverseListRecursive reverses a linked list by calling a recursive helper
// method that operates on node objects.
func (l *List) ReverseListRecursive() {

	current := l.head
	l.head = current.reverseNodesRecursive()
}

// reverseNodesRecursive flips backward a forward-reference pointer to the
// prevailing next pointer.
func (n *node) reverseNodesRecursive() *node {
	if n == nil {
		return nil // Empty list!
	}

	if n.next == nil { // Base case.
		return n
	}

	// Create a forward-reference.
	fwdRef := n.next
	// Recurse on forward-reference.
	newHead := fwdRef.reverseNodesRecursive()
	// Flip forward-reference ptr backwards.
	fwdRef.next = n
	// Ensure subsequent recursive loop base case by
	// terminating prevailing next-ptr.
	n.next = nil

	// Tail-Recursion ensures newHead is passed back down
	// the recursive unwind.
	return newHead
}
