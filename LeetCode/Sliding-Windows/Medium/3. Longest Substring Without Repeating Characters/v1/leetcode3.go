package main

import "fmt"

// Time Complexity : O(n)
// Space Complexity: O(n)

// lengthOfLongestSubstring returns an integer, which is the length of the
// longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0 // Corner case.
	}

	// Recall INDEX of each character.
	isIncluded := make(map[rune]int)

	// Implement a sliding-window to check for duplicates.
	begin := 0
	longest := 0
	for end, char := range s {

		if prior, ok := isIncluded[char]; ok { // Duplicate found.

			// Check if this duplicate is within current window-span.
			if prior >= begin {
				// If so, reposition beginning of window-span.
				begin = prior + 1
			}
		}
		// Update measurement.
		longest = max(longest, end-begin)
		// Overwrite duplicate character index.
		isIncluded[char] = end
	}
	return longest + 1
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

	input = ""
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 0 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "abcabcbb"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 3 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "abc $a"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 5 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "bbbbb"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 1 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "aab"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 2 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "dvdf"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 3 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "tmmzuxt"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 5 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "pwwkew"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 3 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()

	///

	input = "wobgrovw"
	fmt.Printf("Input=%v\n", input)
	fmt.Printf("Want: 6 \t Got: %v\n", lengthOfLongestSubstring(input))
	fmt.Println()
}
