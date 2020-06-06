package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrailingZeroes(t *testing.T) {
	assert.Equal(t, 0, trailingZeroes(3))
	assert.Equal(t, 1, trailingZeroes(5))
	assert.Equal(t, 7, trailingZeroes(30))
}
