package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, "", longestPalindrome(""))
	assert.Equal(t, "bb", longestPalindrome("abbc"))
	assert.Equal(t, "bcdcb", longestPalindrome("abbcdcb"))
}
