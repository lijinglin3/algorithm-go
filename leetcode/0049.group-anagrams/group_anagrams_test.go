package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	assert.ElementsMatch(t, [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
