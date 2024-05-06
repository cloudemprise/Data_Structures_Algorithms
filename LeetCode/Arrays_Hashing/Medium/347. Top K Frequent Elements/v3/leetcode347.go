package main

import (
	"fmt"
)

// Solution 4: Using Bucket Sort.

/*
Logic:

1. Count the frequency of all numbers: Iterate through the input array nums and
	count the frequency of each number using a map.
	This step has a time complexity of O(n) and a space complexity of O(n).
2. Build a list of buckets: Create a list of buckets, where each bucket at index
	i contains the numbers with a frequency of i.
	This step has a time complexity of O(n) and a space complexity of O(n).
3. Extract the k most frequent elements: Traverse the list of buckets from the
	end and extract the k most frequent elements.
	This step has a time complexity of O(k) and a space complexity of O(k).


Time Complexity:
The overall time complexity of the solution is O(n + n + k) = O(n + k), where n
	is the number of elements in the input array. The dominant factor is the
	initial loop to count the frequency of all numbers, followed by extracting
	the k most frequent elements.

Space Complexity:
The space complexity of the solution is O(n) due to the space used for the
	frequency map and the list of buckets, and O(k) for the result list.
*/

func topKFrequent(nums []int, k int) []int {

	// Count the frequency of all nums.
	freqs := make(map[int]int)
	for _, num := range nums {
		freqs[num]++
	} // Time: O(n) / Space: O(n)

	// Create a slice of slices to act as buckets, where the index of each
	// bucket slice represents the frequency of those numbers, whereby index
	// [0] is ignored. max(freq) <= len(nums) + 1.
	buckets := make([][]int, len(nums)+1)
	//DEBUG fmt.Printf("buckets = %#v\n", buckets)
	//DEBUG fmt.Printf("len(buckets) = %v\n", len(buckets))

	// Place the numbers in their respective bucket based on their frequency.
	for n, freq := range freqs {
		buckets[freq] = append(buckets[freq], n)
	} // Time: O(n) / Space: O(n)
	//DEBUG fmt.Printf("buckets = %#v\n", buckets)
	//DEBUG fmt.Printf("len(buckets) = %v\n", len(buckets))

	// Extract the k most frequent elements.
	var result []int
	for i := len(buckets) - 1; i >= 0; i-- { // Start at far-end of slice.
		if len(result) == k {
			break
		}
		//DEBUG fmt.Printf("buckets[%v] = %#v\n", i, buckets[i])

		// No need to check for empty buckets, since appending a nil slice
		// is a no-op. Append empty buckets, i.e. []int(nil) = nil slice,
		// i.e. if len(buckets[i]) > 0 {result = append(result, buckets[i]...)}
		result = append(result, buckets[i]...)
	} // Time: O(k)) / Space: O(k)

	return result
}

///

func main() {

	var nums []int
	var k int

	//*-*

	nums = []int{1, 1, 1, 2, 3, 3, 4, 4}
	k = 3
	fmt.Printf("Input = %v\n", nums)
	fmt.Printf("k = %v \n", k)
	fmt.Printf("Output = %v\n", topKFrequent(nums, k))
	fmt.Println()

	//*-*

	/* nums = []int{1, 1}
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
	fmt.Println() */

}

//*-*
