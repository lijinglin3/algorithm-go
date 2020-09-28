package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNthDigit(t *testing.T) {
	assert.Equal(t, 1, findNthDigit(1))
	assert.Equal(t, 1, findNthDigit(10))
	assert.Equal(t, 5, findNthDigit(100))
	assert.Equal(t, 3, findNthDigit(1000))
}
