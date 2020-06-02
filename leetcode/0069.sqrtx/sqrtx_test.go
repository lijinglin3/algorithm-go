package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 0, mySqrt(0))
	assert.Equal(t, 1, mySqrt(1))
	assert.Equal(t, 1, mySqrt(2))
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 9, mySqrt(99))
	assert.Equal(t, 10, mySqrt(100))
}
