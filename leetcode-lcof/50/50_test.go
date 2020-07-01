package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	assert.Equal(t, byte(' '), firstUniqChar("aabbccdd"))
	assert.Equal(t, byte('e'), firstUniqChar("aabbccdde"))
}
