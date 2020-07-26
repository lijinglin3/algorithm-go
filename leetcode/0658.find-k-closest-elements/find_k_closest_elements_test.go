package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestElements(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	assert.Equal(t, []int{3, 5, 7, 9}, findClosestElements([]int{1, 3, 5, 7, 9}, 4, 6))
}
