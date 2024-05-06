package main

import (
	"fmt"
)

// Solution using two queues.

// Define a q implemented as a queue.
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

// O(n) : requires a queue transversal.
func (s *MyStack) Push(val int) {
	tmp := newQueue()
	tmp.enQueue(val)
	for !s.q.isEmpty() {
		tmp.enQueue(s.q.deQueue())
	}
	s.q = tmp // swap queues
}

// O(1)
func (s *MyStack) Pop() int {
	return s.q.deQueue()
}

// O(1)
func (s *MyStack) Top() int {
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
