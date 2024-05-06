package main

import (
	"container/heap"
	"fmt"
)

// Solution 3: Using a Max-Heap.

/*
Logic Strategy: (Using a Max-Heap)

1. Count the frequency of all numbers in the input array using a map.
2. Build a max-heap based on the count of each number using a custom data
	structure numCount and numCounts.
3. Extract the k most frequent elements from the max-heap.

Time Complexity = O(n) = O(n) + O(n) + O(k)
Space Complexity = O(n) = O(n) + O(n) + O(k)
*/

// Associate each number with their frequency.
type numCount struct {
	num   int
	count int
}

// Create a list of associate objects.
type numCounts []numCount

// Implement heap.Interface interface.
func (h numCounts) Len() int           { return len(h) }
func (h numCounts) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h numCounts) Less(i, j int) bool { return h[i].count > h[j].count }
func (h *numCounts) Push(obj any)      { (*h) = append((*h), obj.(numCount)) }
func (h *numCounts) Pop() any {
	defer func() {
		(*h) = (*h)[:len(*h)-1]
	}()
	return (*h)[len(*h)-1]
}

// topKFrequent returns the k most frequent elements.
func topKFrequent(nums []int, k int) []int {

	// Count the frequency of all nums.
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	} // Time: O(n) / Space: O(n)

	// Build a Max-Heap base on the frequency of each num.
	var nc numCounts
	heap.Init(&nc)
	for num, freq := range freqs {
		heap.Push(&nc, numCount{num, freq})
	} // Time: O(n) / Space: O(n)

	// Extract the k most frequent elements.
	var result []int
	for i := 0; len(result) < k; i++ {
		obj := heap.Pop(&nc).(numCount) // Type cast.
		result = append(result, obj.num)
	} // Time: O(k) / Space: O(k)
	return result
}

///

func main() {

	var nums []int
	var k int

	//*-*

	nums = []int{1, 1, 1, 2, 2, 3}
	k = 2
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output = %v\n", topKFrequent(nums, k))
	fmt.Println()

	//*-*

	nums = []int{1, 1}
	k = 1
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output = %v\n", topKFrequent(nums, k))
	fmt.Println()

	//*-*

	nums = []int{1, 2}
	k = 2
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output = %v\n", topKFrequent(nums, k))
	fmt.Println()

	//*-*

	// Expected: [1,2,5,3,6,7,4,8,10,11]
	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	k = 10
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output:    \t %v\n", topKFrequent(nums, k))
	fmt.Println("Expected:\t [1 2 5 3 6 7 4 8 10 11]")
	fmt.Println()

	//*-*

	// Expected: [1,2,5,3,6,7,4,8,10,11]
	nums = []int{-9, -4, 0, 5, -4, -9, -9, 5, 5, -9, -9, -9, -9}
	k = 3
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output:    \t %v\n", topKFrequent(nums, k))
	//fmt.Println("Expected:\t [1 2 5 3 6 7 4 8 10 11]")
	fmt.Println()

}

//*-*
