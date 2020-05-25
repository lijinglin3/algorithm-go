package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{2, 3}, twoSum([]int{5, 25, 75}, 100))
	assert.Equal(t, []int{2, 3}, twoSum([]int{5, 25, 75, 100}, 100))
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 7, 11, 15}, 9))
}
