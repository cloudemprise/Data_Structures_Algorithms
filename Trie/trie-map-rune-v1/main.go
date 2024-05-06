package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
Efficiency of a Trie.
---------------------

Search: O(K) where K = num of chars in search string.

Insertion: O(K) where K = num of chars in search string.

Space Complexity: O(K) where K = num of chars in search string.
*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/// DEFINITIONS:

// node unit.
type node map[rune]*node

//*-*

// Pointer to root of trie.
type trie struct {
	root *node
}

//*-*

// node constructor.
/* func newNode() *node {
	tmp := make(node)
	return &tmp
} */
func newNode() *node {
	return &node{}
}

//*-*

// trie constructor.
func newTrie() *trie {
	return &trie{
		root: newNode(),
	}
}

//-------------------------------------

// Does a prefix exist within a trie.
func (t *trie) searchPrefix(prefix string) bool {
	index := t.root
	for _, char := range prefix {
		if _, ok := (*index)[char]; !ok {
			return false
		}
		index = (*index)[char]
	}
	return true
}

//*-*

// Insert new word into Trie.
func (t *trie) insert(word string) {
	index := t.root
	for _, char := range word {
		if _, ok := (*index)[char]; !ok {
			new := newNode()
			(*index)[char] = new
		}
		index = (*index)[char]
	}
	// insert end of word indicator
	(*index)['*'] = nil
}

//-------------------------------------

// Insert new word starting at Root Node.
func (n *node) insertWord(word string) {
	/* if len(word) == 0 {
		return // corner case
	} */

	index := n
	for _, char := range word {
		if _, ok := (*index)[char]; !ok {
			// create a new node
			(*index)[char] = newNode()
		}
		// move forward
		index = (*index)[char]
	}
	// insert end of word indicator
	(*index)['*'] = nil
}

//-------------------------------------

/* // Recursive call to get a list of all words in a Trie.
// Pass by Value.

func (t *trie) getWords() []string {
	return t.root.traverseNodes("", []string{})
}

func (n *node) traverseNodes(word string, words []string) []string {
	// loop through all the letters of each node.
	for char, child := range *n {
		if char == '*' { // new word found
			words = append(words, word)
		} else {
			// add letter to word prefix and recurse.
			prefix := word + string(char)
			words = child.traverseNodes(prefix, words)
		}
	}
	return words
} */

//*-*

// Recursive call to get a list of all words in a Trie.
// Pass by Pointer.

func (t *trie) getWords() []string {
	var words []string
	t.root.traverseNodes("", &words)
	return words
}

func (n *node) traverseNodes(word string, words *[]string) {
	// loop through all the letters of each node.
	for char, child := range *n {
		if char == '*' { // new word found
			*words = append(*words, word)
		} else {
			// add letter to word prefix and recurse.
			prefix := word + string(char)
			child.traverseNodes(prefix, words)
		}
	}
}

//-------------------------------------

func main() {

	var words []string
	var myTrie *trie
	var prefix string

	//words := []string{"ace", "act", "bad", "cat"}
	words = []string{"flower", "flow", "flight"}
	fmt.Println("New Trie Creation: ")
	myTrie = newTrie()

	for _, word := range words {
		myTrie.insert(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.searchPrefix(prefix))
	fmt.Println()

	//*-*

	words = []string{"ace", "act", "bad", "cat"}
	fmt.Println("New Trie Creation: ")
	myTrie = newTrie()

	for _, word := range words {
		myTrie.insert(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.searchPrefix(prefix))
	fmt.Println()

	//*-*

	words = []string{""}
	fmt.Println("New Trie Creation: ")
	myTrie = newTrie()

	for _, word := range words {
		myTrie.insert(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.searchPrefix(prefix))
	fmt.Println()

	//*-*

	var zeroWords []string
	fmt.Println("New Trie Creation: ")
	myTrie = newTrie()

	for _, word := range zeroWords {
		myTrie.insert(word)
	}

	fmt.Printf("All Zero Words: %#v\n", myTrie.getWords())
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.searchPrefix(prefix))
	fmt.Println()

	//*-*

}

//*-*
