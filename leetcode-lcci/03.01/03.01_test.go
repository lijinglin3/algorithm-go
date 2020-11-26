package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor(2)
	assert.True(t, c.IsEmpty(0))
	assert.Equal(t, -1, c.Peek(0))
	assert.Equal(t, -1, c.Peek(1))
	assert.Equal(t, -1, c.Peek(2))

	c.Push(0, 1)
	assert.False(t, c.IsEmpty(0))
	assert.Equal(t, 1, c.Peek(0))
	assert.Equal(t, -1, c.Peek(1))
	assert.Equal(t, -1, c.Peek(2))

	c.Push(0, 2)
	assert.False(t, c.IsEmpty(0))
	assert.Equal(t, 2, c.Peek(0))
	assert.Equal(t, -1, c.Peek(1))
	assert.Equal(t, -1, c.Peek(2))

	c.Push(0, 3)
	assert.False(t, c.IsEmpty(0))
	assert.Equal(t, 2, c.Peek(0))
	assert.Equal(t, -1, c.Peek(1))
	assert.Equal(t, -1, c.Peek(2))

	assert.Equal(t, 2, c.Pop(0))
	assert.False(t, c.IsEmpty(0))
	assert.Equal(t, 1, c.Pop(0))
	assert.True(t, c.IsEmpty(0))
	assert.Equal(t, -1, c.Pop(0))
	assert.True(t, c.IsEmpty(0))
}
