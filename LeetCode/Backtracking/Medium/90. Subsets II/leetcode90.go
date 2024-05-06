package main

import (
	"fmt"
	"sort"
)

/*
90. Subsets II - Medium

Given an integer array nums that may contain duplicates, return all possible
subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any
order.


Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

    1 <= nums.length <= 10
    -10 <= nums[i] <= 10
*/

// subsetsWithDup generates all possible unique subsets of the input elements
// where the input element list may contain repeated elements.
// Time Complexity: O(2^n) + O(nlogn)
// Space Complexity: O(2^n)
func subsetsWithDup(nums []int) (subsets [][]int) {

	// Define backtracking closure.
	var backtrack func([]int, int)
	backtrack = func(current []int, idx int) {

		// Clone the current solution & append to result.
		subset := make([]int, len(current))
		copy(subset, current)
		subsets = append(subsets, subset)

		// Loop through all elements of nums.
		for i := idx; i < len(nums); i++ {

			// Ignore duplicates.
			if (i > idx) && (nums[i] == nums[i-1]) {
				continue
			}
			// Add current element to partial solution.
			current = append(current, nums[i])

			// Recurse & proceed to next element of nums.
			backtrack(current, i+1)

			// Remove current element from partial solution.
			current = current[:len(current)-1]
		}
	}

	// sort input numbers.
	sort.Ints(nums)
	soln := make([]int, 0)
	backtrack(soln, 0)
	return subsets
}

//-------------------------------------

func main() {

	//fmt.Println("Hello World")

	/* nums = []int{1, 2, 3}
	fmt.Println(subsetsWithDup(nums)) */

	var nums []int
	var want [][]int
	var got [][]int

	//*-*

	fmt.Println("Test1:")
	nums = []int{1, 2, 3}
	fmt.Printf("nums = %#v \n", nums)
	want = [][]int{
		{},
		{3},
		{2},
		{2, 3},
		{1},
		{1, 3},
		{1, 2},
		{1, 2, 3},
	}
	got = subsetsWithDup(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test1:")
	nums = []int{1, 2, 2}
	fmt.Printf("nums = %#v \n", nums)
	want = [][]int{
		{},
		{1},
		{1, 2},
		{1, 2, 2},
		{2},
		{2, 2},
	}
	got = subsetsWithDup(nums)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

}

//*-*
