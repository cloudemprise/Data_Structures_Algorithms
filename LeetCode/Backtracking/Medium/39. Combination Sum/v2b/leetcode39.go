package main

import (
	"fmt"
)

/*
39. Combination Sum - Medium

Given an array of distinct integers "candidates" and a target integer "target",
return a list of all unique combinations of candidates where the chosen numbers
sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two
combinations are unique if the frequency of at least one of the chosen numbers
is different.

The test cases are generated such that the number of unique combinations that
sum up to target is less than 150 combinations for the given input.


Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]

Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
times. 7 is a candidate, and 7 = 7. These are the only two combinations.

Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:

Input: candidates = [2], target = 1
Output: []


Constraints:

    1 <= candidates.length <= 30
    2 <= candidates[i] <= 40
    All elements of candidates are distinct.
    1 <= target <= 40
*/

// Backtracking recursive closure using an index variable "idx" to control which
// candidate elements to exclude from partial solutions.
func combinationSum(candidates []int, target int) (results [][]int) {

	// Define backtracking closure.
	var backtrack func([]int, int, int)
	backtrack = func(combinations []int, sum int, idx int) {

		if sum == target { // base case
			// Clone current combination.
			result := make([]int, len(combinations))
			copy(result, combinations)
			// Append clone to output.
			results = append(results, result)
			return
		}

		// Loop through all candidate elements.
		for i := idx; i < len(candidates); i++ {

			// Only consider bounded candidate elements.
			if (candidates[i] + sum) <= target {

				// Include current candidate element in partial solution.
				combinations = append(combinations, candidates[i])

				// Recurse and allow repeated elements.
				backtrack(combinations, sum+candidates[i], i)

				// Backtrack: remove extraneous element from partial solution.
				combinations = combinations[:len(combinations)-1]
			}
		}
	}

	// Call backtracking closure.
	soln := make([]int, 0)
	backtrack(soln, 0, 0)
	return results
}

//-------------------------------------

func main() {

	//fmt.Println("Hello World")

	var candidates []int
	var target int
	var want [][]int
	var got [][]int

	fmt.Println("Test1:")
	candidates = []int{2, 3, 6, 7}
	target = 7
	fmt.Printf("candidates = %#v \t target = %#v\n", candidates, target)
	want = [][]int{
		{2, 2, 3},
		{7},
	}
	got = combinationSum(candidates, target)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	candidates = []int{2, 3, 5}
	target = 8
	fmt.Printf("candidates = %#v \t target = %#v\n", candidates, target)
	want = [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	got = combinationSum(candidates, target)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	candidates = []int{2}
	target = 1
	fmt.Printf("candidates = %#v \t target = %#v\n", candidates, target)
	var clone [][]int
	want = clone
	got = combinationSum(candidates, target)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test4:")
	candidates = []int{8, 2, 6, 3}
	target = 13
	fmt.Printf("candidates = %#v \t target = %#v\n", candidates, target)
	want = [][]int{
		{8, 2, 3},
		{2, 2, 2, 2, 2, 3},
		{2, 2, 6, 3},
		{2, 2, 3, 3, 3},
	}
	got = combinationSum(candidates, target)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

}

//*-*
