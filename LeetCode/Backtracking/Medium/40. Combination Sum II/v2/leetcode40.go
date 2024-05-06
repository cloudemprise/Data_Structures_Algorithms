package main

import (
	"fmt"
	"sort"
)

/*
40. Combination Sum II - Medium

Given a collection of candidate numbers (candidates) and a target number
(target), find all unique combinations in candidates where the candidate
numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.


Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]


Constraints:

    1 <= candidates.length <= 100
    1 <= candidates[i] <= 50
    1 <= target <= 30
*/

// Backtracking recursive closure using a cloned copy of the candidate elements
// "elements" to control which candidate elements to exclude from partial solutions.
// Time Complexity: O(n*2^n) Space: O(n*2^n)
func combinationSum2(candidates []int, target int) (results [][]int) {

	// Define backtracking closure.
	var backtrack func([]int, int, int)
	backtrack = func(combinations []int, sum int, start int) {

		if sum > target {
			return
		}

		if sum == target { // base case
			// Clone current combination.
			result := make([]int, len(combinations))
			copy(result, combinations)
			// Append clone to output.
			results = append(results, result)
			return
		}

		for i := start; i < len(candidates); i++ {

			// Ignore repeated candidates.
			if (i > start) && (candidates[i] == candidates[i-1]) {
				continue
			}

			// Include current candidate element in partial solution.
			combinations = append(combinations, candidates[i])

			backtrack(combinations, sum+candidates[i], i+1)

			// Backtrack: remove extraneous element from partial solution.
			combinations = combinations[:len(combinations)-1]
		}
	}

	// Sort input candidates.
	sort.Ints(candidates)
	// Call backtracking closure.
	backtrack([]int{}, 0, 0)

	return results
}

//-------------------------------------

func main() {

	//fmt.Println("Hello World")

	var candidates []int
	var target int
	var want [][]int
	var got [][]int

	//*-*

	fmt.Println("Test1:")
	candidates = []int{10, 1, 2, 7, 6, 1, 5}
	target = 8
	fmt.Printf("candidates = %#v \t target = %#v\n", candidates, target)
	want = [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}
	got = combinationSum2(candidates, target)
	fmt.Printf("Want: %v \t Got: %v\n", want, got)
	fmt.Println()

	//*-*

}

//*-*
