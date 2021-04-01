package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor(2)

	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Push(4)

	assert.Equal(t, 4, c.Pop())
	assert.Equal(t, -1, c.PopAt(2))
	assert.Equal(t, 2, c.PopAt(0))
	assert.Equal(t, 1, c.PopAt(0))
	assert.Equal(t, 3, c.PopAt(0))

	c = Constructor(1)
	c.Push(1)
	c.Push(2)

	assert.Equal(t, 2, c.PopAt(1))
	assert.Equal(t, 1, c.Pop())
	assert.Equal(t, -1, c.Pop())

	c = Constructor(0)
	c.Push(1)
	c.Push(2)
	c.Push(3)

	assert.Equal(t, -1, c.Pop())
	assert.Equal(t, -1, c.Pop())
	assert.Equal(t, -1, c.Pop())
}
