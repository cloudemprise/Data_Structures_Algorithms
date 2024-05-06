package main

import (
	"fmt"
)

// Time complexity: O(n) = O(n) + O(n)
// Space Complexity: O(n)
func longestConsecutive(nums []int) int {
	ln := len(nums)
	if ln < 1 { // corner case.
		return 0
	}

	// Register all nums as hash-table keys.
	cache := make(map[int]bool)
	for _, n := range nums {
		cache[n] = true
	} // Time Complexity: O(n)

	// Search for incremental & decremental sequences.
	max := 0 // Longest sequence.
	for n := range cache {

		// Search for incremental sequence.
		inc := 0
		for j := 1; ; j++ {
			// Sequence found.
			if cache[n+j] {
				inc++              // Increase counter.
				delete(cache, n+j) // Remove key from map.
			} else {
				break // Sequence has ended.
			}
		}

		// Search for decremental sequence.
		dec := 0
		for j := 1; ; j++ {
			// Sequence found.
			if cache[n-j] {
				dec++              // Increase counter.
				delete(cache, n-j) // Remove key from map.
			} else {
				break // Sequence has ended.
			}
		}

		// Check for local maxima.
		if max < dec+inc {
			// Update max so far.
			max = dec + inc
		}

	}
	return max + 1
}

///

func main() {

	var input []int

	input = []int{1, 2, 3, 4}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 4 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{100, 4, 200, 1, 3, 2}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 4 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 9 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 0 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{0}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 1 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{1, 2}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 2 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{1, 2, 0, 1}
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 3 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	//input = []int{-1, -1, 0, 1, 3, 4, 5, 6, 7, 8, 9}
	//sort.Ints(input)
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 7 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

	input = []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}
	//input = []int{-4, -3, -2, -1, 1, 2, 3, 4, 6, 7, 8, 9}
	//sort.Ints(input)
	fmt.Printf("input  = %v\n", input)
	fmt.Printf("Want: 4 Got: %v\n", longestConsecutive(input))
	fmt.Println()

	//*-*

}

//*-*
