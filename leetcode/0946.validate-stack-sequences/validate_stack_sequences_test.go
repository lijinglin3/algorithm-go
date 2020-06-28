package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateStackSequences(t *testing.T) {
	assert.Equal(t, true, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
	assert.Equal(t, false, validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
}
