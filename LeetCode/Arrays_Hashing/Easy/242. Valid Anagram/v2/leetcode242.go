package main

import "fmt"

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		fmt.Println("Words have different lengths.")
		return false
	}

	var cache [26]int

	for _, ch := range s {
		cache[ch-'a']++
	}

	for _, ch := range t {
		cache[ch-'a']--
	}

	for _, index := range cache {
		if index != 0 {
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
