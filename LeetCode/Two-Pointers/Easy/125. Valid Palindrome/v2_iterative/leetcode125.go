package main

import (
	"fmt"
	"unicode"
)

// Time Complexity  = O(n) + O(n) = O(n)
// Space Complexity = O(1) + O(1) = O(1)

// isPalindrome function checks if a string is a palindrome.
func isPalindrome(s string) bool {

	// sanitize input
	var characters []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			characters = append(characters, unicode.ToLower(ch))
		}
	} // Time: O(n) Space: O(n)

	// Search for equality.
	j := len(characters) - 1
	for i := range characters {
		if characters[i] != characters[j-i] {
			return false
		}
	} // Time: O(n) Space: O(1)

	return true // Success.
}

///

func main() {

	var input string

	input = "A man, a plan, a canal: Panama"
	fmt.Printf("Input = %v\n", input)
	fmt.Println("Want: true \t Got: ", isPalindrome(input))
	fmt.Println()

	///

	input = "race a car"
	fmt.Printf("Input = %v\n", input)
	fmt.Println("Want: false \t Got: ", isPalindrome(input))
	fmt.Println()

	///

	input = "Dogma I am God"
	fmt.Printf("Input = %v\n", input)
	fmt.Println("Want: true \t Got: ", isPalindrome(input))
	fmt.Println()

	///

	input = ""
	fmt.Printf("Input = %v\n", input)
	fmt.Println("Want: true \t Got: ", isPalindrome(input))
	fmt.Println()

	///

}
