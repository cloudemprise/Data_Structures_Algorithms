package main

import (
	"fmt"
)

// Time Complexity : O(n)
// Space Complexity: O(n)

// This example illustrates how the logic would form without initialising the
// output list: 'results'.

// nextGreaterElements finds the next greater elements for elements of a
// circular input slice.
func nextGreaterElements(nums []int) []int {
	ln := len(nums)
	results := make([]int, ln) // Preallocate for index management.

	// Implement a classic Decreasing Monotonic Stack of INDICES.
	stack := make([]int, 0, ln)
	for i := 0; i < 2*ln; i++ { // Loop forwards twice for circular array.

		i := i % ln // Wrap-around modulo shadow loop counter.

		// If the current element is greater than the element at the top of
		// stack implies that it is the next greatest element for that element.
		for (len(stack) != 0) && (nums[i] > nums[stack[len(stack)-1]]) {

			results[stack[len(stack)-1]] = nums[i] // Record next greatest element.
			// The element at the top of stack has been solved for and no
			// longer contributes to further solutions, thus discard it.
			stack = stack[:len(stack)-1]
		}

		// The stack is empty for elements that have no solution.
		if len(stack) == 0 {
			results[i] = -1
		}

		stack = append(stack, i) // Keep track of the previous element.
	}
	return results
}

///

func main() {

	var nums []int

	var want []int
	var got []int

	///

	fmt.Println("TEST1:")
	nums = []int{1, 2, 1}
	fmt.Printf("nums: \t %2v \n", nums)

	want = []int{2, -1, 2}
	got = nextGreaterElements(nums)
	fmt.Printf("Want: \t %2v \n Got: \t %2v\n", want, got)
	fmt.Println("====================")

	///

	fmt.Println("TEST2:")
	nums = []int{1, 2, 3, 4, 3}
	fmt.Printf("nums: \t %2v \n", nums)

	want = []int{2, 3, 4, -1, 4}
	got = nextGreaterElements(nums)
	fmt.Printf("Want: \t %2v \n Got: \t %2v\n", want, got)
	fmt.Println("====================")

	///

	fmt.Println("TEST3:")
	nums = []int{4, 3, 1, 2, 3}
	fmt.Printf("nums: \t %2v \n", nums)

	want = []int{-1, 4, 2, 3, 4}
	got = nextGreaterElements(nums)
	fmt.Printf("Want: \t %2v \n Got: \t %2v\n", want, got)
	fmt.Println("====================")

}
