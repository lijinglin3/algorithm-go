package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDays(t *testing.T) {
	assert.Equal(t, 3, minDays([]int{1, 10, 3, 10, 2}, 3, 1))
	assert.Equal(t, -1, minDays([]int{1, 10, 3, 10, 2}, 3, 2))
	assert.Equal(t, 12, minDays([]int{7, 7, 7, 7, 12, 7, 7}, 2, 3))
	assert.Equal(t, 9, minDays([]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2))
	assert.Equal(t, 1000000000, minDays([]int{1000000000, 1000000000}, 1, 1))
}
