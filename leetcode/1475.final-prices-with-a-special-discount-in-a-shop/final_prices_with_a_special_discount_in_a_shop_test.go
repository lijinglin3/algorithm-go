package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalPrices(t *testing.T) {
	assert.Equal(t, []int{}, finalPrices([]int{}))
	assert.Equal(t, []int{1}, finalPrices([]int{1}))
	assert.Equal(t, []int{4, 2, 4, 2, 3}, finalPrices([]int{8, 4, 6, 2, 3}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, finalPrices([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{9, 0, 1, 6}, finalPrices([]int{10, 1, 1, 6}))
}
