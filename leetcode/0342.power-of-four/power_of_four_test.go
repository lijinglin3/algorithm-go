package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.Equal(t, true, isPowerOfFour(16))
	assert.Equal(t, false, isPowerOfFour(5))
}
