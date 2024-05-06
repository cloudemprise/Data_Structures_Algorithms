package main

import (
	"fmt"
)

/*
Introduction to Recursive Programming - Rubio-Sanchez
Chapter 12 - Backtracking - Listing 12.4

Permutations are modeled by a list of indices that references the indices of a
list of elements.
*/

// generateAllPermutations generates all of n! permutations of n distinct elements.
func generateAllPermutations(elems []any, soln []int, exists []bool, idx int) {

	if idx == len(elems) { // base case
		// Print all possible n! permutations.
		printPermutations(elems, soln)
	} else {
		// Generate candidate elements
		for j := range elems {
			// Check if candidate already exists in solution.
			if !exists[j] {
				// Include candidate  in partial solution.
				soln[idx] = j
				// Indicate its inclusion.
				exists[j] = true
				// Expand partial solution at position idx+1
				generateAllPermutations(elems, soln, exists, idx+1)
				// Remove candidate inclusion.
				exists[j] = false
			}
		}
	}
}

// printPermutations is a helper function that prints each solution of the
// generateAllPermutations() function as it is found in each base case.
func printPermutations(elems []any, soln []int) {

	fmt.Printf("{ ")
	for i := 0; i < len(soln); i++ {
		fmt.Printf("%c ", elems[soln[i]].(rune)) // type assertion.
	}
	fmt.Println("}")
}

func main() {

	// Generate the 6 permutations (n!) that is possible to create with
	// 3 distinct elements.
	fmt.Println("All 6 Permutations of 3 elements: { a b c }:")
	elements := []any{'a', 'b', 'c'}
	exists := make([]bool, len(elements))  // allocate enough capacity
	solution := make([]int, len(elements)) // allocate enough capacity
	generateAllPermutations(elements, solution, exists, 0)
	fmt.Println()

	//*-*

	// Generate the 24 permutations (n!) that is possible to create with
	// 4 distinct elements.
	fmt.Println("All 24 Permutations of 4 elements: { a b c d }:")
	elements = []any{'a', 'b', 'c', 'd'}
	exists = make([]bool, len(elements))  // allocate enough capacity
	solution = make([]int, len(elements)) // allocate enough capacity
	generateAllPermutations(elements, solution, exists, 0)
	fmt.Println()

}
