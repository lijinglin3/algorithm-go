package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 687995182, fib(100))
	assert.Equal(t, 586268941, fib(50))
	assert.Equal(t, 1, fib(1))
	assert.Equal(t, 0, fib(0))
}
