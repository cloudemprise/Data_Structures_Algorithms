package main

// stack definition slice
type stack []int

func newStack() *stack {
	return new(stack)
}

///

// Push...
func (s *stack) push(item int) {
	*s = append(*s, item)
}

// Pop...
func (s *stack) pop() int {
	if ln := len(*s); ln > 0 {
		tmp := (*s)[ln-1]
		*s = (*s)[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
}

/* func (s *stack) pop() int {
	if ln := len(*s); ln > 0 {
		defer func() {
			*s = (*s)[:ln-1]
		}()
		return (*s)[ln-1]
	}
	panic("pop from empty stack")
} */

///

func (s *stack) peek() int {
	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	//return !(len(*s) > 0)
	return len(*s) == 0
}

func (s *stack) length() int {
	return len(*s)
}

func (s *stack) reset() {
	*s = (*s)[:0]
	//s = nil
}
