package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRectangleOverlap(t *testing.T) {
	assert.Equal(t, true, isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	assert.Equal(t, false, isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
