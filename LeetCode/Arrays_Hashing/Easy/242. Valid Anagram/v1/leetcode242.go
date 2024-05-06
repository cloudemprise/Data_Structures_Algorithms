package main

import "fmt"

func isAnagram(s string, t string) bool {
	// Check if the two words have same length.
	if len(s) != len(t) {
		fmt.Println("Words have different lengths.")
		return false
	}

	cache := make(map[rune]int) // Account for each character.
	for _, ch := range s {
		cache[ch]++
	}

	// Brute-force each character.
	for _, ch := range t {
		// Check if character exists in cache
		if _, ok := cache[ch]; !ok {
			fmt.Printf("Char %c not found\n", ch)
			return false
		}
		// Found so remove indication of existence.
		cache[ch]--
		// Check for duplicates.
		if cache[ch] < 0 {
			return false
		}
	}
	return true // Passed all checks.
}

///

func main() {
	var s, t string

	fmt.Println()
	s, t = "anagram", "nagaram"
	fmt.Printf("s = %s \t t = %s\n", s, t)
	fmt.Printf("Want: true  \t Got: %v\n", isAnagram(s, t))

	fmt.Println()
	s, t = "rat", "car"
	fmt.Printf("s = %s \t t = %s\n", s, t)
	fmt.Printf("Want: false \t Got: %v\n", isAnagram(s, t))

	fmt.Println()
	s, t = "ab", "a"
	fmt.Printf("s = %s \t\t t = %s\n", s, t)
	fmt.Printf("Want: false \t Got: %v\n", isAnagram(s, t))

}
