package main

import (
	"fmt"
	"log"
)

// Delete Head node.
func (l *List) DeleteHead() error {
	if l.head == nil {
		log.Fatal("DeleteHead() empty list")
	}

	current := l.head
	l.head = current.next
	return nil
}

// Delete Tail node.
func (l *List) DeleteTail() error {
	if l.head == nil {
		return fmt.Errorf("empty list")
	}

	current := l.head
	if current.next == nil {
		l.head = nil
	} else {
		var previousNode *node
		for current.next != nil {
			previousNode = current
			current = current.next
		}
		previousNode.next = nil
	}
	return nil
}

// Delete node at position.
func (l *List) DeleteAtPosition(pos int) {
	if l.head == nil {
		log.Fatal("DeleteAtPosition() : empty list")
	}
	if pos < 1 || pos > l.Length() {
		log.Fatal("DeleteAtPosition() : out of bounds")
	}

	if pos == 1 { // at head so just remove it
		l.head = l.head.next
	} else {
		// Search for one before delete position.
		current := l.head // make a copy
		for (pos - 1) > 1 {
			current = current.next
			pos--
		}
		// Avoid dangling nodes for GC.
		delNode := current.next
		current.next = current.next.next
		delNode.next = nil
	}
}
