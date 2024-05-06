package main

import "fmt"

// This solution hinges on the following constraint:
// Longest Window-Span = (Maximum Frequency Character Count) + (No. of Flips [k])
// and fails for: longest - maxFreq - k > 0, thus reduce window-span for such
// occasion.

// characterReplacement ...

// Solution 1: Search for longest sequence by first iterating right and then
// iterating left. Limiting factors are number of flips, broken sequence or
// edge of byte array.
// Time Complexity: O(n^2)
// Space Complexity: O(1)
func characterReplacement(s string, k int) int {
	ln := len(s)
	var longest int

	// Two pointer sliding window with begin & end.
	for i := 0; i < ln; i++ {
		count := 1
		flips := k

		// Search right.
		for end := i + 1; end < ln; end++ {
			if s[end] == s[i] {
				count++
			} else if flips > 0 {
				count++
				flips--
			} else {
				break // sequence broken.
			}
		}
		// Search left.
		for begin := i - 1; begin >= 0; begin-- {
			if s[begin] == s[i] {
				count++
			} else if flips > 0 {
				count++
				flips--
			} else {
				break // sequence broken.
			}
		}

		longest = max(longest, count)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

///

func main() {

	var input string
	var k int

	input = "ABAB"
	k = 2
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("k=%v\n", k)
	fmt.Printf("Want: 4 \t Got: %v\n", characterReplacement(input, k))
	fmt.Println()

	///

	input = "AABABBA"
	k = 1
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("k=%v\n", k)
	fmt.Printf("Want: 4 \t Got: %v\n", characterReplacement(input, k))
	fmt.Println()

	///

	input = "ABAA"
	k = 0
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("k=%v\n", k)
	fmt.Printf("Want: 2 \t Got: %v\n", characterReplacement(input, k))
	fmt.Println()

	///

	input = "ABBB"
	k = 2
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("k=%v\n", k)
	fmt.Printf("Want: 4 \t Got: %v\n", characterReplacement(input, k))
	fmt.Println()

	///

	input = "BAAAB"
	k = 2
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("k=%v\n", k)
	fmt.Printf("Want: 5 \t Got: %v\n", characterReplacement(input, k))
	fmt.Println()
}
