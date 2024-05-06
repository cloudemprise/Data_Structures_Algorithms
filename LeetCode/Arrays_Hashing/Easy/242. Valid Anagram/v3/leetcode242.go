package main

import "fmt"

/*
Complexity Analysis:
--------------------

The time complexity of this solution is O(n) because of the linear time
complexity of iterating through the input strings to build the frequency map
and the constant space complexity of the map, which is of fixed size regardless
of the input.

The space complexity is also O(1) because the size of the map is fixed, being
based on the alphabet size.
*/

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false // String lengths must match.
	}

	counter := make(map[rune]int)
	for _, v := range s {
		counter[v]++
	}

	for _, v := range t {
		counter[v]--
		if counter[v] < 0 {
			return false
		}
	}

	return true
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
