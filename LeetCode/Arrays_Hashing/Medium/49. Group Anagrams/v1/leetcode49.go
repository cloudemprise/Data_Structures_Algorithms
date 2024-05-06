package main

import (
	"fmt"
)

// Group up the anagrams.
func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 { // corner case
		return [][]string{strs}
	}

	// final output.
	var result [][]string

	// loop through all elements.
	for len(strs) > 0 {

		// 1st temp copy
		temp1 := strs[0]
		// group anagram results.
		grouping := []string{temp1}
		// remove seen element.
		strs = strs[1:]

		for j := 0; j < len(strs); j++ {
			// 2nd temp copy
			temp2 := strs[j]

			if len(temp1) == len(temp2) {
				if isAnagram(temp1, temp2) {
					// anagram found, group and remove.
					grouping = append(grouping, temp2)
					strs = append(strs[:j], strs[j+1:]...)
					j -= 1 // correlate index j
				}
			}

		}
		result = append(result, grouping)
	}
	return result
}

// Check if two words are an anagram.
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cache := make(map[rune]int)
	for _, v := range s {
		cache[v]++
	}
	for _, v := range t {
		cache[v]--
		if cache[v] < 0 {
			return false
		}
	}
	return true
}

///

func main() {

	strs := []string{"tan", "ate", "nat", "bat", "eat", "tea"}
	fmt.Printf("Input string = %v\n", strs)
	fmt.Printf("Output %#v\n", groupAnagrams(strs))
	fmt.Println()

	strs = []string{"", "", ""}
	fmt.Printf("Input string = %v\n", strs)
	fmt.Printf("Output %#v\n", groupAnagrams(strs))
	fmt.Println()

	x := 'a'
	y := 'a'

	fmt.Printf("XOR = %v", x^y)

}
