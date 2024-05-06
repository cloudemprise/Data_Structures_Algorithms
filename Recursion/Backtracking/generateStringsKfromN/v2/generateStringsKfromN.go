package main

import (
	"fmt"
	"math"
)

/*
Data Structures and Algorithic Thinking with Go -Karumanchi
Problem-4 Page 54

Strings are modeled by a slice of elements that tracks each string solution.
This implies element repeats.
*/

// generateStringsKfromN generates all of the k-ary strings from n distinct
// elements.
func generateStringsKfromN(elems []any, soln []any, k int) {

	if len(soln) == k { // base case
		// Print all possible n^k strings.
		printString(soln)
	} else {
		for _, val := range elems {
			// add element to partial solution.
			soln = append(soln, val.(rune))

			generateStringsKfromN(elems, soln, k)

			// remove element to partial solution
			soln = soln[:len(soln)-1]
		}
	}
}

// printString is a helper function that prints each solution of the
// generateStringsKfromN() function as it is found in each base case.
func printString(soln []any) {

	fmt.Printf("{ ")
	for _, s := range soln {
		fmt.Printf("%c ", s.(rune)) // type assertion.
	}
	fmt.Println("}")
}

func main() {

	var elems []any
	var k int
	var numOfSubsets float64

	// Generate the n^k (9) strings that are possible to create with any
	// 2 elements of 3.
	elems = []any{'a', 'b', 'c'}
	k = 2
	numOfSubsets = math.Pow(float64(len(elems)), float64(k))

	fmt.Printf("Input elements = %c, n = %d\n", elems, len(elems))
	fmt.Printf("Let k = %d\n", k)
	fmt.Printf("Generate all the possible %v subsets: \n", int(numOfSubsets))
	generateStringsKfromN(elems, []any{}, k)
	fmt.Println()

	// *-*

	// Generate the n^k (27) strings that are possible to create with any
	// 3 elements of 3.
	elems = []any{'a', 'b', 'c', 'd'}
	k = 2
	numOfSubsets = math.Pow(float64(len(elems)), float64(k))

	fmt.Printf("Input elements = %c, n = %d\n", elems, len(elems))
	fmt.Printf("Let k = %d\n", k)
	fmt.Printf("Generate all the possible %v subsets: \n", int(numOfSubsets))
	generateStringsKfromN(elems, []any{}, k)
	fmt.Println()

	// *-*

}
