package list

/*
A Singly Linked List implementation of an int primitive type.

Operations include:

* Delete 	func (l *List) DeleteHead()
			func (l *List) DeleteTail()
			func (l *List) DeleteWithin(pos int)

* Insert: 	func (l *List) InsertHead(val int)
			func (l *List) InsertTail(val int)
			func (l *List) InsertWithin(pos, val int)

* Length	func (l *List) Length() int

* Middle 	func middleNode1st(head *node) *node
			func middleNode2nd(head *node) *node

* Print		func (l *List) PrintList()
			func (n *node) PrintNodesIterateForward()
			func (n *node) PrintNodesRecurseForward()
			func (n *node) PrintNodesRecurseBackward()

* Reverse 	func (l *List) ReverseList3Ptr()
			func (l *List) ReverseListRecursive()
			func (n *node) reverseNodesRecursive() *node
*/

// Basic Definitions.

type List struct {
	head *node
}

type node struct {
	val  int
	next *node
}

func newList() *List {
	return &List{}
}

func newNode(v int) *node {
	return &node{val: v}
}

///

/* func main() {

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

	// Insertion at position:
	fmt.Println("Testing function InsertWithin")
	fmt.Println("---------------------------------")
	fmt.Println()
	ln := aList.Length()
	// Insert a new node at every 2nd node.
	for i := 1; i < ln*2+2; i += 2 {
		newNum := 99 - i/2
		fmt.Printf("Inserting %v at position %v:\n", newNum, i)
		aList.InsertWithin(i, newNum)
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

	// Test DeleteAtPosition
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
	fmt.Println()

	// ReverseList3Ptr()
	fmt.Println("ReverseList3Ptr():")
	aList.ReverseList3Ptr()
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
	fmt.Println()

	// ReverseListRecursive()
	fmt.Println("ReverseListRecursive():")
	aList.ReverseListRecursive()
	aList.PrintList(printOption)
	fmt.Println("List length = ", aList.Length())
} */
