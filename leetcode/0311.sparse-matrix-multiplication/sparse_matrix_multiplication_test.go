package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, [][]int{
		{7, 0, 0},
		{-7, 0, 3},
	}, multiply([][]int{
		{1, 0, 0},
		{-1, 0, 3},
	}, [][]int{
		{7, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}))
}
