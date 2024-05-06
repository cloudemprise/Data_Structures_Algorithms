package main

import "fmt"

// Solution 2: without string type conversion.
func canConstruct(ransomNote string, magazine string) bool {

	// inventory of runes
	cache := make(map[rune]int, 26)

	// index each character
	for _, ch := range magazine {
		cache[ch]++
	}

	// search for corresponding character
	for _, ch := range ransomNote {
		cache[ch]--
		if cache[ch] < 0 {
			return false
		}
	}
	return true
}

///

func main() {

	var ransomNote, magazine string

	ransomNote, magazine = "a", "b"
	fmt.Printf("Want: false \t Got: %v\n", canConstruct(ransomNote, magazine))

	ransomNote, magazine = "aa", "ab"
	fmt.Printf("Want: false \t Got: %v\n", canConstruct(ransomNote, magazine))

	ransomNote, magazine = "aa", "aab"
	fmt.Printf("Want: true \t Got: %v\n", canConstruct(ransomNote, magazine))
}
