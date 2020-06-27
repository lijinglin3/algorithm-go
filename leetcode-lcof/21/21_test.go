package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	assert.Equal(t, []int{}, exchange([]int{}))
	assert.Equal(t, []int{1}, exchange([]int{1}))
	assert.Equal(t, []int{1, 3, 5}, exchange([]int{1, 3, 5}))
	assert.Equal(t, []int{1, 3, 2, 4}, exchange([]int{1, 2, 3, 4}))
	assert.Equal(t, []int{1, 5, 3, 4, 2}, exchange([]int{1, 2, 3, 4, 5}))
}
