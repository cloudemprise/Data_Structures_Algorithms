package main

import (
	"fmt"
	"sort"
)

/*
Logic Strategy:

1. Count the frequency of all numbers in the input array using a map.
2. Create a hashmap of frequencies, where the key is the frequency and the value
 is a list of numbers with that frequency. Also, maintain a list of all
 frequencies.
3. Sort the list of all frequencies in decreasing order.
4. Extract the results from the frequency map based on the sorted frequencies
 until k elements are obtained.

Time Complexity = O(nlogn) = O(2n) + O(nlogn) + O(k)
Space Complexity = O(n) = O(2n) + O(n) + O(k)
*/

// topKFrequent returns the k most frequent elements.
func topKFrequent(nums []int, k int) []int {

	// Count the frequency of all nums.
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	} // Time: O(n) / Space: O(n)

	// Create a hashmap of rates & a list of all rates.
	type aggregates []int
	rates := make(map[int]aggregates)
	var allRates []int
	for num, freq := range freqs {
		rates[freq] = append(rates[freq], num)
		// A list of all rates.
		allRates = append(allRates, freq) // Unsorted
	} // Time O(n) / Space: O(2n)

	// Sort all rates in decreasing order (k most frequent elements).
	// Time: O(nlogn) / Space: O(n)
	sort.Sort(sort.Reverse(sort.IntSlice(allRates)))

	// Extract the results from these two data.
	var result []int
	currentRate := allRates[0]
	result = append(result, rates[currentRate]...)
	for i := 1; len(result) < k; i++ {
		currentRate = allRates[i]
		if currentRate == allRates[i-1] {
			continue
		}
		result = append(result, rates[allRates[i]]...)
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
