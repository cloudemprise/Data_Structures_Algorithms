package main

import (
	"fmt"
	"io"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

// newNode constructs a new node.
func newNode(v int) *ListNode {
	return &ListNode{Val: v}
}

// newList constructs a new list.
func newList() *List {
	return &List{}
}

///

// List method
// insertTail inserts a new node at the tail of a linked-list.
func (l *List) insertTail(val int) {

	if l.head == nil { // list is empty
		l.head = newNode(val) // create a new node
	} else {
		current := l.head         // make a working copy
		for current.Next != nil { // traverse until the end
			current = current.Next
		}
		current.Next = newNode(val) // insert a new node at tail
	}
}

// ListNode method
// insertTail inserts a new node at the tail of a linked-list.
func (head *ListNode) insertTail(v int) {

	current := head // make a working copy
	for current.Next != nil {
		current = current.Next // traverse until the end
	}
	current.Next = &ListNode{Val: v} // insert a new node at tail
}

///

// ListNode method
// printNodes prints the node values of a linked-list to stdout.
func (head *ListNode) printNodes() {
	if head == nil {
		return
	}

	current := head // create a working copy
	for current != nil {
		fmt.Printf("%v ", current.Val)
		current = current.Next // iterate forward thorough the list
	}
	fmt.Println()
}

// List method
// printList prints the node values of a linked-list to stdout.
func (l *List) printList() {
	if l.head == nil {
		fmt.Println("empty list")
		return
	}

	l.head.printNodes() // call node method
}

///

// List method
// scanList makes use of an io.Writer(bytes.Buffer) to catalog the
// node values  of a linked-list.
func (l *List) scanList(w io.Writer) {
	if l.head == nil {
		return
	}

	current := l.head // create a working copy
	for current != nil {
		fmt.Fprintf(w, "%v ", current.Val)
		current = current.Next // iterate forward thorough the list
	}
	fmt.Println()
}

// ListNode method
// scanNodes makes use of an io.Writer(bytes.Buffer) to catalog the
// node values  of a linked-list.
func (head *ListNode) scanNodes(w io.Writer) {
	if head == nil {
		return
	}

	current := head // create a working copy
	for current != nil {
		fmt.Fprintf(w, "%v ", current.Val)
		current = current.Next // iterate forward thorough the list
	}
	fmt.Println()
}
