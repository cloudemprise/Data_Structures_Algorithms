package main

import (
	"fmt"
	"sort"
)

// Time Complexity : O(nlogn) + O(n^2) = O(n^2)
// Space Complexity: O(n) (for sorting)

// threeSum returns all unique triplets in the array which gives the sum of zero.
// The solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
	ln := len(nums) - 1
	var result [][]int

	// Sort input list.
	sort.Ints(nums) // O(nlogn)

	// Loop over all variables at least once.
	for i := 0; i < ln; i++ {
		// Note: Step-Ahead with [&&] logic.
		// Process 1st element(i=0) but avoid duplicates, as have
		// already considered this combination.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Two-Pointer strategy:
		j, k := i+1, ln
		for j < k {
			if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				result = append(result, []int{
					nums[i],
					nums[j],
					nums[k],
				})
				j++ // Narrow window.
				// Avoid duplicates, as have already considered
				// this combination.
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			}
		}
	} // O(n^2)
	return result
}

///

// all3Combinations prints ALL unique 3 combinations of n nums for illustrative purposes.
func all3Combinations(nums []int) {
	ln := len(nums)

	// collect All unique 3 combinations of n nums.
	combinations := make(map[[3]int]bool)

	var count int
	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			for k := j + 1; k < ln; k++ {
				count++

				result := []int{nums[i], nums[j], nums[k]}
				sort.Ints(result)

				hash := [3]int(result)
				if _, ok := combinations[hash]; !ok {
					combinations[hash] = true
				}
			}
		}
	}

	fmt.Println("The number of all Possible Combinations: ", len(combinations))
	fmt.Println("Calculated with a loop count of: ", count)
	for key := range combinations {
		fmt.Printf("%2v\n", key)
	}
	fmt.Println()
}

///

func main() {

	var input []int

	input = []int{-1, 0, 1, 2, -1, -4}
	fmt.Printf("Input: %v\n", input)
	all3Combinations(input)
	fmt.Printf("Want: [[-1,-1,2],[-1,0,1]] \t Got: %v\n", threeSum(input))
	fmt.Println()

	///

	input = []int{-2, 0, 1, 1, 2}
	fmt.Printf("Input: %v\n", input)
	all3Combinations(input)
	fmt.Printf("Want: [[-2, 0, 2],[-2, 1, 1]] \t Got: %v\n", threeSum(input))
	fmt.Println()

}
