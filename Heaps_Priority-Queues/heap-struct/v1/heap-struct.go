package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/cloudemprise/heap-struct/tree"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >
/*

Review the pseudocode here: https://en.wikipedia.org/wiki/Binary_heap


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
9. getLastValue
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// Implemented as a struct.
type heap struct {
	items []int
}

//*-*

// Initialise an empty heap.
func newHeap() *heap {
	return &heap{}
}

//-------------------------------------

// Return value of Root node.
func (h *heap) getRootValue() int {
	return h.items[0]
}

//*-*

// Return value of Last node.
func (h *heap) getLastValue() int {
	return h.items[len(h.items)-1]
}

//-------------------------------------

// Return index of left child.
func (h *heap) getLeftIdx(idx int) int {
	return (idx * 2) + 1 // >= 1
}

//*-*

// Return index of right child.
func (h *heap) getRightIdx(idx int) int {
	return (idx * 2) + 2 // >= 2
}

//*-*

// Return index of parent node.
func (h *heap) getParentIdx(idx int) int {
	// Integer division
	return (idx - 1) / 2 // >= 0
}

//-------------------------------------

// Insert newNode into heap.
func (h *heap) insert(data int) {
	// append data into lastNode.
	h.items = append(h.items, data)

	// keep track of newNode index.
	idx := len(h.items) - 1

	// execute Trickle-Up Algo:
	// while newNode is not rootNode AND newNode > parentNode.
	for idx > 0 &&
		h.items[idx] > h.items[h.getParentIdx(idx)] {
		// swap nodes.
		h.items[h.getParentIdx(idx)], h.items[idx] =
			h.items[idx], h.items[h.getParentIdx(idx)]
		// update newNode index
		idx = h.getParentIdx(idx)
	}
}

//-------------------------------------

// Check if parent has a larger child.
func (h *heap) hasLargerChild(idx int) bool {

	// check if left child exists.
	if h.getLeftIdx(idx) <= len(h.items)-1 {
		// check if left child is larger than parent.
		if h.items[h.getLeftIdx(idx)] > h.items[idx] {
			return true
		}
	}

	// check if right child exists.
	if h.getRightIdx(idx) <= len(h.items)-1 {
		// check if right child is larger than parent.
		if h.items[h.getRightIdx(idx)] > h.items[idx] {
			return true
		}
	}

	return false // not found.
}

//*-*

// Return index of larger child, assuming it exists.
func (h *heap) getLargerChildIdx(idx int) int {

	// if no right child
	if !(h.getRightIdx(idx) <= len(h.items)-1) {
		// return left child index.
		return h.getLeftIdx(idx)
	}

	// if right child larger than left child value
	if h.items[h.getRightIdx(idx)] > h.items[h.getLeftIdx(idx)] {
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
	if len(h.items) <= 0 {
		panic("empty heap")
	}

	// make lastNode the rootNode.
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	// keep track of trickleNode index.
	idx := 0

	// execute Trickle-Down Algo:
	// while trickleNode has a larger child.
	for h.hasLargerChild(idx) {
		childIdx := h.getLargerChildIdx(idx)
		// swap trickleNode with larger child.
		h.items[childIdx], h.items[idx] = h.items[idx], h.items[childIdx]
		// swap trickleNode index
		idx = childIdx
	}
}

//-------------------------------------

func (h *heap) String() string {
	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < len(h.items); i++ {
		fmt.Fprintf(&b, " %v", h.items[i])
	}
	b.WriteString(" ]")
	return b.String()
}

//*-*

func (h *heap) getValues() []int {
	var values []int
	for i := 0; i < len(h.items); i++ {
		values = append(values, h.items[i])
	}
	return values
}

//-------------------------------------

func main() {

	var myHeap = newHeap()
	var nums []int

	//*-*

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 11; i++ {
		nums = append(nums, r.Intn(20))
	}

	//*-*

	//nums = append(nums, r.Perm(10)...)

	//*-*

	//nums = []int{1, 2, 3, 4, 5, 6, 7}
	//nums = []int{1, 3, 5, 7, 9, 11, 13, 15}

	//*-*

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

	myHeap.delete()
	fmt.Println()
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
