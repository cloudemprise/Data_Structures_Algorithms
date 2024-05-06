package main

import "fmt"

// Time Complexity  = O(n)
// Space Complexity = O(n)

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else if c == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if c == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else if c == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0 // check for dangling
}

///

func main() {

	var input string

	input = "()"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "()[]{}"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "(]"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "([)]"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "{[]}"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "["
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "]"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = ")(){}"
	fmt.Printf("Input: s = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

}
