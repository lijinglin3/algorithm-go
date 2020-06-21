package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, runningSum([]int{1, 1, 1, 1, 1}))
}
