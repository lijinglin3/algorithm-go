package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor([]int{-2, 0, 3, -5, 2, -1})
	assert.Equal(t, -2, c.SumRange(0, 0))
	assert.Equal(t, 1, c.SumRange(0, 2))
	assert.Equal(t, -1, c.SumRange(2, 5))
	assert.Equal(t, -3, c.SumRange(0, 5))
}
