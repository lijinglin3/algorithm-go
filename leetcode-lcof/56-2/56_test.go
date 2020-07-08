package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	assert.Equal(t, 9, singleNumber([]int{9}))
	assert.Equal(t, 1, singleNumber([]int{9, 1, 7, 9, 7, 9, 7}))
}
