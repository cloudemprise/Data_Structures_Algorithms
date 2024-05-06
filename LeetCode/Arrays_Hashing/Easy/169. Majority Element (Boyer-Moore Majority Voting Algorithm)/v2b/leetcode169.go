package main

import (
	"fmt"
	"time"
)

/*
Time Complexity = O(n)
Space Complexity = O(1)
*/

// Boyer-Moore Majority Voting Algorithm (optimized).
// This makes the assumption the a solution exists, thus requiring no
// verification stage.
func majorityElement(nums []int) int {

	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {

		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

///

func main() {

	var nums []int
	var result int
	var start time.Time
	var elapsed time.Duration

	nums = []int{3, 2, 3}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: 3 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	///

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: 2 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

	///

	nums = []int{1, 2, 1, 2, 0}
	fmt.Println("--- Input --- \n", nums)

	start = time.Now()
	result = majorityElement(nums)
	elapsed = time.Since(start)

	fmt.Printf("Want: 2 Got: %v\n", result)
	fmt.Printf("Execution Time : %v\n", elapsed)
	fmt.Println()

}
