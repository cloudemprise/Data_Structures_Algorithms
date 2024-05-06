package main

import "fmt"

// This solution is a Leetcode Refactored Optimisation. The difference is the
// subtle logic between using the inequality of '<=' vs '<' in the Push() method.
// This particular solution allows for the use of the Top() as well as the
// GetMin() methods to formulate the Pop() method, which is of itself a
// substantial simplification of the code.

// Time Complexity  : O(1)
// Space Complexity : O(n)

type MinStack struct {
	data   []int // Values Stack
	minims []int // Minims Stack: indices of the next minimum value.
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1]
}

func (s *MinStack) GetMin() int {
	// dereference Minims Stack top element (index of the current minimum)
	return s.data[s.minims[len(s.minims)-1]]
}

///

// NOTICE the difference between the subtle logic of v1a & v1b: '<=' vs '<' and
// how that effects the resultant Pop() method.

// Push pushes a value on the Values Stack & then checks if it is LESS THAN OR
// EQUAL TO the current minimum. If so, it pushes the top index of Values Stack
// on the Minims Stack.
func (s *MinStack) Push(val int) {
	s.data = append(s.data, val) // push val on Values Stack

	// is val <= minimum? (!!!)
	if len(s.minims) == 0 || val <= s.GetMin() {
		// yes, so push its index on Minims Stack
		s.minims = append(s.minims, len(s.data)-1)
	}
}

// Pop check if the top element of the Values Stack is equal to the current
// minimum value. If so, it then discards top element of the Minims Stack. It
// then discards the top element of the Values Stack.
func (s *MinStack) Pop() {
	// are values equal? (Notice the use of these methods)
	if s.Top() == s.GetMin() {
		// yes, so remove top element of Minims Stack
		s.minims = s.minims[:len(s.minims)-1]
	}
	s.data = s.data[:len(s.data)-1] // remove top element of Values Stack
}

///

func main() {

	myStack := Constructor()
	myStack.Push(-2)
	myStack.Push(0)
	myStack.Push(-3)
	fmt.Printf("getMin: %v\n", myStack.GetMin())
	myStack.Pop()
	fmt.Printf("top: %v\n", myStack.Top())
	fmt.Printf("getMin: %v\n", myStack.GetMin())

	///

	/* myStack := Constructor()
	myStack.Push(0)
	myStack.Push(1)
	myStack.Push(0)
	fmt.Printf("getMin: %v\n", myStack.GetMin())
	myStack.Pop()
	fmt.Printf("top: %v\n", myStack.Top())
	fmt.Printf("getMin: %v\n", myStack.GetMin()) */
}
