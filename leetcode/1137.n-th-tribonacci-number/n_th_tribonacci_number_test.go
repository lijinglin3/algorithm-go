package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTribonacci(t *testing.T) {
	assert.Equal(t, 149, tribonacci(10))
	assert.Equal(t, 66012, tribonacci(20))
	assert.Equal(t, 2082876103, tribonacci(37))
}
