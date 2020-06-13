package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraySum(t *testing.T) {
	example := []int{1, 1, 1}
	assert.Equal(t, 3, subarraySum(example, 1))
	assert.Equal(t, 2, subarraySum(example, 2))
	assert.Equal(t, 1, subarraySum(example, 3))
	assert.Equal(t, 0, subarraySum(example, 4))

	assert.Equal(t, 3, subarraySum2(example, 1))
	assert.Equal(t, 2, subarraySum2(example, 2))
	assert.Equal(t, 1, subarraySum2(example, 3))
	assert.Equal(t, 0, subarraySum2(example, 4))
}
