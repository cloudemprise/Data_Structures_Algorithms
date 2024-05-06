package main

import (
	"fmt"
	"sort"
)

/*
Logic Strategy:

1. Count the frequency of all numbers in the input array using a map.
2. Create a list of associated objects, where each object contains a number and
	its frequency.
3. Sort the list of associated objects by their frequencies in descending order.
4. Extract the k most frequent elements from the sorted list.

Time Complexity = O(nlogn) = O(2n) + O(nlogn) + O(k)
Space Complexity = O(n) = O(3n) + O(k)
*/

// topKFrequent returns the k most frequent elements.
func topKFrequent(nums []int, k int) []int {

	// Count the frequency of all nums.
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	} // Time: O(n) / Space: O(n)

	// Associate each num with their frequency...
	type numCount struct {
		num   int
		count int
	} // by creating a list of associate objects.
	var numCounts []numCount
	for num, freq := range freqs {
		numCounts = append(numCounts, numCount{num, freq})
	} // Time: O(n) / Space: O(n)

	// Sort associate objects by their counts.
	sort.Slice(numCounts, func(i, j int) bool {
		return numCounts[i].count > numCounts[j].count
	}) // Time: O(nlogn) / Space: O(n)

	// Extract the k most frequent elements.
	var result []int
	for i := 0; len(result) < k; i++ {
		result = append(result, numCounts[i].num)
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
