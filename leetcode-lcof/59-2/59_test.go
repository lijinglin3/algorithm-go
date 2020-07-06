package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	q := Constructor()

	assert.Equal(t, -1, q.PopFront())
	assert.Equal(t, -1, q.MaxValue())

	q.PushBack(1)
	assert.Equal(t, 1, q.MaxValue())

	q.PushBack(2)
	assert.Equal(t, 2, q.MaxValue())
	assert.Equal(t, 1, q.PopFront())
	assert.Equal(t, 2, q.MaxValue())

	q.PushBack(1)
	assert.Equal(t, 2, q.MaxValue())
	assert.Equal(t, 2, q.PopFront())
	assert.Equal(t, 1, q.MaxValue())
	assert.Equal(t, 1, q.PopFront())
	assert.Equal(t, -1, q.MaxValue())
}
