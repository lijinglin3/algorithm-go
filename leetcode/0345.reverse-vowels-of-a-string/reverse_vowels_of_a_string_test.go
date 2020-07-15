package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, "holle", reverseVowels("hello"))
	assert.Equal(t, "leotcede", reverseVowels("leetcode"))
	assert.Equal(t, "LEOTCEDE", reverseVowels("LEETCODE"))
}
