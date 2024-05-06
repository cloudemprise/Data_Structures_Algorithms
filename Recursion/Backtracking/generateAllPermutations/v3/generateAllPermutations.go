package main

import (
	"fmt"
)

/*
Backtracking Recursive Closure.

Permutations are modeled by a list of indices that references the indices of a
list of elements.
*/

// generateAllPermutations generates all of n! permutations from n distinct elements.
func generateAllPermutations(elems []any) (results [][]any) {

	var backtrack func([]any, []bool, int)
	backtrack = func(solution []any, exists []bool, idx int) {

		if idx == len(elems) { // base case
			// Record all possible n! permutations.
			result := make([]any, len(solution))
			copy(result, solution)
			results = append(results, result)
			return
		}
		// Generate candidate elements
		for i := range elems {
			// Check if candidate already exists in solution.
			if !exists[i] {
				// Include candidate  in partial solution.
				solution[idx] = elems[i]
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
	cache := make([]bool, len(elems))
	// To hold partial solutions.
	soln := make([]any, len(elems))
	// Call backtracking recursive closure.
	backtrack(soln, cache, 0)

	return results
}

// printPermutations type asserts and prints the result of calling generateAllPermutations.
func printPermutations(results [][]any) {

	for _, result := range results {
		fmt.Printf("%q\n", result)
	}
}

func main() {

	var elements []any

	// Generate the 6 permutations (n!) that is possible to create with
	// 3 distinct elements.
	fmt.Println("All 6 Permutations of 3 elements: { a b c }:")
	elements = []any{'a', 'b', 'c'}
	printPermutations(generateAllPermutations(elements))
	fmt.Println()

	//*-*

	// Generate the 24 permutations (n!) that is possible to create with
	// 4 distinct elements.
	fmt.Println("All 24 Permutations of 4 elements: { a b c d }:")
	elements = []any{'a', 'b', 'c', 'd'}
	printPermutations(generateAllPermutations(elements))
	fmt.Println()

}
