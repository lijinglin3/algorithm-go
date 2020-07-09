package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()

	assert.Equal(t, true, c.Empty())
	assert.Equal(t, -1, c.Peek())
	assert.Equal(t, -1, c.Pop())

	c.Push(1)
	assert.Equal(t, false, c.Empty())
	assert.Equal(t, 1, c.Peek())

	c.Push(2)
	assert.Equal(t, false, c.Empty())
	assert.Equal(t, 1, c.Peek())
	assert.Equal(t, 1, c.Pop())
	assert.Equal(t, 2, c.Peek())
	assert.Equal(t, 2, c.Pop())
	assert.Equal(t, -1, c.Peek())
	assert.Equal(t, -1, c.Pop())
	assert.Equal(t, true, c.Empty())
}
