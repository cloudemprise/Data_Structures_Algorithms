package main

import (
	"fmt"
)

// Time Complexity : O(n)
// Space Complexity: O(n)

// This example can be seen as implementing The Next Greater Element but in reverse.

// nextGreaterElements finds the next greater elements for elements of a
// circular input slice.
func nextGreaterElements(nums []int) []int {
	ln := len(nums)

	results := make([]int, ln) // Preallocate for index management.
	// Initialise to avoid having to include internal logic
	// within the Next Greater Element algorithm.
	for i := range results {
		results[i] = -1
	}

	// Implement a Decreasing Monotonic Stack of INDICES.
	stack := make([]int, 0)
	for i := (2 * ln) - 1; i >= 0; i-- { // Loop backwards twice for circular array.

		/*
			// For DEBUG:
			fmt.Printf("[%v] : stack: \t [", i)
			for j := 0; j < len(stack); j++ {
				fmt.Printf("%2v", nums[stack[j]])
			}
			fmt.Printf(" ] \t\t current = %v \n", nums[i%ln])
		*/

		// For the 1st half of the loop just build a decreasing monotonic stack,
		// then for the 2nd half search the stack until we find an element for
		// which the current element is the Next Greater Element,
		// where current = nums[i%ln].
		for (len(stack)) > 0 && (nums[i%ln] >= nums[stack[len(stack)-1]]) {
			stack = stack[:len(stack)-1]
		}

		// Only store results for the 2nd half of the loop.
		if i < ln && len(stack) > 0 {
			results[i] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, i%ln) // Keep track of the next element.

		// For DEBUG:
		//fmt.Printf("stack: \t %+v \n", *stack)
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
