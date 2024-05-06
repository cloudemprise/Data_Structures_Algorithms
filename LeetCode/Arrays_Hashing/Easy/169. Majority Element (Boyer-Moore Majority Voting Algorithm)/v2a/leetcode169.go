package main

import (
	"fmt"
	"time"
)

/*
Time Complexity = O(n)
Space Complexity = O(1)
*/

// Boyer-Moore Majority Voting Algorithm
func majorityElement(nums []int) int {

	// Initialise state variables.
	candidate := nums[0]
	poll := 1

	// 1st Pass: Boyer-Moore Algorithm.
	for i := 1; i < len(nums); i++ {

		if nums[i] == candidate {
			poll++
		} else {
			poll--
			if poll == 0 {
				candidate = nums[i]
				poll++
			}
		}
	}
	// 2nd Pass: Verify.
	count := 0
	for _, num := range nums {
		if candidate == num {
			count++
		}
	}

	if count >= len(nums)/2 {
		return candidate
	} else {
		// Maximum value for the int type
		return int(^uint(0) >> 1)
	}
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
