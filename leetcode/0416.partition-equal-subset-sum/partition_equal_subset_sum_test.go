package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	assert.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
	assert.Equal(t, false, canPartition([]int{1}))
	assert.Equal(t, false, canPartition([]int{1, 2, 3, 5}))
}
