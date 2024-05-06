package main

import (
	"container/heap"
	"fmt"
	"sort"
	"time"
)

/*
Summary:
Quick Select algorithm to find the Kth largest element in the array
*/

/*
Time complexity:
O(N + klogn), where n is the number of elements in the input array nums. The
time complexity is dominated by the heap initialization, which takes O(n) time,
and the subsequent k heap operations, each of which takes O(logn) time.

space complexity:
O(n), where n is the number of elements in the input array nums. This is due to
the space used by the max-heap to store the elements of the input array.
*/

// Golang Standard Lib: container/heap (https://pkg.go.dev/container/heap)
// Implement the heap.Interface interface:

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(val any)      { (*h) = append((*h), val.(int)) }
func (h *maxHeap) Pop() any {
	defer func() { (*h) = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}

///

// findKthLargest returns the Kth largest element of an unordered array.
func findKthLargest(nums []int, k int) int {
	h := maxHeap(nums) // type conversion.
	heap.Init(&h)      // Heapify nums.

	// Search for target.
	for i := k - 1; i > 0; i-- {
		heap.Pop(&h)
	}

	return heap.Pop(&h).(int)
}

///

func main() {

	var nums []int
	var k, result int
	var start time.Time
	var elapsed time.Duration

	nums = []int{3, 2, 1, 5, 6, 4}
	k = 2
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 5 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 4 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

	nums = []int{-1, 2, 0}
	k = 2
	fmt.Println("--- Unsorted --- \n", nums)
	fmt.Printf("k = %v\n", k)

	start = time.Now()
	result = findKthLargest(nums, k)
	elapsed = time.Since(start)

	fmt.Printf("Want: 0 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)

	sort.Ints(nums)
	fmt.Println("--- Sorted --- \n", nums)
	fmt.Println()

	///

}
