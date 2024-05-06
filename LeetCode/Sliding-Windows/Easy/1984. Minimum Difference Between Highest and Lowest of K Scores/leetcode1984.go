package main

import (
	"fmt"
	"sort"
)

/*
Time Complexity: O(nlogn) + O(n) = O(nlogn)
Space Complexity: O(1)
*/

// minimumDifference returns the minimum difference between the highest & lowest.
func minimumDifference(nums []int, k int) int {
	/* if k < 2 || k > len(nums) {
		return 0
	} // Invalid input. */

	// Sort input list, non-decreasing order.
	sort.Ints(nums) // O(nlogn)

	// Initial difference between highest & lowest.
	diff := nums[k-1] - nums[0]

	// Implement Sliding Window & minimise difference.
	min := diff
	for i := k; i < len(nums); i++ {

		// Recalculate difference between highest & lowest.
		diff = nums[i] - nums[i-(k-1)]

		if diff < min {
			min = diff
		}
	} // O(n)
	return min
}

///

func main() {
	var input []int
	var k int

	input = []int{9, 4, 1, 7}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 2 \t Got: %v\n", minimumDifference(input, k))
	fmt.Println()

	input = []int{90}
	k = 1
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 0 \t Got: %v\n", minimumDifference(input, k))
	fmt.Println()

	input = []int{87063, 61094, 44530, 21297, 95857, 93551, 9918}
	k = 6
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 74560 \t Got: %v\n", minimumDifference(input, k))
	fmt.Println()
}
