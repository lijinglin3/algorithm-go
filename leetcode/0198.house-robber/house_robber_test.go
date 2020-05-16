package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 0, rob([]int{}))
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
}
