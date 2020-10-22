package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPermutePalindrome(t *testing.T) {
	assert.Equal(t, true, canPermutePalindrome("tactcoa"))
	assert.Equal(t, false, canPermutePalindrome("leetcode"))
}
