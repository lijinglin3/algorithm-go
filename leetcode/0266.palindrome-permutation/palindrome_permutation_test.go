package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPermutePalindrome(t *testing.T) {
	assert.Equal(t, false, canPermutePalindrome("code"))
	assert.Equal(t, true, canPermutePalindrome("aab"))
	assert.Equal(t, true, canPermutePalindrome("carerac"))
}
