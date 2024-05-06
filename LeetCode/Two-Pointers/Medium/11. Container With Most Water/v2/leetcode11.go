package main

import (
	"fmt"
)

/*
Time Complexity
The time complexity of this solution is O(n), where n is the number of elements
in the input array. The algorithm performs a single pass through the array,
updating the maximum area in each step.

Space Complexity
The space complexity of this solution is O(1), indicating that the amount of
memory used by the algorithm does not increase with the size of the input array.
*/

// maxArea returns the maximum amount of water a container can store designated
// by two vectors defined by a series of array heights of length n.
func maxArea(h []int) int {

	// Two-Pointer Strategy.
	lhs, rhs := 0, len(h)-1

	// Search for maxArea by modulating width of water container.
	maxArea := 0
	for width := len(h) - 1; width > 0; width-- {

		// Reduce window base on heights(doesn't matter).
		if h[lhs] < h[rhs] {
			maxArea = max(maxArea, h[lhs]*width)
			lhs++
		} else {
			maxArea = max(maxArea, h[rhs]*width)
			rhs--
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

///

func main() {

	var input []int

	input = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("Want: 49 \t Got: %v\n", maxArea(input))

	///

	input = []int{1, 2, 1}
	fmt.Printf("Want: 2 \t Got: %v\n", maxArea(input))

}
