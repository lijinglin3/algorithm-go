package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	assert.Equal(t, true, repeatedSubstringPattern("ababab"))
	assert.Equal(t, false, repeatedSubstringPattern("ababa"))
	assert.Equal(t, false, repeatedSubstringPattern("abcaba"))
}
