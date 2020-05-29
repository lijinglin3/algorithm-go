package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIsomorphic(t *testing.T) {
	assert.Equal(t, false, isIsomorphic("123", "1234"))
	assert.Equal(t, false, isIsomorphic("foo", "bar"))
	assert.Equal(t, false, isIsomorphic("abca", "abcb"))
	assert.Equal(t, true, isIsomorphic("add", "egg"))
}
