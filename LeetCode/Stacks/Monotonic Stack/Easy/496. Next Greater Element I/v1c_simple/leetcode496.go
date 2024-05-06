package main

import (
	"fmt"
)

// This version just uses a slightly different algorithm for constructing of the
// output list by preallocating & (-1) initialization.

// Time Complexity  = O(n * m)
// Space Complexity = O(n)

/*
Logic Overview:
===============

A decreasing monotonic stack implies:

				stack[i] < stack[i-1], thus

				pop() from stack while:

				current > stack[i-1]

this reveals that 'current' is the next greater value of stack[i-1] if it exists.

Thus by using the reverse logic of constructing a decreasing monotonic stack we
are asking the question: "Is there a number on the stack or have we already seen
a number in the list for which 'current' is its nextGreater value?"

If so then record this pair for later use.

Or in more general terms:

The idea is to store the elements for which we have to find the next greater
element on a stack and while traversing an array. If we find a greater element,
we will pair it with the elements from the stack till the top element of the
stack is less than the current element value.
*/

// nextGreaterElement returns a list of the next greatest elements of nums1[i]
// that occur within the superset of nums2.
func nextGreaterElement(nums1 []int, nums2 []int) []int {

	// Store the nextGreater value (if exists) of each
	// nums1[i] that occurs within the list nums2.
	nextGreater := make(map[int]int) // Space: O(n)

	// Build a decreasing Monotonic Stack and use the reverse logic to determine
	// whether there is a value for which 'current' is the next greater value.
	// See here for the more general case:
	// https://www.geeksforgeeks.org/next-greater-element/
	var stack []int
	for _, current := range nums2 {

		nextGreater[current] = -1 // preallocate default value

		for (len(stack) > 0) && (stack[len(stack)-1] < current) {

			nextGreater[stack[len(stack)-1]] = current // record this pair

			// The element at the top of stack has been solved for and no
			// longer contributes to further solutions, thus discard it.
			stack = stack[:len(stack)-1]
		} // Time: O(m)
		stack = append(stack, current) // push on stack
	} // Time: O(n)

	// Now build the output list:
	result := make([]int, len(nums1)) // Preallocate for faster index management.
	for i, num := range nums1 {       // Cross check those that have a value.

		result[i] = nextGreater[num]
	} // Time: O(n)

	return result
}

///

func main() {

	var nums1 []int
	var nums2 []int

	var want []int
	var got []int

	//*-*

	// This is a good example to run in debug mode to see the logic in action.
	fmt.Println("Test1:")
	nums1 = []int{3, 2, 5, 1, 4}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{4, 5, 2, 1, 3}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 3, -1, 3, 5}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test2:")
	nums1 = []int{4, 1, 2}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{1, 3, 4, 2}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 3, -1}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test3:")
	nums1 = []int{4, 1, 2}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{1, 2, 3, 4}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{-1, 2, 3}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

	fmt.Println("Test4:")
	nums1 = []int{2, 1, 3}
	fmt.Printf("nums1 = %v \n", nums1)
	nums2 = []int{2, 3, 1}
	fmt.Printf("nums2 = %v \n", nums2)

	want = []int{3, -1, -1}
	got = nextGreaterElement(nums1, nums2)
	fmt.Printf("Want: %v \n Got: %v\n", want, got)
	fmt.Println()

	//*-*

}
