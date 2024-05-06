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
	stack := make([]int, 0)

	for _, t := range tokens {

		if t == "+" || t == "-" || t == "*" || t == "/" {

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result int
			if t == "+" {
				result = a + b
			} else if t == "-" {
				result = a - b
			} else if t == "*" {
				result = a * b
			} else if t == "/" {
				result = a / b
			}
			stack = append(stack, result)
		} else {
			num, _ := strconv.Atoi(t)
			stack = append(stack, num)
		}
	}
	return stack[0]
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
