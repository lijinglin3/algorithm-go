package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -3, divide(10, -3))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
