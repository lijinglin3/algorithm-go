package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	assert.Equal(t, 0, longestSubarray([]int{0, 0, 0}))
	assert.Equal(t, 2, longestSubarray([]int{1, 1, 1}))
	assert.Equal(t, 3, longestSubarray([]int{1, 1, 0, 1}))
	assert.Equal(t, 3, longestSubarray([]int{1, 1, 0, 1, 0}))
}
