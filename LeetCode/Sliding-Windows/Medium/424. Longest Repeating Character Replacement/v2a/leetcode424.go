package main

import "fmt"

// Sliding Window Technique using a hashmap to remember the frequency count of
// each character as well as the highest character count.

// Time Complexity : O(n)
// Space Complexity: O(1)

// This solution hinges on the following constraint:
// Longest Window-Span = (Maximum Frequency Character Count) + (No. of Flips [k])
// and fails for: maxLength - highestCount - k > 0, thus reduce window-span for such
// occasion.

// characterReplacement finds the length of the longest substring containing the
// same letter, after performing k number of character replacements.
func characterReplacement(s string, k int) int {

	// Keep track of the highest character count...
	highestCount := 0
	// by recalling the count of each character.
	charCount := make(map[byte]int) // O(1)

	// Also, keep track of the maximum sliding-window length.
	maxLength := 0

	// Implement a sliding-window to check for maxLength...
	left := 0
	for right := 0; right < len(s); right++ {

		// Increment the current character count...
		charCount[s[right]]++
		// and check if it is the highest count.
		highestCount = max(highestCount, charCount[s[right]])

		// Problem constraint dictates that:
		// maxLength = highestCount + k
		// Check if the Window Constraint has been violated.
		if (right-left+1)-(highestCount+k) > 0 { // If so...
			charCount[s[left]]-- // remove character from consideration...
			left++               // and shrink window-span.
		}

		// Check for maximum sliding-window length.
		maxLength = max(maxLength, right-left+1)
	} // O(n)
	return maxLength
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
