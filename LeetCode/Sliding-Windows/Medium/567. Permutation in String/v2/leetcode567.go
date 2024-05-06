package main

import "fmt"

// Solution 2a: Using arrays to record matches characters.

// Time Complexity  = O(n)
// Space Complexity = O(n)

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
	// This is done by counting the matched elements of both
	// sigils. Which would equal 26 if a match is found.
	matches := 0
	for i := 0; i < 26; i++ {
		if s1Sigil[i] == s2Sigil[i] {
			matches++
		}
	}
	if matches == 26 {
		return true
	}

	// Implement a sliding window of length s1, running
	// across s2 while comparing sigil matches.
	for i := ln1; i < ln2; i++ {

		var idx byte

		// Process added character.
		idx = s2[i] - 'a'
		s2Sigil[idx]++
		// Check for a match.
		if s2Sigil[idx] == s1Sigil[idx] {
			matches++
		} else if s2Sigil[idx] == s1Sigil[idx]+1 {
			matches-- // Created a mismatch.
		}

		// Process removed character.
		idx = s2[i-ln1] - 'a'
		s2Sigil[idx]--
		// Check for a match.
		if s2Sigil[idx] == s1Sigil[idx] {
			matches++
		} else if s2Sigil[idx] == s1Sigil[idx]-1 {
			matches-- // Created a mismatch.
		}

		if matches == 26 {
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
