package main

import "fmt"

/*
Nested loop, inefficient solution.
Time Complexity = O(n^2)
Space Complexity = O(n)
*/

// Construct map first then search for remainder
func twoSum(nums []int, target int) []int {
	var output []int
	for i := 0; i <= len(nums)-1; i++ {
		remainder := target - nums[i]
		// Search for remainder.
		for j := 0; j <= len(nums)-1; j++ {
			if i == j {
				continue // Bypass same element.
			}
			if nums[j] == remainder {
				// Solution found.
				output = []int{i, j}
				return output
			}
		}
	}
	return []int{} // No solution.
}

///

func main() {

	var input []int
	var target int

	input = []int{2, 7, 11, 15}
	target = 9
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

	input = []int{3, 2, 4}
	target = 6
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

	input = []int{3, 3}
	target = 6
	fmt.Printf("\nInput = %v\n", input)
	fmt.Printf("target = %v\n", target)
	fmt.Println(twoSum(input, target))

}
