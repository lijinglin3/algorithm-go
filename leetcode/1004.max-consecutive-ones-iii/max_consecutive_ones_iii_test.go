package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	assert.Equal(t, 6, longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
}
