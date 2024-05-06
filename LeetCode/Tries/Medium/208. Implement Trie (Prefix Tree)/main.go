package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
208. Implement Trie (Prefix Tree)
Medium

A trie (pronounced as "try") or prefix tree is a tree data structure used to
efficiently store and retrieve keys in a dataset of strings. There are various
applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

    Trie() Initializes the trie object.

    void insert(String word) Inserts the string word into the trie.

    boolean search(String word) Returns true if the string word is in the trie
	(i.e., was inserted before), and false otherwise.

    boolean startsWith(String prefix) Returns true if there is a previously
	inserted string word that has the prefix prefix, and false otherwise.


Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]

Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True


Constraints:

    1 <= word.length, prefix.length <= 2000
    word and prefix consist only of lowercase English letters.
    At most 3 * 104 calls in total will be made to insert, search, and
	startsWith.
*/

/*
Intuitive Understanding:
------------------------

*/

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{
		children: make(map[rune]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	index := this
	for _, char := range word {
		if _, ok := index.children[char]; !ok {
			newNode := Constructor()
			index.children[char] = &newNode
		}
		index = index.children[char]
	}
	// end-of-word indicator.
	index.children['*'] = nil
}

func (this *Trie) Search(word string) bool {
	index := this
	for _, char := range word {
		if _, ok := index.children[char]; !ok {
			return false
		}
		index = index.children[char]
	}
	// check for end-of-word indicator.
	if _, ok := index.children['*']; ok {
		return true
	} else {
		return false // not a whole word.
	}
}

func (this *Trie) StartsWith(prefix string) bool {
	index := this
	for _, char := range prefix {
		if _, ok := index.children[char]; !ok {
			return false
		}
		index = index.children[char]
	}
	return true // prefix exists.
}

//-------------------------------------

//*-*

// Recursive call to get a list of all words in a Trie.
// Pass by Value.

func (t *Trie) getWords(word string, words []string) []string {
	// loop through all the letters of each node.
	for char, child := range t.children {
		if char == '*' { // new word found
			words = append(words, word)
		} else {
			// add letter to word prefix and recurse.
			prefix := word + string(char)
			words = child.getWords(prefix, words)
		}
	}
	return words
}

//-------------------------------------
//-------------------------------------

func main() {

	var words []string
	var myTrie Trie
	var prefix string

	//words := []string{"ace", "act", "bad", "cat"}
	words = []string{"flower", "flow", "flight"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.Insert(word)
	}

	wordCollection := []string{}
	fmt.Printf("All Words: %#v\n", myTrie.getWords("", wordCollection))
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"ace", "act", "bad", "cat"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.Insert(word)
	}

	wordCollection = []string{}
	fmt.Printf("All Words: %#v\n", myTrie.getWords("", wordCollection))
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{""}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.Insert(word)
	}

	wordCollection = []string{}
	fmt.Printf("All Words: %#v\n", myTrie.getWords("", wordCollection))
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.Search(prefix))
	fmt.Println()

	//*-*

	var zeroWords []string
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range zeroWords {
		myTrie.Insert(word)
	}

	wordCollection = []string{}
	fmt.Printf("All Words: %#v\n", myTrie.getWords("", wordCollection))
	prefix = "cat"
	fmt.Printf("Does %#v exisits within Trie? %#v\n",
		prefix, myTrie.Search(prefix))
	fmt.Println()

	//*-*

}

//*-*
