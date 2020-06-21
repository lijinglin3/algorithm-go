package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {
	q := Constructor([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	assert.Equal(t, 1, q.GetValue(0, 2))
	q.UpdateSubrectangle(0, 0, 3, 2, 5)
	assert.Equal(t, 5, q.GetValue(0, 2))
	q.UpdateSubrectangle(3, 0, 3, 2, 10)
	assert.Equal(t, 10, q.GetValue(3, 1))
	assert.Equal(t, 5, q.GetValue(0, 2))
}
