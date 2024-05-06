package main

import "fmt"

// > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > > >

/*
211. Design Add and Search Words Data Structure
Medium

Design a data structure that supports adding new words and finding if a string
matches any previously added string.

Implement the WordDictionary class:

    WordDictionary() Initializes the object.
    void addWord(word) Adds word to the data structure, it can be matched later.
    bool search(word) Returns true if there is any string in the data structure
	that matches word or false otherwise. word may contain dots '.' where dots
	can be matched with any letter.

Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True


Constraints:

    1 <= word.length <= 25
    word in addWord consists of lowercase English letters.
    word in search consist of '.' or lowercase English letters.
    There will be at most 2 dots in word for search queries.
    At most 104 calls will be made to addWord and search.
*/

/*
Intuitive Understanding:
------------------------

https://www.youtube.com/watch?v=BTf05gs_8iU

*/

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

// < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < < <

type node map[rune]*node

type WordDictionary struct {
	root *node
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &node{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	index := this.root
	for _, char := range word {
		if _, ok := (*index)[char]; !ok {
			(*index)[char] = &node{}
		}
		index = (*index)[char]
	}
	(*index)['*'] = nil // end-of-word indicator.
}

func (this *WordDictionary) Search(word string) bool {
	index := this.root
	return index.searchNodes(word)
}

//*-*

// Solution 3:
func (n *node) searchNodes(word string) bool {
	if n == nil {
		// word is longer than node search path.
		return false
	}

	index := n
	for i, char := range word {
		// normal recursive search starts here.
		if _, ok := (*index)[char]; !ok {
			if char != '.' {
				return false
			}
			// check if child nodes have solution.
			for _, child := range *index {
				if child.searchNodes(word[i+1:]) {
					// solution found, exit out.
					return true
				}
			}
			return false // no solution found.
		}
		// char exists so move forward.
		index = (*index)[char]
	}

	// check end-of-word indicator, i,e, '*'
	if _, ok := (*index)['*']; !ok {
		return false
	} else {
		return true
	}
}

//*-*

/* // Solution 2:
func (n *node) searchNodes(word string, idx int) bool {
	if n == nil {
		// word is longer than node search path.
		return false
	}

	index := n
	for i := idx; i < len(word); i++ {
		switch char := rune(word[i]); char {
		case '.':
			// check if child nodes have solution.
			for _, child := range *index {
				if child.searchNodes(word, i+1) {
					// solution found, exit out.
					return true
				}
			}
			return false // no solution found.
		default:
			// normal recursive search here.
			if _, ok := (*index)[char]; !ok {
				return false
			}
			// char exists so move forward.
			index = (*index)[char]
		}
	}

	// check end-of-word indicator, i,e, '*'
	if _, ok := (*index)['*']; !ok {
		return false
	} else {
		return true
	}
} */

//*-*

/* // Solution 1:
func (n *node) searchNodes(word string, idx int) bool {
	if n == nil {
		// word is longer than node search path.
		return false
	}

	index := n
	for i := idx; i < len(word); i++ {
		char := rune(word[i])
		if char == '.' {
			// check all child nodes for solution.
			for _, child := range *index {
				if child.searchNodes(word, i+1) {
					// solution found, exit out.
					return true
				}
			}
			return false // no solution found.
		} else {
			// normal recursive search here.
			if _, ok := (*index)[char]; !ok {
				return false
			}
			// char exists so move forward.
			index = (*index)[char]
		}
	}

	// check end-of-word indicator, i,e, '*'
	if _, ok := (*index)['*']; !ok {
		return false
	} else {
		return true
	}
} */

//-------------------------------------

func main() {

	var words []string
	var myTrie WordDictionary
	var prefix string

	/*
		words = []string{"flower", "flow", "flight"}
		fmt.Println("New Trie Creation: ")
		myTrie = Constructor()

		for _, word := range words {
			myTrie.AddWord(word)
		}

		fmt.Printf("All Words: %#v\n", myTrie.getWords())
		prefix = "cat"
		fmt.Printf("Does %#v exisits within Trie? %#v\n",
			prefix, myTrie.Search(prefix))
		fmt.Println()

		//*-*

		words = []string{"ace", "act", "bad", "cat"}
		fmt.Println("New Trie Creation: ")
		myTrie = Constructor()

		for _, word := range words {
			myTrie.AddWord(word)
		}

		fmt.Printf("All Words: %#v\n", myTrie.getWords())
		prefix = "cat"
		fmt.Printf("Does %#v exisits within Trie? %#v\n",
			prefix, myTrie.Search(prefix))
		fmt.Println()

		//*-*

		words = []string{""}
		fmt.Println("New Trie Creation: ")
		myTrie = Constructor()

		for _, word := range words {
			myTrie.AddWord(word)
		}

		fmt.Printf("All Words: %#v\n", myTrie.getWords())
		prefix = "cat"
		fmt.Printf("Does %#v exisits within Trie? %#v\n",
			prefix, myTrie.Search(prefix))
		fmt.Println()

		//*-*

		var zeroWords []string
		fmt.Println("New Trie Creation: ")
		myTrie = Constructor()

		for _, word := range zeroWords {
			myTrie.AddWord(word)
		}

		fmt.Printf("All Words: %#v\n", myTrie.getWords())
		prefix = "cat"
		fmt.Printf("Does %#v exisits within Trie? %#v\n",
			prefix, myTrie.Search(prefix))
		fmt.Println() */

	//*-*

	words = []string{"bad", "dad", "mad"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "bad"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: true \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"bad", "dad", "mad"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "pad"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"bad", "dad", "mad"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = ".ad"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: true \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"bad", "dad", "mad"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "b.."
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: true \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"a", "a"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "aa"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"a", "a"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = ".a"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"a", "a"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "a."
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"at", "and", "an", "add"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "a"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

	//*-*

	words = []string{"at", "and", "an", "add"}
	fmt.Println("New Trie Creation: ")
	myTrie = Constructor()

	for _, word := range words {
		myTrie.AddWord(word)
	}

	fmt.Printf("All Words: %#v\n", myTrie.getWords())
	prefix = "a"
	fmt.Printf("Does %#v exisits within Trie?\n", prefix)
	fmt.Printf("Want: false \t Got: %v\n", myTrie.Search(prefix))
	fmt.Println()

}

//-------------------------------------

// Recursive call to get a list of all words in a Trie.
// Pass by Value.

func (t *WordDictionary) getWords() []string {
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
}

//*-*
