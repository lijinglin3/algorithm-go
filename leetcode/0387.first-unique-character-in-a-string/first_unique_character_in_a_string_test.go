package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	assert.Equal(t, -1, firstUniqChar("aaaa"))
	assert.Equal(t, 0, firstUniqChar("leetcode"))
	assert.Equal(t, 2, firstUniqChar("loveleetcode"))
}
