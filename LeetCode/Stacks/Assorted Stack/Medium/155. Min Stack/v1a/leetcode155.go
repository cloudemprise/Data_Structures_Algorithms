package main

import "fmt"

// This solution involves implementing a Minims Stack that runs concurrently
// alongside a Values Stack, where the index of the next minimum value is stored
// on this Minims Stack so that the GetMin() method can directly reference the
// current minimum value of the Values Stack in O(1) time.

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

// Push pushes a value on the Values Stack & then checks if it is a new minimum.
// If so, it pushes the top index of Values Stack on the Minims Stack.
func (s *MinStack) Push(val int) {
	s.data = append(s.data, val) // push val on Values Stack

	// is val now a new minimum?
	if len(s.minims) == 0 || val < s.data[s.minims[len(s.minims)-1]] {
		// yes, so push its index on Minims Stack
		s.minims = append(s.minims, len(s.data)-1)
	}
	// NOTE:
	// only the initial minimum is remembered
	// repeated minimums are ignored because of < inequality
}

// Pop discards the top Minims Stack element if it is currently pointing to the
// top of the Values Stack element. It then discards the top element of the
// Values Stack.
func (s *MinStack) Pop() {
	// is top element of Minims Stack pointing to top element of Values Stack?
	if s.minims[len(s.minims)-1] == len(s.data)-1 {
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
