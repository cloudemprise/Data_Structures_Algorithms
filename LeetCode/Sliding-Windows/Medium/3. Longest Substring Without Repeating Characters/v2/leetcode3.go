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

	// Recall IF character has already been included.
	isIncluded := make(map[byte]bool)

	// Implement a sliding-window to check for duplicates.
	begin, end := 0, 0
	longest := 0
	for end < len(s) {

		if isIncluded[s[end]] { // Duplicate found.
			// i.e. s[end] == s[begin] !!!

			// NB! Thus the purpose of this condition is to shrink lhs of the
			// window-span, but to also remove the duplicate character from the
			/// hash-map (in this iteration) so as to allow for the inclusion of
			// said duplicate character in the next loop iteration. It's a two
			// step process.

			isIncluded[s[begin]] = false // Mark duplicate character as false
			// but could also do: isIncluded[s[end]] = false.
			begin++ // and shrink window.
		} else {
			isIncluded[s[end]] = true         // Mark as included.
			longest = max(longest, end-begin) // Update measurement.
			end++                             // Expand window for next check.
		}
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
