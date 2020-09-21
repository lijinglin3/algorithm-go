package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{
		{-1, 0, 0, 1},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
	}, fourSum([]int{0, 1, 0, -1, 0, -2, 2, 2}, 0))
}
