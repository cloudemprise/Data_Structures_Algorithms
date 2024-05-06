package main

import (
	"fmt"
)

/*
78. Subsets - Medium

Given an integer array nums of unique nums, return all possible
subsets
(the power set).

The solution set must not contain duplicate subsets. Return the solution in
any order.


Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
    All the numbers of nums are unique.
*/

// subsets generates all of the subsets of n distinct numbers by
// way of a recursive closure.
func subsets(nums []int) (subsets [][]int) {

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

//-------------------------------------

func main() {

	//fmt.Println("Hello World")

	var nums []int

	nums = []int{1, 2, 3}

	fmt.Println(subsets(nums))

}

//*-*
