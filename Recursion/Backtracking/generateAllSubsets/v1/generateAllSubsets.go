package main

import (
	"fmt"
)

/*
Introduction to Recursive Programming - Rubio-Sanchez
Chapter 12 - Backtracking - Listing 12.2

Subsets are modeled by a boolean list of true false values.
*/

// generateAllSubsets generate all of the subsets of n distinct elements.
// func generateAllSubsets(solution []bool, elements []any) {
func generateAllSubsets(elems []any, soln []bool, idx int) {

	if idx == len(elems) { // base case
		// Print all possible 2^n subsets.
		printSubset(elems, soln)
		return
	} else {
		// Generate candidate elements
		for _, val := range []bool{false, true} {
			// Include candidate  in partial solution.
			soln[idx] = val
			// Expand partial solution at position idx+1
			generateAllSubsets(elems, soln, idx+1)
		}
	}
}

// printSubset is a helper function that prints each solution of the
// generateAllSubsets() function as it is found in each base case.
func printSubset(elems []any, soln []bool) {

	fmt.Printf("{ ")
	for i := 0; i < len(soln); i++ {
		if soln[i] {
			fmt.Printf("%c ", elems[i].(rune)) // type assertion.
		}
	}
	fmt.Println("}")
}

func main() {

	// Generate the 8 subsets (2^3) that is possible to create with
	// 3 distinct elements.
	fmt.Println("All 8 Subsets of 3 elements: { a b c }:")
	elements := []any{'a', 'b', 'c'}
	solution := make([]bool, len(elements))
	generateAllSubsets(elements, solution, 0)
	fmt.Println()

	// Generate the 16 subsets (2^4) that is possible to create with
	// 4 distinct elements.
	fmt.Println("All 16 Subsets of 4 elements: { a b c d }:")
	elements = []any{'a', 'b', 'c', 'd'}
	solution = make([]bool, len(elements))
	generateAllSubsets(elements, solution, 0)
	fmt.Println()
}
