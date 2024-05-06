package main

type Node struct {
	Value    string
	Children map[string]*Node
}

type Trie struct {
	Root *Node
}

func NewPrefixTrie() *Trie {
	return &Trie{
		Root: &Node{
			Value:    "",
			Children: make(map[string]*Node),
		},
	}
}

func (t *Trie) Insert(word string) {

	temp := t.Root

	for _, el := range word {

		charStr := string(el)

		if _, ok := temp.Children[charStr]; !ok {
			temp.Children[charStr] = &Node{
				Value:    charStr,
				Children: make(map[string]*Node),
			}
		}

		temp = temp.Children[charStr]
	}
}

func getMinStr(strs []string) string {
	minStr := strs[0]

	for _, el := range strs {
		if len(string(el)) < len(minStr) {
			minStr = string(el)
		}
	}

	return minStr
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var result string

	prefixTrie := NewPrefixTrie()

	for _, el := range strs {
		prefixTrie.Insert(el)
	}

	temp := prefixTrie.Root
	minStr := getMinStr(strs)

	for _, el := range minStr {
		strEl := string(el)

		if len(temp.Children) == 1 {
			result += strEl
			temp = temp.Children[strEl]
		} else {
			return result
		}
	}

	return result
}
