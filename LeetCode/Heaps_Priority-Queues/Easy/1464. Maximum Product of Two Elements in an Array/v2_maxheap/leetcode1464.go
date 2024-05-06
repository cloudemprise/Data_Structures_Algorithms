package main

import (
	"container/heap"
)

// Use stdlib container/heap to implement
// a Max-Heap ADT.

// maxHeap is a max-heap of ints.
type maxHeap []int

// implement sort.Interface
func (h maxHeap) Len() int {
	return len(h)
}

// implement sort.Interface
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// implement sort.Interface
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// implement heap.Interface.
// reslice requires ptr-receiver.
func (h *maxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

// implement heap.Interface.
// reslice requires ptr-receiver.
func (h *maxHeap) Pop() any {
	defer func() {
		*h = (*h)[:len(*h)-1]
	}()
	return (*h)[len(*h)-1]
}

func maxProduct(nums []int) int {
	h := maxHeap(nums) // type conversion.
	heap.Init(&h)
	return (heap.Pop(&h).(int) - 1) * (heap.Pop(&h).(int) - 1)
}

///

func main() {

	PrintResult()
}
