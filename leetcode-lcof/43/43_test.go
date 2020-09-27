package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDigitOne(t *testing.T) {
	assert.Equal(t, 0, countDigitOne(0))
	assert.Equal(t, 1, countDigitOne(1))
	assert.Equal(t, 1, countDigitOne(5))
	assert.Equal(t, 2, countDigitOne(10))
}
