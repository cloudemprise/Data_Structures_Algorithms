package main

import "fmt"

// "Permutation in String" problem is solved via a sliding-window algorithm that
// directly compares array values.

// Time Complexity  = O(n)
// Space Complexity = O(1)

// checkInclusion checks whether s2 contains a permutation of s1.
func checkInclusion(s1 string, s2 string) bool {
	ln1, ln2 := len(s1), len(s2)

	if ln1 > ln2 {
		return false // Corner case.
	}

	// Count character occurrences in arrays.
	type alphabet [26]int
	var s1Sigil, s2Sigil alphabet

	// Build initial window-spans of length s1.
	for i := 0; i < ln1; i++ {
		s1Sigil[s1[i]-'a']++
		s2Sigil[s2[i]-'a']++
	}

	// Check for an initial solution.
	if s1Sigil == s2Sigil {
		return true
	}

	// Implement a sliding window of length s1, running
	// across s2 while comparing sigils.
	for i := ln1; i < ln2; i++ {

		// Slide the window.
		s2Sigil[s2[i-ln1]-'a']--
		s2Sigil[s2[i]-'a']++

		// Check if they are the same.
		if s1Sigil == s2Sigil {
			return true
		}
	}
	return false
}

///

func main() {

	var s1, s2 string

	s1 = "ab"
	s2 = "eidbaooo"
	fmt.Printf("s1=%v s2=%v\n", s1, s2)
	fmt.Printf("Want: true \t Got: %v\n", checkInclusion(s1, s2))
	fmt.Println()

	///

	s1 = "ab"
	s2 = "eidboaoo"
	fmt.Printf("s1=%v s2=%v\n", s1, s2)
	fmt.Printf("Want: false \t Got: %v\n", checkInclusion(s1, s2))
	fmt.Println()

	///

	s1 = "a"
	s2 = "ab"
	fmt.Printf("s1=%v s2=%v\n", s1, s2)
	fmt.Printf("Want: true \t Got: %v\n", checkInclusion(s1, s2))
	fmt.Println()

	///

	s1 = "abc"
	s2 = "bbbca"
	fmt.Printf("s1=%v s2=%v\n", s1, s2)
	fmt.Printf("Want: true \t Got: %v\n", checkInclusion(s1, s2))
	fmt.Println()

	///

	s1 = "hello"
	s2 = "ooolleoooleh"
	fmt.Printf("s1=%v s2=%v\n", s1, s2)
	fmt.Printf("Want: false \t Got: %v\n", checkInclusion(s1, s2))
	fmt.Println()
}
