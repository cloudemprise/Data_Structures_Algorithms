package stack

// Reference:
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/generics

// Stack definition struct
type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

///

// Push...
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop but instead of panic return bool.
func (s *Stack[T]) Pop() (T, bool) {
	if ln := len(s.items); ln > 0 {
		defer func() {
			s.items = s.items[:ln-1]
		}()
		return s.items[ln-1], true
	} else {
		var zero T
		return zero, false
	}
}

/* func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		v := new(T)
		return *v, false
	}

	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return v, true
} */

///

func (s *Stack[T]) Peek() (any, bool) {
	if ln := len(s.items); ln > 0 {
		return s.items[ln-1], true
	} else {
		var zero T
		return zero, false
	}
}

func (s *Stack[T]) IsEmpty() bool {
	//return !(len(s.items) > 0)
	return len(s.items) == 0
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}

func (s *Stack[T]) Reset() {
	s.items = s.items[:0]
}

///

/* func main() {

	// Push a bunch of characters(runes) onto a Stack:
	chars := "abcdefghij"

	stack := NewStack[rune]()

	for _, char := range []rune(chars) {
		stack.Push(char)
		val, _ := stack.Peek()
		fmt.Printf("\nPush: \t %c ", val)
	}

	for i := 5; i > 0; i-- {
		val, _ := stack.Pop()
		fmt.Printf("\nPop: \t %c ", val)
	}

	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
	stack.Reset()
	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
} */
