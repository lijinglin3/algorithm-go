package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	assert.Equal(t, 0, addDigits(0))
	assert.Equal(t, 1, addDigits(1))
	assert.Equal(t, 9, addDigits(9))
	assert.Equal(t, 1, addDigits(100))

	assert.Equal(t, 0, addDigits2(0))
	assert.Equal(t, 1, addDigits2(1))
	assert.Equal(t, 9, addDigits2(9))
	assert.Equal(t, 1, addDigits2(100))
}
