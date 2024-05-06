package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	// inventory of characters
	cache := make(map[byte]int)
	// type convertion: string to byte slice
	chars := []byte(magazine)
	// take inventory of each character
	for _, char := range chars {
		cache[char]++
	}

	// type convertion: string to byte slice
	chars = []byte(ransomNote)
	// check for corresponding characters
	for _, char := range chars {
		cache[char]--
		if cache[char] < 0 {
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
