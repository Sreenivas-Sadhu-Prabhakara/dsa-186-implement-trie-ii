package main

import "fmt"

// Trie counts words rather than merely marking them, so duplicates,
// prefix counts and erasure all work.
type Trie struct {
	children  [26]*Trie
	wordCount int // words ending exactly here
	passCount int // words passing through this node
}

// NewTrie returns an empty counting trie.
func NewTrie() *Trie { return &Trie{} }

// Insert adds one occurrence of word.
func (t *Trie) Insert(word string) {
	n := t
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if n.children[c] == nil {
			n.children[c] = &Trie{}
		}
		n = n.children[c]
		n.passCount++
	}
	n.wordCount++
}

func (t *Trie) walk(s string) *Trie {
	n := t
	for i := 0; i < len(s); i++ {
		n = n.children[s[i]-'a']
		if n == nil {
			return nil
		}
	}
	return n
}

// CountWordsEqualTo returns how many times word was inserted (minus erases).
func (t *Trie) CountWordsEqualTo(word string) int {
	if n := t.walk(word); n != nil {
		return n.wordCount
	}
	return 0
}

// CountWordsStartingWith returns how many stored words have the prefix.
func (t *Trie) CountWordsStartingWith(prefix string) int {
	if n := t.walk(prefix); n != nil {
		return n.passCount
	}
	return 0
}

// Erase removes one occurrence of word; false if it isn't stored.
func (t *Trie) Erase(word string) bool {
	if t.CountWordsEqualTo(word) == 0 {
		return false
	}
	n := t
	for i := 0; i < len(word); i++ {
		n = n.children[word[i]-'a']
		n.passCount--
	}
	n.wordCount--
	return true
}

func main() {
	tr := NewTrie()
	tr.Insert("apple")
	tr.Insert("apple")
	tr.Insert("apps")
	fmt.Println(tr.CountWordsEqualTo("apple"), tr.CountWordsStartingWith("app"))
	tr.Erase("apple")
	fmt.Println(tr.CountWordsEqualTo("apple"))
}
