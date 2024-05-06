package stack

// Stack definition struct
type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{}
}

///

// Push...
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop...
func (s *Stack) Pop() rune {
	if ln := len(s.items); ln > 0 {
		tmp := s.items[ln-1]
		s.items = s.items[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
}

/* func (s *Stack) Pop() rune {
	if ln := len(s.items); ln > 0 {
		defer func() {
			s.items = s.items[:ln-1]
		}()
		return s.items[ln-1]
	}
	panic("pop from empty stack")
} */

///

func (s *Stack) Peek() rune {
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

///

/* func main() {

	// Push a bunch of characters(runes) onto a Stack:
	letters := "abcdefghij"

	stack := NewStack()

	for _, letter := range []rune(letters) {
		stack.Push(letter)
		fmt.Printf("\nPush %v ", stack.Peek())
	}

	for i := 5; i > 0; i-- {
		fmt.Printf("\nPop %v ", stack.Pop())
	}

	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
	stack.Reset()
	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
} */
