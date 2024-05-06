package main

import "fmt"

// Time Complexity  = O(n)
// Space Complexity = O(n)

func isValid(str string) bool {

	ln := len(str)
	/* 	if ln < 2 || ln%2 != 0 {
		return false // eliminate invalid
	} */
	if ln%2 != 0 {
		return false // eliminate invalid
	}

	pairs := map[byte]byte{ // shortcut
		'(': ')',
		'[': ']',
		'{': '}',
	}

	s := NewStack()
	for i := ln - 1; i >= 0; i-- { // process string in reverse order
		switch str[i] {
		case ')', ']', '}':
			s.Push(str[i])
		case '(', '[', '{':
			// ensure don't Pop from empty stack
			if s.IsEmpty() || s.Pop() != pairs[str[i]] {
				return false
			}
		}
	}
	return s.IsEmpty() // check for dangling
}

///

func main() {

	var input string

	input = "()"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "()[]{}"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "(]"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "([)]"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "{[]}"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: true \t Got: %v\n\n", isValid(input))

	input = "["
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = "]"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

	input = ")(){}"
	fmt.Printf("Input: str = %v\n", input)
	fmt.Printf("Want: false \t Got: %v\n\n", isValid(input))

}
