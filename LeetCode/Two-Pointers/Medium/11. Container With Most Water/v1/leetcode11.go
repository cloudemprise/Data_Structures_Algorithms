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

	// Search for maxArea
	maxArea := 0
	for lhs < rhs {
		area := calcArea(lhs+1, rhs+1, h[lhs], h[rhs])

		if area > maxArea {
			maxArea = area
		}

		// Reduce window base on heights(doesn't matter).
		if h[lhs] > h[rhs] {
			rhs--
		} else {
			lhs++
		}
	}
	return maxArea
}

// calcArea calculates the amount of water(area) that can be contained within two vectors.
func calcArea(x1, x2, y1, y2 int) int {
	width := x2 - x1
	return width * min(y1, y2) // minimum height.
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
