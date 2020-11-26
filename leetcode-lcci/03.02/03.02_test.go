package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()

	c.Push(-2)
	assert.Equal(t, -2, c.Top())
	assert.Equal(t, -2, c.GetMin())

	c.Push(0)
	assert.Equal(t, 0, c.Top())
	assert.Equal(t, -2, c.GetMin())

	c.Push(-3)
	assert.Equal(t, -3, c.Top())
	assert.Equal(t, -3, c.GetMin())

	c.Pop()
	assert.Equal(t, 0, c.Top())
	assert.Equal(t, -2, c.GetMin())

	c.Pop()
	assert.Equal(t, -2, c.Top())
	assert.Equal(t, -2, c.GetMin())
}
