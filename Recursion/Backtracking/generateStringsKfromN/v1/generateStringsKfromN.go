package main

import (
	"fmt"
	"math"
)

/*
Data Structures and Algorithic Thinking with Go -Karumanchi
Problem-4 Page 54

Strings are modeled by a list of indices that references the indices of a
element list. This implies element repeats.
*/

// generateStringsKfromN generates all of the k-ary strings of n distinct
// elements.
func generateStringsKfromN(elems []any, soln []int, idx int, k int) {
	if idx == k { // base case
		// Print all possible n^k strings.
		printString(elems, soln, k)
		return
	}
	for j := 0; j < len(elems); j++ {
		soln[idx] = j // overwrite at the index of the solution slice.
		generateStringsKfromN(elems, soln, idx+1, k)
	}
}

// printString is a helper function that prints each solution of the
// generateStringsKfromN() function as it is found in each base case.
func printString(elems []any, soln []int, k int) {
	for i := 0; i < k; i++ {
		fmt.Printf("%c", elems[soln[i]].(rune))
	}
	fmt.Printf("\n")
}

func main() {

	var elems []any
	var soln []int
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
	soln = make([]int, k) // preallocate for index approach.
	generateStringsKfromN(elems, soln, 0, k)
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
	soln = make([]int, k) // preallocate for index approach.
	generateStringsKfromN(elems, soln, 0, k)
	fmt.Println()

	// *-*

}
