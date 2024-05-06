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

// isPalindrome function checks if a string is a palindrome.
func isPalindrome(s string) bool {

	// sanitize input
	var characters []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			characters = append(characters, unicode.ToLower(ch))
		}
	}

	// Recursive closure.
	var check4Palindrome func(string) bool
	check4Palindrome = func(str string) bool {
		ln := len(str)
		if ln <= 1 {
			return true
		} else {
			return (str[0] == str[ln-1]) && check4Palindrome(str[1:ln-1])
		}
	}
	return check4Palindrome(string(characters)) // type conversion.
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
