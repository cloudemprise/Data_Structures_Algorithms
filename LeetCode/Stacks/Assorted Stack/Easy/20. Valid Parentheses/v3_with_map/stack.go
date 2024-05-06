package main

// Stack definition struct
type Stack struct {
	items []byte
}

func NewStack() *Stack {
	return &Stack{}
}

///

// Push...
func (s *Stack) Push(item byte) {
	s.items = append(s.items, item)
}

// Pop...
func (s *Stack) Pop() byte {
	if ln := len(s.items); ln > 0 {
		tmp := s.items[ln-1]
		s.items = s.items[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
}

/* func (s *Stack) Pop() byte {
	if ln := len(s.items); ln > 0 {
		defer func() {
			s.items = s.items[:ln-1]
		}()
		return s.items[ln-1]
	}
	panic("pop from empty stack")
} */

///

func (s *Stack) Peek() byte {
	if ln := len(s.items); ln > 0 {
		return s.items[ln-1]
	}
	panic("peek into empty stack")
}

func (s *Stack) IsEmpty() bool {
	//return !(len(s.items) > 0)
	return len(s.items) == 0
}

func (s *Stack) Length() int {
	return len(s.items)
}

func (s *Stack) Reset() {
	//s.items = s.items[:0]
	s.items = nil
}
