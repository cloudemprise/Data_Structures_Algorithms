package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func newNode(v int) *Node {
	return &Node{Val: v}
}

type List struct {
	head *Node
}

func newList() *List {
	return &List{}
}

func (l *List) insertNode(n *Node) {
	if l.head == nil {
		l.head = n
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = n
	}
}

/* func (l *List) insertTail(v int) {
	if l.head == nil {
		l.head = newNode(v)
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode(v)
	}
} */

/* func (l *List) scanList(w io.Writer) {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Fprintf(w, "%v ", current.Val)
		current = current.Next
	}
	fmt.Println()
} */

/* func (head *Node) insertTail(v int) {
	if head == nil {
		head = &Node{Val: v}
	} else {
		current := head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &Node{Val: v}
	}
} */

/* func (head *Node) scanNodes(w io.Writer) {
	if head == nil {
		return
	}
	current := head
	for current != nil {
		fmt.Fprintf(w, "%v ", current.Val)
		current = current.Next
	}
	fmt.Println()
} */

/* func (head *Node) printNodes() {
	if head == nil {
		return
	}
	current := head
	for current != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next
	}
	fmt.Println()
} */

/* func (l *List) printList() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf(" %v", current.Val)
		current = current.Next
	}
	fmt.Println()
} */

// listPrint prints the elements of a linked list along with the index their
// corresponding random pointers.
func (l *List) listPrint() {

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
