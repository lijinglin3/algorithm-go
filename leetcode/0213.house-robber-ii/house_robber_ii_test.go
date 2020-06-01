package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 0, rob([]int{}))
	assert.Equal(t, 1, rob([]int{1}))
	assert.Equal(t, 3, rob([]int{2, 3, 2}))
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
}
