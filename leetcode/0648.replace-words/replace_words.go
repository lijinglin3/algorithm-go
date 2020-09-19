package leetcode

import (
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	type Trie struct {
		children map[rune]*Trie
		val      string
	}

	trie := &Trie{children: make(map[rune]*Trie)}
	for _, d := range dictionary {
		tmp := trie
		for _, v := range d {
			if tmp.children[v] == nil {
				tmp.children[v] = &Trie{children: make(map[rune]*Trie)}
			}
			tmp = tmp.children[v]
		}
		tmp.val = d
	}

	fields := strings.Fields(sentence)
	for i, f := range fields {
		tmp := trie
		for _, v := range f {
			if tmp.children[v] == nil {
				break
			}
			tmp = tmp.children[v]
			if tmp.val != "" {
				fields[i] = tmp.val
				break
			}
		}
	}
	return strings.Join(fields, " ")
}
