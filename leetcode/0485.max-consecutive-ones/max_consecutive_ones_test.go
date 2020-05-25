package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 1, findMaxConsecutiveOnes([]int{1}))
	assert.Equal(t, 4, findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1}))
}
