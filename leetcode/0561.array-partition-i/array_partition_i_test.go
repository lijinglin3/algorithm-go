package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayPairSum(t *testing.T) {
	assert.Equal(t, 9, arrayPairSum([]int{3, 4, 1, 2, 5, 6}))
}
