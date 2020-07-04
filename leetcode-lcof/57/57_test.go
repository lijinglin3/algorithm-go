package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	assert.ElementsMatch(t, []int{2, 7}, twoSum([]int{1, 2, 7, 11, 15}, 9))
	assert.ElementsMatch(t, nil, twoSum([]int{2}, 9))
}
