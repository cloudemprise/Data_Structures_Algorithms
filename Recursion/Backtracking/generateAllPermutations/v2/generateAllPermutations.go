package main

import (
	"fmt"
)

/*
Introduction to Recursive Programming - Rubio-Sanchez
Chapter 12 - Backtracking - Listing 12.4

Permutations are modeled by a slice of elements that track each permutation
solution.
*/

// generateAllPermutations generates all of n! permutations of n distinct elements.
func generateAllPermutations(elems []any, soln []any, exists map[any]bool) {

	if len(soln) == len(elems) { // base case
		// Print all possible n! permutations.
		printPermutations(elems, soln)
	} else {
		// Generate candidate elements
		for _, e := range elems {
			// Check if candidate already exists in solution.
			if !exists[e] {
				// Include candidate  in partial solution.
				soln = append(soln, e)
				// Indicate its inclusion.
				exists[e] = true
				// Expand partial solution at position idx+1
				generateAllPermutations(elems, soln, exists)
				// Remove candidate inclusion.
				exists[e] = false
				soln = soln[:len(soln)-1]
			}
		}
	}
}

// printPermutations is a helper function that prints each solution of the
// generateAllPermutations() function as it is found in each base case.
func printPermutations(elems []any, soln []any) {

	fmt.Printf("{ ")
	for i := 0; i < len(soln); i++ {
		fmt.Printf("%c ", soln[i].(rune)) // type assertion.
	}
	fmt.Println("}")
}

func main() {

	// Generate the 6 permutations (n!) that is possible to create with
	// 3 distinct elements.
	fmt.Println("All 6 Permutations of 3 elements: { a b c }:")
	elements := []any{'a', 'b', 'c'}
	exists := make(map[any]bool, len(elements)) // allocate enough capacity
	solution := make([]any, 0)                  // allocate no capacity
	generateAllPermutations(elements, solution, exists)
	fmt.Println()

	//*-*

	// Generate the 24 permutations (n!) that is possible to create with
	// 4 distinct elements.
	fmt.Println("All 24 Permutations of 4 elements: { a b c d }:")
	elements = []any{'a', 'b', 'c', 'd'}
	exists = make(map[any]bool, len(elements)) // allocate enough capacity
	solution = make([]any, 0)                  // allocate no capacity
	generateAllPermutations(elements, solution, exists)
	fmt.Println()

}
