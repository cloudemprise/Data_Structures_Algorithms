package main

import "fmt"

/*
How to achieve Monotonic Increasing Stack?

To create a Monotonic Increasing Stack, start with an empty stack, then, while
iterating through elements in a sequence, keep removing elements from the stack
as long as they are smaller than the current element, and push the current
element onto the stack. This process ensures the stack maintains a strictly
increasing order from bottom to top.
*/

func increasingStack(nums []int) {

	stack := []int{}
	for _, n := range nums {

		for len(stack) > 0 && stack[len(stack)-1] > n {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}

	fmt.Printf("The Array: %#v\n", nums)
	fmt.Printf("The Stack: %#v\n", stack)
}

///

func main() {

	input := []int{1, 4, 5, 3, 12, 10}
	increasingStack(input)
}
