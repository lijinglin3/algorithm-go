package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	assert.Equal(t, -1, s.Peek())

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 1, s.Peek())
	s.Pop()
	assert.Equal(t, 2, s.Peek())
	s.Pop()
	assert.Equal(t, 3, s.Peek())
	s.Pop()
	assert.Equal(t, -1, s.Peek())
	s.Pop()
}
