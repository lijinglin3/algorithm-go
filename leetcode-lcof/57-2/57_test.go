package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContinuousSequence(t *testing.T) {
	assert.Equal(t, [][]int{{2, 3, 4}, {4, 5}}, findContinuousSequence(9))
	assert.Equal(t, [][]int{{1, 2}}, findContinuousSequence(3))
	assert.Equal(t, [][]int{}, findContinuousSequence(2))
	assert.Equal(t, [][]int{}, findContinuousSequence(1))
}
