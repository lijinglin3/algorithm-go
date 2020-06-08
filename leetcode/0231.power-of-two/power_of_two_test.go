package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	assert.Equal(t, false, isPowerOfTwo(-1))
	assert.Equal(t, false, isPowerOfTwo(0))
	assert.Equal(t, true, isPowerOfTwo(1))
	assert.Equal(t, true, isPowerOfTwo(16))
	assert.Equal(t, false, isPowerOfTwo(218))

	assert.Equal(t, false, isPowerOfTwo2(-1))
	assert.Equal(t, false, isPowerOfTwo2(0))
	assert.Equal(t, true, isPowerOfTwo2(1))
	assert.Equal(t, true, isPowerOfTwo2(16))
	assert.Equal(t, false, isPowerOfTwo2(218))
}
