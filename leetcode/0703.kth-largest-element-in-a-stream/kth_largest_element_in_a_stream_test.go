package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	c1 := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, c1.Add(3))
	assert.Equal(t, 5, c1.Add(5))
	assert.Equal(t, 5, c1.Add(10))
	assert.Equal(t, 8, c1.Add(9))
	assert.Equal(t, 8, c1.Add(4))

	c2 := Constructor(1, []int{})
	assert.Equal(t, -3, c2.Add(-3))
	assert.Equal(t, -2, c2.Add(-2))
	assert.Equal(t, -2, c2.Add(-4))
	assert.Equal(t, 0, c2.Add(0))
	assert.Equal(t, 4, c2.Add(4))
}
