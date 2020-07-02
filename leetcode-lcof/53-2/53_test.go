package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, 0, missingNumber([]int{1}))
	assert.Equal(t, 2, missingNumber([]int{0, 1, 3}))
	assert.Equal(t, 8, missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 9}))
	assert.Equal(t, 9, missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	assert.Equal(t, 0, missingNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
