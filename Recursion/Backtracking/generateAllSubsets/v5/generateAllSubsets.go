package main

import (
	"fmt"
)

/*
Subsets are modeled as a subsets queue. Dynamic Programming.
*/

// generateAllSubsets generates all of the subsets of n distinct elements by
// way of a (depth first search)??? algorithm.
func generateAllSubsets(elements []any) (subsets [][]any) {

	// Intialise with the empty set.
	emptySet := []any{}
	subsets = append(subsets, emptySet)

	for _, elem := range elements {

		for _, subset := range subsets {

			// Create a new clone.
			clone := make([]any, len(subset))
			copy(clone, subset)

			// Include current element.
			clone = append(clone, elem)

			// Append new clone.
			subsets = append(subsets, clone)
		}
	}
	return subsets
}

func main() {

	var elements []any

	// Generate the 8 subsets (2^3) that is possible to create with
	// 3 distinct elements.
	elements = []any{'a', 'b', 'c'}
	fmt.Printf("All 8 Subsets of 3 elements: %q\n", elements)
	fmt.Printf("%q\n", generateAllSubsets(elements))
	fmt.Println()

	// Generate the 16 subsets (2^4) that is possible to create with
	// 4 distinct elements.
	elements = []any{'a', 'b', 'c', 'd'}
	fmt.Printf("All 16 Subsets of 4 elements: %q\n", elements)
	fmt.Printf("%q\n", generateAllSubsets(elements))
	fmt.Println()

}
