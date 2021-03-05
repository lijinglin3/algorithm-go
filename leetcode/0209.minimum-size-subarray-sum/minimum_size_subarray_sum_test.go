package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSubArrayLen(t *testing.T) {
	assert.Equal(t, 0, minSubArrayLen(7, []int{}))
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3, 9}))
	assert.Equal(t, 0, minSubArrayLen(100, []int{2, 3, 1, 2, 4, 3, 1}))
}
