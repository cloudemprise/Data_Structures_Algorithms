package stack

// Stack definition slice
type Stack []int

func NewStack() *Stack {
	return new(Stack)
}

///

// Push...
func (s *Stack) Push(item int) {
	*s = append(*s, item)
}

// Pop...
func (s *Stack) Pop() int {
	if ln := len(*s); ln > 0 {
		tmp := (*s)[ln-1]
		*s = (*s)[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
}

/* func (s *Stack) Pop() int {
	if ln := len(*s); ln > 0 {
		defer func() {
			*s = (*s)[:ln-1]
		}()
		return (*s)[ln-1]
	}
	panic("pop from empty stack")
} */

///

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	//return !(len(*s) > 0)
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) Reset() {
	*s = (*s)[:0]
	//s = nil
}

///

/* func main() {

	stack := NewStack()

	//stack.Push(10)
	//stack.Push(10)
	//stack.Push(10)
	//stack.Push(10)
	//fmt.Printf("Size of stack = %v \n", len(stack))
	//fmt.Println(stack.Pop())
	//fmt.Printf("Size of stack = %v \n", len(stack))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for _, num := range nums {
		stack.Push(num)
		fmt.Printf("\nPush %v ", stack.Peek())
	}

	for i := 5; i > 0; i-- {
		fmt.Printf("\nPop %v ", stack.Pop())
	}

	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
	stack.Reset()
	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
} */
