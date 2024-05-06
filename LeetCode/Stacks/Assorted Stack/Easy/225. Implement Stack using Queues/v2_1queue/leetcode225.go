package main

import (
	"fmt"
)

// Solution using ONLY one queue.

// Define a stack implemented as a queue.
type MyStack struct {
	q *queue
}

func newStack() *MyStack {
	return &MyStack{
		&queue{},
	}
}

func (s *MyStack) IsEmpty() bool {
	return s.q.isEmpty()
}

// O(1)
func (s *MyStack) Push(val int) {
	s.q.enQueue(val)
}

// O(n) : requires complete queue transversal.
func (s *MyStack) Pop() int {
	for i := 0; i < s.q.length()-1; i++ {
		// reinsert back, except last index
		s.q.enQueue(s.q.deQueue())
	}
	return s.q.deQueue() // last index
}

// O(n) : requires complete queue transversal.
func (s *MyStack) Top() int {
	for i := 0; i < s.q.length()-1; i++ {
		// reinsert back, except last index
		s.q.enQueue(s.q.deQueue())
	}
	defer func() { // requeue last index
		s.q.enQueue(s.q.deQueue())
	}()
	return s.q.peek()
}

///

func main() {

	tStack := newStack()

	fmt.Println("Instantiate a stack: ")
	fmt.Printf("%#v\n", tStack)
	fmt.Println()

	fmt.Println("Push=1:")
	tStack.Push(1)
	fmt.Printf("%#v\n", tStack)
	fmt.Println()

	fmt.Println("Push=2:")
	tStack.Push(2)
	fmt.Printf("%#v\n", tStack)
	fmt.Println()

	fmt.Println("Push=3:")
	tStack.Push(3)
	fmt.Printf("%#v\n", tStack)
	fmt.Println()

	fmt.Printf("Top()\t = %v", tStack.Top())
	fmt.Println()

	fmt.Printf("Pop()\t = %v", tStack.Pop())
	fmt.Println()

	fmt.Printf("IsEmpty()\t = %v", tStack.IsEmpty())
	fmt.Println()

	fmt.Printf("Pop()\t = %v", tStack.Pop())
	fmt.Println()

	fmt.Printf("Pop()\t = %v", tStack.Pop())
	fmt.Println()

	fmt.Printf("IsEmpty()\t = %v", tStack.IsEmpty())
	fmt.Println()
}
