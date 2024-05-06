package stack

type Stack struct {
	items []any
}

func New() *Stack {
	return &Stack{}
}

//*-*

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

//*-*

/* // Another way of doing Pop()
func (s *Stack) Pop() any {
	if ln := len(s.items); ln > 0 {
		tmp := s.items[ln-1]
		s.items = s.items[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
} */

func (s *Stack) Pop() any {
	if ln := len(s.items); ln > 0 {
		defer func() {
			s.items = s.items[:ln-1]
		}()
		return s.items[ln-1]
	}
	panic("pop from empty stack")
}

//*-*

func (s *Stack) SizeOf() int {
	return len(s.items)
}

//*-*

func (s *Stack) Top() any {
	if ln := len(s.items); ln > 0 {
		return s.items[ln-1]
	}
	panic("peek into empty stack")
}

//*-*

func (s *Stack) IsEmpty() bool {
	return !(len(s.items) > 0)
}

//*-*

func (s *Stack) Reset() {
	//s.items = s.items[:0]
	s.items = nil
}
