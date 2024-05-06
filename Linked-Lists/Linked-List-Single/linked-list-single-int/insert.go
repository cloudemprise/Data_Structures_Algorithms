package list

import (
	"log"
)

// InsertHead inserts a new node with the given value at Head of list.
func (l *List) InsertHead(val int) {
	newNode := &node{ // Allocation
		val: val,
	}

	if l.head == nil { // Empty list
		l.head = newNode
	} else {
		newNode.next = l.head // Swap out
		l.head = newNode      // Insert
	}
}

// InsertTail inserts a new node at the tail of list.
func (l *List) InsertTail(val int) {
	// Allocate a new node.
	newNode := &node{
		val: val,
	}

	if l.head == nil { // list is empty
		l.head = newNode
	} else {
		current := l.head         // make a working copy
		for current.next != nil { // traverse until the end
			current = current.next
		}
		current.next = newNode // insert at tail
	}
}

// InsertWithin inserts a new node with the given value  at a particular
// position within a linked list, i.e. 1 <= pos <= len(List) + 1.
func (l *List) InsertWithin(pos, val int) {
	if pos < 1 {
		log.Fatal("InsertWithin() : out of bounds")
	}
	if pos > l.Length()+1 { // (+1) Total includes the new node.
		log.Fatal("InsertWithin() : out of bounds")
	}

	newNode := &node{ // Allocation
		val: val,
	}

	if pos == 1 { // Insert at Head
		newNode.next = l.head
		l.head = newNode
	} else {
		// NOTE!
		// This algorithm requires locating the node in front of the Position Of
		// Insertion, so as to swap & replace the associated links that
		// introduce the new node. Therefore:
		pos--
		current := l.head
		for pos > 1 { // while pos > 1
			current = current.next
			pos--
		}
		newNode.next = current.next
		current.next = newNode
	}
}
