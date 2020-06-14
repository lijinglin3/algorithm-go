package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("rat", "car"))
}
