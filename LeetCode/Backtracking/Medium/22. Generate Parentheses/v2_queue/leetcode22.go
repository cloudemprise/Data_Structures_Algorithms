package main

import (
	"fmt"
)

type state struct {
	str   string
	open  int
	close int
}

// generateParenthesis generates all n combinations of well-formed parentheses.
func generateParenthesis(n int) (results []string) {

	// Instantiate & initialise a queue.
	var queue []state
	queue = append(queue, state{"", 0, 0})

	// Infinite loop.
	for {

		// Process each queue element only once through every turn.
		ln := len(queue)
		for i := 0; i < ln; i++ {

			current := queue[i]

			if current.close < current.open {

				queue = append(queue, state{
					str:   current.str + ")",
					open:  current.open,
					close: current.close + 1,
				})
			}

			if current.open < n {

				queue = append(queue, state{
					str:   current.str + "(",
					open:  current.open + 1,
					close: current.close,
				})
			}
		}

		// Process queue until no more modifications
		// required to scratch elements.
		if len(queue) == ln {
			break // base case.
		} else {
			// Remove scratch elements from queue
			// for this turn.
			queue = queue[ln:]
		}
	}
	// Create output string
	for i := 0; i < len(queue); i++ {
		results = append(results, queue[i].str)
	}
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
