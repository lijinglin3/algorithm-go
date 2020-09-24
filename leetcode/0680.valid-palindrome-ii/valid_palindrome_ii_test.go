package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidPalindrome(t *testing.T) {
	assert.Equal(t, true, validPalindrome("aba"))
	assert.Equal(t, true, validPalindrome("abca"))
	assert.Equal(t, false, validPalindrome("abcda"))
}
