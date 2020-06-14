package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	assert.Equal(t, false, isPowerOfThree(0))
	assert.Equal(t, false, isPowerOfThree(26))
	assert.Equal(t, true, isPowerOfThree(27))
}
