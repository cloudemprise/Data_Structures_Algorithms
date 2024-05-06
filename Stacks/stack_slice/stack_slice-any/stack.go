package stack

// Stack definition slice
type Stack []any

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(item any) {
	*s = append(*s, item)
}

// Pop...
func (s *Stack) Pop() any {
	if ln := len(*s); ln > 0 {
		tmp := (*s)[ln-1]
		*s = (*s)[:ln-1]
		return tmp
	}
	panic("pop from empty stack")
}

/* func (s *Stack) Pop() any {
	if ln := len(*s); ln > 0 {
		defer func() {
			*s = (*s)[:ln-1]
		}()
		return (*s)[ln-1]
	}
	panic("pop from empty stack")
} */

func (s *Stack) Peek() any {
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
		fmt.Printf("\nPush: \t %v ", stack.Peek().(int))
	}

	fmt.Printf("\n---")

	for i := 5; i > 0; i-- {
		fmt.Printf("\nPop: \t %v ", stack.Pop().(int))
	}

	fmt.Printf("\nSize of Stack = %v\n", stack.Length())

	fmt.Printf("\nPerforming a Reset():")
	stack.Reset()
	fmt.Printf("\nSize of Stack = %v\n", stack.Length())
} */
