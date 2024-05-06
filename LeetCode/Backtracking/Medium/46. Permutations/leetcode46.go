package main

import (
	"fmt"
)

/*
46. Permutations - Medium

Given an array nums of distinct integers, return all the possible permutations.
You can return the answer in any order.


Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

    1 <= nums.length <= 6
    -10 <= nums[i] <= 10
    All the integers of nums are unique.
*/

// permute generates all of n! permutations from n distinct elements.
func permute(nums []int) (results [][]int) {

	// Define a backtracking recursive closure.
	var backtrack func([]int, []bool, int)
	backtrack = func(solution []int, exists []bool, idx int) {

		if idx == len(nums) { // base case
			// Record each possible permutation.
			result := make([]int, len(solution))
			copy(result, solution)
			results = append(results, result)
			return
		}

		// Generate candidate elements
		for i := range nums {
			// Check if candidate already exists in solution.
			if !exists[i] {
				// Include candidate  in partial solution.
				solution[idx] = nums[i]
				// Indicate its inclusion.
				exists[i] = true
				// Expand partial solution at position idx+1
				backtrack(solution, exists, idx+1)
				// Remove candidate inclusion.
				exists[i] = false
			}
		}
	}
	// Indicate element inclusion in partial solution.
	cache := make([]bool, len(nums))
	// To hold partial solutions.
	soln := make([]int, len(nums))
	// Call backtracking recursive closure.
	backtrack(soln, cache, 0)

	return results
}

//-------------------------------------

func main() {

	//fmt.Println("Hello World")

	var nums []int

	nums = []int{1, 2, 3}

	fmt.Println(permute(nums))

}

//*-*
