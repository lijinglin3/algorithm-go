package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int(nil), threeSum([]int{0, 1}))
	assert.ElementsMatch(t, [][]int{{-1, 0, 1}}, threeSum([]int{-1, 0, 1}))
	assert.ElementsMatch(t, [][]int{{-1, 0, 1}, {-1, -1, 2}, {-4, 2, 2}}, threeSum([]int{-1, 0, 0, 1, 2, 2, -1, -4}))

}
