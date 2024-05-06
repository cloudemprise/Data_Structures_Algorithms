package main

import (
	"log"
)

// Insert node at Head of list.
func (l *List) InsertHead(num int) {
	newNode := &node{ // allocate
		num: num,
	}

	if l.head == nil { // empty list
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
}

// Insert node at Tail of list.
func (l *List) InsertTail(num int) {
	newNode := &node{ // allocate
		num: num,
	}

	if l.head == nil { // empty list
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}
}

// Insert node at position by shifting into one place before insert position.
func (l *List) InsertAtPosition(pos, num int) {
	if pos < 1 {
		log.Fatal("InsertAtPosition() : out of bounds")
	}
	if pos > l.Length()+1 { // (+1) inc. new node
		log.Fatal("InsertAtPosition() : out of bounds")
	}

	if pos == 1 { // At Head
		l.InsertHead(num)
	} else if pos == l.Length()+1 { // At Tail.
		l.InsertTail(num)
	} else { // Not at head
		newNode := &node{ // allocate
			num: num,
		}
		// Search for one before insert position.
		current := l.head // make a copy
		for (pos - 1) > 1 {
			current = current.next
			pos--
		}
		// Insert newNode ahead of current node.
		newNode.prev = current
		newNode.next = current.next
		current.next = newNode
		newNode.next.prev = newNode
	}
}

/* // Insert node at position by shifting
func (l *List) InsertAtPosition(position, num int) {
	if position < 1 {
		log.Fatal("InsertAtPosition() : out of bounds")
	}
	if position > l.Length()+1 { // (+1) inc. new node
		log.Fatal("InsertAtPosition() : out of bounds")
	}

	if position == 1 { // at head
		l.InsertHead(num)
	} else { // at !head

		newNode := &node{ // allocate
			num: num,
		}

		current := l.head // make a copy
		fmt.Printf("pos=%v\n", position)
		for position > 1 { // search insert position
			current = current.next
			position--
			fmt.Printf("pos=%v\n", position)

		}
		fmt.Printf("pos=%v\n", position)

		// Insert newNode ahead of current.
		newNode.prev = current
		newNode.next = current.next
		current.next = newNode
		newNode.next.prev = newNode

	}
} */
