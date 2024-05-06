package main

import (
	"fmt"
)

// generateAllSubsets generates all of the subsets of n distinct numbers by
// way of a recursive closure.
func generateAllSubsets(nums []int) (subsets [][]int) {

	// genSubsets is an anonymous backtracking recursive clousure.
	var genSubsets func([]int, int)
	genSubsets = func(current []int, idx int) {

		// Clone current subset & append to subsets result.
		subset := make([]int, len(current))
		copy(subset, current)
		subsets = append(subsets, subset)

		// Loop through elements of nums.
		for i := idx; i < len(nums); i++ {

			// Add current element to partial solution.
			current = append(current, nums[i])

			// Recurse & increment starting loop counter.
			genSubsets(current, i+1)

			// Remove current element from partial solution.
			current = current[:len(current)-1]
		}
	}

	// Call recursive clousure.
	curr := make([]int, 0)
	genSubsets(curr, 0)

	return subsets
}

func main() {

	var nums []int

	// Generate the 8 subsets (2^3) that is possible to create with
	// 3 distinct nums.
	nums = []int{1, 2, 3}
	fmt.Printf("All 8 Subsets of 3 nums: %d\n", nums)
	fmt.Printf("%d\n", generateAllSubsets(nums))
	fmt.Println()

	/* // Generate the 16 subsets (2^4) that is possible to create with
	// 4 distinct nums.
	nums = []int{1, 2, 3, 4}
	fmt.Printf("All 16 Subsets of 4 nums: %d\n", nums)
	fmt.Printf("%d\n", generateAllSubsets(nums))
	fmt.Println() */

}
