package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorOperation(t *testing.T) {
	assert.Equal(t, 8, xorOperation(5, 0))
	assert.Equal(t, 8, xorOperation(4, 3))
	assert.Equal(t, 7, xorOperation(1, 7))
	assert.Equal(t, 2, xorOperation(10, 5))
}
