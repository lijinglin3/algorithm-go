package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressString(t *testing.T) {
	assert.Equal(t, "", compressString(""))
	assert.Equal(t, "abbccd", compressString("abbccd"))
	assert.Equal(t, "a2b1c5a3", compressString("aabcccccaaa"))
}
