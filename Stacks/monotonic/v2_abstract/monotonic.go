package main

import "fmt"

/*
NOTE that this implementation assumes that all elements in the input array are
distinct. If there are duplicates, we need to modify the implementation to
handle them correctly.
*/

// Time Complexity : O(N)
// Space Complexity: O(N)

/*
How to achieve Monotonic Increasing Stack?

To create a Monotonic Increasing Stack, start with an empty stack, then, while
iterating through elements in a sequence, keep removing elements from the stack
as long as they are LARGER than the current element, and push the current
element onto the stack. This process ensures the stack maintains a strictly
increasing order from bottom to top.

+ increasing monotonic stack ==> stack[i] > stack[i-1] ==> current > stack[i-1],
thus pop for current < stack[i-1]:
*/

// increasingStack builds a increasing monotonic stack from a list of input values.
func increasingStack(nums []int) {

	s := newStack()
	for _, current := range nums {

		for s.length() > 0 && s.peek() > current {
			s.pop()
		}
		s.push(current)
	}

	fmt.Printf("The Array: %#v\n", nums)
	fmt.Printf("The Stack: %#v\n", s)
}

///

/*
How to achieve Monotonic Decreasing Stack?

To create a Monotonic Decreasing Stack, start with an empty stack, then, while
iterating through elements in a sequence, keep removing elements from the stack
as long as they are SMALLER than the current element, and push the current
element onto the stack. This process ensures the stack maintains a strictly
increasing order from bottom to top.

+ decreasing monotonic stack ==> stack[i] < stack[i-1] ==> current < stack[i-1],
thus pop for current > stack[i-1]:
*/

// decreasingStack builds a decreasing monotonic stack from a list of input values.
func decreasingStack(nums []int) {

	s := newStack()
	for _, current := range nums {

		for s.length() > 0 && s.peek() < current {
			s.pop()
		}
		s.push(current)
	}

	fmt.Printf("The Array: %#v\n", nums)
	fmt.Printf("The Stack: %#v\n", s)
}

///

func main() {

	input := []int{1, 5, 8, 3, 2, 6, 7, 4}
	increasingStack(input)
	fmt.Println()
	decreasingStack(input)

}
