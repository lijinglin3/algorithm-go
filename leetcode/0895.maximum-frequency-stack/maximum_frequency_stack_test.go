package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c := Constructor()

	for _, v := range []int{5, 7, 5, 7, 4, 5} {
		c.Push(v)
	}

	assert.Equal(t, 5, c.Pop())
	assert.Equal(t, 7, c.Pop())
	assert.Equal(t, 5, c.Pop())
	assert.Equal(t, 4, c.Pop())
	assert.Equal(t, 7, c.Pop())
	assert.Equal(t, 5, c.Pop())
}
