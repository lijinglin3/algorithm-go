package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStraight(t *testing.T) {
	assert.Equal(t, true, isStraight([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, true, isStraight([]int{0, 0, 3, 4, 7}))
	assert.Equal(t, false, isStraight([]int{0, 0, 3, 3, 7}))
}
