package main

import (
	"fmt"
	"unicode"
)

// Time Complexity  = O(n)
// Space Complexity = O(1)

// isPalindrome function checks if a string is a palindrome.
func isPalindrome(s string) bool {

	var i, j = 0, len(s) - 1

	for i < j {
		x := unicode.ToLower(rune(s[i]))
		y := unicode.ToLower(rune(s[j]))
		switch {
		// when x is neither a letter && neither a digit
		case !unicode.IsLetter(x) && !unicode.IsDigit(x):
			i++
		// when y is neither a letter && neither a digit
		case !unicode.IsLetter(y) && !unicode.IsDigit(y):
			j--
		case x == y:
			i++
			j--
		default:
			return false
		}
	}
	return true
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
