package main

import (
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
O(n) + O(m) + O(m) where m=max-min+1

Space complexity:
O(m) where m=max-min+1


The time complexity of the findKthLargest function is O(n + 2m), where n is the
number of elements in the input array nums, and m is the difference between the
maximum and minimum elements in the array. The time complexity is dominated by
the two loops that iterate through the nums array and the count array, both
taking O(n) and O(m) time respectively.
.
The space complexity of the findKthLargest function is O(m), where m is the
difference between the maximum and minimum elements in the input array nums. This
is due to the space used by the count array, which has a length of max-min+1.
*/

///

// findKthLargest returns the Kth largest element of an unordered array.
func findKthLargest(nums []int, k int) int {

	// Find the max & min values.
	max, min := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	} // Time: O(n) / Space: O(1)

	// Allocate slice of length max-min+1
	count := make([]int, max-min+1)
	for i := 0; i < len(nums); i++ {
		// Use slice index to indicate value existence.
		count[nums[i]-min]++
	} // Time: O(m) / Space: O(m) where m=max-min+1

	// Search for the kth largest index.
	var result int
	for i := len(count) - 1; i >= 0 && k > 0; i-- {
		// If the value exists...
		if count[i] > 0 {
			// note its value.
			result = i + min
			// minus how many are in existence.
			k -= count[i]
		}
	} // Time: O(m) / Space: O(1) where m=max-min+1

	return result
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
