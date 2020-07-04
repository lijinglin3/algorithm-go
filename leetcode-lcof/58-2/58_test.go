package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLeftWords(t *testing.T) {
	assert.Equal(t, "cdefgab", reverseLeftWords("abcdefg", 2))
}
