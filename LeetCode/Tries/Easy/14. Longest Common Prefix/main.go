package main

import (
	"fmt"
	"strings"
)

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
14. Longest Common Prefix
Easy

Write a function to find the longest common prefix string amongst an array of
strings.

If there is no common prefix, return an empty string "".


Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

    1 <= strs.length <= 200
    0 <= strs[i].length <= 200
    strs[i] consists of only lowercase English letters.
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/* // Solution 5b: Using stdlib function strings.HasPrefix() together with a Binary
//             Search Algorithm to search for common prefixes.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // corner case
	}

	// clousure to determine if a prefix is common.
	var isCommonPrefix func(idx int) bool
	isCommonPrefix = func(idx int) bool {
		prefix := strs[0][:idx]
		for i := 1; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], prefix) {
				// prefix doesn't exist at this length.
				return false
			}
		}
		return true
	}

	// find length of shortest string.
	ln := len(strs[0])
	for _, str := range strs {
		if len(str) < ln {
			ln = len(str)
		}
	}

	// perform binary search on string indices.
	low, high := 0, ln-1
	mid := high / 2

	for low <= high {
		if isCommonPrefix(mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = high/2 + low
	}
	return strs[0][:mid]
} */

//-------------------------------------

// Solution 5a: Using stdlib function strings.HasPrefix() together with a Binary
//             Search Algorithm to search for common prefixes.

/* func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // corner case
	}

	// find length of shortest string.
	ln := len(strs[0])
	for _, str := range strs {
		if len(str) < ln {
			ln = len(str)
		}
	}

	// perform binary search on string indices.
	low, high := 0, ln-1
	mid := high / 2

	for low <= high {
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = high/2 + low
	}
	return strs[0][:mid]
}

// Helper function
func isCommonPrefix(strs []string, idx int) bool {
	prefix := strs[0][:idx]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			// prefix doesnt exist at this length.
			return false
		}
	}
	return true
} */

//-------------------------------------

/* // Solution 4b: Building and searching a Trie (Simplified)

// trie node unit.
type node map[rune]*node

//*-*

//func newNode() *node {
//	return &node{}
//}

//*-*

// Insert new word starting at Root Node.
func (n *node) insertWord(word string) {
	index := n
	for _, char := range word {
		if _, ok := (*index)[char]; !ok {
			// insert a new node
			(*index)[char] = &node{}
		}
		// move forward
		index = (*index)[char]
	}
	// insert end of word indicator
	(*index)['*'] = nil
}

//*-*

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // corner case
	}

	// Build a trie
	root := &node{} // allocate mem
	for _, word := range strs {
		root.insertWord(word)
	}

	// find where words no longer share a common child,
	// OR if prefix is a whole word.
	var result string
	index := root
	for len(*index) == 1 {
		for key := range *index {
			if key == '*' { // prefix is a whole word
				return result
			}
			result = result + string(key)
			index = (*index)[key]
		}
	}
	return result
} */

//-------------------------------------

/* // Solution 4a: Building and searching a Trie (Simplified)

// trie node unit.
type node map[rune]node

func newNode() node {
	return make(map[rune]node)
}

//*-*

// Insert new word into trie.
func insertWord(root node, word string) {
	index := root
	for _, char := range word {
		if _, ok := index[char]; !ok {
			// create a new node.
			index[char] = newNode()
		}
		// move to the next char
		index = index[char]
	}
	// insert end of word indicator = *
	index['*'] = nil
}

//*-*

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // corner case
	}

	root := newNode()

	// Build a trie
	for _, word := range strs {
		insertWord(root, word)
	}

	// find where words no longer share a common child,
	// OR where if prefix is a whole word.
	var result string
	index := root
	for len(index) == 1 {
		for key := range index {
			if key == '*' { // prefix is a whole word
				return result
			}
			result = result + string(key)
			index = index[key]
		}
	}
	return result
} */

//-------------------------------------

/* // Solution 3: Building and searching a Trie using a struct.

// trie node unit.
type node struct {
	children map[rune]*node
}

//*-*

// Insert new word into trie.
func insertWord(root *node, word string) {
	index := root
	for _, char := range word {
		if _, ok := index.children[char]; !ok {
			// create a new node.
			index.children[char] = &node{
				make(map[rune]*node),
			}
		}
		// move to the next char
		index = index.children[char]
	}
	// insert end of word indicator = *
	index.children['*'] = nil
}

//*-*

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // corner case
	}

	root := &node{ // allocate mem.
		make(map[rune]*node),
	}

	// Build a trie
	for _, word := range strs {
		insertWord(root, word)
	}

	// find where words no longer share a common child,
	// OR where if prefix is a whole word.
	var result string
	index := root
	for len(index.children) == 1 {
		for key := range index.children {
			if key == '*' { // prefix is a whole word
				return result
			}
			result = result + string(key)
			index = index.children[key]
		}
	}
	return result
} */

//-------------------------------------

/* // Solution 2: Vertical Scanning, more efficient than Soln 1.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// loop through all indices of strs[0].
	for i := 0; i < len(strs[0]); i++ {
		// loop through each element of strs.
		for j := 1; j < len(strs); j++ {
			// if element is too short OR chars don't match.
			if i > len(strs[j])-1 || strs[j][i] != strs[0][i] {
				// return the previous successful sweep.
				return string(strs[0][:i])
			}
		}
	}
	return string(strs[0]) // all words match each other.
} */

/*
Complexity Analysis

Time complexity : O(S) , where S is the sum of all characters in all strings.

In the worst case there will be n equal strings with length m and the algorithm
performs S=m*n character comparisons.
Even though the worst case is still the same as Approach 1, in the best case
there are at most (n*minLen) comparisons where (minLen) is the length of the
shortest string in the array.

Space complexity : O(1). We only used constant extra space.
*/

//-------------------------------------

// Solution 1: Horizontal Scanning, easiest to remember.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // 1st string
	// loop through all the strings
	for i := 1; i < len(strs); i++ {
		// while the index of the prefix != 0
		for strings.Index(strs[i], prefix) != 0 {
			// remove last letter
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return "" // no prefix exists
			}
		}
	}
	return prefix
}

/*
Complexity Analysis

Time complexity : O(S), where S is the sum of all characters in all strings.

In the worst case all n strings are the same. The algorithm compares the string
S1 with the other strings [S2...Sn]. There are S character comparisons, where S
is the sum of all characters in the input array.

Space complexity : O(1). We only used constant extra space.
*/

//-------------------------------------

func main() {

	var strs []string

	/* 	strs = []string{"flower", "flow", "flight"}
	   	fmt.Println(strs)
	   	fmt.Println("Longest Common Prefix: ", longestCommonPrefix(strs)) */

	//strs = []string{"a", "ab"}
	//strs = []string{"a"}
	//strs = []string{""}
	strs = []string{"flower", "flow", "flight", "fl"}
	fmt.Println(strs)
	fmt.Println("Longest Common Prefix: ", longestCommonPrefix(strs))

}

//*-*
