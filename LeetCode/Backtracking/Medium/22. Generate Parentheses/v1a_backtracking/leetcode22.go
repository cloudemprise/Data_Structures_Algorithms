package main

import (
	"fmt"
	"strings"
)

// generateParenthesis generates all n combinations of well-formed parentheses.
func generateParenthesis(n int) (results []string) {

	var backtrack func(string, int, int)
	backtrack = func(current string, open int, close int) {

		if len(current) == n*2 {
			// Old Copy idiom.
			//sCopy := make([]byte, len(current))
			//copy(sCopy, current)
			//results = append(results, string(sCopy))

			// New Copy idiom.
			sCopy := strings.Clone(current)
			results = append(results, sCopy)
			return
		}

		// More open than closed brackets.
		if close < open {
			backtrack(current+")", open, close+1)
		}

		// Add an open bracket.
		if open < n {
			backtrack(current+"(", open+1, close)
		}
	}

	backtrack("", 0, 0)
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
