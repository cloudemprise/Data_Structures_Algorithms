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

	s := NewStack()
	for i := 0; i < ln; i++ { // process string in forward order

		switch str[i] {
		case '(', '[', '{':
			s.Push(str[i])
		case ')':
			if s.IsEmpty() { // don't Pop from empty stack
				return false
			} else if s.Pop() != '(' {
				return false
			}
		case ']':
			if s.IsEmpty() { // don't Pop from empty stack
				return false
			} else if s.Pop() != '[' {
				return false
			}
		case '}':
			if s.IsEmpty() { // don't Pop from empty stack
				return false
			} else if s.Pop() != '{' {
				return false
			}
		default:
			continue // ignore other characters or...
			/* return false // something is wrong */
		}
	}

	return s.IsEmpty() // check for dangling

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
