package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanArrange(t *testing.T) {
	assert.Equal(t, true, canArrange([]int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5))
	assert.Equal(t, false, canArrange([]int{1, 2, 3, 4, 5, 6}, 10))
	assert.Equal(t, true, canArrange([]int{1, 2, 3, 4, 5, 6}, 7))
	assert.Equal(t, true, canArrange([]int{-1, 1, -2, 2, -3, 3, -4, 4}, 3))
	assert.Equal(t, true, canArrange([]int{-4, -7, 5, 2, 9, 1, 10, 4, -8, -3}, 3))
}
