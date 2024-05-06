package list

import (
	"fmt"
)

// PrintList prints the nodes of a singly linked list to stdout.
func (l *List) PrintList(option int) {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}

	switch option {
	case 1:
		l.head.printNodesRecurseForward()
	case 2:
		l.head.printNodesRecurseBackward()
	default:
		l.head.printNodesIterateForward()
	}
	fmt.Println()
}

// printNodesIterateForward prints nodes while iterating forward.
func (n *node) printNodesIterateForward() {
	current := n
	for current != nil { // move past tail node
		fmt.Printf(" %d ", current.val)
		current = current.next
	}
}

// printNodesRecurseForward prints nodes while recursing forward.
func (n *node) printNodesRecurseForward() {
	if n == nil {
		return
	}
	fmt.Printf(" %d ", n.val)
	n.next.printNodesRecurseForward()
}

// printNodesRecurseForward prints nodes while recursing backward.
func (n *node) printNodesRecurseBackward() {
	if n == nil {
		return
	}
	n.next.printNodesRecurseBackward()
	fmt.Printf(" %d ", n.val)
}
