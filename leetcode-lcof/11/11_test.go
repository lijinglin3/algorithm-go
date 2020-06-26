package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArray(t *testing.T) {
	assert.Equal(t, 3, minArray([]int{3}))
	assert.Equal(t, 3, minArray([]int{3, 4, 5}))
	assert.Equal(t, 1, minArray([]int{3, 1, 3}))
	assert.Equal(t, 1, minArray([]int{10, 1, 10, 10, 10, 10, 10}))
	assert.Equal(t, 3, minArray([]int{3, 3, 3, 3, 3, 3, 3, 4, 5}))
	assert.Equal(t, 0, minArray([]int{3, 4, 5, 0, 1, 2}))
	assert.Equal(t, 0, minArray([]int{2, 2, 2, 0, 1}))
	assert.Equal(t, 0, minArray([]int{2, 2, 2, 2, 2, 0, 1}))
}
