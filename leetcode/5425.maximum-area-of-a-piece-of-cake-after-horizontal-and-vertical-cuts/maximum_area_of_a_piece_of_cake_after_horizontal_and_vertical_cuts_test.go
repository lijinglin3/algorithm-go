package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 4, maxArea(5, 4, []int{1, 2, 3}, []int{1, 3}))
}
