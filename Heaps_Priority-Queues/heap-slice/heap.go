package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cloudemprise/heap/tree"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >
/*
Heap ADT required operations
----------------------------

1. getParentIdx
2. getLeftIdx
3. getRightIdx
4. hasLargerChild
5. getLargerChildIdx
6. insert : Trickle-Up Algo
7. delete : Trickle-Down Algo
8. getRootValue
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/* // implemented as a struct.
type heap struct {
	items []int
} */

// implemented as a slice
type heap []int

//*-*

// initialise an empty heap.
func newHeap() *heap {
	return &heap{}
}

//-------------------------------------

// return value of Root node.
func (h *heap) getRootValue() int {
	return (*h)[0]
}

//*-*

// return value of Last node.
func (h *heap) getLastValue() int {
	return (*h)[len(*h)-1]
}

//-------------------------------------

// return index of left child.
func (h *heap) getLeftIdx(idx int) int {
	return (idx * 2) + 1 // >= 1
}

//*-*

// return index of right child.
func (h *heap) getRightIdx(idx int) int {
	return (idx * 2) + 2 // >= 2
}

//*-*

// return index of parent node.
func (h *heap) getParentIdx(idx int) int {
	// int division
	return (idx - 1) / 2 // >= 0
}

//-------------------------------------

func (h *heap) insert(data int) {
	// append data into lastNode
	*h = append(*h, data)

	// keep track of newNode index
	idx := len(*h) - 1

	// execute Trickle-Up Algo:
	// while newNode is not rootNode AND newNode > parentNode.
	for idx > 0 &&
		(*h)[idx] > (*h)[h.getParentIdx(idx)] {
		// swap nodes.
		(*h)[h.getParentIdx(idx)], (*h)[idx] =
			(*h)[idx], (*h)[h.getParentIdx(idx)]
		// update newNode index
		idx = h.getParentIdx(idx)
	}
}

//-------------------------------------

func (h *heap) hasLargerChild(idx int) bool {

	// check if left child exists.
	if h.getLeftIdx(idx) <= len(*h)-1 {
		// if so, check if left child is larger than parent.
		if (*h)[h.getLeftIdx(idx)] > (*h)[idx] {
			return true
		}
	}

	// check if right child exists.
	if h.getRightIdx(idx) <= len(*h)-1 {
		// if so, check if right child is larger than parent.
		if (*h)[h.getRightIdx(idx)] > (*h)[idx] {
			return true
		}
	}

	return false // not found
}

//*-*

// return index of larger child, assuming it exists.
func (h *heap) getLargerChildIdx(idx int) int {

	// if no right child.
	if !(h.getRightIdx(idx) <= len(*h)-1) {
		// return left child index.
		return h.getLeftIdx(idx)
	}

	// if right child larger than left child value
	if (*h)[h.getRightIdx(idx)] > (*h)[h.getLeftIdx(idx)] {
		// return right child index.
		return h.getRightIdx(idx)
	} else {
		// return left child index.
		return h.getLeftIdx(idx)
	}
}

//*-*

func (h *heap) delete() {

	// can not delete from empty heap.
	if len(*h) <= 0 {
		panic("empty heap")
	}

	// make lastNode the rootNode.
	(*h)[0] = (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]

	// keep track of trickleNode index.
	idx := 0

	// execute Trickle-Down Algo:
	// while trickleNode has a larger child.
	for h.hasLargerChild(idx) {
		childIdx := h.getLargerChildIdx(idx)
		// swap trickelNode with larger child.
		(*h)[childIdx], (*h)[idx] = (*h)[idx], (*h)[childIdx]
		// swap trickleNode index.
		idx = childIdx
	}
}

//-------------------------------------

// to satisfy Stringer interface
func (h *heap) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < len(*h); i++ {
		fmt.Fprintf(&b, " %v", (*h)[i])
	}
	b.WriteString(" ]")
	return b.String()
}

//*-*

// to get a slice of heap node values.
func (h *heap) getValues() []int {
	var values []int
	for i := 0; i < len(*h); i++ {
		values = append(values, (*h)[i])
	}
	return values
}

//-------------------------------------

func main() {

	var myHeap = newHeap()
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	var nums []int

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 9; i++ {
		nums = append(nums, r.Intn(100))
	}
	//nums = append(nums, r.Perm(10)...)
	fmt.Printf("\n")

	fmt.Printf("Insert the following into a Heap: ")
	for _, num := range nums {
		fmt.Printf(">%v ", num+1)
		myHeap.insert(num + 1)
	}

	fmt.Println()
	fmt.Println("Heap backing-array = ", myHeap)

	//*-*

	myTree := tree.NewTree()
	heapValues := myHeap.getValues()

	for _, value := range heapValues {
		myTree.Insert(value)
	}

	fmt.Println()
	fmt.Println("2D Print Out: ")
	myTree.PrintTree2D()

	//*-*
	fmt.Printf("\n//-------------------------------------\n\n")
	//*-*

	myHeap.delete()
	fmt.Println("Remove element from Heap by deleting root node: ")
	fmt.Println("Heap backing-array = ", myHeap)

	myTree.Reset()
	heapValues = myHeap.getValues()

	for _, value := range heapValues {
		myTree.Insert(value)
	}

	fmt.Println()
	fmt.Println("2D Print Out: ")
	myTree.PrintTree2D()
}

//*-*
