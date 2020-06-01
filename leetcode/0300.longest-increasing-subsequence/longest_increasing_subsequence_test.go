package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 6, lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	assert.Equal(t, 1, lengthOfLIS([]int{0}))
	assert.Equal(t, 0, lengthOfLIS([]int{}))
}
