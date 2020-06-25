package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.Equal(t, true, searchMatrix(matrix, 5))
	assert.Equal(t, true, searchMatrix(matrix, 18))
	assert.Equal(t, false, searchMatrix(matrix, 20))
	assert.Equal(t, false, searchMatrix([][]int{}, 20))
	assert.Equal(t, false, searchMatrix([][]int{{}}, 20))
}
