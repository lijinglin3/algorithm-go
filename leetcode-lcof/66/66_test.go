package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructArr(t *testing.T) {
	assert.Equal(t, []int{}, constructArr([]int{}))
	assert.Equal(t, []int{1}, constructArr([]int{5}))
	assert.Equal(t, []int{120, 60, 40, 30, 24}, constructArr([]int{1, 2, 3, 4, 5}))
}
