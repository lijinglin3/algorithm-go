package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanMakeArithmeticProgression(t *testing.T) {
	assert.Equal(t, false, canMakeArithmeticProgression([]int{1, 2, 4}))
	assert.Equal(t, true, canMakeArithmeticProgression([]int{16, 20}))
	assert.Equal(t, true, canMakeArithmeticProgression([]int{-68, -96, -12, -40, 16}))
}
