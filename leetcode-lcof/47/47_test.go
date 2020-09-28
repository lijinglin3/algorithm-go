package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	assert.Equal(t, 12,
		maxValue([][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}))
}
