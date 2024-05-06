package main

import (
	"fmt"
	"log"
	"math"
)

/*
A Singly Linked List inplementation of an int primitive type.

Operations include:

* Insert: 	func (l *List) InsertHead(num int)
			func (l *List) InsertTail(num int)
			func (l *List) InsertAtPosition(position, num int)

* Delete 	func (l *List) DeleteHead() error
			func (l *List) DeleteTail() error
			func (l *List) DeleteAtPosition(position int)

* Print		func (l *List) PrintList()
			func (n *node) PrintNodesIterateForward()
			func (n *node) PrintNodesRecurseForward()
			func (n *node) PrintNodesRecurseBackward()

* Length	func (l *List) Length() int

* Reverse 	func (l *List) ReverseList3Ptr()
			func (l *List) ReverseListRecursive()
			func (n *node) ReverseNodesRecursive() *node
*/

// Head of a singly linked list.
type List struct {
	head *node
	tail *node
}

type node struct {
	num  int
	prev *node
	next *node
}

// Length traverses a list forward and returns the node count.
func (l *List) Length() int {
	var length int
	current := l.head
	for current != nil {
		current = current.next
		length++
	}
	return length
}

///

func main() {

	// Print algorithm choice.
	const printOption = 0
	const insertOption = 1

	// Create linked list.
	fmt.Println("Creating a list")
	fmt.Println("---------------")
	fmt.Println()

	nums := []int{1, 2, 3, 4, 5}
	aList := &List{}
	for _, n := range nums {
		// Choose to insert at Head or Tail.
		switch insertOption {
		case 1:
			fmt.Printf("Insert at Head: %v\n", n)
			aList.InsertHead(n)
		default:
			fmt.Printf("Insert at Tail: %v\n", n)
			aList.InsertTail(n)
		}
	}
	fmt.Println()

	// Print linked list characteristics.
	fmt.Println("Reference List:")
	aList.PrintList(printOption)
	fmt.Printf("List length = %v\n", aList.Length())
	fmt.Println()
	fmt.Println()

	// Insertion at postition:
	fmt.Println("Testing function InsertAtPosition")
	fmt.Println("---------------------------------")
	fmt.Println()
	ln := aList.Length()
	// Insert a new node at every 2nd node.
	for i := 1; i < ln*2+2; i += 2 {
		newNum := 99 - i/2
		fmt.Printf("Inserting %v at position %v:\n", newNum, i)
		aList.InsertAtPosition(i, newNum)
		aList.PrintList(printOption)
		fmt.Println("List length =", aList.Length())
		fmt.Println()
	}
	fmt.Println()

	// DeleteHead()
	// ------------
	fmt.Println("Test function DeleteHead:")
	fmt.Println("------------------------")

	var err error
	err = aList.DeleteHead()
	if err != nil {
		log.Fatal(err)
	}

	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
	fmt.Println()
	fmt.Println()

	// DeleteTail()
	// ------------
	fmt.Println("Test function DeleteTail:")
	fmt.Println("------------------------")

	err = aList.DeleteTail()
	if err != nil {
		log.Fatal(err)
	}
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
	fmt.Println()
	fmt.Println()

	// Test DeleteAtPostion
	// --------------------
	fmt.Println("Test function DeleteAtPosition:")
	fmt.Println("------------------------")
	fmt.Println()

	// DeleteAtPosition() at Head
	fmt.Println("DeleteAtPosition(1):")
	aList.DeleteAtPosition(1)
	aList.PrintList(printOption)
	fmt.Println("List length =", aList.Length())
	fmt.Println()

	// DeleteAtPosition() at Head
	fmt.Println("DeleteAtPosition(aList.Length()):")
	aList.DeleteAtPosition(aList.Length())
	aList.PrintList(printOption)
	fmt.Println("List length =", aList.Length())
	fmt.Println()

	// DeleteAtPosition() at In-between
	fmt.Println("DeleteAtPosition(aList.Length()/2):")
	midPos := math.Ceil(float64(aList.Length()) / 2.0)
	aList.DeleteAtPosition(int(midPos))
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
	fmt.Println()
	fmt.Println()

	// Test Reverse List functions
	// --------------------
	fmt.Println("Test Reverse List functions:")
	fmt.Println("---------------------------")
	fmt.Println("TBD")

	/* // ReverseList3Ptr()
	fmt.Println("ReverseList3Ptr():")
	aList.ReverseList3Ptr()
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
	fmt.Println() */

	/* // ReverseListRecursive()
	fmt.Println("ReverseListRecursive():")
	aList.ReverseListRecursive()
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length()) */

}
