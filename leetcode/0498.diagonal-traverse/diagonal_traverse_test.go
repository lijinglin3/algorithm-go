package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiagonalOrder(t *testing.T) {
	assert.Equal(t, ([]int)(nil), findDiagonalOrder([][]int{}))
	assert.Equal(t, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}, findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
