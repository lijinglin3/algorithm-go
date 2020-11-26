package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()
	assert.True(t, c.Empty())

	c.Push(1)
	assert.False(t, c.Empty())
	assert.Equal(t, 1, c.Peek())

	c.Push(2)
	assert.False(t, c.Empty())
	assert.Equal(t, 1, c.Peek())

	assert.Equal(t, 1, c.Pop())
	assert.False(t, c.Empty())
	assert.Equal(t, 2, c.Peek())

	assert.Equal(t, 2, c.Pop())
	assert.True(t, c.Empty())
}
