package main

import (
	"fmt"
	"unicode"
)

/*
Time Complexity
The time complexity of the solution is O(n), where n is the length of the input
string. This is because the solution involves iterating over each character in
the string to check for palindrome properties.

Space Complexity
The space complexity of the solution is also O(n). This is due to the creation
of a new string to store the alphanumeric characters from the input string, as
well as the recursive function call stack.
*/

/*
Boolean Arithmetic
| x y | x.y |  | !x !y | !x.!y |
| 0 0 |  0  |  |  1 1  |   1   | <- Logic Used
| 0 1 |  0  |  |  1 0  |   0   |
| 1 0 |  0  |  |  0 1  |   0   |
| 1 1 |  1  |  |  0 0  |   0   |
*/

// isPalindrome function checks if a string is a palindrome.
func isPalindrome(s string) bool {

	for len(s) > 0 && !isAlphanumeric(rune(s[0])) {
		s = s[1:]
	}

	for len(s) > 0 && !isAlphanumeric(rune(s[len(s)-1])) {
		s = s[:len(s)-1]
	}

	if len(s) <= 1 {
		return true
	} else {
		lo := unicode.ToLower(rune(s[0]))
		hi := unicode.ToLower(rune(s[len(s)-1]))
		return (lo == hi) && isPalindrome(s[1:len(s)-1])
	}
}

func isAlphanumeric(ch rune) bool {
	if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
		return true
	}
	return false
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
