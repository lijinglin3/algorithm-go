package leetcode

// Trie ...
type Trie struct {
	children map[rune]*Trie
	flag     bool
}

// Constructor initialize your data structure here.
func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

// Insert inserts a word into the trie.
func (t *Trie) Insert(word string) {
	tmp := t
	for _, v := range word {
		if tmp.children[v] == nil {
			tmp.children[v] = &Trie{children: make(map[rune]*Trie)}
		}
		tmp = tmp.children[v]
	}
	tmp.flag = true
}

// Search returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	tmp := t
	for _, v := range word {
		if tmp.children[v] == nil {
			return false
		}
		tmp = tmp.children[v]
	}
	return tmp.flag
}

// StartsWith returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	tmp := t
	for _, v := range prefix {
		if tmp.children[v] == nil {
			return false
		}
		tmp = tmp.children[v]
	}
	return true
}
