package main

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

// A max-heap implementation.
type maxHeap []int

/*
// Implement heap.Interface interface:
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}
*/

/*
// Implement the sort.Interface interface.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(val any)      { (*h) = append((*h), val.(int)) }
func (h *maxHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

//-------------------------------------

func main() {

}

//*-*
