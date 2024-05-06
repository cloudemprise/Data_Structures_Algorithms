package main

import (
	"fmt"
)

/*
Subsets are modeled by a boolean list of true false values.
*/

// generateAllSubsets generate all of the subsets of n distinct elements by
// way of a recursive closure.
func generateAllSubsets(elements []any) (subsets []string) {

	// genSubsets is an anonymous backtracking recursive clousure.
	var genSubsets func([]bool, int)
	genSubsets = func(solution []bool, idx int) {
		if idx == len(elements) { // base case
			// Create ouput string from solution.
			var strSoln string
			for i := 0; i < len(solution); i++ {
				if solution[i] {
					// type assertion + conversion.
					strSoln += string(elements[i].(rune))
				}
			}
			subsets = append(subsets, strSoln)
			return
		}

		// Search & create subsets.
		for _, val := range []bool{false, true} {
			// add element to partial solution.
			solution[idx] = val

			genSubsets(solution, idx+1)
		}
	}

	// Call recursive clousure.
	soln := make([]bool, len(elements))
	genSubsets(soln, 0)

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
