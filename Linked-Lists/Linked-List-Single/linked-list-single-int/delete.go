package list

// DeleteHead removes the Head node if one exists.
func (l *List) DeleteHead() {

	if l.head != nil { // Check for empty list.

		current := l.head
		l.head = current.next
	}
}

// DeleteTail removes the Tail node if one exists.
func (l *List) DeleteTail() {

	if l.head != nil { // Ignore empty lists.

		current := l.head
		if current.next == nil { // Only one node in list.
			l.head = nil
		} else {
			var previousNode *node // Keep track of previous node.
			for current.next != nil {
				previousNode = current
				current = current.next
			}
			previousNode.next = nil // Discard Tail.
		}
	}
}

// DeleteWithin removes a node (if it exists, otherwise is ignored) at a
// particular position within a linked list, i.e. 1 <= pos <= len(List)
func (l *List) DeleteWithin(pos int) {

	/* if l.head == nil {
		log.Fatal("DeleteWithin() : empty list")
	}
	if pos < 1 || pos > l.Length() {
		log.Fatal("DeleteWithin() : out of bounds")
	} */

	// Ignore empty lists & out of bounds.
	if l.head != nil && pos > 0 && pos <= l.Length() {

		if pos == 1 { // At head so just remove it.
			l.head = l.head.next
		} else {
			// NOTE!
			// This algorithm requires locating the node in front of the
			// Position-Of-Deletion, so as to swap out the associated link that
			// removes the node. Therefore:
			pos--
			current := l.head
			for pos > 1 { // while pos > 1
				current = current.next
				pos--
			}
			delNode := current.next
			current.next = current.next.next
			// Avoid dangling nodes for Garbage Collector.
			delNode.next = nil
		}
	}

}
