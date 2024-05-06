package main

import (
	"fmt"
)

/* // generateParenthesis generates all n combinations of well-formed parentheses.
func generateParenthesis(n int) (results []string) {

	var backtrack func([]byte, int, int)
	backtrack = func(current []byte, open int, close int) {

		if (open + close) == n*2 { // base case.
			results = append(results, string(current))
			return
		}

		// More open than closed brackets.
		if close < open {
			current[open+close] = ')'
			backtrack(current, open, close+1)
		}

		// Add an open bracket.
		if open < n {
			current[open+close] = '('
			backtrack(current, open+1, close)
		}
	}

	curr := make([]byte, n*2)
	backtrack(curr, 0, 0)
	return results
} */

// generateParenthesis generates all n combinations of well-formed parentheses.
func generateParenthesis(n int) (results []string) {

	current := make([]byte, n*2)
	var open, close int

	var backtrack func()
	backtrack = func() {

		if (open + close) == n*2 { // base case.
			results = append(results, string(current))
			return
		}

		// More open than closed brackets.
		if close < open {
			current[open+close] = ')'
			close++
			backtrack()
			close--
		}

		// Add an open bracket.
		if open < n {
			current[open+close] = '('
			open++
			backtrack()
			open--
		}
	}

	backtrack()
	return results
}

///

func main() {

	var n int
	var want []string
	var got []string

	fmt.Println("Test1:")
	n = 3
	fmt.Printf("n = %d \n", n)
	want = []string{
		"()()()",
		"()(())",
		"(())()",
		"(()())",
		"((()))",
	}
	got = generateParenthesis(n)
	fmt.Printf("Want: %q \nGot : %q\n", want, got)
	fmt.Println()
}
