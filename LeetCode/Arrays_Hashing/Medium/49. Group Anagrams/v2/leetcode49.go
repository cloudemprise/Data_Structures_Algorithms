package main

import (
	"fmt"
)

/*
Time Complexity is O(n * k), where n is the number of strings and k is the
maximum length of a string. This is because the function iterates through each
word to generate its key, which takes O(k) time, and then groups the words based
on the key.

Space Complexity is also O(n * k) due to the use of the map to store the groups
of anagrams.
*/

// Use hashmap to store every letter's frequency and group those who have the same frequency.
func groupAnagrams(strs []string) [][]string {

	// int array as key to map.
	type key [26]int
	// group anagrams within map.
	groups := make(map[key][]string)

	for _, word := range strs {

		// generate a key for each word.
		var wordKey key
		for _, char := range word {
			wordKey[char-'a']++
		}

		// group words with same key.
		groups[wordKey] = append(groups[wordKey], word)
	}

	// generate the output.
	var results [][]string
	for _, anagrams := range groups {
		results = append(results, anagrams)
	}
	return results
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
