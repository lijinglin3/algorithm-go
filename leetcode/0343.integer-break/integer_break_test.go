package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerBreak(t *testing.T) {
	assert.Equal(t, 1, integerBreak(2))
	assert.Equal(t, 36, integerBreak(10))
}
