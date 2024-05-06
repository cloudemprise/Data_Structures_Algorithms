package main

import (
	"fmt"
	"math"
)

/*
Time complexity = O(n), where n is the number of elements.

Space complexity is O(1) as we are not using any extra space.
*/

// findMaxAverage uses a sliding window of size k to find the maximum sum of a
// sub-array, and then returns the average.
func findMaxAverage(nums []int, k int) float64 {
	if k > len(nums) {
		return math.Inf(-1)
	}

	// Tally the 1st k elements.
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Implement a Sliding Window.
	max := sum
	for i := k; i < len(nums); i++ {
		// Add the next element but subtract the previous element.
		sum = sum + nums[i] - nums[i-k]

		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}

///

func main() {

	var input []int
	var k int

	input = []int{1, 12, -5, -6, 50, 3}
	k = 4
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 12.750 \t Got: %v\n", findMaxAverage(input, k))
	fmt.Println()

	input = []int{5}
	k = 1
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 5.0 \t Got: %v\n", findMaxAverage(input, k))
	fmt.Println()

	input = []int{-5, -3}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: -4 \t Got: %v\n", findMaxAverage(input, k))
	fmt.Println()

	input = []int{-1}
	k = 1
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: -1 \t Got: %v\n", findMaxAverage(input, k))
	fmt.Println()

	input = []int{-1}
	k = 2
	fmt.Printf("Input = %v\n", input)
	fmt.Printf("k = %v\n", k)
	fmt.Printf("Want: 0 \t Got: %v\n", findMaxAverage(input, k))
	fmt.Println()
}
