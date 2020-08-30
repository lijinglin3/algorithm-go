package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 3, findMaxConsecutiveOnes([]int{1, 0, 0, 1, 1, 0}))
}
