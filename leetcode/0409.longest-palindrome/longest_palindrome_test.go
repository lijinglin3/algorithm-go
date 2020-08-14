package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
	assert.Equal(t, 8, longestPalindrome("aaccccdd"))
}
