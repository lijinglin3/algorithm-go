package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSumCount(t *testing.T) {
	assert.Equal(t, 2, fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
