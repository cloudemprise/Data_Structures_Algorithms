package main

import (
	"fmt"
	"strconv"
)

/*
Time complexity: O(n), where n is the number of tokens in the input list. The
algorithm makes a single pass through the list of tokens, performing O(1)
operations for each token.

Space complexity: O(n), where n is the number of tokens in the input list.
The space used by the stack is linear in the number of tokens.
*/

// evalRPN returns the integer result of evaluating an arithmetic string
// expression in Reverse Polish Notation.
func evalRPN(tokens []string) int {

	s := newStack()
	for _, t := range tokens {

		switch t {
		case "+":
			s.push(s.pop() + s.pop()) // push result back onto stack
		case "-":
			val1, val2 := s.pop(), s.pop() // NOTE: operand precedence
			//val1 := s.pop()
			//val2 := s.pop()
			s.push(val2 - val1) // push result back onto stack
		case "/":
			val1, val2 := s.pop(), s.pop() // NOTE: operand precedence
			//val1 := s.pop()
			//val2 := s.pop()
			s.push(val2 / val1) // push result back onto stack
		case "*":
			s.push(s.pop() * s.pop()) // push result back onto stack
		default:
			num, _ := strconv.Atoi(t)
			s.push(num)
		}
	}
	return s.pop()
}

///

func main() {

	var input []string

	input = []string{"2", "1", "+", "3", "*"}
	fmt.Printf("Input=%v\n", input)
	// Explanation: ((2 + 1) * 3) = 9
	fmt.Printf("Want: 9 Got:%v\n", evalRPN(input))
	fmt.Println()

	///

	input = []string{"4", "13", "5", "/", "+"}
	fmt.Printf("Input=%v\n", input)
	// Explanation: (4 + (13 / 5)) = 6
	fmt.Printf("Want: 6 Got:%v\n", evalRPN(input))
	fmt.Println()

	///

	input = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Printf("Input=%v\n", input)
	// ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
	fmt.Printf("Want: 22 Got:%v\n", evalRPN(input))
	fmt.Println()
}
