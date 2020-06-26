package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	assert.Equal(t, 3, findMin([]int{3}))
	assert.Equal(t, 3, findMin([]int{3, 4, 5}))
	assert.Equal(t, 1, findMin([]int{3, 1, 3}))
	assert.Equal(t, 1, findMin([]int{10, 1, 10, 10, 10, 10, 10}))
	assert.Equal(t, 3, findMin([]int{3, 3, 3, 3, 3, 3, 3, 4, 5}))
	assert.Equal(t, 0, findMin([]int{3, 4, 5, 0, 1, 2}))
	assert.Equal(t, 0, findMin([]int{2, 2, 2, 0, 1}))
	assert.Equal(t, 0, findMin([]int{2, 2, 2, 2, 2, 0, 1}))
}
